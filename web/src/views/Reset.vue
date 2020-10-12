<template>
  <div class="wrapper">
    <div class="page-header bg-default">
      <div
        class="page-header-image"
        style="background-image: url('img/ill/reset.svg');"
      ></div>
      <div class="container">
        <div class="col-lg-5 col-md-8 mx-auto">
          <div class="card bg-secondary shadow border-0">
            <div class="card-header">
              <img class="card-img" src="img/ill/bg5-1.svg" alt="Card image" />
              <h4 class="card-title text-center">Reset Password</h4>
            </div>
            <div class="card-body px-lg-5 py-lg-5">
              <div class="text-center text-muted mb-4">
                <small>Enter email address to reset password</small>
              </div>
              <form role="form">
                <div class="form-group mb-3">
                  <base-input
                    alternative
                    placeholder="Email"
                    addonLeftIcon="ni ni-email-83"
                  ></base-input>
                </div>
                <div class="text-center">
                  <base-button type="primary" class="my-4">Send</base-button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  bodyClass: "reset-page"
};
</script>
<style></style>
