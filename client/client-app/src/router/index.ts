import Vue from 'vue'
import VueRouter, { RouteConfig } from 'vue-router'
import HomeView from '../views/HomeView.vue'
import Top from '../components/pages/Top.vue'

Vue.use(VueRouter)

const routes: Array<RouteConfig> = [
  {
    path: '/',
    name: 'home',
    component: HomeView
  },
  {
    path: '/top',
    name: 'top',
    component: Top
  }
]

const router = new VueRouter({
  mode: 'history',
  routes
})

export default router
