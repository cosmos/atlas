<template>
  <header class="header-3 bg-dark">
    <base-nav class="navbar-main" type="" effect="dark" expand>
      <router-link slot="brand" class="navbar-brand mr-lg-5" to="/">
        <a href="javascript:;" class="navbar-brand text-white">Creative Tim</a>
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

      <ul class="navbar-nav ml-auto">
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
            Contact Us
          </a>
        </li>
      </ul>
    </base-nav>
    <div class="page-header header-filter">
      <div class="content-center">
        <div class="row">
          <div class="col-lg-5 col-md-8 col-12 mx-auto positioned">
            <h4 class="title text-white text-uppercase ml-0">
              Build stunning websites and apps
            </h4>
            <div class="info info-horizontal ml-0">
              <div
                class="icon icon-shape bg-white shadow rounded-circle text-default"
              >
                <i class="ni ni-active-40"></i>
              </div>
              <div class="description">
                <h6 class="info-title text-uppercase text-white pl-0">
                  Connect with founders
                </h6>
                <p class="text-white opacity-8">
                  As we live, our hearts turn colder. Cause pain is what we go
                  through as we become older and we get insulted
                </p>
              </div>
            </div>
            <div class="info info-horizontal ml-0">
              <div
                class="icon icon-shape bg-white shadow rounded-circle text-default"
              >
                <i class="ni ni-trophy"></i>
              </div>
              <div class="description">
                <h6 class="info-title text-uppercase text-white pl-0">
                  Learn from experts
                </h6>
                <p class="text-white opacity-8">
                  We get back stabbed by friends. It becomes harder for us to
                  give others a hand. We get our heart broken by people we love.
                </p>
              </div>
            </div>
            <div class="info info-horizontal ml-0">
              <div
                class="icon icon-shape bg-white shadow rounded-circle text-default"
              >
                <i class="ni ni-paper-diploma"></i>
              </div>
              <div class="description">
                <h6 class="info-title text-uppercase text-white pl-0">
                  Get practical advice
                </h6>
                <p class="text-white opacity-8">
                  When we lose family over time. What else could rust the heart
                  more over time? Blackgold. It becomes harder...
                </p>
              </div>
            </div>
          </div>
          <div class="col-md-12 col-md-6">
            <vue-glide
              :startAt="0"
              :perView="4"
              :perTouch="1"
              :focutAt="2"
              type="carousel"
            >
              <vue-glide-slide>
                <img
                  src="img/theme/victor-garcia.jpg"
                  height="500"
                  width="450"
                />
              </vue-glide-slide>
              <vue-glide-slide>
                <img
                  src="img/theme/sacha-styles.jpg"
                  height="500"
                  width="450"
                />
              </vue-glide-slide>
              <vue-glide-slide>
                <img src="img/theme/doyoun-seo.jpg" height="500" width="450" />
              </vue-glide-slide>
              <vue-glide-slide>
                <img
                  src="img/theme/ayo-ogunseinde.jpg"
                  height="500"
                  width="450"
                />
              </vue-glide-slide>
              <vue-glide-slide>
                <img
                  src="img/theme/sofia-kuniakina.jpg"
                  height="500"
                  width="450"
                />
              </vue-glide-slide>
              <template slot="control">
                <button
                  class="glide__arrow glide__arrow--left"
                  data-glide-dir="<"
                >
                  <i class="ni ni-bold-left"></i>
                </button>
                <button
                  class="glide__arrow glide__arrow--right"
                  data-glide-dir=">"
                >
                  <i class="ni ni-bold-right"></i>
                </button>
              </template>
            </vue-glide>
          </div>
        </div>
      </div>
    </div>
  </header>
</template>
<script>
import BaseNav from "@/components/BaseNav";
import CloseButton from "@/components/CloseButton";
import { Glide, GlideSlide } from "vue-glide-js";
import "vue-glide-js/dist/vue-glide.css";

export default {
  components: {
    BaseNav,
    CloseButton,
    [Glide.name]: Glide,
    [GlideSlide.name]: GlideSlide
  }
};
</script>
<style>
.glide__slides {
  margin-left: 50%;
}
</style>
