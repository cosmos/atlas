import Vue from "vue";
import Router from "vue-router";
import Header from "./layout/AppHeader";
import Footer from "./layout/AppFooter";
import Starter from "./views/Starter.vue";

Vue.use(Router);

export default new Router({
  routes: [
    {
      path: "/",
      name: "starter",
      components: {
        header: Header,
        default: Starter,
        footer: Footer
      }
    }
  ]
});
