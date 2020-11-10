import axios from "axios";

const BASE_URI = process.env.VUE_APP_ATLAS_API_ADDR;
const client = axios.create({
  baseURL: BASE_URI,
  json: true,
  withCredentials: true
});

const APIClient = {
  getUser() {
    return this.perform("get", "/me");
  },

  getUserTokens() {
    return this.perform("get", "/me/tokens");
  },

  updateUser(user) {
    return this.perform("put", "/me", user);
  },

  logoutUser() {
    return this.perform("post", "/session/logout");
  },

  createUserToken(name) {
    return this.perform("put", "/me/tokens", { name: name });
  },

  revokeUserToken(token) {
    return this.perform("delete", `/me/tokens/${token.id}`);
  },

  getUserByName(name) {
    return this.perform("get", `/users/${name}`);
  },

  getUserModules(name) {
    return this.perform("get", `/users/${name}/modules`);
  },

  getModule(id) {
    return this.perform("get", `/modules/${id}`);
  },

  getModules(pageURI) {
    return this.perform("get", `/modules${pageURI}`);
  },

  starModule(id) {
    return this.perform("put", `/modules/${id}/star`);
  },

  unstarModule(id) {
    return this.perform("put", `/modules/${id}/unstar`);
  },

  searchModules(query, pageURI) {
    return this.perform("get", `/modules/search${pageURI}&q=${query}`);
  },

  confirmEmail(token) {
    return this.perform("put", `/me/confirm/${token}`);
  },

  inviteModuleOwner(user, moduleID) {
    return this.perform("put", "/me/invite", {
      user: user,
      module_id: moduleID
    });
  },

  acceptModuleOwnerInvite(token) {
    return this.perform("put", `/me/invite/accept/${token}`);
  },

  async perform(method, resource, data) {
    return client({ method, url: resource, data, headers: {} }).then(req => {
      return req.data;
    });
  }
};

export default APIClient;
