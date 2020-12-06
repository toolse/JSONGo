import Vue from 'vue';
import App from './App.vue';

import HeyUI from 'heyui';
Vue.use(HeyUI);

import 'heyui/themes/index.less';
import '@/assets/css/dome.css';
import router from './router'

new Vue({
    router,
    render: h => h(App)
}).$mount('#app');

Vue.config.productionTip = false;
// Vue.config.devtools = false;
