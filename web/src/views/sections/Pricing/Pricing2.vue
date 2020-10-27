<template>
  <div class="pricing-2" id="pricing-2">
    <div class="container">
      <div class="row">
        <div class="col-md-6 mx-auto text-center mb-3">
          <h3 class="display-3">Pick the best plan for you</h3>
          <p class="lead ">
            You have Free Unlimited Updates and Premium Support on each package.
          </p>
          <div class="section-space"></div>
        </div>
      </div>
      <div class="space-100"></div>
      <div class="row">
        <div class="nav-wrapper mx-auto mr-3">
          <tabs pills fill type="info" class="ml-auto">
            <tab-pane active>
              <template slot="label">
                <i class="ni ni-satisfied mr-2"></i> Mothly
              </template>
            </tab-pane>
            <tab-pane>
              <template slot="label">
                <i class="ni ni-collection mr-2"></i> Yearly
              </template>
            </tab-pane>
          </tabs>
        </div>
      </div>
      <div class="row">
        <div class="col-md-4">
          <card
            class="card-pricing card-background text-center"
            type="info"
            headerClasses="bg-transparent"
            footerClasses="bg-transparent"
            style="background-image: url('img/ill/pricing_bg.svg"
            data-background="image"
          >
            <template slot="body">
              <h6 class="card-category text-white text-uppercase">Starter</h6>
              <h1 class="card-title"><small class="text-white">$</small>199</h1>
              <ul>
                <li><b>10</b> Projects</li>
                <li><b>1</b> Team Members</li>
                <li><b>5</b> Personal Contacts</li>
                <li><b>500</b> Messages</li>
              </ul>
              <a href="javascript:;" class="btn btn-white btn-sm mt-3">
                Get Started
              </a>
            </template>
          </card>
        </div>
        <div class="col-md-4">
          <card
            class="card-pricing card-background text-center"
            type="warning"
            headerClasses="bg-transparent"
            footerClasses="bg-transparent"
            style="background-image: url('img/ill/pricing_bg.svg"
            data-background="image"
          >
            <template slot="body">
              <h6 class="card-category text-white text-uppercase">Pro</h6>
              <h1 class="card-title"><small class="text-white">$</small>399</h1>
              <ul>
                <li><b>100</b> Projects</li>
                <li><b>10</b> Team Members</li>
                <li><b>50</b> Personal Contacts</li>
                <li><b>500</b> Messages</li>
              </ul>
              <a href="javascript:;" class="btn btn-white btn-sm mt-3">
                Get Started
              </a>
            </template>
          </card>
        </div>
        <div class="col-md-4">
          <card
            class="card-pricing card-background text-center"
            type="success"
            headerClasses="bg-transparent"
            footerClasses="bg-transparent"
            style="background-image: url('img/ill/pricing_bg.svg"
            data-background="image"
          >
            <template slot="body">
              <h6 class="card-category text-white text-uppercase">
                Enterprise
              </h6>
              <h1 class="card-title"><small class="text-white">$</small>599</h1>
              <ul>
                <li><b>1k</b> Projects</li>
                <li><b>100</b> Team Members</li>
                <li><b>500</b> Personal Contacts</li>
                <li><b>Unlimited</b> Messages</li>
              </ul>
              <a href="javascript:;" class="btn btn-white btn-sm mt-3">
                Get Started
              </a>
            </template>
          </card>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import Tabs from "@/components/Tabs/Tabs.vue";
import TabPane from "@/components/Tabs/TabPane.vue";

export default {
  components: {
    Tabs,
    TabPane
  }
};
</script>
<style></style>
