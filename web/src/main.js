import Vue from "vue";

import App from "./App.vue";
import Argon from "./plugins/argon-kit";
import router from "./router";

Vue.config.productionTip = false;
Vue.use(Argon);

new Vue({
  router,
  render: h => h(App)
}).$mount("#app");
