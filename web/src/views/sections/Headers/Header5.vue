<template>
  <header class="header-5">
    <base-nav class="navbar-main bg-white" type="" effect="white" expand>
      <router-link slot="brand" class="navbar-brand mr-lg-5" to="/">
        <a href="javascript:;" class="navbar-brand text-primary"
          >Creative Tim</a
        >
      </router-link>

      <div class="row" slot="content-header" slot-scope="{ closeMenu }">
        <div class="col-6 collapse-brand">
          <a href="javascript:void(0)">
            <img src="img/brand/blue.png" />
          </a>
        </div>
        <div class="col-6 collapse-close">
          <close-button @click="closeMenu"></close-button>
        </div>
      </div>

      <ul class="navbar-nav mx-auto">
        <li class="nav-item">
          <a class="nav-link" href="javascript:;">
            Home
          </a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="javascript:;">
            About Us
          </a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="javascript:;">
            Products
          </a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="javascript:;">
            Contact Us
          </a>
        </li>
      </ul>
      <ul class="nav navbar-nav navbar-right">
        <li class="nav-item">
          <a class="nav-link" href="https://twitter.com/CreativeTim">
            <i class="fa fa-twitter"></i>
          </a>
        </li>
        <li class="nav-item">
          <a class="nav-link" href="https://www.facebook.com/CreativeTim">
            <i class="fa fa-facebook-square"></i>
          </a>
        </li>
        <li class="nav-item">
          <a
            class="nav-link"
            href="https://www.instagram.com/CreativeTimOfficial"
          >
            <i class="fa fa-instagram"></i>
          </a>
        </li>
      </ul>
    </base-nav>
    <div class="page-header">
      <div class="container">
        <div class="row align-items-center text-left">
          <div class="col-lg-5 col-12">
            <h1 class="display-3">
              Design with us,<span class="text-primary">
                Develop Anything.</span
              >
            </h1>
            <p class="lead pb-4">
              We aim high at being focused on building relationships with our
              clients and community. Using our creative gifts drives this
              foundation.
            </p>
            <div class="row row-input">
              <div class="col-sm-8 col-12">
                <base-input placeholder="Your email"></base-input>
              </div>
              <div class="col-sm-4 col-12 pl-lg-0">
                <base-button type="primary" class="btn-block">More</base-button>
              </div>
            </div>
          </div>
          <div class="col-lg-7 col-12 pl-0">
            <img class="ml-lg-5" src="img/ill/bg6-2.svg" width="100%" />
          </div>
        </div>
      </div>
    </div>
  </header>
</template>
<script>
import BaseNav from "@/components/BaseNav";
import CloseButton from "@/components/CloseButton";

export default {
  components: {
    BaseNav,
    CloseButton
  }
};
</script>
<style></style>
