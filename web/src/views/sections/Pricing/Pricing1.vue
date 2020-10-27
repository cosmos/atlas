<template>
  <div class="pricing-1" id="pricing-1">
    <div class="container">
      <div class="row">
        <div class="col-md-6 mx-auto text-center mb-5">
          <h2 class="title">Pick the best plan for you</h2>
          <h4 class="description ">
            You have Free Unlimited Updates and Premium Support on each package.
          </h4>
          <div class="section-space"></div>
        </div>
      </div>
      <div class="row">
        <tabs pills fill type="success" class="ml-auto">
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
      <div class="row">
        <div class="col-md-4">
          <card
            class="card-pricing text-center"
            type="gradient-info"
            headerClasses="bg-transparent"
            footerClasses="bg-transparent"
          >
            <template slot="header">
              <h6 class="text-uppercase ls-1 text-white py-3 mb-0">Starter</h6>
            </template>
            <template slot="body">
              <div class="display-2 text-white">$22</div>
              <span class=" text-white">per month</span>
              <ul class="list-unstyled my-4">
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape bg-white shadow rounded-circle text-info"
                      >
                        <i class="ni ni-book-bookmark"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm text-white"
                        >Complete documentation</span
                      >
                    </div>
                  </div>
                </li>
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape bg-white shadow rounded-circle text-info"
                      >
                        <i class="ni ni-diamond"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm text-white"
                        >Working materials in Sketch</span
                      >
                    </div>
                  </div>
                </li>
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape bg-white shadow rounded-circle text-info"
                      >
                        <i class="ni ni-chart-pie-35"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm text-white"
                        >2GB cloud storage</span
                      >
                    </div>
                  </div>
                </li>
              </ul>
              <base-button link class="text-white" type=""
                >Get started</base-button
              >
            </template>
            <template slot="footer">
              <a href="#!" class=" text-white">Request a demo</a>
            </template>
          </card>
        </div>
        <div class="col-md-4">
          <card
            class="card-pricing text-center"
            type="gradient-warning"
            headerClasses="bg-transparent"
            footerClasses="bg-transparent"
          >
            <template slot="header">
              <h6 class="text-uppercase ls-1 text-white py-3 mb-0">Pro</h6>
            </template>
            <template slot="body">
              <div class="display-2 text-white">$45</div>
              <span class=" text-white">per month</span>
              <ul class="list-unstyled my-4">
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape bg-white shadow rounded-circle text-warning"
                      >
                        <i class="ni ni-book-bookmark"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm text-white"
                        >Complete documentation</span
                      >
                    </div>
                  </div>
                </li>
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape bg-white shadow rounded-circle text-warning"
                      >
                        <i class="ni ni-diamond"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm text-white"
                        >Working materials in Sketch</span
                      >
                    </div>
                  </div>
                </li>
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape bg-white shadow rounded-circle text-warning"
                      >
                        <i class="ni ni-chart-pie-35"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm text-white"
                        >2GB cloud storage</span
                      >
                    </div>
                  </div>
                </li>
              </ul>
              <base-button link class="text-white" type=""
                >Get started</base-button
              >
            </template>
            <template slot="footer">
              <a href="#!" class=" text-white">Request a demo</a>
            </template>
          </card>
        </div>
        <div class="col-md-4">
          <card
            class="card-pricing text-center"
            type="gradient-success"
            headerClasses="bg-transparent"
            footerClasses="bg-transparent"
          >
            <template slot="header">
              <h6 class="text-uppercase ls-1 text-white py-3 mb-0">
                Enterprise
              </h6>
            </template>
            <template slot="body">
              <div class="display-2 text-white">$68</div>
              <span class=" text-white">per month</span>
              <ul class="list-unstyled my-4">
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape bg-white shadow rounded-circle text-success"
                      >
                        <i class="ni ni-book-bookmark"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm text-white"
                        >Complete documentation</span
                      >
                    </div>
                  </div>
                </li>
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape bg-white shadow rounded-circle text-success"
                      >
                        <i class="ni ni-diamond"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm text-white"
                        >Working materials in Sketch</span
                      >
                    </div>
                  </div>
                </li>
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape bg-white shadow rounded-circle text-success"
                      >
                        <i class="ni ni-chart-pie-35"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm text-white"
                        >Unlimited cloud storage</span
                      >
                    </div>
                  </div>
                </li>
              </ul>
              <base-button link class="text-white" type=""
                >Get started</base-button
              >
            </template>
            <template slot="footer">
              <a href="#!" class=" text-white">Request a demo</a>
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
