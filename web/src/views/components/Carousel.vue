<template>
  <div class="section" style="background-image: url('img/ill/1.svg');">
    <div class="container py-md">
      <div class="row justify-content-between align-items-center">
        <div class="col-lg-6 mb-lg-auto">
          <div
            class="rounded shadow-lg overflow-hidden transform-perspective-left"
          >
            <b-carousel id="carousel1" controls indicators>
              <!-- Text slides with image -->
              <b-carousel-slide
                img-src="img/theme/img-1-1200x1000.jpg"
              ></b-carousel-slide>
              <b-carousel-slide
                img-src="img/theme/img-2-1200x1000.jpg"
              ></b-carousel-slide>
              <b-carousel-slide
                img-src="img/theme/img-1-1200x1000.jpg"
              ></b-carousel-slide>
            </b-carousel>
          </div>
        </div>
        <div class="col-lg-5 mb-5 mb-lg-0">
          <h1 class="font-weight-light">Bootstrap carousel</h1>
          <p class="lead mt-4">
            Argon Design System comes with four pre-built pages to help you get
            started faster. You can change the text and images and you're good
            to go.
          </p>
          <a href="javascript:void(0)" class="btn btn-white mt-4"
            >See all components</a
          >
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { BCarousel } from "bootstrap-vue/esm/components/carousel/carousel";
import { BCarouselSlide } from "bootstrap-vue/esm/components/carousel/carousel-slide";

export default {
  components: {
    BCarousel,
    BCarouselSlide
  }
};
</script>
<style></style>
