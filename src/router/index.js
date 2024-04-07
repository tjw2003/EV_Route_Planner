import Vue from 'vue';
import Router from 'vue-router';
import Planner from '@/views/Planner.vue';
import Login from '@/views/Login.vue';
import Register from'@/views/Register.vue'

Vue.use(Router);

const routes = [
  {
    path: '/Planner',
    name: 'Planner',
    component: Planner
  },
  {
    path: '/login',
    name: 'Login',
    component: Login
  },
  {
    path: '/register',
    name: 'Register',
    component: Register
  }, 
];

const router = new Router({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
});

export default router;
