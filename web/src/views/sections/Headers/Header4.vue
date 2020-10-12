<template>
  <header class="header-4 skew-separator">
    <div class="header-wrapper">
      <base-nav class="navbar-main" type="" effect="dark" expand>
        <router-link slot="brand" class="navbar-brand mr-lg-5" to="/">
          <a href="javascript:;" class="navbar-brand">Creative Tim</a>
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
      <div class="page-header header-video">
        <div class="overlay"></div>
        <video
          playsinline="playsinline"
          autoplay="autoplay"
          muted="muted"
          loop="loop"
        >
          <source
            src="https://www.oberlo.com/wp-content/uploads/2018/10/1.-Pixabay.mp4"
            type="video/mp4"
          />
        </video>
        <div class="container text-center">
          <div class="row">
            <div class="col-lg-7 mx-auto">
              <h1 class="video-text">People stories</h1>
              <h1 class="display-3 text-white">The time is right now!</h1>
              <base-button
                tag="a"
                type="warning"
                href="javascript:void(0)"
                class="mt-3"
                icon="ni ni-button-play"
              >
                Play more
              </base-button>
            </div>
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
