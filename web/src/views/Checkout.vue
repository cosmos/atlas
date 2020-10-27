<template>
  <div class="wrapper">
    <div class="section-shaped my-0 skew-separator skew-mini">
      <div class="page-header page-header-small header-filter">
        <div
          class="page-header-image"
          style="background-image: url('img/pages/photo-15.jpg');"
        ></div>
        <div class="container">
          <div class="header-body text-center mb-7">
            <div class="row justify-content-center">
              <div class="col-xl-5 col-lg-6 col-md-8 px-5">
                <h1 class="text-white">Checkout</h1>
                <p class="text-lead text-white">
                  We’ll show you the checkout page optimization tips that’ll
                  make your visitors want.
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="upper">
      <div class="container">
        <div class="row">
          <div class="col-lg-4">
            <div class="container">
              <h3 class="title text-white mt-3">Order summary</h3>
              <div class="row">
                <div class="card">
                  <div class="card-body">
                    <div class="media align-items-center mb-3">
                      <div class="col-md-5 col-6">
                        <img
                          src="img/pages/gucci.png"
                          alt="Rounded image"
                          class="img-fluid"
                        />
                      </div>
                      <div class="media-body">
                        <h2 class="h6">Shorts</h2>
                        <small class="d-block opacity-8">Small</small>
                      </div>
                      <div class="media-body text-right">
                        <span>$29</span>
                      </div>
                    </div>
                    <div class="media align-items-center">
                      <div class="col-md-5 col-6">
                        <img
                          src="img/pages/jacket.png"
                          alt="Rounded image"
                          class="img-fluid"
                        />
                      </div>
                      <div class="media-body">
                        <h2 class="h6 mb-0">Jacket</h2>
                      </div>
                      <div class="media-body text-right">
                        <span>$999</span>
                      </div>
                    </div>
                    <hr class="line-info mb-3" />
                    <form class="code-validate">
                      <label>Discount</label>
                      <div class="input-group">
                        <base-input
                          placeholder="Discout Code"
                          class="mb-0"
                        ></base-input>
                        <div class="input-group-append">
                          <base-button nativeType="submit" type="info"
                            >Apply</base-button
                          >
                        </div>
                      </div>
                    </form>
                    <hr class="line-info mb-3" />
                    <div class="media align-items-center">
                      <h3 class="h6 opacity-8 mr-3">Subtotal</h3>
                      <div class="media-body text-right">
                        <span>$1038</span>
                      </div>
                    </div>
                    <div class="media align-items-center">
                      <h3 class="h6 opacity-8 mr-3">Shipping</h3>
                      <div class="media-body text-right">
                        <span>$5.8</span>
                      </div>
                    </div>
                    <hr class="line-info mb-3" />
                    <div class="media align-items-center">
                      <h3 class="h6">Total</h3>
                      <div class="media-body text-right">
                        <span class="font-weight-semi-bold">$1045.8</span>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
          <div class="col-lg-8">
            <card bodyClasses="p-0">
              <template slot="body">
                <div class="container">
                  <h3 class="title mt-3">Billing address</h3>
                  <div class="row">
                    <div class="col-md-6">
                      <base-input
                        type="text"
                        name="firstName"
                        placeholder="Cristopher"
                        required
                        label="First name"
                      ></base-input>
                    </div>
                    <div class="col-md-6">
                      <base-input
                        type="text"
                        name="lastName"
                        placeholder="Thompson"
                        aria-label="Thompson"
                        required
                        label="Last name"
                      ></base-input>
                    </div>
                  </div>
                  <br />
                  <div class="row">
                    <div class="col-md-6">
                      <base-input
                        type="email"
                        name="emailAddress"
                        placeholder="thompson@gmail.com"
                        aria-label="thompson@gmail.com"
                        required
                        label="Email address"
                      ></base-input>
                    </div>
                    <div class="col-md-6">
                      <base-input
                        type="text"
                        name="emailAddress"
                        placeholder="+4 (0762) 230991"
                        aria-label="+4 (0762) 230991"
                        label="Phone"
                      ></base-input>
                    </div>
                  </div>
                  <br />
                  <div class="row">
                    <div class="col-md-8">
                      <base-input
                        type="text"
                        name="streetAddress"
                        placeholder="420 Long Beach, CA"
                        aria-label="420 Long Beach, CA"
                        required
                        label="Street address"
                      ></base-input>
                    </div>
                    <div class="col-md-4">
                      <base-input
                        type="text"
                        name="postcode"
                        placeholder="340112"
                        aria-label="340112"
                        required
                        label="Postcode/Zip"
                      ></base-input>
                    </div>
                  </div>
                  <br />
                  <div class="row">
                    <div class="col-md-4">
                      <base-input
                        type="text"
                        placeholder="YC7B 3UT"
                        aria-label="YC7B 3UT"
                        label="Apt, suite, etc."
                      ></base-input>
                    </div>
                    <div class="col-md-4">
                      <base-input
                        type="text"
                        name="cityAddress"
                        placeholder="London"
                        aria-label="London"
                        required
                        label="City"
                      ></base-input>
                    </div>
                    <div class="col-md-4">
                      <div class="mb-4">
                        <label class="labels">
                          Country
                          <span class="text-danger">*</span>
                        </label>
                        <choices-single
                          :options="optionsMultiple"
                          v-model="selected"
                        >
                          <option disabled value="0">Select one</option>
                        </choices-single>
                      </div>
                    </div>
                  </div>
                  <h4 class="title">Payment method</h4>
                  <tabs pills type="info" tabNavClasses="btn-group">
                    <tab-pane label="Credit Card">
                      <div class="row">
                        <div class="col-md-12">
                          <base-input
                            type="text"
                            name="cardNumber"
                            placeholder="**** **** **** ***"
                            aria-label="**** **** **** ***"
                            label="Card number"
                          ></base-input>
                        </div>
                      </div>
                      <br />
                      <div class="row">
                        <div class="col-md-8">
                          <base-input
                            type="text"
                            name="cardHolder"
                            placeholder="Jack Wayley"
                            aria-label="Jack Wayley"
                            label="Card holder"
                          ></base-input>
                        </div>
                        <div class="col-md-2">
                          <base-input
                            type="text"
                            name="cardExpirationDate"
                            placeholder="MM/YY"
                            aria-label="MM/YY"
                            label="Expiration"
                          ></base-input>
                        </div>
                        <div class="col-md-2">
                          <base-input
                            type="number"
                            name="cardCVC"
                            placeholder="***"
                            aria-label="***"
                            label="CVC"
                          ></base-input>
                        </div>
                      </div>
                      <div
                        class="d-flex justify-content-between align-items-center"
                      >
                        <a href="javascript:;"
                          ><span class="fa fa-angle-left mr-2"></span> Return to
                          cart</a
                        >
                        <base-button nativeType="submit" type="info" size="sm"
                          >Order now</base-button
                        >
                      </div>
                    </tab-pane>

                    <tab-pane label="Paypal">
                      <h5>Send payment to:</h5>
                      <base-input
                        type="text"
                        placeholder="Name or email"
                        class="mb-0"
                      ></base-input>
                      <br />
                      <base-button nativeType="submit" type="info" size="sm">
                        Pay with PayPal
                      </base-button>
                    </tab-pane>
                  </tabs>
                </div>
              </template>
            </card>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import ChoicesSingle from "@/components/SingleSelect";
import Tabs from "@/components/Tabs/Tabs.vue";
import TabPane from "@/components/Tabs/TabPane.vue";

export default {
  bodyClass: "checkout-page",
  components: {
    ChoicesSingle,
    Tabs,
    TabPane
  },
  data() {
    return {
      selected: 2,
      optionsMultiple: [
        { id: 1, text: "Czech Republic" },
        { id: 2, text: "Denmark" },
        { id: 3, text: "Dominican Republic" },
        { id: 4, text: "Iraq" },
        { id: 5, text: "Israel" },
        { id: 6, text: "Italy" },
        { id: 7, text: "Jamaica" },
        { id: 8, text: "Japan" },
        { id: 9, text: "Madagascar" },
        { id: 10, text: "Malta" },
        { id: 11, text: "Norway" },
        { id: 12, text: "Poland" },
        { id: 13, text: "Japan" },
        { id: 14, text: "Portugal" },
        { id: 15, text: "Romania" },
        { id: 16, text: "Russian Federation" },
        { id: 17, text: "Saint Lucia" },
        { id: 18, text: "Samoa" },
        { id: 19, text: "San Marino" },
        { id: 20, text: "Saudi Arabia" },
        { id: 21, text: "Spain" },
        { id: 22, text: "Switzerland" },
        { id: 23, text: "Sweden" },
        { id: 24, text: "Turkey" },
        { id: 25, text: "Uganda" },
        { id: 26, text: "Ukraine" },
        { id: 27, text: "United Arab Emirates" },
        { id: 28, text: "United Kingdom" },
        { id: 29, text: "United States" },
        { id: 30, text: "Viet Nam" }
      ]
    };
  }
};
</script>
<style>
.checkout-page .nav.btn-group li {
  flex: 1 1 auto;
  padding: 0;
  text-align: center;
}

.checkout-page .nav.btn-group li a {
  padding: 0.25rem 0.5rem;
  margin: 0;
  box-shadow: none;
  text-transform: uppercase;
  border: 1px solid transparent;
}

.checkout-page .nav.btn-group li:first-child a {
  border-top-right-radius: 0;
  border-bottom-right-radius: 0;
}

.checkout-page .nav.btn-group li:last-child a {
  border-top-left-radius: 0;
  border-bottom-left-radius: 0;
}

.checkout-page .nav.btn-group li a:not(.active) {
  border: 1px solid #11cdef;
}
</style>
