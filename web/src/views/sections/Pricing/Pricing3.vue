<template>
  <section
    class="pricing-3"
    style="background-image: url('img/ill/1.svg')"
    id="pricing-4"
  >
    <div class="container">
      <div class="row">
        <div class="col-lg-8 mx-auto text-center my-5">
          <h3 class="display-3">Some of Our Pricing Plans</h3>
        </div>
      </div>
      <div class="row">
        <div class="col-md-5 pr-md-0">
          <card
            class="card-pricing text-center"
            type="default"
            headerClasses="bg-transparent"
            bodyClasses="px-lg-6"
            footerClasses="bg-transparent"
          >
            <template slot="header">
              <h4 class="text-uppercase ls-1 text-primary py-3 mb-0">
                Bravo pack
              </h4>
            </template>
            <template slot="body">
              <div class="display-2 text-white">$49</div>
              <span class=" text-white">per application</span>
              <ul class="list-unstyled my-4">
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape icon-shape-primary shadow rounded-circle"
                      >
                        <i class="ni ni-folder-17 text-white"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm">Complete documentation</span>
                    </div>
                  </div>
                </li>
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape icon-shape-primary shadow rounded-circle"
                      >
                        <i class="ni ni-camera-compact text-white"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm"
                        >Working materials in Sketch</span
                      >
                    </div>
                  </div>
                </li>
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape icon-shape-primary shadow rounded-circle"
                      >
                        <i class="ni ni-chart-pie-35 text-white"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm">2GB cloud storage</span>
                    </div>
                  </div>
                </li>
              </ul>
              <base-button type="primary" class="text-white"
                >START FREE TRIAL</base-button
              >
            </template>
            <template slot="footer">
              <a href="#!" class="text-muted">Request a demo</a>
            </template>
          </card>
        </div>
        <div class="col-md-7 pl-md-0">
          <card
            class="card-pricing text-center my-5"
            type=""
            headerClasses="bg-transparent"
            bodyClasses="px-lg-6"
            footerClasses="bg-transparent"
          >
            <template slot="header">
              <h4 class="text-uppercase ls-1 text-primary py-3 mb-0">
                Enterprise
              </h4>
            </template>
            <template slot="body">
              <div class="card-description">
                What colors are good, before somebody tells you you shouldn’t
                like pink because that’s for girls, or you’d instantly become a
                gay two-year-old.
              </div>
              <table class="table table-bordered mt-3">
                <tbody>
                  <tr>
                    <td>24/7 Support</td>
                    <td>Design Tools</td>
                  </tr>
                  <tr>
                    <td>10K emails</td>
                    <td>Private Brand</td>
                  </tr>
                </tbody>
              </table>
            </template>
            <template slot="footer">
              <a href="#!" class="text-primary">Create new account</a>
            </template>
          </card>
        </div>
      </div>
    </div>
  </section>
</template>
<script></script>
<style></style>
