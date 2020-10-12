<template>
  <div
    class="section section-hero section-shaped"
    style="padding-bottom: 0; padding-top: 0;"
  >
    <div class="page-header">
      <div class="container shape-container d-flex align-items-center py-lg">
        <div class="col px-0">
          <div class="row align-items-center justify-content-center">
            <div class="col-lg-8 text-center">
              <img src="img/brand/logo-hero.png" class="img-fluid" />
              <p class="lead" style="margin-bottom: 24px;">
                <b class="display-3">Cosmos SDK Module Registry</b>
              </p>
              <div class="row">
                <div class="col-md-8 center">
                  <form v-on:submit="queryModules">
                    <base-input
                      v-model="searchCriteria"
                      addonLeftIcon="fa fa-search"
                      placeholder="Search"
                    ></base-input>
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
export default {
  data() {
    return {
      searchCriteria: ""
    };
  }
};
</script>

<style></style>
