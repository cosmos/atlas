<template>
  <div class="pricing-4" id="pricing-5">
    <div class="container">
      <div class="row">
        <div class="col-md-6 ml-auto mr-auto text-center">
          <h2 class="title">Pick the best plan for you</h2>
          <h4 class="description">
            You have Free Unlimited Updates and Premium Support on each package.
          </h4>
        </div>
      </div>
      <div class="row mt-5">
        <div class="col-md-4">
          <card
            class="card-pricing card-coin"
            footerClasses="text-center bg-transparent"
          >
            <template slot="header">
              <img src="img/theme/vue.jpg" class="img-center shadow" />
            </template>
            <template slot="body">
              <div class="row">
                <div class="col-md-12 text-center">
                  <h5 class="text-uppercase">Standard Pack</h5>
                  <span>£ 39</span>
                  <hr class="bg-info" />
                </div>
              </div>
              <div class="row">
                <ul class="list-group">
                  <li class="list-group-item">50 messages</li>
                  <li class="list-group-item">100 emails</li>
                  <li class="list-group-item">No Support</li>
                </ul>
              </div>
            </template>
            <template slot="footer">
              <base-button type="primary" class="mb-3">Get plan</base-button>
            </template>
          </card>
        </div>
        <div class="col-md-4">
          <card
            class="card-pricing card-coin"
            footerClasses="text-center bg-transparent"
          >
            <template slot="header">
              <img src="img/theme/react.jpg" class="img-center shadow" />
            </template>
            <template slot="body">
              <div class="row">
                <div class="col-md-12 text-center">
                  <h5 class="text-uppercase">Premium Pack</h5>
                  <span>£ 49</span>
                  <hr class="bg-info" />
                </div>
              </div>
              <div class="row">
                <ul class="list-group">
                  <li class="list-group-item">100 messages</li>
                  <li class="list-group-item">1K emails</li>
                  <li class="list-group-item">Premium Support</li>
                </ul>
              </div>
            </template>
            <template slot="footer">
              <base-button type="primary" class="mb-3">Get plan</base-button>
            </template>
          </card>
        </div>
        <div class="col-md-4">
          <card
            class="card-pricing card-coin"
            footerClasses="text-center bg-transparent"
          >
            <template slot="header">
              <img src="img/theme/sketch.jpg" class="img-center shadow" />
            </template>
            <template slot="body">
              <div class="row">
                <div class="col-md-12 text-center">
                  <h5 class="text-uppercase">Platinum Pack</h5>
                  <span>£ 69</span>
                  <hr class="bg-info" />
                </div>
              </div>
              <div class="row">
                <ul class="list-group">
                  <li class="list-group-item">350 messages</li>
                  <li class="list-group-item">10K emails</li>
                  <li class="list-group-item">24/7 Support</li>
                </ul>
              </div>
            </template>
            <template slot="footer">
              <base-button type="primary" class="mb-3">Get plan</base-button>
            </template>
          </card>
        </div>
      </div>
    </div>
  </div>
</template>
<script></script>
<style></style>
