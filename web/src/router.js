import Vue from "vue";
import Router from "vue-router";
import AppFooter from "./layout/AppFooter";
import AppHeader from "./layout/AppHeader";
import store from "./plugins/store";
import Account from "./views/Account.vue";
import Components from "./views/Components.vue";
import Error from "./views/Error.vue";
import ConfirmEmailPage from "./views/ConfirmEmailPage.vue";
import AcceptOwnerInvitePage from "./views/AcceptOwnerInvitePage.vue";
import ProfilePage from "./views/ProfilePage.vue";
import ModulePage from "./views/Module.vue";
import ModulesPage from "./views/Modules.vue";
import NodeExplorer from "./views/NodeExplorer.vue";
import SearchResults from "./views/SearchResults.vue";

Vue.use(Router);

const ifAuthenticated = (to, from, next) => {
  if (store.getters.isAuthenticated) {
    next();
    return;
  }

  next("/");
};

export default new Router({
  linkExactActiveClass: "active",
  mode: "history",
  routes: [
    {
      path: "/",
      name: "home",
      components: { header: AppHeader, default: Components, footer: AppFooter },
      props: { header: { showSearch: false } }
    },
    {
      path: "/search",
      name: "search",
      components: {
        header: AppHeader,
        default: SearchResults,
        footer: AppFooter
      },
      props: { header: { showSearch: true } }
    },
    {
      path: "/account",
      name: "account",
      components: { header: AppHeader, default: Account, footer: AppFooter },
      props: { header: { showSearch: true } },
      beforeEnter: ifAuthenticated
    },
    {
      path: "/profile/:name",
      name: "profile",
      components: {
        header: AppHeader,
        default: ProfilePage,
        footer: AppFooter
      },
      props: { header: { showSearch: true } }
    },
    {
      path: "/modules/:id",
      name: "modules",
      components: { header: AppHeader, default: ModulePage, footer: AppFooter },
      props: { header: { showSearch: true } }
    },
    {
      path: "/modules/:id/:version",
      name: "modulesVersioned",
      components: { header: AppHeader, default: ModulePage, footer: AppFooter },
      props: { header: { showSearch: true } }
    },
    {
      path: "/modules",
      name: "browse",
      components: {
        header: AppHeader,
        default: ModulesPage,
        footer: AppFooter
      },
      props: { header: { showSearch: true } }
    },
    {
      path: "/nodes",
      name: "nodeExplorer",
      components: {
        header: AppHeader,
        default: NodeExplorer,
        footer: AppFooter
      },
      props: { header: { showSearch: true } }
    },
    {
      path: "/confirm/:token",
      name: "confirmEmail",
      components: {
        header: AppHeader,
        default: ConfirmEmailPage,
        footer: AppFooter
      },
      props: { header: { showSearch: false } }
    },
    {
      path: "/accept/:token",
      name: "acceptOwnerInvite",
      components: {
        header: AppHeader,
        default: AcceptOwnerInvitePage,
        footer: AppFooter
      },
      props: { header: { showSearch: false } }
    },
    {
      path: "*",
      name: "error",
      components: { header: AppHeader, default: Error },
      props: { header: { showSearch: true } }
    }
  ],
  scrollBehavior: to => {
    if (to.hash) {
      return { selector: to.hash };
    } else {
      return { x: 0, y: 0 };
    }
  }
});
