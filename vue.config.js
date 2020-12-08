module.exports = {
    devServer: {
        port: 8000
    },
    // 禁止生成map镜像文件
    productionSourceMap: false,
    // 判断环境：生产环境production、开发环境development
    publicPath: process.env.NODE_ENV === 'production' ? './' : '/',
    pluginOptions: {
        'style-resources-loader': {
            preProcessor: 'less',
            patterns: ['./node_modules/heyui/themes/var.less']
        }
    }
};
