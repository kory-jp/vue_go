import Vue from 'vue'
import Vuex from 'vuex'
import { account } from './modules/accounts'

Vue.use(Vuex)

export default new Vuex.Store({
  // state,
  // getters: {
  // },
  // mutations: {
  // },
  // actions: {
  // },
  modules: {
    account
  }
})
