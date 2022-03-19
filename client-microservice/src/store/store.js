import Vue from 'vue'
import Vuex from 'vuex'

import setAuthToken from '../utils/setAuthToken.js';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    user: null,
    token: null,
  },
  mutations: {
    setUser(state, user) {
      state.user = user;
    },
    setToken(state, token) {
      state.token = token;
    },
  },
  actions: {
    setUser({ commit }, user) {
      commit('setUser', user);
    },
    setToken({ commit }, token) {
      if (token) {
        localStorage.setItem('token', token);
      } else {
        localStorage.removeItem('token');
      }
      setAuthToken(token);
      commit('setToken', token);
    },
  },
  getters: {
    user(state) {
      return state.user;
    },
    token(state) {
      return state.token;
    },
  }
})
