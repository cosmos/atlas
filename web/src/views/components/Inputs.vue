<template>
  <div class="section pb-0 section-components">
    <div class="container mb-5">
      <!-- Inputs -->
      <h3 class="h4 text-success font-weight-bold mb-4">Inputs</h3>
      <div class="mb-3">
        <small class="text-uppercase font-weight-bold">Form controls</small>
      </div>
      <div class="row">
        <div class="col-lg-4 col-sm-6">
          <base-input placeholder="Regular"></base-input>
        </div>
        <div class="col-lg-4 col-sm-6">
          <base-input placeholder="Regular" disabled></base-input>
        </div>
        <div class="col-lg-4 col-sm-6">
          <base-input placeholder="Success" :valid="true"></base-input>
        </div>
      </div>

      <div class="row">
        <div class="col-lg-4 col-sm-6">
          <base-input
            placeholder="Regular"
            addonLeftIcon="ni ni-zoom-split-in"
          ></base-input>
        </div>
        <div class="col-lg-4 col-sm-6">
          <base-input
            placeholder="Birthday"
            addonRightIcon="ni ni-zoom-split-in"
          ></base-input>
        </div>
        <div class="col-lg-4 col-sm-6">
          <base-input placeholder="Error input" :valid="false"></base-input>
        </div>
      </div>
    </div>
    <div class="py-5 bg-secondary">
      <div class="container">
        <div class="mb-3">
          <small class="text-uppercase font-weight-bold"
            >Form controls (alternative)</small
          >
        </div>
        <div class="row">
          <div class="col-lg-4 col-sm-6">
            <base-input placeholder="Regular" alternative></base-input>
          </div>
          <div class="col-lg-4 col-sm-6">
            <base-input placeholder="Regular" disabled alternative></base-input>
          </div>
          <div class="col-lg-4 col-sm-6">
            <base-input
              placeholder="Success"
              :valid="true"
              alternative
            ></base-input>
          </div>
        </div>

        <div class="row">
          <div class="col-lg-4 col-sm-6">
            <base-input
              placeholder="Regular"
              addonLeftIcon="ni ni-zoom-split-in"
              alternative
            ></base-input>
          </div>
          <div class="col-lg-4 col-sm-6">
            <base-input
              placeholder="Birthday"
              addonRightIcon="ni ni-zoom-split-in"
              alternative
            ></base-input>
          </div>
          <div class="col-lg-4 col-sm-6">
            <base-input
              placeholder="Error input"
              :valid="false"
              alternative
            ></base-input>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script></script>
<style></style>
