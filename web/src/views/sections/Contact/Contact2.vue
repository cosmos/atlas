<template>
  <div class="contactus-2">
    <gmap-map
      id="map"
      class="map"
      :center="center"
      :zoom="13"
      :options="options"
      map-type-id="terrain"
    >
      <gmap-marker :position="center"> </gmap-marker>
    </gmap-map>
    <div class="col-lg-6 col-md-10">
      <form role="form" id="contact-form" method="post">
        <card class="card-contact card-raised" headerClasses="text-center">
          <template slot="header">
            <h4 class="card-title">Contact Us</h4>
          </template>
          <template slot="body">
            <div class="row">
              <div class="col-md-6">
                <div class="info info-horizontal">
                  <div
                    class="icon icon-shape icon-shape-primary shadow rounded-circle text-primary"
                  >
                    <i class="ni ni-mobile-button"></i>
                  </div>
                  <div class="description">
                    <h5 class="info-title mt-2">Give us a ring</h5>
                    <p class="ml-3">
                      Michael Jordan<br />
                      +40 762 321 762<br />
                      Mon - Fri, 8:00-22:00
                    </p>
                  </div>
                </div>
              </div>
              <div class="col-md-6">
                <div class="info info-horizontal">
                  <div
                    class="icon icon-shape icon-shape-primary shadow rounded-circle text-primary"
                  >
                    <i class="ni ni-square-pin"></i>
                  </div>
                  <div class="description">
                    <h5 class="info-title mt-2">Find us at:</h5>
                    <p class="ml-3">
                      Mihail Kogalnic 8,<br />
                      7652 Bucharest,<br />
                      Romania
                    </p>
                  </div>
                </div>
              </div>
            </div>
            <div class="row">
              <div class="col-md-6">
                <div class="form-group">
                  <label>First name</label>
                  <base-input
                    placeholder="First Name..."
                    addonLeftIcon="ni ni-circle-08"
                  ></base-input>
                </div>
              </div>
              <div class="col-md-6">
                <div class="form-group">
                  <label>Email address</label>
                  <base-input
                    placeholder="Email here..."
                    addonLeftIcon="ni ni-email-83"
                  ></base-input>
                </div>
              </div>
            </div>
            <div class="row mt-3">
              <div class="col-md-6">
                <base-checkbox class="mt-2">I'm not a robot</base-checkbox>
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
      </form>
    </div>
  </div>
</template>
<script>
import { API_KEY } from "@/constants";
import Vue from "vue";
import * as VueGoogleMaps from "vue2-google-maps";
Vue.use(VueGoogleMaps, {
  load: {
    key: API_KEY
  }
});

export default {
  data() {
    return {
      center: {
        lat: 40.748817,
        lng: -73.985428
      },
      options: {
        styles: [
          {
            featureType: "administrative",
            elementType: "labels.text.fill",
            stylers: [
              {
                color: "#444444"
              }
            ]
          },
          {
            featureType: "landscape",
            elementType: "all",
            stylers: [
              {
                color: "#f2f2f2"
              }
            ]
          },
          {
            featureType: "poi",
            elementType: "all",
            stylers: [
              {
                visibility: "off"
              }
            ]
          },
          {
            featureType: "road",
            elementType: "all",
            stylers: [
              {
                saturation: -100
              },
              {
                lightness: 45
              }
            ]
          },
          {
            featureType: "road.highway",
            elementType: "all",
            stylers: [
              {
                visibility: "simplified"
              }
            ]
          },
          {
            featureType: "road.arterial",
            elementType: "labels.icon",
            stylers: [
              {
                visibility: "off"
              }
            ]
          },
          {
            featureType: "transit",
            elementType: "all",
            stylers: [
              {
                visibility: "off"
              }
            ]
          },
          {
            featureType: "water",
            elementType: "all",
            stylers: [
              {
                color: "#C5CBF5"
              },
              {
                visibility: "on"
              }
            ]
          }
        ]
      }
    };
  }
};
</script>
<style></style>
