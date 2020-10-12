<template>
  <div class="wrapper">
    <header class="header-4 skew-separator">
      <div class="header-wrapper">
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
                  icon="ni ni-button-play"
                  type="warning"
                  class="mt-3 mb-sm-0"
                  >Play more</base-button
                >
              </div>
            </div>
          </div>
        </div>
      </div>
    </header>
    <!-- First Feature Section -->
    <Feature></Feature>
    <!-- Second Feature Section -->
    <Feature2></Feature2>
    <!-- Blog Section -->
    <Blog></Blog>
    <br />
    <br />
    <br />
    <!-- Pricing Section -->
    <Pricing></Pricing>
  </div>
</template>
<script>
import Feature from "./sections/Features/Feature6.vue";
import Feature2 from "./sections/Features/Feature1.vue";
import Blog from "./sections/Blogs/Blog2.vue";
import Pricing from "./sections/Pricing/Pricing5.vue";

export default {
  components: {
    Feature,
    Feature2,
    Blog,
    Pricing
  }
};
</script>
<style></style>
