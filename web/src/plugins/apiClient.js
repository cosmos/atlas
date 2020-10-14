import axios from 'axios';

const BASE_URI = process.env.VUE_APP_ATLAS_API_ADDR;
const client   = axios.create({baseURL: BASE_URI, json: true});

const APIClient = {
  getUser() {
    return this.perform('get', '/me');
  },

  updateUser(user) {
    return this.perform('put', '/me', user);
  },

  logoutUser() {
    return this.perform('post', '/session/logout');
  },

  // updateKudo(repo) {
  //   return this.perform('put', `/kudos/${repo.id}`, repo);
  // },

  async perform(method, resource, data) {
    return client({method, url: resource, data, headers: {}}).then(req => {
      return req.data;
    });
  }
};

export default APIClient;
