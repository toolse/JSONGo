<template>
    <Layout class="layout-demo-1-vue">
        <HHeader> JSON / XML Convert</HHeader>
        <Content>
            <Row :space="20">
                <Cell width="12">
                    <codemirror v-model="codeContent" :options="cmOptions" />
                </Cell>
                <Cell width="12">
                    <div class="h-panel">
                        <div class="h-panel-bar">
                            <span class="h-panel-title">输出结果</span>
                            <!-- <Button color="blue" @click="sourcecode = ''">清空</Button>
                            <Button color="primary" @click="convert">复制</Button>
                            <Button color="primary" @click="convert">转换</Button> -->
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
export default {
    components: {},
    data() {
        return {
            codeContent: '',
            sourcecode: '',
            loading: false,
            cmOptions: {
                tabSize: 4,
                mode: 'json/go/xml',
                theme: 'base16-dark',
                lineNumbers: true,
                line: true,
                highlightDifferences: true
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
    height: 100%;
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
        padding: 20px;
        .h-row {
            .CodeMirror {
                height: 80vh;
            }
            .h-btn-group {
                float: right;
            }
            .h-panel {
                height: 80vh;
                overflow: hidden; // auto
                .h-panel-body {
                    overflow: auto;
                    height: inherit;
                    padding: 0;
                    pre {
                        height: inherit;
                        .hljs {
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
    }
}
</style>
