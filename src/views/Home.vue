<template>
    <Layout class="layout-demo-1-vue">
        <HHeader>JSON / XML Convert</HHeader>
        <Content>
            <Row :space="20">
                <Cell width="12">
                    <codemirror v-model="codeContent" :options="cmOptions" />
                </Cell>
                <Cell width="12">
                    <div class="h-panel">
                        <div class="h-panel-bar">
                            <span class="h-panel-title">输出结果</span>
                            <div class="h-btn-group">
                                <button class="h-btn" @click="convert"><i class="h-icon-refresh"></i><span>转换</span></button>
                                <button class="h-btn" @click="copyCode"><i class="h-icon-complete"></i><span>复制</span></button>
                                <button class="h-btn" @click="sourcecode = ''"><i class="h-icon-trash"></i><span>清空</span></button>
                            </div>
                        </div>
                        <div class="h-panel-body">
                            <pre v-highlightjs="sourcecode"><code class="go"></code></pre>
                        </div>
                    </div>
                </Cell>
            </Row>
        </Content>
        <HFooter>© 2020-2021 JSON / XML Convert All right reserved</HFooter>
    </Layout>
</template>

<script>
const qs = require('qs');
// 引入JS语言文件
import 'codemirror/mode/xml/xml';
import 'codemirror/mode/javascript/javascript';
import '../assets/js/setting';

export default {
    data() {
        return {
            codeContent: '',
            sourcecode: '',
            cmOptions: {
                tabSize: 4, // tab
                styleActiveLine: true, // 高亮选中行
                lineNumbers: true, // 显示行号
                styleSelectedText: true,
                line: true,
                foldGutter: true, // 块槽
                gutters: ['CodeMirror-lint-markers', 'CodeMirror-linenumbers', 'CodeMirror-foldgutter'],
                highlightSelectionMatches: { showToken: /\w/, annotateScrollbar: true }, // 可以启用该选项来突出显示当前选中的内容的所有实例
                modes: [
                    // 模式, 可查看 codemirror/mode 中的所有模式
                    'application/json',
                    {
                        ext: 'xml'
                    }
                ],
                // hint.js options
                hintOptions: {
                    // 当匹配只有一项的时候是否自动补全
                    completeSingle: false
                },
                // 快捷键 可提供三种模式 sublime、emacs、vim
                keyMap: 'sublime',
                matchBrackets: true,
                showCursorWhenSelecting: true,
                theme: 'gruvbox-dark', // 主题
                extraKeys: { Ctrl: 'autocomplete' } // 可以用于为编辑器指定额外的键绑定，以及keyMap定义的键绑定
            }
        };
    },
    methods: {
        // 复制
        copyCode() {
            if (this.sourcecode.length === 0) {
                return this.$Message.warn('没有内容');
            }
            this.$Clipboard({
                text: this.sourcecode
            });
        },
        // 开始转换
        async convert() {
            try {
                if (this.codeContent.length === 0) {
                    return this.$Message.warn('内容不能为空');
                }
                const { data: res } = await this.$axios.post(
                    '/postOriginContent',
                    qs.stringify({
                        content: this.codeContent
                    })
                );
                if (res.content === 'unknown content') {
                    return this.$Message.error(res.content);
                }
                this.sourcecode = res.content;
                this.$Message.success('转换成功');
            } catch (error) {
                return this.$Message.error(`${error}`);
            }
        }
    }
};
</script>

<style lang="less">
@primary1-color: fade(@primary-color, 88%);
@primary2-color: fade(@primary-color, 80%);

.layout-demo-1-vue {
    .h-layout-header {
        background: @primary1-color;
        letter-spacing: 1px;
        color: white;
        box-shadow: 0 0 5px #d3d3d3;
        text-align: center;
        font-size: 18px;
        font-weight: bold;
    }
    .h-layout-content {
        height: 100%;
        margin: 20px;
        .h-row {
            .CodeMirror {
                border-radius: 4px;
                height: 75vh;
            }
            .h-panel {
                height: 75vh;
                overflow: hidden;
                position: relative;
                .h-panel-bar {
                    width: 100%;
                    position: absolute;
                    background: white;
                    display: flex;
                    justify-content: space-between;
                }
                .h-panel-body {
                    overflow: auto;
                    height: inherit;
                    padding: 0;
                    pre {
                        height: inherit;
                        .hljs {
                            padding: 60px 0 0 10px;
                            height: inherit;
                        }
                    }
                }
            }
        }
    }
    .h-layout-footer {
        background: @primary2-color;
        color: #fff;
        line-height: 64px;
        height: 64px;
        text-align: center;
        position: fixed;
        left: 0;
        right: 0;
        bottom: 0;
    }
}
</style>
