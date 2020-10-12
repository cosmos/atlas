import axios from "axios";

const BASE_URI = process.env.VUE_APP_ATLAS_API_ADDR;
const client = axios.create({ baseURL: BASE_URI, json: true });

const APIClient = {
  startSession() {
    return this.perform("get", "/session/start");
  },

  // createKudo(repo) {
  //   return this.perform('post', '/kudos', repo);
  // },

  // deleteKudo(repo) {
  //   return this.perform('delete', `/kudos/${repo.id}`);
  // },

  // updateKudo(repo) {
  //   return this.perform('put', `/kudos/${repo.id}`, repo);
  // },

  // getKudos() {
  //   return this.perform('get', '/kudos');
  // },

  // getKudo(repo) {
  //   return this.perform('get', `/kudo/${repo.id}`);
  // },

  async perform(method, resource, data) {
    // let accessToken = await Vue.prototype.$auth.getAccessToken()
    return client({
      method,
      url: resource,
      data,
      headers: {
        //  Authorization: `Bearer ${accessToken}`
      }
    }).then(req => {
      return req.data;
    });
  }
};

export default APIClient;
