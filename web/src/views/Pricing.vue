<template>
  <div class="wrapper">
    <div class="page-header header-2">
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
    <div class="main mt-5">
      <!-- Pricing Section -->
      <Pricing></Pricing>
      <br />
      <br />
      <!-- Feature Section -->
      <Feature></Feature>
      <!-- Contact Section -->
      <Contact></Contact>
      <!-- Project Section -->
      <Project></Project>
      <br />
      <br />
      <div class="subscribe-line">
        <div class="container">
          <div class="row">
            <div class="col-lg-6 col-md-10 mx-auto">
              <div class="text-center">
                <h2 class="title">Subscribe to our Newsletter</h2>
                <p class="description">
                  You'll be informed about updates, special offers and planned
                  changes..
                </p>
              </div>
              <div class="card card-raised card-form-horizontal">
                <div class="card-body">
                  <form method="" action="">
                    <div class="row">
                      <div class="col-sm-8">
                        <base-input
                          icon="ni ni-email-83"
                          placeholder="Your Email..."
                        ></base-input>
                      </div>
                      <div class="col-sm-4">
                        <base-button type="primary">Subscribe</base-button>
                      </div>
                    </div>
                  </form>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import Pricing from "./sections/Pricing/Pricing3.vue";
import Feature from "./sections/Features/Feature1.vue";
import Contact from "./sections/Contact/Contact3.vue";
import Project from "./sections/Projects/Project2.vue";

export default {
  bodyClass: "pricing-page",
  components: {
    Pricing,
    Feature,
    Contact,
    Project
  }
};
</script>
<style></style>
