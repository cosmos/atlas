<template>
  <div class="wrapper">
    <header class="header-5">
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
                  <base-button type="primary" class="btn-block"
                    >More</base-button
                  >
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

    <Feature></Feature>
    <Blog></Blog>
    <Pricing></Pricing>
    <Team></Team>
  </div>
</template>
<script>
import Feature from "./sections/Features/Feature1.vue";
import Blog from "./sections/Blogs/Blog6.vue";
import Pricing from "./sections/Pricing/Pricing5.vue";
import Team from "./sections/Teams/Team1.vue";

export default {
  bodyClass: "about-us",
  components: {
    Feature,
    Blog,
    Pricing,
    Team
  }
};
</script>
<style></style>
