<template>
  <div class="wrapper">
    <div class="section-shaped my-0 skew-separator skew-mini">
      <div class="page-header page-header-small header-filter">
        <div
          class="page-header-image"
          style="background-image: url('img/pages/nathan-dumlao.jpg');"
        ></div>
        <div class="container">
          <div class="header-body text-center mb-7">
            <div class="row justify-content-center">
              <div class="col-xl-5 col-lg-6 col-md-8 px-5">
                <h1 class="text-white">Invoice</h1>
                <p class="text-lead text-white">
                  Everyone’s born confident, and everything’s taken away from
                  you
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <section class="upper">
      <div class="container">
        <div class="row">
          <div class="col-md-10 mx-auto">
            <div class="card card-invoice">
              <div class="card-header text-center">
                <div class="row justify-content-between">
                  <div class="col-md-4 text-left">
                    <img
                      class="mb-2 w-100"
                      src="img/brand/blue.png"
                      alt="Logo"
                    />
                    <h6>
                      St. Independence Embankment, 050105 Bucharest, Romania
                    </h6>
                    <small class="d-block text-muted"
                      >tel: +4 (074) 1090873</small
                    >
                  </div>
                  <div class="col-lg-3 col-md-5 text-left mt-3">
                    <h4 class="mb-1">Billed to:</h4>
                    <span class="d-block">James Thompson</span>
                    <p>
                      Bld Mihail Kogalniceanu,<br />
                      7652 Bucharest,<br />
                      Romania
                    </p>
                  </div>
                </div>
                <br />
                <div class="row justify-content-md-between">
                  <div class="col-md-4">
                    <h3 class="mt-3 text-left">
                      Invoice no <br /><small class="mr-2">#0453119</small>
                    </h3>
                  </div>
                  <div class="col-lg-4 col-md-5">
                    <div class="row mt-5">
                      <div class="col-md-6">Invoice date:</div>
                      <div class="col-md-6">06/03/2019</div>
                    </div>
                    <div class="row">
                      <div class="col-md-6">Due date:</div>
                      <div class="col-md-6">11/03/2019</div>
                    </div>
                  </div>
                </div>
              </div>
              <div class="card-body">
                <div class="row">
                  <div class="col-12">
                    <div class="table-responsive">
                      <table class="table text-right">
                        <thead class="bg-default">
                          <tr>
                            <th scope="col" class="text-right text-white">
                              Item
                            </th>
                            <th scope="col" class="text-right text-white">
                              Qty
                            </th>
                            <th scope="col" class="text-right text-white">
                              Rate
                            </th>
                            <th scope="col" class="text-right text-white">
                              Amount
                            </th>
                          </tr>
                        </thead>
                        <tbody>
                          <tr>
                            <td>Premium Support</td>
                            <td>1</td>
                            <td>$ 9.00</td>
                            <td>$ 9.00</td>
                          </tr>
                          <tr>
                            <td>BLK Design System PRO</td>
                            <td>3</td>
                            <td>$ 100.00</td>
                            <td>$ 300.00</td>
                          </tr>
                          <tr>
                            <td>Parts for service</td>
                            <td>1</td>
                            <td>$ 89.00</td>
                            <td>$ 89.00</td>
                          </tr>
                        </tbody>
                        <tfoot>
                          <tr>
                            <th class="h4">Total</th>
                            <th colspan="3" class="text-right h4">$ 750</th>
                          </tr>
                        </tfoot>
                      </table>
                    </div>
                  </div>
                </div>
              </div>
              <div class="card-footer text-right">
                <div class="col-md-5 ml-auto">
                  <h5>Thank you!</h5>
                  <p class="description">
                    If you encounter any issues related to the invoice you can
                    contact us at:
                  </p>
                  <h6 class="d-block">
                    email:
                    <small class="text-muted">support@cretive-tim.com</small>
                  </h6>
                </div>
              </div>
            </div>
          </div>
          <div class="col-md-3 ml-auto">
            <base-button type="primary" class="btn-print mt-3 mb-6">
              Print
            </base-button>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
<script>
export default {
  bodyClass: "invoice-page"
};
</script>
<style></style>
