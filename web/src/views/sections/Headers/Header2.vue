<template>
  <header class="header-2 skew-separator">
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
      <div
        class="page-header-image"
        style="background-image: url('img/ill/p8.svg')"
      ></div>
      <div class="container pt-300">
        <div class="row">
          <div class="col-md-8 mx-auto text-center">
            <h2 class="display-2">Trello lets you work</h2>
          </div>
        </div>
        <div class="row">
          <div class="floating-box bg-default">
            <div class="box text-center">
              <div
                class="icon icon-shape bg-success icon-xl rounded-circle text-white"
              >
                <i class="ni ni-spaceship"></i>
              </div>
            </div>
            <h2 class="lead text-white p-5">
              As a result of growing greenhouse gas emissions, climate models
              predict that our planet will get significantly warmer, that
              ecosystems will be changed or destroyed, and that enormous human
              and economic costs will be incurred. These scenarios aren’t
              guaranteed, but avoiding them will be very hard. We’re trying to
              take small steps to mitigate our impact.
            </h2>
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
