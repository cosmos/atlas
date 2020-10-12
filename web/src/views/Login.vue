<template>
  <div class="wrapper">
    <div class="section-shaped my-0 skew-separator skew-mini">
      <div class="page-header page-header-small header-filter">
        <div
          class="page-header-image"
          style="background-image: url('img/sections/unsplashs.jpg');"
        ></div>
        <div class="container">
          <div class="header-body text-center mb-7">
            <div class="row justify-content-center">
              <div class="col-xl-5 col-lg-6 col-md-8 px-5">
                <h1 class="text-white">Welcome!</h1>
                <p class="text-lead text-white">
                  Use these awesome forms to login or create new account in your
                  project for free.
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <section class="upper">
      <div class="container">
        <div class="col-lg-5 col-md-8 mx-auto">
          <div class="card bg-secondary shadow border-0">
            <div class="card-header bg-white pb-5">
              <div class="text-muted text-center mb-3">
                <small>Login with</small>
              </div>
              <div class="btn-wrapper text-center">
                <base-button
                  tag="a"
                  href="javascript:void(0);"
                  type="neutral"
                  class="btn-icon"
                >
                  <span class="btn-inner--icon"
                    ><img src="img/icons/common/github.svg"
                  /></span>
                  <span class="btn-inner--text">Github</span>
                </base-button>
                <base-button
                  tag="a"
                  href="javascript:void(0);"
                  type="neutral"
                  class="btn-icon"
                >
                  <span class="btn-inner--icon"
                    ><img src="img/icons/common/google.svg"
                  /></span>
                  <span class="btn-inner--text">Google</span>
                </base-button>
              </div>
            </div>
            <div class="card-body px-lg-5 py-lg-5">
              <div class="text-center text-muted mb-4">
                <small>Or login with credentials</small>
              </div>
              <form role="form">
                <div class="form-group mb-3">
                  <base-input
                    alternative
                    placeholder="Email"
                    addonLeftIcon="ni ni-email-83"
                  ></base-input>
                </div>
                <div class="form-group">
                  <base-input
                    alternative
                    placeholder="Password"
                    addonLeftIcon="ni ni-lock-circle-open"
                  ></base-input>
                </div>
                <div class="custom-control-alternative custom-checkbox pl-0">
                  <base-checkbox alternative>Remember me</base-checkbox>
                </div>
                <div class="text-center">
                  <base-button type="primary" class="my-4">Login</base-button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
<script>
export default {
  bodyClass: "login-page"
};
</script>
<style></style>
