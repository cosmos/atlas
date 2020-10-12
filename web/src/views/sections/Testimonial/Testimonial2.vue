<template>
  <div class="testimonials-2">
    <div class="container">
      <div class="row">
        <div class="col-md-6 ml-auto mr-auto text-center">
          <h2 class="title">What Clients Say</h2>
        </div>
      </div>
      <div class="row">
        <b-carousel
          id="carousel-testimonials2"
          height="100vh"
          class="slide carousel-fade carousel-testimonials pointer-event"
          ref="carouselTestimonials2"
        >
          <b-carousel-slide class="carousel-inner">
            <div class="info text-left bg-danger shadow-lg">
              <div
                class="icon icon-shape bg-white text-danger shadow rounded-circle mb-4"
              >
                <i class="fa fa-quote-right"></i>
              </div>
              <p class="description text-white">
                It really matters and then like it really doesn’t matter. What
                matters is the people who are sparked by it. And the people who
                are like offended by it, it doesn’t matter. Because it's about
                motivating the doers. Because I’m here to follow my dreams and
                inspire other people to follow their dreams.
              </p>
              <div class="author">
                <img
                  src="img/faces/team-5.jpg"
                  alt="..."
                  class="avatar img-raised"
                />
                <span>Mike John</span>
              </div>
            </div>
            <card class="card-blog card-background" data-animation="true">
              <div
                class="full-background"
                style="background-image: url('img/sections/austin.jpg')"
              ></div>
              <template slot="body">
                <div class="content-bottom ml-auto">
                  <a
                    href="javascript:void(0)"
                    class="btn btn-round btn-simple btn-neutral"
                    ><i class="ni ni-book-bookmark"></i> Read more</a
                  >
                </div>
              </template>
            </card>
          </b-carousel-slide>
          <b-carousel-slide class="carousel-inner">
            <div class="info text-left bg-info shadow-lg">
              <div
                class="icon icon-shape bg-white text-info shadow rounded-circle mb-4"
              >
                <i class="fa fa-quote-right"></i>
              </div>
              <p class="description text-white">
                Take up one idea. Make that one idea your life - think of it,
                dream of it, live on that idea. Let the brain, muscles, nerves,
                every part of your body, be full of that idea, and just leave
                every other idea alone. This is the way to success. A single
                rose can be my garden... a single friend, my world.
              </p>
              <div class="author">
                <img
                  src="img/faces/team-3.jpg"
                  alt="..."
                  class="avatar img-raised"
                />
                <span>Andrew Lino</span>
              </div>
            </div>
            <card class="card-blog card-background" data-animation="true">
              <div
                class="full-background"
                style="background-image: url('img/sections/unsplashs.jpg')"
              ></div>
              <template slot="body">
                <div class="content-bottom ml-auto">
                  <a
                    href="javascript:void(0)"
                    class="btn btn-round btn-simple btn-neutral"
                    ><i class="ni ni-book-bookmark"></i> Read more</a
                  >
                </div>
              </template>
            </card>
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
      this.$refs.carouselTestimonials2.prev();
    },
    next() {
      this.$refs.carouselTestimonials2.next();
    }
  }
};
</script>
<style>
.carousel-testimonials .carousel-caption {
  width: 100%;
  left: 0;
}
</style>
