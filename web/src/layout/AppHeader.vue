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
        <base-dropdown class="nav-item">
          <a
            role="button"
            slot="title"
            href="#"
            class="nav-link"
            data-toggle="dropdown"
          >
            <i class="ni ni-single-copy-04 d-lg-none"></i>
            <span class="nav-link-inner--text">Browse</span>
          </a>
          <router-link to="/about" class="dropdown-item">
            <i class="ni ni-tie-bow text-warning"></i>
            About-us
          </router-link>
          <router-link to="/blog-post" class="dropdown-item">
            <i class="ni ni-align-center text-info"></i>
            Blog Post
          </router-link>
          <router-link to="/blog-posts" class="dropdown-item">
            <i class="ni ni-chart-bar-32 text-yellow"></i>
            Blog Posts
          </router-link>
          <router-link to="/contact-us" class="dropdown-item">
            <i class="ni ni-square-pin text-danger"></i>
            Contact Us
          </router-link>
          <router-link to="/landing-page" class="dropdown-item">
            <i class="ni ni-planet text-purple"></i>
            Landing Page
          </router-link>
          <router-link to="/pricing" class="dropdown-item">
            <i class="ni ni-money-coins text-success"></i>
            Pricing
          </router-link>
          <router-link to="/ecommerce" class="dropdown-item">
            <i class="ni ni-box-2 text-pink"></i>
            Ecommerce Page
          </router-link>
          <router-link to="/product-page" class="dropdown-item">
            <i class="ni ni-bag-17 text-primary"></i>
            Product Page
          </router-link>
          <router-link to="/profile-page" class="dropdown-item">
            <i class="ni ni-circle-08 text-info"></i>
            Profile Page
          </router-link>
          <router-link to="/error" class="dropdown-item">
            <i class="ni ni-button-power text-warning"></i>
            404 Error Page
          </router-link>
          <router-link to="/500-error" class="dropdown-item">
            <i class="ni ni-ungroup text-yellow"></i>
            500 Error Page
          </router-link>
        </base-dropdown>
        <base-dropdown class="nav-item">
          <a
            role="button"
            slot="title"
            href="#"
            class="nav-link"
            data-toggle="dropdown"
          >
            <i class="ni ni-tablet-button d-lg-none"></i>
            <span class="nav-link-inner--text">Stuff</span>
          </a>
          <router-link to="/account" class="dropdown-item">
            <i class="ni ni-lock-circle-open text-muted"></i>
            Account Settings
          </router-link>
          <router-link to="/login" class="dropdown-item">
            <i class="ni ni-tv-2 text-danger"></i>
            Login Page
          </router-link>
          <router-link to="/register" class="dropdown-item">
            <i class="ni ni-air-baloon text-pink"></i>
            Register Page
          </router-link>
          <router-link to="/reset" class="dropdown-item">
            <i class="ni ni-atom text-info"></i>
            Reset Page
          </router-link>
          <router-link to="/invoice" class="dropdown-item">
            <i class="ni ni-bullet-list-67 text-success"></i>
            Invoice Page
          </router-link>
          <router-link to="/checkout" class="dropdown-item">
            <i class="ni ni-basket text-orange"></i>
            Checkout Page
          </router-link>
          <router-link to="/chat-page" class="dropdown-item">
            <i class="ni ni-chat-round text-primary"></i>
            Chat Page
          </router-link>
        </base-dropdown>
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
