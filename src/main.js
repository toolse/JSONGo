import Vue from 'vue';
import App from './App.vue';

// UI框架
import HeyUI from 'heyui';
Vue.use(HeyUI);
import 'heyui/themes/index.less';
import '@/assets/css/dome.css';

// vue-highlightjs
import VueHighlightJS from 'vue-highlightjs';
Vue.use(VueHighlightJS);
import './assets/styles/gruvbox-dark.css';

import VueCodemirror from 'vue-codemirror';
import 'codemirror/lib/codemirror.css';
Vue.use(VueCodemirror);

// Axios
import axios from 'axios';
axios.defaults.baseURL = 'http://localhost:8080';
Vue.prototype.$axios = axios;

import router from './router';

new Vue({
    router,
    render: h => h(App)
}).$mount('#app');

Vue.config.productionTip = false;
// Vue.config.devtools = false;
