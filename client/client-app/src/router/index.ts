import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Top from '../components/pages/Top.vue'
import Login from '@/components/pages/auth/Login.vue'
import Register from '@/components/pages/auth/Register.vue'
import Articles from '@/components/pages/articles/Articles.vue'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'top',
    component: Top
  },
  {
    path: '/register',
    name: 'register',
    component: Register,
  },
  {
    path: '/login',
    name: 'login',
    component: Login
  },
  {
    path: '/articles',
    name: 'articles',
    component: Articles
  }
]

const router = new VueRouter({
  mode: 'history',
  routes
})

export default router
