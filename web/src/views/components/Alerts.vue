<template>
  <div class="container">
    <h3 class="h4 text-success font-weight-bold mb-4 mt-5">Alerts</h3>
    <base-alert icon="ni ni-like-2" dismissible type="success">
      <div slot="text">
        <strong>Success!</strong> This is a success alert—check it out!
      </div>
    </base-alert>
    <base-alert icon="ni ni-bell-55" dismissible type="info">
      <div slot="text">
        <strong>Info!</strong> This is a info alert—check it out!
      </div>
    </base-alert>
    <base-alert icon="ni ni-bell-55" dismissible type="warning">
      <div slot="text">
        <strong>Warning!</strong> This is a warning alert—check it out!
      </div>
    </base-alert>
    <base-alert icon="ni ni-support-16" dismissible type="danger">
      <div slot="text">
        <strong>Danger!</strong> This is an error alert—check it out!
      </div>
    </base-alert>
    <p class="text-uppercase font-weight-bold mt-5 mb-3">Toasts</p>
    <div class="row">
      <div class="col-lg-4 col-md-6">
        <toast>
          <template slot="header-content">
            <img src="img/tim.png" class="rounded mr-2" alt="..." />
            <strong class="mr-auto">Creative Tim</strong>
            <small>11 mins ago</small>
          </template>
          <div slot="text">
            Hello, world! This is a toast message.
          </div>
        </toast>
      </div>
      <div class="col-lg-4 col-md-6">
        <toast class="bg-info">
          <template slot="header-content">
            <img src="img/tim.png" class="rounded mr-2" alt="..." />
            <strong class="mr-auto text-white">Creative Tim</strong>
            <small class="text-white">just now</small>
          </template>
          <div slot="text" class="text-white">
            See? Just like this.
          </div>
        </toast>
      </div>
      <div class="col-lg-4 col-md-6">
        <toast class="bg-warning">
          <template slot="header-content">
            <img src="img/tim.png" class="rounded mr-2" alt="..." />
            <strong class="mr-auto text-white">Creative Tim</strong>
            <small class="text-white">Active</small>
          </template>
          <div slot="text" class="text-white">
            Heads up, toasts will stack automatically
          </div>
        </toast>
      </div>
    </div>
  </div>
</template>
<script>
import Toast from "@/components/Toast";

export default {
  components: {
    Toast
  }
};
</script>
<style></style>
