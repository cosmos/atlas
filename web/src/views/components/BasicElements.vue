<template>
  <section class="section section-components pb-0" id="section-components">
    <div class="container">
      <div class="row justify-content-center">
        <div class="col-lg-12">
          <h2 class="mb-5">
            <span>Basic Elements</span>
          </h2>
          <h3 class="h4 text-success font-weight-bold mb-4">Buttons</h3>
          <div class="mb-3 mt-5">
            <small class="text-uppercase font-weight-bold"
              >Pick your color</small
            >
          </div>
          <base-button type="primary">Primary</base-button>
          <base-button type="info">Info</base-button>
          <base-button type="success">Success</base-button>
          <base-button type="danger">Danger</base-button>
          <base-button type="warning">Warning</base-button>
          <base-button type="default">Default</base-button>
          <base-button type="secondary">Secondary</base-button>
          <div class="mb-3 mt-5">
            <small class="text-uppercase font-weight-bold">Rounded</small>
          </div>
          <base-button type="primary" class="btn-round">Primary</base-button>
          <base-button type="info" class="btn-round">Info</base-button>
          <base-button type="success" class="btn-round">Success</base-button>
          <base-button type="danger" class="btn-round">Danger</base-button>
          <base-button type="warning" class="btn-round">Warning</base-button>
          <base-button type="default" class="btn-round">Default</base-button>
          <base-button type="secondary" class="btn-round"
            >Secondary</base-button
          >
          <div class="mb-3 mt-5">
            <small class="text-uppercase font-weight-bold">Outline</small>
          </div>
          <base-button type="primary" outline>Primary</base-button>
          <base-button type="info" outline>Info</base-button>
          <base-button type="success" outline>Success</base-button>
          <base-button type="danger" outline>Danger</base-button>
          <base-button type="warning" outline>Warning</base-button>
          <base-button type="default" outline>Default</base-button>
          <base-button type="secondary" outline>Secondary</base-button>
          <div class="mb-3 mt-5">
            <small class="text-uppercase font-weight-bold"
              >Outline Rounded</small
            >
          </div>
          <base-button type="primary" class="btn-round" outline
            >Primary</base-button
          >
          <base-button type="info" class="btn-round" outline>Info</base-button>
          <base-button type="success" class="btn-round" outline
            >Success</base-button
          >
          <base-button type="danger" class="btn-round" outline
            >Danger</base-button
          >
          <base-button type="warning" class="btn-round" outline
            >Warning</base-button
          >
          <base-button type="default" class="btn-round" outline
            >Default</base-button
          >
          <base-button type="secondary" class="btn-round" outline
            >Secondary</base-button
          >
          <div class="mb-3 mt-5">
            <small class="text-uppercase font-weight-bold">Links</small>
          </div>
          <base-button
            tag="a"
            href="javascript:void()"
            type=""
            class="btn-link text-primary"
            >Primary</base-button
          >
          <base-button
            tag="a"
            href="javascript:void()"
            type=""
            class="btn-link text-info"
            >Info</base-button
          >
          <base-button
            tag="a"
            href="javascript:void()"
            type=""
            class="btn-link text-success"
            >Success</base-button
          >
          <base-button
            tag="a"
            href="javascript:void()"
            type=""
            class="btn-link text-danger"
            >Danger</base-button
          >
          <base-button
            tag="a"
            href="javascript:void()"
            type=""
            class="btn-link text-warning"
            >Warning</base-button
          >
          <base-button
            tag="a"
            href="javascript:void()"
            type=""
            class="btn-link text-default"
            >Default</base-button
          >
          <base-button
            tag="a"
            href="javascript:void()"
            type=""
            class="btn-link text-secondary"
            >Secondary</base-button
          >
          <div class="mb-3 mt-5">
            <small class="text-uppercase font-weight-bold">Gradients</small>
          </div>
          <base-button type="" class="btn-gradient-primary"
            >Primary</base-button
          >
          <base-button type="" class="btn-gradient-info">Info</base-button>
          <base-button type="" class="btn-gradient-success"
            >Success</base-button
          >
          <base-button type="" class="btn-gradient-danger">Danger</base-button>
          <base-button type="" class="btn-gradient-warning"
            >Warning</base-button
          >
          <base-button type="" class="btn-gradient-default"
            >Default</base-button
          >
          <base-button type="" class="btn-gradient-secondary"
            >Secondary</base-button
          >
          <div class="mb-3 mt-5">
            <small class="text-uppercase font-weight-bold"
              >Pick your style</small
            >
          </div>
          <base-button type="primary">Button</base-button>
          <base-button class="btn-icon" type="primary" icon="ni ni-bag-17"
            >Left Icon</base-button
          >
          <base-button class="btn-icon" type="primary">
            Right Icon
            <span class="btn-inner--icon ml-2"
              ><i class="ni ni-bag-17"></i
            ></span>
          </base-button>
          <base-button type="primary" icon="ni ni-bag-17"></base-button>
          <div class="mb-3 mt-5">
            <small class="text-uppercase font-weight-bold"
              >Pick your size</small
            >
          </div>
          <base-button type="primary" size="sm">Small</base-button>
          <base-button type="primary">Regular</base-button>
          <base-button type="primary" size="lg">Large Button</base-button>
          <div class="mb-3 mt-5">
            <small class="text-uppercase font-weight-bold">Floating</small>
          </div>
          <base-button
            type="primary"
            icon="ni ni-bag-17"
            rounded
            iconOnly
            size="sm"
          ></base-button>
          <base-button
            type="primary"
            icon="ni ni-bag-17"
            rounded
            iconOnly
          ></base-button>
          <base-button
            type="primary"
            icon="ni ni-bag-17"
            rounded
            iconOnly
            size="lg"
          ></base-button>

          <div class="row">
            <div class="col-md-6">
              <div class="mb-3 mt-5">
                <small class="text-uppercase font-weight-bold"
                  >Active &amp; Disabled</small
                >
              </div>
              <div class="row">
                <div class="col-md-6">
                  <base-button class="btn-block active" type="primary"
                    >Active</base-button
                  >
                </div>
                <div class="col-md-6">
                  <base-button class="btn-block disabled" type="primary"
                    >Disabled</base-button
                  >
                </div>
              </div>
            </div>
            <div class="col-md-6">
              <div class="mb-3 mt-5">
                <small class="text-uppercase font-weight-bold"
                  >Block Level</small
                >
              </div>
              <div class="row">
                <base-button class="btn-block" type="primary"
                  >Primary</base-button
                >
                <base-button class="btn-block" type="info">Info</base-button>
              </div>
            </div>
          </div>

          <div class="mb-3 mt-5">
            <small class="text-uppercase font-weight-bold"
              >Social buttons</small
            >
          </div>
          <base-button
            class="btn-icon btn-facebook"
            type=""
            icon="fa fa-facebook"
            >Facebook</base-button
          >
          <base-button class="btn-icon btn-twitter" type="" icon="fa fa-twitter"
            >Twitter</base-button
          >
          <base-button
            class="btn-icon btn-instagram"
            type=""
            icon="fa fa-instagram"
            >Instagram</base-button
          >
          <base-button class="btn-icon btn-github" type="" icon="fa fa-github"
            >Github</base-button
          >
          <base-button
            class="btn-icon btn-pinterest"
            type=""
            icon="fa fa-pinterest"
            >Pinterest</base-button
          >
          <base-button class="btn-icon btn-youtube" type="" icon="fa fa-youtube"
            >Youtube</base-button
          >
          <base-button class="btn-icon btn-vimeo" type="" icon="fa fa-vimeo"
            >Vimeo</base-button
          >
          <base-button class="btn-icon btn-slack" type="" icon="fa fa-slack"
            >Slack</base-button
          >
          <base-button
            class="btn-icon btn-dribbble"
            type=""
            icon="fa fa-dribbble"
            >Dribbble</base-button
          >
          <base-button class="btn-icon btn-reddit" type="" icon="fa fa-reddit"
            >Reddit</base-button
          >
          <base-button class="btn-icon btn-tumblr" type="" icon="fa fa-tumblr"
            >Tumblr</base-button
          >
          <base-button
            class="btn-icon btn-linkedin"
            type=""
            icon="fa fa-linkedin"
            >Linkedin</base-button
          >

          <div class="row">
            <div class="col-md-6">
              <div class="mb-3 mt-5">
                <small class="text-uppercase font-weight-bold">W/ Icons</small>
              </div>
              <base-button type="facebook" iconOnly
                ><i slot="icon" class="fa fa-facebook"></i
              ></base-button>
              <base-button type="twitter" iconOnly
                ><i slot="icon" class="fa fa-twitter"></i
              ></base-button>
              <base-button type="instagram" iconOnly
                ><i slot="icon" class="fa fa-instagram"></i
              ></base-button>
              <base-button type="pinterest" iconOnly
                ><i slot="icon" class="fa fa-pinterest"></i
              ></base-button>
              <base-button type="youtube" iconOnly
                ><i slot="icon" class="fa fa-youtube"></i
              ></base-button>
              <base-button type="vimeo" iconOnly
                ><i slot="icon" class="fa fa-vimeo"></i
              ></base-button>
              <base-button type="slack" iconOnly
                ><i slot="icon" class="fa fa-slack"></i
              ></base-button>
              <base-button type="dribbble" iconOnly
                ><i slot="icon" class="fa fa-dribbble"></i
              ></base-button>
            </div>
            <div class="col-md-6">
              <div class="mb-3 mt-5">
                <small class="text-uppercase font-weight-bold"
                  >Rounded w/ Icons</small
                >
              </div>
              <base-button
                type="gradient-facebook"
                icon="fa fa-facebook"
                iconOnly
                rounded
              ></base-button>
              <base-button
                type="gradient-twitter"
                icon="fa fa-twitter"
                iconOnly
                rounded
              ></base-button>
              <base-button
                type="gradient-instagram"
                icon="fa fa-instagram"
                iconOnly
                rounded
              ></base-button>
              <base-button
                type="gradient-pinterest"
                icon="fa fa-pinterest"
                iconOnly
                rounded
              ></base-button>
              <base-button
                type="gradient-youtube"
                icon="fa fa-youtube"
                iconOnly
                rounded
              ></base-button>
              <base-button
                type="gradient-vimeo"
                icon="fa fa-vimeo"
                iconOnly
                rounded
              ></base-button>
              <base-button
                type="gradient-slack"
                icon="fa fa-slack"
                iconOnly
                rounded
              ></base-button>
              <base-button
                type="gradient-dribbble"
                icon="fa fa-dribbble"
                iconOnly
                rounded
              ></base-button>
            </div>
          </div>

          <div class="row">
            <div class="col-md-6">
              <div class="mb-3 mt-5">
                <small class="text-uppercase font-weight-bold"
                  >W/ Gradient</small
                >
              </div>
              <base-button
                type="gradient-facebook"
                icon="fa fa-facebook"
                iconOnly
              ></base-button>
              <base-button
                type="gradient-twitter"
                icon="fa fa-twitter"
                iconOnly
              ></base-button>
              <base-button
                type="gradient-instagram"
                icon="fa fa-instagram"
                iconOnly
              ></base-button>
              <base-button
                type="gradient-pinterest"
                icon="fa fa-pinterest"
                iconOnly
              ></base-button>
              <base-button
                type="gradient-youtube"
                icon="fa fa-youtube"
                iconOnly
              ></base-button>
              <base-button
                type="gradient-vimeo"
                icon="fa fa-vimeo"
                iconOnly
              ></base-button>
              <base-button
                type="gradient-slack"
                icon="fa fa-slack"
                iconOnly
              ></base-button>
              <base-button
                type="gradient-dribbble"
                icon="fa fa-dribbble"
                iconOnly
              ></base-button>
            </div>
            <div class="col-md-6">
              <div class="mb-3 mt-5">
                <small class="text-uppercase font-weight-bold">Link</small>
              </div>
              <base-button
                type=""
                class="btn-facebook btn-simple"
                icon="fa fa-facebook"
              ></base-button>
              <base-button
                type=""
                class="btn-twitter btn-simple"
                icon="fa fa-twitter"
              ></base-button>
              <base-button
                type=""
                class="btn-instagram btn-simple"
                icon="fa fa-instagram"
              ></base-button>
              <base-button
                type=""
                class="btn-pinterest btn-simple"
                icon="fa fa-pinterest"
              ></base-button>
              <base-button
                type=""
                class="btn-youtube btn-simple"
                icon="fa fa-youtube"
              ></base-button>
              <base-button
                type=""
                class="btn-vimeo btn-simple"
                icon="fa fa-vimeo"
              ></base-button>
              <base-button
                type=""
                class="btn-slack btn-simple"
                icon="fa fa-slack"
              ></base-button>
              <base-button
                type=""
                class="btn-dribbble btn-simple"
                icon="fa fa-dribbble"
              ></base-button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
</template>
<script></script>
<style></style>
