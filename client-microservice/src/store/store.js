import Vue from 'vue'
import Vuex from 'vuex'

import setAuthToken from '../utils/setAuthToken.js';
import { setLocalStorageLanguage, getLocalStorageLanguage } from '../utils/userLanguage.js';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    language: getLocalStorageLanguage(),
    user: null,
    token: null,
  },
  mutations: {
    setLanguage(state, language) {
      state.language = language;
      setLocalStorageLanguage(state.language);
    },
    setUser(state, user) {
      state.user = user;
    },
    setToken(state, token) {
      state.token = token;
    },
  },
  actions: {
    setLanguage({ commit }, language) {
      commit('setLanguage', language);
    },
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
    language(state) {
      return state.language;
    },
    user(state) {
      return state.user;
    },
    token(state) {
      return state.token;
    },
  }
})
