import Vue from 'vue'
import App from './App.vue'
import VueRouter from 'vue-router' // 导入 Vue Router 模块

import Element from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css';
import Login from './views/login.vue'
import Planner from './views/Planner.vue'
import SignUp from './views/Signup.vue'

Vue.config.productionTip = false
Vue.use(Element)
Vue.use(VueRouter) // 使用 Vue Router 模块

// 创建路由实例，并配置路由规则
const router = new VueRouter({
  routes: [
    {
      path: '/login',
      component: Login
    },
    {
      path: '/planner',
      component: Planner
    },
    {
      path: '/signup',
      component: SignUp
    },
  ]
})

new Vue({
  render: h => h(App),
  router // 将路由器添加到 Vue 实例中
}).$mount('#app')
