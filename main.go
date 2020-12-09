package main

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"flag"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/m-zajac/json2go"
	"time"

	"github.com/miku/zek"
	"go/format"
	"log"
	"net/http"
	"strings"
)

const (
	JSON    = 1
	XML     = 2
	UNKNOWN = -1
)

var (
	port string
)

func main() {
	flag.StringVar(&port, "port", ":8080", "-port=:8080")
	flag.Parse()
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	// 开启跨域
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET"},
		AllowHeaders:     []string{"Origin"},
		AllowCredentials: true,
		ExposeHeaders:    []string{"Content-Length"},
		MaxAge:           12 * time.Hour,
	}))
	router.POST("/postOriginContent", contentHandler)
	if err := router.Run(port); err != nil {
		log.Fatal(err)
	}
}

func contentHandler(g *gin.Context) {
	content := g.PostForm("content")
	if content == "" {
		err := g.AbortWithError(http.StatusMethodNotAllowed, errors.New("content cannot be empty"))
		if err != nil {
			log.Println(err)
		}
		return
	}
	content = strings.TrimSpace(content)
	content = parseContent(content)
	if content == "" {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"msg": "content parse failed",
		})
	}

	g.JSON(http.StatusOK, gin.H{
		"content": content,
		"length":  len(content),
	})
}

func parseContent(content string) string {
	var ret string
	switch formatGuessing(content) {
	case JSON:
		res, err := genStructFromJSON(content)
		if err != nil {
			ret = err.Error()
		} else {
			ret = res
		}
	case XML:
		res, err := genStructFromXML(content)
		if err != nil {
			ret = err.Error()
		} else {
			ret = res
		}
	case UNKNOWN:
		ret = "unknown content"
	}
	return ret
}

func formatGuessing(v string) int {
	var tmp interface{}
	if err := json.Unmarshal([]byte(v), &tmp); err == nil {
		return JSON
	}
	if err := xml.Unmarshal([]byte(v), &tmp); err == nil {
		return XML
	}
	return UNKNOWN
}

func genStructFromJSON(json string) (string, error) {
	parser := json2go.NewJSONParser("AutoGenerator")
	if err := parser.FeedBytes([]byte(json)); err != nil {
		return "", err
	}
	return parser.String(), nil
}

func genStructFromXML(xml string) (string, error) {
	root := &zek.Node{}
	if _, err := root.ReadFrom(strings.NewReader(xml)); err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	sw := zek.NewStructWriter(buf)
	if err := sw.WriteNode(root); err != nil {
		return "", err
	}
	b, err := format.Source(buf.Bytes())
	if err != nil {
		return "", err
	}
	return string(b), nil
}
