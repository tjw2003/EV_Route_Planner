import Vue from 'vue'
import App from './App.vue'
import router from './router'; // 引入路由配置文件
import Element from 'element-ui'
import axios from 'axios';
import 'element-ui/lib/theme-chalk/index.css';
Vue.config.productionTip = false
Vue.use(Element)

new Vue({
  router,
  render: h => h(App),
}).$mount('#app')

const axiosInstance = axios.create({
  baseURL: 'http://localhost:8888', // 后端 API 的基础 URL
  headers: {
    'Content-Type': 'application/json', // 设置请求头为 JSON
    'Access-Control-Allow-Origin': 'http://localhost:8080', // 允许的来源
  },
});
// 将 Axios 实例绑定到 Vue 实例的原型上，方便在组件中使用
Vue.prototype.$axios = axiosInstance;
