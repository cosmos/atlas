<template>
  <div class="wrapper">
    <div
      class="page-header header-filter page-header-small skew-separator skew-mini"
    >
      <div
        class="page-header-image"
        style="background-image: url('img/pages/boxed-water.jpg');"
      ></div>
      <div class="container">
        <div class="row">
          <div class="col-lg-7 col-md-7 mr-auto text-left">
            <h1 class="title text-white">Our products</h1>
            <br />
            <h4 class="category text-white opacity-8">
              This is the best way to find your favorite stuff
            </h4>
          </div>
        </div>
      </div>
      <!-- SVG separator -->
      <div class="separator separator-bottom separator-skew">
        <svg
          x="0"
          y="0"
          viewBox="0 0 2560 100"
          preserveAspectRatio="none"
          version="1.1"
          xmlns="http://www.w3.org/2000/svg"
        >
          <polygon class="fill-white" points="2560 0 2560 100 0 100"></polygon>
        </svg>
      </div>
    </div>
    <div class="section section-item">
      <div class="container">
        <div class="row">
          <div class="col-lg-6 col-md-12 ">
            <b-carousel id="carousel1" ref="carousel1">
              <!-- Text slides with image -->
              <b-carousel-slide img-src="img/pages/shirt.png">
              </b-carousel-slide>
              <b-carousel-slide img-src="img/pages/shorts.png">
              </b-carousel-slide>
              <b-carousel-slide img-src="img/pages/tshirt.png">
              </b-carousel-slide>
              <a
                class="carousel-control-prev"
                v-on:click="prev"
                role="button"
                data-slide="prev"
              >
                <button
                  type="button"
                  class="btn btn-warning btn-icon btn-round btn-sm"
                  name="button"
                >
                  <i class="ni ni-bold-left"></i>
                </button>
              </a>

              <a
                class="carousel-control-next"
                v-on:click="next"
                role="button"
                data-slide="next"
              >
                <button
                  type="button"
                  class="btn btn-warning btn-icon btn-round btn-sm"
                  name="button"
                >
                  <i class="ni ni-bold-right"></i>
                </button>
              </a>
            </b-carousel>
          </div>
          <div class="col-lg-6 col-md-12 mx-auto">
            <h2 class="title">Givenchy</h2>
            <div class="stats">
              <div class="stars text-warning">
                <i class="fa fa-star"></i>
                <i class="fa fa-star"></i>
                <i class="fa fa-star"></i>
                <i class="fa fa-star"></i>
                <i class="fa fa-star"></i>
                <p class="d-inline">(76 customer reviews)</p>
              </div>
            </div>
            <br />
            <h2 class="main-price">$3,390</h2>
            <h6 class="category">Description</h6>
            <p class="description">
              Eres' daring 'Grigri Fortune' swimsuit has the fit and coverage of
              a bikini in a one-piece silhouette. This fuchsia style is crafted
              from the label's sculpting peau douce fabric and has flattering
              cutouts through the torso and back. Wear yours with mirrored
              sunglasses on vacation.
            </p>
            <br />
            <div class="row pick-size">
              <div class="col-lg-4 col-md-4">
                <label>Quantity</label>
                <div class="input-group">
                  <div class="input-group-btn">
                    <base-button
                      type="warning"
                      outline
                      class="btn-round"
                      size="sm"
                      @click.native="decreaseQuantity()"
                    >
                      <i class="ni ni-fat-delete"></i>
                    </base-button>
                  </div>
                  <input
                    type="text"
                    id="myNumber"
                    class="form-control input-number"
                    :value="quantity"
                  />
                  <div class="input-group-btn">
                    <base-button
                      type="warning"
                      outline
                      class="btn-round"
                      size="sm"
                      @click.native="increaseQuantity()"
                    >
                      <i class="ni ni-fat-add"></i>
                    </base-button>
                  </div>
                </div>
              </div>
              <div class="col-lg-4 col-md-4 col-sm-6">
                <label>Select color</label>
                <choices-single :options="options1" v-model="selected">
                  <option disabled value="0">Select one</option>
                </choices-single>
              </div>
              <div class="col-lg-4 col-md-4 col-sm-6">
                <label>Select size</label>
                <choices-single :options="options2" v-model="selected">
                  <option disabled value="0">Select one</option>
                </choices-single>
              </div>
            </div>
            <br />
            <div class="row justify-content-start">
              <base-button type="warning" class="ml-3">
                Add to Cart &nbsp;<i class="ni ni-cart"></i>
              </base-button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <br />
    <br />
    <br />
    <br />
    <br />
    <!-- Testimonial Section -->
    <Testimonial></Testimonial>
    <div
      class="section related-products bg-secondary skew-separator skew-top mt-7"
    >
      <div class="container">
        <div class="col-md-8">
          <h2 class="title">You may also like</h2>
        </div>
        <div class="row">
          <div class="col-lg-3 col-md-6">
            <card class="card-product">
              <template slot="image">
                <a href="#pablo">
                  <img class="img rounded" src="img/pages/tshirt.png" />
                </a>
              </template>
              <template slot="body">
                <h6 class="category text-warning">Trending</h6>
                <h4 class="card-title">
                  <a href="javascript:;" class="card-link">Gucci</a>
                </h4>
                <div class="card-description opacity-8 mt-2">
                  Dolce &amp; Gabbana's 'Greta' tote has been crafted in Italy
                  from hard-wearing red textured-leather.
                </div>
                <div class="card-footer">
                  <div class="price-container">
                    <span class="price">€1,459</span>
                  </div>
                  <base-button
                    type="warning"
                    iconOnly
                    size="sm"
                    class="btn-round pull-right"
                  >
                    <i class="ni ni-favourite-28"></i>
                  </base-button>
                </div>
              </template>
            </card>
          </div>
          <div class="col-lg-3 col-md-6">
            <card class="card-product">
              <template slot="image">
                <a href="#pablo">
                  <img class="img rounded" src="img/pages/shorts.png" />
                </a>
              </template>
              <template slot="body">
                <h6 class="category text-warning">Popular</h6>
                <h4 class="card-title">
                  <a href="javascript:;" class="card-link">Balmain</a>
                </h4>
                <div class="card-description opacity-8 mt-2">
                  Balmain's mid-rise skinny jeans are cut with stretch to ensure
                  they retain their second-skin fit but move comfortably.
                </div>
                <div class="card-footer">
                  <div class="price-container">
                    <span class="price">€459</span>
                  </div>
                  <base-button
                    type="warning"
                    iconOnly
                    size="sm"
                    class="btn-round pull-right"
                  >
                    <i class="ni ni-favourite-28"></i>
                  </base-button>
                </div>
              </template>
            </card>
          </div>
          <div class="col-lg-3 col-md-6">
            <card class="card-product">
              <template slot="image">
                <a href="#pablo">
                  <img class="img rounded" src="img/pages/shirt.png" />
                </a>
              </template>
              <template slot="body">
                <h6 class="category text-warning">Popular</h6>
                <h4 class="card-title">
                  <a href="javascript:;" class="card-link">Balenciaga</a>
                </h4>
                <div class="card-description opacity-8 mt-2">
                  Balenciaga's black textured-leather wallet is finished with
                  the label's iconic 'Giant' studs. This is where you can...
                </div>
                <div class="card-footer">
                  <div class="price-container">
                    <span class="price">€559</span>
                  </div>
                  <base-button
                    type="warning"
                    iconOnly
                    size="sm"
                    class="btn-round pull-right"
                  >
                    <i class="ni ni-favourite-28"></i>
                  </base-button>
                </div>
              </template>
            </card>
          </div>
          <div class="col-lg-3 col-md-6">
            <card class="card-product">
              <template slot="image">
                <a href="#pablo">
                  <img class="img rounded" src="img/pages/jeans.png" />
                </a>
              </template>
              <template slot="body">
                <h6 class="category text-warning">Trending</h6>
                <h4 class="card-title">
                  <a href="javascript:;" class="card-link">D &amp; G</a>
                </h4>
                <div class="card-description opacity-8 mt-2">
                  Dolce &amp; Gabbana's 'Greta' tote has been crafted in Italy
                  from hard-wearing red textured-leather.
                </div>
                <div class="card-footer">
                  <div class="price-container">
                    <span class="price">€1,359</span>
                  </div>
                  <base-button
                    type="warning"
                    iconOnly
                    size="sm"
                    class="btn-round pull-right"
                  >
                    <i class="ni ni-favourite-28"></i>
                  </base-button>
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
import { BCarousel } from "bootstrap-vue/esm/components/carousel/carousel";
import { BCarouselSlide } from "bootstrap-vue/esm/components/carousel/carousel-slide";
import Testimonial from "./sections/Testimonial/Testimonial3.vue";
import ChoicesSingle from "@/components/SingleSelect";

export default {
  bodyClass: "product-page",
  components: {
    BCarousel,
    BCarouselSlide,
    Testimonial,
    ChoicesSingle
  },
  data() {
    return {
      quantity: 1,
      selected: 2,
      options1: [
        { id: 1, text: "Black" },
        { id: 2, text: "Grey" },
        { id: 3, text: "White" }
      ],
      options2: [
        { id: 1, text: "S" },
        { id: 2, text: "M" },
        { id: 3, text: "L" },
        { id: 4, text: "XS" },
        { id: 5, text: "XL" }
      ]
    };
  },
  methods: {
    prev() {
      this.$refs.carousel1.prev();
    },
    next() {
      this.$refs.carousel1.next();
    },
    increaseQuantity() {
      this.quantity++;
    },
    decreaseQuantity() {
      if (this.quantity > 1) {
        this.quantity--;
      }
    }
  }
};
</script>
<style>
.product-page .carousel-item {
  height: 60vh;
}

.product-page .carousel a.carousel-control-next {
  right: -20px;
}

.product-page .carousel a.carousel-control-prev {
  left: -20px;
}
</style>
