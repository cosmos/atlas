<template>
  <header class="header-global">
    <base-nav
      class="navbar-main headroom"
      id="navbar-main"
      :type="navbarType ? navbarType : ''"
      effect="dark"
      expand
    >
      <router-link slot="brand" class="navbar-brand mr-lg-5" to="/">
        <img src="img/brand/white.png" alt="logo" href="" target="_self" />
      </router-link>

      <div class="row" slot="content-header" slot-scope="{ closeMenu }">
        <div class="col-6 collapse-brand">
          <a href="" target="_self">
            <img src="img/brand/white.png" />
          </a>
        </div>
        <div class="col-6 collapse-close">
          <close-button @click="closeMenu"></close-button>
        </div>
      </div>

      <div
        class="navbar-nav navbar-nav-hover align-items-lg-center"
        style="width: -webkit-fill-available;"
      >
        <form v-on:submit="queryModules" style="width: inherit;">
          <base-input
            v-if="showSearch"
            v-model="searchCriteria"
            addonLeftIcon="fa fa-search"
            placeholder="Search"
            style="margin-bottom:0px;"
          ></base-input>
        </form>
      </div>

      <ul class="navbar-nav navbar-nav-hover align-items-lg-center ml-lg-auto">
        <router-link
          class="nav-link"
          :to="{ name: 'browse' }"
          style="color: rgba(255, 255, 255, 0.95);"
        >
          Browse
        </router-link>
        <a
          class="nav-link"
          role="button"
          :href="sessionStartURL"
          v-if="!isAuthenticated"
          >Login</a
        >
        <base-dropdown class="nav-item" v-if="isAuthenticated">
          <a
            role="button"
            slot="title"
            href="#"
            class="nav-link"
            data-toggle="dropdown"
          >
            <i class="ni ni-tablet-button d-lg-none"></i>
            <span class="nav-link-inner--text">Account</span>
          </a>
          <router-link
            :to="{ name: 'profile', params: { name: user.name } }"
            class="dropdown-item"
          >
            <i class="fa fa-user-circle text-muted"></i>
            Profile
          </router-link>
          <router-link to="/account" class="dropdown-item">
            <i class="fa fa-cog text-muted"></i>
            Account
          </router-link>
          <a href="#" class="dropdown-item" v-on:click="logout">
            <i class="fa fa-sign-out text-muted"></i>
            Logout
          </a>
        </base-dropdown>
      </ul>
    </base-nav>
  </header>
</template>

<script>
import BaseNav from "@/components/BaseNav";
import CloseButton from "@/components/CloseButton";
import BaseDropdown from "@/components/BaseDropdown";
import Headroom from "headroom.js";

export default {
  components: {
    BaseNav,
    CloseButton,
    BaseDropdown
  },
  props: {
    showSearch: Boolean,
    navbarType: String
  },
  computed: {
    user() {
      return this.$store.getters.userRecord;
    },

    isAuthenticated() {
      return this.$store.getters.isAuthenticated;
    }
  },
  methods: {},
  created() {
    this.$store.dispatch("getUser");
  },
  data() {
    return {
      searchCriteria: "",
      sessionStartURL: process.env.VUE_APP_ATLAS_API_ADDR + "/session/start"
    };
  },
  mounted: function() {
    let headroom = new Headroom(document.getElementById("navbar-main"), {
      offset: 300,
      tolerance: {
        up: 30,
        down: 30
      }
    });
    headroom.init();
  }
};
</script>

<style>
.navbar-main.headroom {
  z-index: 9999;
}
</style>
