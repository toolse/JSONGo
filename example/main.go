package main

import (
	"fmt"
	"github.com/m-zajac/json2go"
	"io/ioutil"
	"log"
)

//func main() {
//	var reader io.Reader = os.Stdin
//	root := &zek.Node{}
//	if _, err := root.ReadFrom(reader); err != nil {
//		log.Fatal(err)
//	}
//	var buf bytes.Buffer
//	sw := zek.NewStructWriter(&buf)
//	sw.WithComments = true
//	sw.WithJSONTags = true
//	sw.Strict = false
//	sw.UniqueExamples = false
//	sw.Compact = true
//
//	if err := sw.WriteNode(root); err != nil {
//		log.Fatal(err)
//	}
//
//	b, err := format.Source(buf.Bytes())
//	if err != nil {
//		log.Fatal(err)
//	}
//	fmt.Println(string(b))
//}

func main() {
	data, err := ioutil.ReadFile("./example.json")
	if err != nil {
		log.Fatal(err)
	}
	parser := json2go.NewJSONParser("dock")
	err = parser.FeedBytes(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(parser.String())
}
