<template>
  <div
    class="contactus-1 bg-default"
    style="background-image: url('img/ill/1.svg')"
  >
    <div class="container">
      <div class="row">
        <div
          class="col-lg-5 col-md-5 d-flex justify-content-center flex-column"
        >
          <h2 class="title text-white">Get in Touch</h2>
          <h4 class="description text-white">
            You need more information? Check what other persons are saying about
            our product. They are very happy with their purchase.
          </h4>
          <div class="info info-horizontal">
            <div
              class="icon icon-shape icon-shape-primary shadow rounded-circle text-white"
            >
              <i class="ni ni-square-pin"></i>
            </div>
            <div class="description">
              <h4 class="info-title text-white">Find us at the office</h4>
              <p class="description ml-3 text-white">
                Bld Mihail Kogalniceanu, nr. 8,<br />
                7652 Bucharest,<br />
                Romania
              </p>
            </div>
          </div>
          <div class="info info-horizontal">
            <div
              class="icon icon-shape icon-shape-primary shadow rounded-circle text-white"
            >
              <i class="ni ni-mobile-button"></i>
            </div>
            <div class="description">
              <h4 class="info-title text-white">Give us a ring</h4>
              <p class="description ml-3 text-white">
                Michael Jordan<br />
                +40 762 321 762<br />
                Mon - Fri, 8:00-22:00
              </p>
            </div>
          </div>
        </div>
        <div class="col-lg-5 col-md-7 ml-auto mr-auto">
          <card headerClasses="text-center">
            <template slot="header" class="text-center">
              <h4 class="card-title">Contact Us</h4>
            </template>
            <template slot="body">
              <div class="row">
                <div class="col-md-6">
                  <div class="form-group">
                    <label>First name</label>
                    <base-input
                      placeholder="First name"
                      addonLeftIcon="ni ni-circle-08"
                    ></base-input>
                  </div>
                </div>
                <div class="col-md-6 pl-2">
                  <div class="form-group">
                    <label>Last name</label>
                    <base-input
                      placeholder="Last name"
                      addonLeftIcon="ni ni-collection"
                    ></base-input>
                  </div>
                </div>
              </div>
              <div class="form-group">
                <label>Email address</label>
                <base-input
                  placeholder="Email here..."
                  addonLeftIcon="ni ni-email-83"
                ></base-input>
              </div>
              <div class="form-group">
                <label>Your message</label>
                <textarea
                  name="message"
                  class="form-control"
                  id="message"
                  rows="6"
                ></textarea>
              </div>
              <div class="row">
                <div class="col-md-6">
                  <base-checkbox>I'm not a robot</base-checkbox>
                </div>
                <div class="col-md-6">
                  <base-button
                    nativeType="submit"
                    type="primary"
                    class="pull-right"
                    >Send Message</base-button
                  >
                </div>
              </div>
            </template>
          </card>
        </div>
      </div>
    </div>
  </div>
</template>
<script></script>
<style></style>
