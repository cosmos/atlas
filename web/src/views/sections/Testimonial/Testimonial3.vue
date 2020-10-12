<template>
  <div class="testimonials-3">
    <div class="container-fluid">
      <div class="row">
        <div class="col-md-6 mx-auto text-center">
          <h3 class="display-3">Join our world</h3>
        </div>
      </div>
      <div class="row">
        <div class="col-lg-3 col-md-8 col-10 positioned">
          <h3 class="display-3">See our customers opinion</h3>
          <p class="lead">
            Meet Wingman, a robust suite of styled pages and components, powered
            by Bootstrap 4. The ideal starting point for product landing pages,
            stylish web-apps and complete company websites.
          </p>
          <base-button tag="a" href="javascript:;" type="primary" class="mt-5"
            >Contact us</base-button
          >
        </div>
        <div class="col-md-12 col-md-6">
          <vue-glide class="testimonial-glide" type="carousel" :perView="4">
            <vue-glide-slide>
              <div class="info text-left">
                <p class="description">
                  "Take up one idea. Make that one idea your life - think of it,
                  dream of it, live on that idea. Let the brain, muscles,
                  nerves, every part of your body, be full of that idea, and
                  just leave every other idea alone. This is the way to success.
                  A single rose can be my garden... a single friend, my world."
                </p>
                <div class="author">
                  <img
                    src="img/faces/team-1.jpg"
                    alt="..."
                    class="avatar img-raised"
                  />
                  <span>Andrew Lino</span>
                </div>
              </div>
            </vue-glide-slide>
            <vue-glide-slide>
              <div class="info text-left">
                <p class="description">
                  Artist is a term applied to a person who engages in an
                  activity deemed to be an art. An artist also may be defined
                  unofficially as "a person who expresses him- or herself
                  through a medium". He is a descriptive term applied to a
                  person who engages in an activity deemed to be an art."
                </p>
                <div class="author">
                  <img
                    src="img/faces/team-2.jpg"
                    alt="..."
                    class="avatar img-raised"
                  />
                  <span>Michael Jenos</span>
                </div>
              </div>
            </vue-glide-slide>
            <vue-glide-slide>
              <div class="info text-left">
                <p class="description">
                  "The simplest visual description uses ordinary words to convey
                  what the writer sees. First he or she must look at the subject
                  – slowly, carefully, and repeatedly, if possible – to identify
                  the parts that make the whole"
                </p>
                <div class="author">
                  <img
                    src="img/faces/team-3.jpg"
                    alt="..."
                    class="avatar img-raised"
                  />
                  <span>Rose Arthur</span>
                </div>
              </div>
            </vue-glide-slide>
            <vue-glide-slide>
              <div class="info text-left">
                <p class="description">
                  "Finding temporary housing for your dog should be as easy as
                  renting an Airbnb. That’s the idea behind Rover, which raised
                  $65 million to expand its pet sitting and dog-walking
                  businesses..Finding temporary housing for your dog should be
                  as easy as renting an Airbnb."
                </p>
                <div class="author">
                  <img
                    src="img/faces/team-4.jpg"
                    alt="..."
                    class="avatar img-raised"
                  />
                  <span>Lora Jimi</span>
                </div>
              </div>
            </vue-glide-slide>
            <vue-glide-slide>
              <div class="info text-left">
                <p class="description">
                  "Venture investment in U.S. startups rose sequentially in the
                  second quarter of 2017, boosted by large, late-stage
                  financings and a few outsized early-stage rounds in tech and
                  healthcare..enture investment in U.S. startups rose
                  sequentially in the second quarter of 2017, boosted by large."
                </p>
                <div class="author">
                  <img
                    src="img/faces/team-5.jpg"
                    alt="..."
                    class="avatar img-raised"
                  />
                  <span>Mike Ranson</span>
                </div>
              </div>
            </vue-glide-slide>
            <template slot="control">
              <button
                class="glide__arrow glide__arrow--left"
                data-glide-dir="<"
              >
                <i class="ni ni-bold-left text-default"></i>
              </button>
              <button
                class="glide__arrow glide__arrow--right"
                data-glide-dir=">"
              >
                <i class="ni ni-bold-right text-default"></i>
              </button>
            </template>
          </vue-glide>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { Glide, GlideSlide } from "vue-glide-js";
import "vue-glide-js/dist/vue-glide.css";

export default {
  components: {
    [Glide.name]: Glide,
    [GlideSlide.name]: GlideSlide
  }
};
</script>
<style></style>
