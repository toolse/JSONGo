<template>
    <Layout class="layout-demo-1-vue">
        <HHeader>
            <img src="../assets/logo.png" />
            <Button color="primary">在线解析</Button>
            <Button color="primary">JSON压缩转义</Button>
            <Button color="primary">JS混淆加密</Button>
            <DropdownMenu button class="h-btn-primary" @click="trigger" :datas="param">JSON生成实体类</DropdownMenu>
            <Button color="primary">代码对比</Button>
        </HHeader>
        <Content>
            <Row :space-x="19" :space-y="5">
                <Cell width="12">
                    <textarea v-model="textarea" placeholder="在此填入json字符串或XML字符串..."></textarea>
                </Cell>
                <Cell width="12"> <textarea v-model="textarea" placeholder="在此填入json字符串或XML字符串..."></textarea> </Cell>
            </Row>
        </Content>
        <HFooter>© 2020-2021 JSON-GO All right reserved</HFooter>
    </Layout>
</template>

<script>
export default {
    data() {
        return {
            param: [
                { title: 'JSON生成Java实体类', key: 'Java' },
                { title: 'JSON生成C#实体类', key: 'C#' }
            ],
            textarea: ''
        };
    },
    methods: {
        trigger(code) {
            this.$Message.info(`您点击了${code}`);
        },
        copyText() {
            this.$Clipboard({
                text: this.text
            });
        }
    }
};
</script>

<style lang="less" scoped>
@primary1-color: fade(@primary-color, 88%);
@primary2-color: fade(@primary-color, 80%);

.layout-demo-1-vue {
    height: 100%;
    .h-layout-header {
        display: flex;
        align-items: center;
        box-shadow: 0 0 5px #d3d3d3;
        img {
            height: 32px;
            margin: 0 40px;
            cursor: pointer;
        }
        button,
        .h-btn-primary {
            margin-right: 20px;
            box-shadow: none;
        }
    }
    .h-layout-content {
        padding: 20px;
        text-align: center;
        textarea {
            height: 80vh;
            width: 100%;
            border: 4px solid #e3f5ed;
            border-radius: 4px;
            resize: none;
        }
        .h-col {
            height: 100%;
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
