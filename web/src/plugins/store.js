import Vue from "vue";
import Vuex from "vuex";
import APIClient from "./apiClient";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {
    user: { authenticated: false, record: {}, tokens: [] }
  },

  getters: {
    isAuthenticated: state => {
      return localStorage.isLoggedIn === "1" || state.user.authenticated;
    },

    userRecord: state => {
      return state.user.record;
    },

    userTokens: state => {
      return state.user.tokens;
    }
  },

  mutations: {
    setUserAuthenticated(state, authenticated) {
      state.user.authenticated = authenticated;
    },

    setUser(state, record) {
      state.user.record = record;
    },

    setUserTokens(state, tokens) {
      state.user.tokens = tokens;
    }
  },

  actions: {
    getUser(context) {
      APIClient.getUser()
        .then(resp => {
          context.commit("setUser", resp);
          context.commit("setUserAuthenticated", true);
          localStorage.isLoggedIn = "1";
        })
        .catch(err => {
          console.log(err);
          context.commit("setUser", {});
          context.commit("setUserAuthenticated", false);
          localStorage.removeItem("isLoggedIn");
        });
    },

    updateUser(context, user) {
      return new Promise((resolve, reject) => {
        APIClient.updateUser(user)
          .then(resp => {
            context.commit("setUser", resp);
            resolve();
          })
          .catch(err => {
            console.log(err);
            if (err.response) {
              reject(err.response.data.error);
            }

            reject(err);
          });
      });
    },

    getUserTokens(context) {
      APIClient.getUserTokens()
        .then(resp => {
          context.commit("setUserTokens", resp);
        })
        .catch(err => {
          console.log(err);
          context.commit("setUserTokens", []);
        });
    },

    createUserToken(context, name) {
      return new Promise((resolve, reject) => {
        APIClient.createUserToken(name)
          .then(resp => {
            let tokens = context.getters.userTokens;
            tokens.push(resp);
            context.commit("setUserTokens", tokens);
            resolve();
          })
          .catch(err => {
            console.log(err);
            if (err.response) {
              reject(err.response.data.error);
            }

            reject(err);
          });
      });
    },

    revokeUserToken(context, token) {
      return new Promise((resolve, reject) => {
        APIClient.revokeUserToken(token)
          .then(resp => {
            let tokens = context.getters.userTokens.filter(
              token => token.id != resp.id
            );
            context.commit("setUserTokens", tokens);
            resolve();
          })
          .catch(err => {
            console.log(err);
            if (err.response) {
              reject(err.response.data.error);
            }

            reject(err);
          });
      });
    },

    logoutUser(context, router) {
      APIClient.logoutUser().finally(() => {
        context.commit("setUser", {});
        context.commit("setUserAuthenticated", false);
        localStorage.removeItem("isLoggedIn");

        // refresh page/component
        router.go();
      });
    }
  }
});
