<template>
  <div class="contactus-4">
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
    <div class="container">
      <div class="row">
        <div class="col-md-5">
          <h1 class="title text-white">Get in Touch</h1>
          <h4 class="description text-white">
            Do you need more information? Please contact us to find more about
            our products and services.
          </h4>
        </div>
        <div class="col-md-12 m-auto">
          <div class="card card-contact card-raised">
            <div class="row">
              <div class="col-lg-8 col-md-7 pr-md-0">
                <form role="form" class="p-3" id="contact-form" method="post">
                  <div class="card-header">
                    <h4 class="card-title">Send us a message</h4>
                  </div>
                  <div class="card-body">
                    <div class="row">
                      <div class="col-md-6">
                        <div class="form-group">
                          <label>First name</label>
                          <base-input
                            placeholder="First Name.."
                            addonLeftIcon="ni ni-circle-08"
                          ></base-input>
                        </div>
                      </div>
                      <div class="col-md-6">
                        <div class="form-group">
                          <label>Last name</label>
                          <base-input
                            placeholder="Last Name..."
                            addonLeftIcon="ni ni-collection"
                          ></base-input>
                        </div>
                      </div>
                    </div>
                    <div class="form-group">
                      <label>Email address</label>
                      <base-input
                        placeholder="Email Here"
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
                        <base-checkbox class="mb-3"
                          >I'm not a robot</base-checkbox
                        >
                      </div>
                      <div class="col-md-6">
                        <base-button
                          nativeType="submit"
                          type="info"
                          class="pull-right"
                          >Send Message</base-button
                        >
                      </div>
                    </div>
                  </div>
                </form>
              </div>
              <div class="col-lg-4 col-md-5 pl-md-0">
                <div class="info text-left bg-info">
                  <h4 class="card-title text-white">Contact information</h4>
                  <div class="info info-horizontal mt-lg-5">
                    <div
                      class="icon icon-shape bg-white rounded-circle text-info"
                    >
                      <i class="ni ni-square-pin"></i>
                    </div>
                    <div class="description">
                      <p class="info-title text-white mt-2">
                        345 Street 2, Bucharest
                      </p>
                    </div>
                  </div>
                  <div class="info info-horizontal">
                    <div
                      class="icon icon-shape bg-white rounded-circle text-info"
                    >
                      <i class="ni ni-mobile-button"></i>
                    </div>
                    <div class="description">
                      <p class="info-title text-white mt-2">
                        +16(3412) 421 241
                      </p>
                    </div>
                  </div>
                  <div class="info info-horizontal">
                    <div
                      class="icon icon-shape bg-white rounded-circle text-info"
                    >
                      <i class="ni ni-email-83"></i>
                    </div>
                    <div class="description">
                      <p class="info-title text-white mt-2">
                        contact@yoursite.com
                      </p>
                    </div>
                  </div>
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
