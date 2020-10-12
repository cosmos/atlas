<template>
  <div class="wrapper">
    <Contact></Contact>
    <div class="main">
      <div class="container-fluid">
        <div class="row mt-5 mb-4 pt-5">
          <div class="col-md-8 ml-auto mr-auto text-center mt-5">
            <badge type="info">Leave a message</badge>
            <h1 class="title">Tell us more about <b>yourself</b></h1>
            <h4 class="desc">
              Whether you have questions or you would just like to say hello,
              contact us.
            </h4>
          </div>
        </div>
        <div class="row">
          <div class="col-md-6 mx-auto ">
            <form
              role="form"
              class="card bg-secondary p-3"
              id="contact-form"
              method="post"
            >
              <div class="card-body">
                <div class="row">
                  <div class="col-md-6">
                    <div class="form-group">
                      <label>First name</label>
                      <base-input
                        alternative
                        placeholder="First Name..."
                        addonLeftIcon="ni ni-circle-08"
                      ></base-input>
                    </div>
                  </div>
                  <div class="col-md-6">
                    <div class="form-group">
                      <label>Last name</label>
                      <base-input
                        alternative
                        placeholder="Last Name..."
                        addonLeftIcon="ni ni-tag"
                      ></base-input>
                    </div>
                  </div>
                </div>
                <div class="form-group">
                  <label>Email address</label>
                  <base-input
                    alternative
                    placeholder="Email Here..."
                    addonLeftIcon="ni ni-email-83"
                  ></base-input>
                </div>
                <div class="form-group">
                  <label>Your message</label>
                  <textarea
                    name="message"
                    class="form-control form-control-alternative"
                    id="message"
                    rows="6"
                  ></textarea>
                </div>
                <div class="row">
                  <div class="col-md-6 ml-auto">
                    <base-button
                      nativeType="submit"
                      type="primary"
                      class="pull-right"
                      >Send Message</base-button
                    >
                  </div>
                </div>
              </div>
            </form>
          </div>
        </div>
      </div>
    </div>
    <gmap-map
      id="map-contactus-3"
      class="map skew-separator skew-top"
      style="height: 500px; position: relative; overflow: hidden;"
      :center="center"
      :zoom="13"
      :options="options"
      map-type-id="terrain"
    >
      <gmap-marker :position="center"> </gmap-marker>
    </gmap-map>
  </div>
</template>
<script>
import Contact from "./sections/Contact/Contact3.vue";
import { API_KEY } from "@/constants";
import Vue from "vue";
import * as VueGoogleMaps from "vue2-google-maps";
Vue.use(VueGoogleMaps, {
  load: {
    key: API_KEY
  }
});

export default {
  bodyClass: "contact-page",
  components: {
    Contact
  },
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
