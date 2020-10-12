import Vue from 'vue';
import Vuex from 'vuex';
import APIClient from './apiClient';

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    user: {authenticated: false, name: '', avatarURL: ''},
  },
  getters: {
    // By default, Vuex getters accept two arguments.
    //
    // state — the state object for our application;
    // getters — the store.getters object, meaning that we can call other getters in our store.
    //
    // Example:
    //
    // lastName(state, getters) {
    //   return state.user.fullName.replace(getters.firstName, '');
    // },
    //
    // You can also pass arguments to getters:
    //
    // prefixedName: (state, getters) => (prefix) => {
    //   return prefix + getters.lastName;
    // }
    isAuthenticated: state => !!state.user.authenticated,
  },
  // Note: Mutations must be synchronous!
  mutations: {
    setUserAuthenticated(state, authenticated) {
      state.user.authenticated = authenticated;
    },
    setUserName(state, name) {
      state.user.name = name;
    },
    setUserAvatarURL(state, url) {
      state.user.avatarURL = url;
    }
  },
  actions: {
    getUser(context) {
      APIClient.getUser()
          .then((resp) => {
            context.commit('setUserAuthenticated', true);
            context.commit('setUserName', resp.name);
            context.commit('setUserAvatarURL', resp.avatar_url);
          })
          .catch(err => {
            console.log(err);
          })
    },
  }
});
