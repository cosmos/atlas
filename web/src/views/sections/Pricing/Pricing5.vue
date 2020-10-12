<template>
  <div class="pricing-5" id="pricing-6">
    <div class="container">
      <div class="row">
        <div class="col-md-4">
          <h3 class="display-3 mt-3">Choose a plan for your next project</h3>
          <tabs
            v-model="pricing5Tab"
            pills
            type="primary"
            class="my-4"
            no-content-space
          >
            <tab-pane label="Cheaper"></tab-pane>
            <tab-pane label="Expensive"></tab-pane>
          </tabs>
          <p class="lead mt-0">
            You have Free Unlimited Updates and Premium Support on each package.
            You also have 20 days to request a refund.
          </p>
        </div>
        <div class="col-md-7 ml-auto mr-auto">
          <div class="tab-content tab-space">
            <div
              v-if="pricing5Tab === 'Expensive'"
              class="tab-pane active"
              id="personal"
            >
              <div class="row">
                <div class="col-md-6">
                  <card
                    class="card-pricing text-center"
                    type=""
                    headerClasses="bg-transparent"
                    footerClasses="bg-transparent"
                  >
                    <template slot="header">
                      <h6 class="text-uppercase ls-1 py-3 mb-0">GOLD</h6>
                    </template>
                    <template slot="body">
                      <div class="display-2">$100</div>
                      <span>per month</span>
                      <ul class="list-unstyled my-4">
                        <li class="align-items-center">
                          <b class="text-primary">200GB</b>
                          <span>File Storage</span>
                        </li>
                        <li class="align-items-center">
                          <b class="text-primary">Unlimited</b>
                          <span>Users</span>
                        </li>
                        <li class="align-items-center">
                          <b class="text-primary">Premium</b>
                          <span>Support</span>
                        </li>
                      </ul>
                    </template>
                    <template slot="footer">
                      <a href="#!" class="text-primary">Request a demo</a>
                    </template>
                  </card>
                </div>
                <div class="col-md-6">
                  <card
                    class="card-pricing text-center"
                    type=""
                    headerClasses="bg-transparent"
                    footerClasses="bg-transparent"
                  >
                    <template slot="header">
                      <h6 class="text-uppercase ls-1 py-3 mb-0">Platinum</h6>
                    </template>
                    <template slot="body">
                      <div class="display-2">$135</div>
                      <span>per month</span>
                      <ul class="list-unstyled my-4">
                        <li class="align-items-center">
                          <b class="text-primary">500GB</b>
                          <span>File Storage</span>
                        </li>
                        <li class="align-items-center">
                          <b class="text-primary">Unlimited</b>
                          <span>Users</span>
                        </li>
                        <li class="align-items-center">
                          <b class="text-primary">No time</b>
                          <span>Tracking</span>
                        </li>
                      </ul>
                    </template>
                    <template slot="footer">
                      <a href="#!" class="text-primary">Request a demo</a>
                    </template>
                  </card>
                </div>
              </div>
            </div>
            <div v-else class="tab-pane active" id="commercial">
              <div class="row">
                <div class="col-md-6">
                  <card
                    class="card-pricing text-center"
                    type=""
                    headerClasses="bg-transparent"
                    footerClasses="bg-transparent"
                  >
                    <template slot="header">
                      <h6 class="text-uppercase ls-1 py-3 mb-0">Standard</h6>
                    </template>
                    <template slot="body">
                      <div class="display-2">$25</div>
                      <span>per month</span>
                      <ul class="list-unstyled my-4">
                        <li class="align-items-center">
                          <b class="text-primary">20GB</b>
                          <span>File Storage</span>
                        </li>
                        <li class="align-items-center">
                          <b class="text-primary">15</b> <span>Users</span>
                        </li>
                        <li class="align-items-center">
                          <b class="text-primary">false</b> <span>Support</span>
                        </li>
                      </ul>
                    </template>
                    <template slot="footer">
                      <a href="#!" class="text-primary">Request a demo</a>
                    </template>
                  </card>
                </div>
                <div class="col-md-6">
                  <card
                    class="card-pricing text-center"
                    type=""
                    headerClasses="bg-transparent"
                    footerClasses="bg-transparent"
                  >
                    <template slot="header">
                      <h6 class="text-uppercase ls-1 py-3 mb-0">Premium</h6>
                    </template>
                    <template slot="body">
                      <div class="display-2">$59</div>
                      <span>per month</span>
                      <ul class="list-unstyled my-4">
                        <li class="align-items-center">
                          <b class="text-primary">50GB</b>
                          <span>File Storage</span>
                        </li>
                        <li class="align-items-center">
                          <b class="text-primary">100</b> <span>Users</span>
                        </li>
                        <li class="align-items-center">
                          <b class="text-primary">Premium</b>
                          <span>Support</span>
                        </li>
                      </ul>
                    </template>
                    <template slot="footer">
                      <a href="#!" class="text-primary">Request a demo</a>
                    </template>
                  </card>
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
import Tabs from "@/components/Tabs/Tabs.vue";
import TabPane from "@/components/Tabs/TabPane.vue";
export default {
  components: {
    Tabs,
    TabPane
  },
  data() {
    return {
      pricing5Tab: "Expensive"
    };
  }
};
</script>
<style></style>
