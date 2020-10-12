import axios from 'axios';

const BASE_URI = process.env.VUE_APP_ATLAS_API_ADDR;
const client   = axios.create({baseURL: BASE_URI, json: true});

const APIClient = {
  getUser() {
    return this.perform('get', '/user');
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
    return client({method, url: resource, data, headers: {}}).then(req => {
      return req.data;
    });
  }
};

export default APIClient;
