<template>
    <Layout class="layout-demo-1-vue">
        <HHeader>
            <img src="../assets/logo.png" />
            <Button color="primary">GREEN按钮</Button>
            <Button color="primary">GREEN按钮</Button>
            <Button color="primary">GREEN按钮</Button>
            <DropdownMenu button class="h-btn-primary" @click="trigger" :datas="param">开始</DropdownMenu>
            <Button color="primary">GREEN按钮</Button>
        </HHeader>
        <Content>Content</Content>
        <HFooter>Footer</HFooter>
    </Layout>
</template>

<script>
export default {
    data() {
        return {
            param: [
                { title: '测试1', key: 'test1' },
                { title: '测试2', key: 'test2' },
                { title: '测试3', key: 'test3' }
            ]
        };
    },
    methods: {
        trigger(code) {
            this.$Message.info(`您点击了${code}`);
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
        img {
            height: 32px;
            margin: 0 25px;
        }
        button,
        .h-btn-primary {
            margin-right: 15px;
        }
    }

    .h-layout-footer {
        background: @primary2-color;
        color: #fff;
        line-height: 64px;
        height: 64px;
        text-align: center;
    }
    .h-layout-content {
        background: @primary-color;
        color: #fff;
        min-height: 120px;
        line-height: 120px;
        text-align: center;
    }
}
</style>
