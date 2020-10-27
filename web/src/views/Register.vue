<template>
  <div class="wrapper">
    <div class="page-header bg-default">
      <div
        class="page-header-image"
        style="background-image: url('img/ill/register_bg.png');"
      ></div>
      <div
        class="container"
        v-bind:class="{ 'right-panel-active': isActive }"
        id="container"
      >
        <div class="form-container">
          <form action="javascript:;">
            <h2>Create Account</h2>
            <div class="social-container">
              <base-button type="facebook" size="sm"
                ><span class="btn-inner--icon"
                  ><i class="fa fa-facebook"></i></span
              ></base-button>

              <base-button type="instagram" size="sm"
                ><span class="btn-inner--icon"
                  ><i class="fa fa-instagram"></i></span
              ></base-button>

              <base-button type="twitter" size="sm"
                ><span class="btn-inner--icon"
                  ><i class="fa fa-twitter"></i></span
              ></base-button>
            </div>
            <span class="text-default mb-4"
              >or use your email for registration</span
            >
            <div class="form-group mb-0">
              <base-input
                alternative
                placeholder="Name"
                addonLeftIcon="ni ni-circle-08"
              ></base-input>
            </div>
            <div class="form-group mb-0">
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
            <base-button type="primary">Sign Up</base-button>
          </form>
        </div>
        <div class="form-container sign-in-container">
          <form action="#" role="form">
            <h2>Sign in</h2>
            <div class="social-container">
              <base-button type="facebook" size="sm"
                ><span class="btn-inner--icon"
                  ><i class="fa fa-facebook"></i></span
              ></base-button>

              <base-button type="instagram" size="sm"
                ><span class="btn-inner--icon"
                  ><i class="fa fa-instagram"></i></span
              ></base-button>

              <base-button type="twitter" size="sm"
                ><span class="btn-inner--icon"
                  ><i class="fa fa-twitter"></i></span
              ></base-button>
            </div>
            <span class="text-default mb-4">or use your account</span>
            <div class="form-group mb-0">
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
            <a href="javascript:;">Forgot your password?</a>
            <base-button type="primary" class="mt-3">Sign In</base-button>
          </form>
        </div>
        <div class="overlay-container">
          <div class="overlay">
            <div class="overlay-panel overlay-left">
              <h1 class="text-white">Welcome Back!</h1>
              <p>
                To keep connected with us please login with your personal info
              </p>
              <base-button type="neutral" size="sm" @click="isActive = false"
                >Sign In</base-button
              >
            </div>
            <div class="overlay-panel overlay-right">
              <h1 class="text-white">Hello, Friend!</h1>
              <p>Enter your personal details and start journey with us</p>
              <base-button type="neutral" size="sm" @click="isActive = true"
                >Sign Up</base-button
              >
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  bodyClass: "register-page",
  data() {
    return {
      isActive: false
    };
  }
};
</script>
<style></style>
