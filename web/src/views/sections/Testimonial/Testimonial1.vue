<template>
  <div
    class="testimonials-1"
    style="background-image: url('img/ill/testimonial_bg.svg');"
  >
    <div class="container">
      <div class="row">
        <div class="col-md-6 ml-auto mr-auto text-center">
          <h2 class="title">What Clients Say</h2>
        </div>
      </div>
      <b-carousel
        id="carousel-testimonials"
        height="150vh"
        class="carousel-team"
        ref="carouselTestimonials"
      >
        <b-carousel-slide class="carousel-inner">
          <div class="row text-left">
            <div class="col-md-5 ml-auto mt-5">
              <h3 class="card-title">Sarah Smith</h3>
              <h3 class="text-danger">• • •</h3>
              <h4 class="lead">
                Take up one idea. Make that one idea your life - think of it,
                dream of it, live on that idea. Let the brain, muscles, nerves,
                every part of your body, be full of that idea, and just leave
                every other idea alone. This is the way to success. A single
                rose can be my garden... a single friend, my world.
              </h4>
              <base-button
                tag="a"
                href="javascript:void(0)"
                class="mt-3"
                type="danger"
                >Read more</base-button
              >
            </div>
            <div class="col-md-5 p-5 ml-auto">
              <div class="p-3">
                <img
                  class="img-fluid rounded shadow transform-perspective-right"
                  src="img/faces/fezbot.jpg"
                  alt="First slide"
                />
              </div>
            </div>
          </div>
        </b-carousel-slide>
        <b-carousel-slide class="carousel-inner">
          <div class="row text-left">
            <div class="col-md-5 mt-5 ml-auto">
              <h3 class="card-title">Isaac Hunter</h3>
              <h3 class="text-info">• • •</h3>
              <h4 class="lead">
                Take up one idea. Make that one idea your life - think of it,
                dream of it, live on that idea. Let the brain, muscles, nerves,
                every part of your body, be full of that idea, and just leave
                every other idea alone. This is the way to success. A single
                rose can be my garden... a single friend, my world.
              </h4>
              <base-button
                tag="a"
                href="javascript:void(0)"
                class="mt-3"
                type="info"
                >Read more</base-button
              >
            </div>
            <div class="col-md-5 p-5 ml-auto">
              <div class="p-3">
                <img
                  class="img-fluid rounded shadow transform-perspective-right"
                  src="img/faces/brooke-cagle.jpg"
                  alt="First slide"
                />
              </div>
            </div>
          </div>
        </b-carousel-slide>
        <a
          class="carousel-control-prev text-primary"
          v-on:click="prev"
          role="button"
          data-slide="prev"
        >
          <i class="ni ni-bold-left"></i>
          <span class="sr-only">Previous</span>
        </a>
        <a
          class="carousel-control-next text-primary"
          v-on:click="next"
          role="button"
          data-slide="next"
        >
          <i class="ni ni-bold-right"></i>
          <span class="sr-only">Next</span>
        </a>
      </b-carousel>
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
  },
  methods: {
    prev() {
      this.$refs.carouselTestimonials.prev();
    },
    next() {
      this.$refs.carouselTestimonials.next();
    }
  }
};
</script>
<style>
.testimonials-1 .carousel-caption {
  width: 100%;
  left: 0 !important;
}

.testimonials-1 .carousel-control-prev,
.testimonials-1 .carousel-control-next {
  z-index: 9999;
}
</style>
