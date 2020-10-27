<template>
  <div class="wrapper">
    <div class="section section-hero section-shaped">
      <div class="page-header">
        <div
          class="page-header-image"
          style="background-image: url('img/ill/presentation_bg.png');"
        ></div>
        <div
          class="container-fluid shape-container d-flex align-items-center py-lg"
        >
          <div class="col px-0">
            <div class="row">
              <div class="col-lg-4 ml-5">
                <img
                  src="img/brand/green.png"
                  style="width: 200px;"
                  class="img-fluid"
                />
                <badge type="danger" class="ml-2">PRO</badge>
                <p class="lead">
                  Start the Development with this<br />
                  <b>Badass Bootstrap 4 Design System.</b>
                </p>
                <div class="btn-wrapper mt-5">
                  <base-button
                    tag="a"
                    href="https://www.creative-tim.com/product/vue-argon-design-system-pro"
                    type="success"
                    icon="fa fa-shopping-cart"
                    class="mb-3 mb-sm-0"
                  >
                    Purchase
                  </base-button>
                </div>
                <div class="mt-5">
                  <small class="font-weight-bold mb-0 mr-2"
                    >*proudly coded by</small
                  >
                  <img
                    src="img/brand/creativetim-black-slim.png"
                    style="height: 28px;"
                  />
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <section class="section-info">
      <div class="container">
        <div class="row">
          <div class="col-lg-4 col-md-6">
            <div class="info text-left">
              <div
                class="icon icon-lg icon-shape icon-shape-primary shadow rounded-circle"
              >
                <i class="ni ni-user-run"></i>
              </div>
              <h6 class="info-title text-uppercase text-primary pl-0">
                Huge number of components
              </h6>
              <p class="description opacity-8">
                Every element that you need in a product comes built in as a
                component. All components fit perfectly with each other and can
                take variations in colour.
              </p>
              <a href="javascript:void(0);" class="text-primary"
                >More about us
                <i class="ni ni-bold-right text-primary"></i>
              </a>
            </div>
          </div>
          <div class="col-lg-4 col-md-6">
            <div class="info text-left" style="margin-top:-50px;">
              <div
                class="icon icon-lg icon-shape icon-shape-success shadow rounded-circle"
              >
                <i class="ni ni-atom"></i>
              </div>
              <h6 class="info-title text-uppercase text-success pl-0">
                Multi-Purpose Sections
              </h6>
              <p class="description opacity-8">
                Putting together a page has never been easier than matching
                together sections. From team presentation to pricing options,
                you can easily customise and built your pages.
              </p>
              <a href="javascript:void(0);" class="text-primary"
                >Learn about our products
                <i class="ni ni-bold-right text-primary"></i>
              </a>
            </div>
          </div>
          <div class="col-lg-4 col-md-6">
            <div class="info text-left">
              <div
                class="icon icon-lg icon-shape icon-shape-warning shadow rounded-circle"
              >
                <i class="ni ni-istanbul"></i>
              </div>
              <h6 class="info-title text-uppercase text-warning pl-0">
                Multiple Example Pages
              </h6>
              <p class="description opacity-8">
                If you want to get inspiration or just show something directly
                to your clients, you can jump start your development with our
                pre-built example pages.
              </p>
              <a href="javascript:void(0);" class="text-primary"
                >Check our documentation
                <i class="ni ni-bold-right text-primary"></i>
              </a>
            </div>
          </div>
        </div>
      </div>
    </section>
    <section class="section-basic-components">
      <div class="container">
        <div class="row">
          <div class="col-lg-5 col-md-10 mb-md-5">
            <h1 class="display-3">
              Basic Elements<span class="text-primary">
                The core elements of your website</span
              >
            </h1>
            <p class="lead">
              We re-styled every Bootstrap 4 element to match the Argon Design
              System style. All the Bootstrap 4 components that you need in a
              development have been re-design with the new look. Besides the
              numerous basic elements, we have also created additional classes.
              All these items will help you take your project to the next level.
            </p>
          </div>
          <div class="col-lg-6 col-md-12">
            <div class="image-container">
              <img
                class="table-img"
                src="img/presentation-page/table.png"
                alt=""
              />
              <img
                class="coloured-card-btn-img"
                src="img/presentation-page/card-btn.png"
                alt=""
              />
              <img
                class="coloured-card-img"
                src="img/presentation-page/card-orange.png"
                alt=""
              />
              <img
                class="linkedin-btn-img"
                src="img/presentation-page/slack-btn.png"
                alt=""
              />
              <img class="w-100" src="img/ill/example-3.svg" />
            </div>
          </div>
        </div>
      </div>
    </section>
    <section class="section-cards mb-5">
      <div class="content-center">
        <div class="container-fluid">
          <div class="row">
            <div class="col-lg-5 col-md-10 col-12 mx-auto text-center">
              <h1 class="display-3">
                Unconventional cards<span class="text-primary">
                  One card for every problem</span
                >
              </h1>
              <p class="lead">
                We love cards and everybody on the web seems to. We have gone
                above and beyond with options for you to organise your
                information. From cards designed for blog posts, to product
                cards or user profiles, you will have many options to choose
                from. All the cards follow the material principles and have a
                design that stands out.
              </p>
            </div>
            <div class="col-md-12 col-md-6">
              <vue-glide
                class="presentation-cards"
                :startAt="0"
                :perView="4"
                :perTouch="1"
                :focutAt="2"
                type="carousel"
              >
                <vue-glide-slide>
                  <img
                    src="img/presentation-page/card-blog.png"
                    height="400"
                    width="350"
                  />
                </vue-glide-slide>
                <vue-glide-slide>
                  <img
                    src="img/presentation-page/card-pricing-1.png"
                    height="400"
                    width="350"
                  />
                </vue-glide-slide>
                <vue-glide-slide>
                  <img
                    src="img/presentation-page/card-background.png"
                    height="400"
                    width="350"
                  />
                </vue-glide-slide>
                <vue-glide-slide>
                  <img
                    src="img/presentation-page/card-pricing-2.png"
                    height="400"
                    width="350"
                  />
                </vue-glide-slide>
                <vue-glide-slide>
                  <img
                    src="img/presentation-page/card-profile-2.png"
                    height="400"
                    width="350"
                  />
                </vue-glide-slide>
                <template slot="control">
                  <button
                    class="glide__arrow glide__arrow--left"
                    data-glide-dir="<"
                  >
                    <i class="ni ni-bold-left text-default"></i>
                  </button>
                  <button
                    class="glide__arrow glide__arrow--right"
                    data-glide-dir=">"
                  >
                    <i class="ni ni-bold-right text-default"></i>
                  </button>
                </template>
              </vue-glide>
            </div>
          </div>
        </div>
      </div>
    </section>
    <section class="section-content">
      <div class="container">
        <div class="row">
          <div class="col-lg-6 mt-md-5 order-md-2 order-lg-1">
            <div class="image-container">
              <img
                class="img shadow rounded img-comments w-100"
                src="img/presentation-page/content-2.png"
              />
              <img
                class="img shadow rounded img-blog mt-5 w-100"
                src="img/presentation-page/content-1.png"
              />
            </div>
          </div>
          <div class="col-lg-6 mx-auto order-md-1">
            <div class="section-description">
              <h1 class="display-3">
                Content Areas<span class="text-danger">
                  For Areas that Need More Space</span
                >
              </h1>
              <p class="lead">
                We took into consideration multiple use cases and came up with
                some specific areas for you. If you need elements such as
                tables, comments, description areas, tabs and many others, we've
                got you covered. They're beautiful and easy to use for the end
                user navigating your website.
              </p>
            </div>
          </div>
        </div>
      </div>
    </section>
    <section class="section-sections bg-secondary">
      <div class="container">
        <div class="col-md-8 mx-auto">
          <div class="section-description text-center">
            <h2 class="display-2">Sections you will love</h2>
            <p class="lead">
              Build pages in no time using pre-made sections! From headers to
              footers, you will be able to choose the best combination for your
              project. We have created multiple sections for you to put together
              and customise into pixel perfect example pages.
            </p>
            <a
              href="sections.html"
              target="_blank"
              class="btn btn-primary btn-round"
              >View All Sections</a
            >
          </div>
        </div>
      </div>
      <div class="container-fluid">
        <div class="section-cols">
          <div class="row mt-5 mb-5">
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/header-1.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/header-2.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/header-3.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/header-4.jpg"
                alt=""
              />
            </div>
          </div>
          <div class="row mt-5 mb-5">
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/feature-1.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/feature-2.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/feature-3.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/feature-4.jpg"
                alt=""
              />
            </div>
          </div>
          <div class="row mt-5 mb-5">
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/blog-1.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/blog-2.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/blog-3.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/blog-4.jpg"
                alt=""
              />
            </div>
          </div>
          <div class="row mt-5 mb-5">
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/blog-5.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/team-1.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/team-2.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/team-3.jpg"
                alt=""
              />
            </div>
          </div>
          <div class="row mt-5 mb-5">
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page//sections/team-4.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/team-5.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/project-1.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/project-2.jpg"
                alt=""
              />
            </div>
          </div>
          <div class="row mt-5 mb-5">
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/project-3.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/pricing-1.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/pricing-2.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/pricing-3.jpg"
                alt=""
              />
            </div>
          </div>
          <div class="row mt-5 mb-5">
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/pricing-4.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/pricing-5.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/testimonial-1.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/testimonial-2.jpg"
                alt=""
              />
            </div>
          </div>
          <div class="row mt-5 mb-5">
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/testimonial-3.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/contact-1.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/contact-2.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3">
              <img
                class="shadow"
                src="img/presentation-page/sections/contact-3.jpg"
                alt=""
              />
            </div>
          </div>
          <div class="row mt-5 mb-5">
            <div class="col-md-3 ml-auto">
              <img
                class="shadow"
                src="img/presentation-page/sections/contact-4.jpg"
                alt=""
              />
            </div>
            <div class="col-md-3 mr-auto">
              <img
                class="shadow"
                src="img/presentation-page/sections/table-1.jpg"
                alt=""
              />
            </div>
          </div>
        </div>
      </div>
    </section>
    <section class="section-examples">
      <div class="container">
        <div class="col-md-8 mx-auto text-center">
          <div class="section-description">
            <h2 class="display-2">Example Pages</h2>
            <p class="lead">
              Forget about building everything from scratch. From landing pages
              to e-commerce or blog pages, you will be able jump start your
              development. Show your clients a quick prototype and get inspired
              for your next project!
            </p>
          </div>
        </div>
        <div class="row mt-5">
          <div class="col-md-4">
            <h5 class="title">About Us</h5>
            <div class="card">
              <router-link to="/about" target="_blank">
                <img src="img/presentation-page/pages/about-us.png" alt="" />
              </router-link>
            </div>
            <h5 class="title">Landing Page</h5>
            <div class="card">
              <router-link to="/landing-page" target="_blank">
                <img
                  src="img/presentation-page/pages/landing-page.png"
                  alt=""
                />
              </router-link>
            </div>
            <h5 class="title">Login Page</h5>
            <div class="card">
              <router-link to="/login" target="_blank">
                <img src="img/presentation-page/pages/login-page.png" alt="" />
              </router-link>
            </div>
            <h5 class="title">Register Page</h5>
            <div class="card">
              <router-link to="/register" target="_blank">
                <img
                  src="img/presentation-page/pages/register-page.png"
                  alt=""
                />
              </router-link>
            </div>
            <h5 class="title">Error 500</h5>
            <div class="card">
              <router-link to="/500-error" target="_blank">
                <img src="img/presentation-page/pages/500-error.png" alt="" />
              </router-link>
            </div>
            <h5 class="title">Error 404</h5>
            <div class="card">
              <router-link to="/error" target="_blank">
                <img src="img/presentation-page/pages/404-error.png" alt="" />
              </router-link>
            </div>
            <h5 class="title">Invoice Page</h5>
            <div class="card">
              <router-link to="/invoice" target="_blank">
                <img
                  src="img/presentation-page/pages/invoice-page.png"
                  alt=""
                />
              </router-link>
            </div>
          </div>
          <div class="col-md-4">
            <h5 class="title">Blog Post</h5>
            <div class="card">
              <router-link to="/blog-post" target="_blank">
                <img src="img/presentation-page/pages/blog-post.png" alt="" />
              </router-link>
            </div>
            <h5 class="title">Product Page</h5>
            <div class="card">
              <router-link to="/product-page" target="_blank">
                <img
                  src="img/presentation-page/pages/product-page.png"
                  alt=""
                />
              </router-link>
            </div>
            <h5 class="title">Profile Page</h5>
            <div class="card">
              <router-link to="/profile-page" target="_blank">
                <img
                  src="img/presentation-page/pages/profile-page.png"
                  alt=""
                />
              </router-link>
            </div>
            <h5 class="title">Account Settings</h5>
            <div class="card">
              <router-link to="/account" target="_blank">
                <img
                  src="img/presentation-page/pages/account-settings.png"
                  alt=""
                />
              </router-link>
            </div>
            <h5 class="title">Reset Page</h5>
            <div class="card">
              <router-link to="/reset" target="_blank">
                <img src="img/presentation-page/pages/reset-page.png" alt="" />
              </router-link>
            </div>
            <h5 class="title">Checkout Page</h5>
            <div class="card">
              <router-link to="/checkout" target="_blank">
                <img
                  src="img/presentation-page/pages/checkout-page.png"
                  alt=""
                />
              </router-link>
            </div>
            <h5 class="title">Chat Page</h5>
            <div class="card">
              <router-link to="/chat-page" target="_blank">
                <img src="img/presentation-page/pages/chat-page.png" alt="" />
              </router-link>
            </div>
          </div>
          <div class="col-md-4">
            <h5 class="title">Pricing Page</h5>
            <div class="card">
              <router-link to="/pricing" target="_blank">
                <img
                  src="img/presentation-page/pages/pricing-page.png"
                  alt=""
                />
              </router-link>
            </div>
            <h5 class="title">Ecommerce</h5>
            <div class="card">
              <router-link to="/ecommerce" target="_blank">
                <img
                  src="img/presentation-page/pages/ecommerce-page.png"
                  alt=""
                />
              </router-link>
            </div>
            <h5 class="title">Blog Posts</h5>
            <div class="card">
              <router-link to="/blog-posts" target="_blank">
                <img src="img/presentation-page/pages/blog-posts.png" alt="" />
              </router-link>
            </div>
            <h5 class="title">Contact Us</h5>
            <div class="card">
              <router-link to="/contact-us" target="_blank">
                <img src="img/presentation-page/pages/contact-us.png" alt="" />
              </router-link>
            </div>
          </div>
        </div>
      </div>
    </section>
    <section class="section-patterns">
      <div class="container-fluid">
        <div class="row">
          <div class="col-lg-4 col-md-10 col-12 mx-auto text-center align">
            <h1 class="display-3">
              Wonderful Patterns<span class="text-primary">
                Different styles, colors and emotions</span
              >
            </h1>
            <p class="lead">
              Devices mockups, Branding mockups, UI/UX Tools, Illustrations and
              much more. Free and premium. Use Paaatterns together with powerful
              design system and speed up your workflow
            </p>
            <a
              href="https://www.ls.graphics/?status=accepted&amp;expires=1574163072&amp;seller=15046&amp;affiliate=102023&amp;link=1681&amp;p_tok=05cf15f3-a34b-4dd4-aa6b-a4e8652ed45b"
              target="_blank"
              rel="nofollow"
              class="btn btn-primary"
              >View more</a
            >
          </div>
          <div class="col-lg-6 col-md-12">
            <img
              class="w-50 pattern-1 shadow"
              src="img/presentation-page/layer-1.jpg"
              alt=""
            />
            <img
              class="w-50 pattern-2 shadow"
              src="img/presentation-page/layer-2.jpg"
              alt=""
            />
            <img
              class="w-50 pattern-3 shadow"
              src="img/presentation-page/layer-3.jpg"
              alt=""
            />
            <img
              class="w-50 pattern-4 shadow"
              src="img/presentation-page/layer-4.jpg"
              alt=""
            />
          </div>
        </div>
      </div>
    </section>
    <section class="section-free-demo bg-secondary skew-separator">
      <div class="container">
        <div class="row">
          <div class="col-lg-7 col-md-12">
            <div class="section-description">
              <h3 class="display-3">Free Demo</h3>
              <p class="lead mb-4">
                Do you want to test and see the benefits of this Design System
                before purchasing it? You can give the demo version a try. It
                features enough basic components for you to get a feel of the
                design and also test the quality of the code. Get it free on
                GitHub!
              </p>
              <div class="github-buttons">
                <a
                  href="https://github.com/creativetimofficial/vue-argon-design-system"
                  target="_blank"
                  rel="nofollow"
                  class="btn btn-primary btn-round"
                  >View Demo on Github</a
                >
                <div class="mt-2">
                  <template>
                    <github-button
                      href="https://github.com/creativetimofficial/vue-argon-design-system"
                      data-icon="octicon-star"
                      data-size="large"
                      data-show-count="true"
                      aria-label="Star creativetimofficial/ct-vue-argon-design-system-pro on GitHub"
                      >Star</github-button
                    >
                  </template>
                </div>
              </div>
            </div>
          </div>
          <div class="col-lg-5 col-md-12">
            <div class="github-background-container">
              <i class="fa fa-github"></i>
            </div>
          </div>
        </div>
        <div class="row">
          <div class="col-lg-4 col-md-6 pt-5">
            <card class="card-pricing card-background">
              <template slot="body">
                <h2 class="card-title text-primary text-left ml-2">
                  Free Demo
                </h2>
                <ul>
                  <li class="text-left"><strong>70</strong> Components</li>
                  <li class="text-left"><strong>3</strong> Example Pages</li>
                  <li class="text-left">
                    <div class="badge badge-circle badge-danger">
                      <i class="fa fa-times text-white"></i>
                    </div>
                    Uncoventional cards
                  </li>
                  <li class="text-left">
                    <div class="badge badge-circle badge-danger">
                      <i class="fa fa-times text-white"></i>
                    </div>
                    Sections
                  </li>
                  <li class="text-left">
                    <div class="badge badge-circle badge-danger">
                      <i class="fa fa-times text-white"></i>
                    </div>
                    Photoshop for Prototype
                  </li>
                  <li class="text-left">
                    <div class="badge badge-circle badge-danger">
                      <i class="fa fa-times text-white"></i>
                    </div>
                    Premium Support
                  </li>
                </ul>
              </template>
            </card>
          </div>
          <div class="col-lg-4 col-md-6 pt-5">
            <card class="card-pricing card-background">
              <template slot="body">
                <h2 class="card-title text-primary text-left ml-2">
                  PRO Version
                </h2>
                <ul>
                  <li class="text-left"><strong>1100+</strong> Components</li>
                  <li class="text-left"><strong>17</strong> Example Pages</li>
                  <li class="text-left">
                    <div class="badge badge-circle badge-success">
                      <i class="ni ni-check-bold text-white"></i>
                    </div>
                    Uncoventional cards
                  </li>
                  <li class="text-left">
                    <div class="badge badge-circle badge-success">
                      <i class="ni ni-check-bold text-white"></i>
                    </div>
                    Sections
                  </li>
                  <li class="text-left">
                    <div class="badge badge-circle badge-success">
                      <i class="ni ni-check-bold text-white"></i>
                    </div>
                    Photoshop for Prototype
                  </li>
                  <li class="text-left">
                    <div class="badge badge-circle badge-success">
                      <i class="ni ni-check-bold text-white"></i>
                    </div>
                    Premium Support
                  </li>
                </ul>
              </template>
            </card>
          </div>
        </div>
      </div>
    </section>
    <section class="section-icons mb-5">
      <div class="container-fluid">
        <div class="row">
          <div class="col-lg-3 col-sm-2">
            <div class="icons-nucleo">
              <i class="first-left-icon ni ni-send text-primary"></i>
              <i
                class="second-left-icon ni ni-alert-circle-exc text-warning"
              ></i>
              <i class="third-left-icon ni ni-cart text-info"></i>
              <i class="fourth-left-icon ni ni-bold text-default"></i>
              <i class="fifth-left-icon ni ni-headphones text-danger"></i>
              <i class="sixth-left-icon ni ni-satisfied text-success"></i>
              <i class="seventh-left-icon ni ni-cart text-pink"></i>
              <i class="eighth-left-icon ni ni-spaceship text-info"></i>
              <i class="ninth-left-icon ni ni-sound-wave text-warning"></i>
              <i class="tenth-left-icon ni ni-heart-2 text-danger"></i>
            </div>
          </div>
          <div class="col-lg-6 col-sm-8 text-center">
            <h2 class="display-2">Custom Icons</h2>
            <p class="lead">
              Vue Argon Design System PRO comes with 100 custom demo icons made
              by our friends from
              <a
                href="https://nucleoapp.com/?ref=1712"
                target="_blank"
                rel="nofollow"
                >NucleoApp</a
              >. The official package contains over 20.729 icons which are
              looking great in combination with Vue Argon Design System PRO.
              Make sure you check all of them and use those that you like the
              most.
            </p>
            <br />
            <base-button
              tag="a"
              href="https://www.creative-tim.com/learning-lab/vue/icons/argon-design-system"
              type="primary"
              class="btn-round"
              target="_blank"
              >View Demo Icons</base-button
            >
            <base-button
              tag="a"
              outline
              href="https://nucleoapp.com/?ref=1712"
              type="primary"
              class="btn-round"
              target="_blank"
              >View All Icons</base-button
            >
          </div>
          <div class="col-lg-3 col-sm-2">
            <div class="icons-nucleo icons-nucleo-right text-success">
              <i class="first-right-icon ni ni-palette text-warning"></i>
              <i class="second-right-icon ni ni-tie-bow text-primary"></i>
              <i class="third-right-icon ni ni-pin text-info"></i>
              <i class="fourth-right-icon ni ni-key-25 text-purple"></i>
              <i class="fifth-right-icon ni ni-istanbul text-danger"></i>
              <i class="sixth-right-icon ni ni-bus-front-12 text-warning"></i>
              <i class="seventh-right-icon ni ni-image-02 text-success"></i>
              <i class="eighth-right-icon ni ni-world text-info"></i>
              <i class="ninth-right-icon ni ni-puzzle-10 text-primary"></i>
              <i class="tenth-right-icon ni ni-atom text-default"></i>
            </div>
          </div>
        </div>
      </div>
    </section>
    <section class="section-features bg-secondary">
      <div class="container-fluid">
        <div class="row">
          <div class="col-lg-6 col-md-12 pt-5">
            <div class="info info-horizontal">
              <div
                class="icon icon-shape icon-shape-warning rounded-circle text-white"
              >
                <i class="ni ni-money-coins text-warning"></i>
              </div>
              <div class="description pl-4 pt-2">
                <h5 class="title">Save Time &amp; Money</h5>
                <p>
                  Creating your design from scratch with dedicated designers can
                  be very expensive. Using Vue Argon Design System PRO you don't
                  have to worry about customising the basic Bootstrap 4 design
                  or its components.
                </p>
              </div>
            </div>
            <div class="info info-horizontal">
              <div
                class="icon icon-shape icon-shape-info rounded-circle text-white"
              >
                <i class="ni ni-bold text-info"></i>
              </div>
              <div class="description pl-4 pt-2">
                <h5 class="title">Bootstrap 4 &amp; Flexbox</h5>
                <p>
                  It is built over Bootstrap 4, it's fully responsive and has
                  all the benefits of the flexbox for the layout, grid system
                  and components. This is a huge advantage when you work with
                  columns.
                </p>
              </div>
            </div>
            <div class="info info-horizontal">
              <div
                class="icon icon-shape icon-shape-danger rounded-circle text-white"
              >
                <i class="ni ni-paper-diploma text-danger"></i>
              </div>
              <div class="description pl-4 pt-2">
                <h5 class="title">Fast Prototype</h5>
                <p>
                  Using Vue Argon Design System PRO you can create hundreds of
                  components combinations within seconds and present them to
                  your client. You just need to change some classes and colors.
                </p>
              </div>
            </div>
          </div>
          <div class="col-lg-6 col-md-12">
            <div class="image-container">
              <img class="w-100" src="img/presentation-page/ipad.png" alt="" />
            </div>
          </div>
        </div>
      </div>
    </section>
    <section class="section-testimonials mt-5">
      <div class="container">
        <div class="row">
          <div class="col-md-8 ml-auto mr-auto text-center">
            <h2 class="display-2 mb-5">Trusted by 880,000+ People</h2>
            <p class="lead">
              The UI Kits, Templates and Dashboards that we've created are used
              by
              <strong class="text-primary">880,000+ web developers</strong> in
              over <strong class="text-primary">1,500,000 Web Projects</strong>.
              This is what some of them think:
            </p>
          </div>
        </div>
        <div class="row">
          <div class="col-md-6 col-8 mx-auto">
            <b-carousel
              id="carousel"
              height="100vh"
              class="slide pointer-event"
              ref="carousel"
              indicators
            >
              <b-carousel-slide class="carousel-inner" height="100vh">
                <div class="card card-testimonial card-plain">
                  <div class="card-avatar">
                    <a href="#pablo">
                      <img
                        class="img img-raised rounded"
                        src="https://s3.amazonaws.com/creativetim_bucket/photos/154001/thumb.JPG?1501184024"
                      />
                    </a>
                  </div>
                  <div class="card-body text-center">
                    <p class="card-description text-dark">
                      "Awesome Design and very well organized code structure!
                      Also, it contains numerous elements using which achieving
                      the perfect or required template can be done with ease.
                      Great job!"
                    </p>
                    <h4 class="card-title">Stefan</h4>
                    <h6 class="category text-muted">Web Designer</h6>
                    <div class="card-footer">
                      <i class="fa fa-star text-warning"></i>
                      <i class="fa fa-star text-warning"></i>
                      <i class="fa fa-star text-warning"></i>
                      <i class="fa fa-star text-warning"></i>
                      <i class="fa fa-star text-warning"></i>
                    </div>
                  </div>
                </div>
              </b-carousel-slide>
              <b-carousel-slide class="carousel-inner">
                <div class="card card-testimonial card-plain">
                  <div class="card-avatar">
                    <a href="#pablo">
                      <img
                        class="img img-raised rounded"
                        src="https://s3.amazonaws.com/uifaces/faces/twitter/chadengle/128.jpg"
                      />
                    </a>
                  </div>
                  <div class="card-body text-center">
                    <p class="card-description text-dark">
                      "It looks great and its somewhat futuristics cant wait to
                      use it on a project here in nigeria i'm sure it would put
                      me ahead.. I cant wait to hv enough money to buy ur
                      product."
                    </p>
                    <h4 class="card-title">Mr. Bones</h4>
                    <h6 class="category text-muted">Web Designer</h6>
                    <div class="card-footer">
                      <i class="fa fa-star text-warning"></i>
                      <i class="fa fa-star text-warning"></i>
                      <i class="fa fa-star text-warning"></i>
                      <i class="fa fa-star text-warning"></i>
                      <i class="fa fa-star text-warning"></i>
                    </div>
                  </div>
                </div>
              </b-carousel-slide>
              <b-carousel-slide class="carousel-inner">
                <div class="card card-testimonial card-plain">
                  <div class="card-avatar">
                    <a href="#pablo">
                      <img
                        class="img img-raised rounded"
                        src="https://s3.amazonaws.com/creativetim_bucket/photos/123124/thumb.?1480480048"
                      />
                    </a>
                  </div>
                  <div class="card-body text-center">
                    <p class="card-description text-dark">
                      "Everything is perfect. Codes are really organized. It's
                      easy to edit for my own purposes. It's great that it is
                      built on top of Bootstrap 4."<br /><br />
                    </p>
                    <h4 class="card-title">DOĞA</h4>
                    <h6 class="category text-muted">Web Developer</h6>
                    <div class="card-footer">
                      <i class="fa fa-star text-warning"></i>
                      <i class="fa fa-star text-warning"></i>
                      <i class="fa fa-star text-warning"></i>
                      <i class="fa fa-star text-warning"></i>
                      <i class="fa fa-star text-warning"></i>
                    </div>
                  </div>
                </div>
              </b-carousel-slide>
            </b-carousel>
          </div>
          <a
            class="carousel-control-prev text-primary"
            v-on:click="prev"
            role="button"
            data-slide="prev"
          >
            <i class="ni ni-bold-left"></i>
            <span class="sr-only">Previous</span>
          </a>
          <a
            class="carousel-control-next text-primary"
            v-on:click="next"
            role="button"
            data-slide="next"
          >
            <i class="ni ni-bold-right"></i>
            <span class="sr-only">Next</span>
          </a>
        </div>
      </div>
    </section>
    <div class="section section-pricing" id="sectionBuy">
      <div class="container">
        <div class="row our-clients">
          <div class="col-lg-3 col-md-6 col-6">
            <img class="w-50" src="img/presentation-page/harvard.jpg" alt="" />
          </div>
          <div class="col-lg-3 col-md-6 col-6">
            <img
              class="w-50"
              src="img/presentation-page/microsoft.jpg"
              alt=""
            />
          </div>
          <div class="col-lg-3 col-md-6 col-6">
            <img class="w-50" src="img/presentation-page/vodafone.jpg" alt="" />
          </div>
          <div class="col-lg-3 col-md-6 col-6">
            <img class="w-50" src="img/presentation-page/stanford.jpg" alt="" />
          </div>
        </div>
        <div class="row mt-5" id="pricing">
          <div class="col-md-8 mx-auto text-center mt-5">
            <h2 class="title">
              Ready to grab
              <span class="text-success">Vue Argon Design System PRO</span>?
            </h2>
          </div>
          <div class="col-lg-8 col-md-12 mx-auto">
            <p class="description mb-5 text-center">
              You have <b>Free Unlimited Updates</b> and
              <b>Premium Support</b> on each package. You also have
              <b>30 days</b> to request a refund if you're not happy with your
              purchase.
              <br />
              Read more about
              <b
                ><a href="https://www.creative-tim.com/license" target="_blank"
                  >licenses here</a
                ></b
              >.
            </p>
            <div class="row">
              <div class="col-md-6 pt-lg-5">
                <card
                  class="card-pricing card-background"
                  style="background-image: url('img/ill/pattern_pricing6.svg"
                >
                  <template slot="body">
                    <h6 class="card-category text-danger text-uppercase">
                      Freelancer License
                    </h6>
                    <h1 class="card-title text-default"><small>$</small>109</h1>
                    <ul>
                      <li>
                        <div class="d-flex align-items-center">
                          <div class="text-left w-50">
                            <i class="ni ni-check-bold"></i>
                          </div>
                          <div class="text-left w-75">
                            <span class="pl-2 ml-3 text-sm">Documentation</span>
                          </div>
                        </div>
                      </li>
                      <li>
                        <div class="d-flex align-items-center">
                          <div class="text-left w-50">
                            <i class="ni ni-check-bold"></i>
                          </div>
                          <div class="text-left w-75">
                            <span class="pl-2 ml-3 text-sm">Full Code</span>
                          </div>
                        </div>
                      </li>
                      <li>
                        <div class="d-flex align-items-center">
                          <div class="text-left w-50">
                            <i class="ni ni-fat-delete"></i>
                          </div>
                          <div class="text-left w-75">
                            <span class="pl-2 ml-3 text-sm">Design Files</span>
                          </div>
                        </div>
                      </li>
                      <li>
                        <div class="d-flex align-items-center">
                          <div class="text-left w-50">
                            Domains:
                          </div>
                          <div class="text-left w-75">
                            <span class="pl-2 text-sm ml-3">1/product</span>
                          </div>
                        </div>
                      </li>
                      <li>
                        <div class="d-flex align-items-center">
                          <div class="text-left w-50">
                            Team size:
                          </div>
                          <div class="text-left w-75">
                            <span class="pl-2 text-sm ml-3">1 person</span>
                          </div>
                        </div>
                      </li>
                      <li>
                        <div class="d-flex align-items-center">
                          <div class="text-left w-50">
                            Tech Support:
                          </div>
                          <div class="text-left w-75">
                            <span class="pl-2 text-sm ml-3">6 months</span>
                          </div>
                        </div>
                      </li>
                      <li>
                        <div class="d-flex align-items-center">
                          <div class="text-left w-50">
                            Free Updates:
                          </div>
                          <div class="text-left w-75">
                            <span class="pl-2 text-sm ml-3">6 months</span>
                          </div>
                        </div>
                      </li>
                    </ul>
                    <base-button
                      tag="a"
                      href="https://www.creative-tim.com/product/vue-argon-design-system-pro"
                      type="primary"
                      class="mt-4"
                    >
                      Buy Now!
                    </base-button>
                  </template>
                </card>
              </div>
              <div class="col-md-6 pt-lg-5">
                <card
                  class="card-pricing card-background"
                  style="background-image: url('img/ill/pattern_pricing6.svg"
                >
                  <template slot="body">
                    <h6 class="card-category text-danger text-uppercase">
                      Company License
                    </h6>
                    <h1 class="card-title text-default"><small>$</small>299</h1>
                    <ul>
                      <li>
                        <div class="d-flex align-items-center">
                          <div class="text-left w-50">
                            <i class="ni ni-check-bold"></i>
                          </div>
                          <div class="text-left w-75">
                            <span class="pl-2 ml-3 text-sm">Documentation</span>
                          </div>
                        </div>
                      </li>
                      <li>
                        <div class="d-flex align-items-center">
                          <div class="text-left w-50">
                            <i class="ni ni-check-bold"></i>
                          </div>
                          <div class="text-left w-75">
                            <span class="pl-2 ml-3 text-sm">Full Code</span>
                          </div>
                        </div>
                      </li>
                      <li>
                        <div class="d-flex align-items-center">
                          <div class="text-left w-50">
                            <i class="ni ni-check-bold"></i>
                          </div>
                          <div class="text-left w-75">
                            <span class="pl-2 ml-3 text-sm">Design Files</span>
                          </div>
                        </div>
                      </li>
                      <li>
                        <div class="d-flex align-items-center">
                          <div class="text-left w-50">
                            Domains:
                          </div>
                          <div class="text-left w-75">
                            <span class="pl-2 text-sm ml-3">unlimited</span>
                          </div>
                        </div>
                      </li>
                      <li>
                        <div class="d-flex align-items-center">
                          <div class="text-left w-50">
                            Team size:
                          </div>
                          <div class="text-left w-75">
                            <span class="pl-2 text-sm ml-3">5-10 people</span>
                          </div>
                        </div>
                      </li>
                      <li>
                        <div class="d-flex align-items-center">
                          <div class="text-left w-50">
                            Tech Support:
                          </div>
                          <div class="text-left w-75">
                            <span class="pl-2 text-sm ml-3">12 months</span>
                          </div>
                        </div>
                      </li>
                      <li>
                        <div class="d-flex align-items-center">
                          <div class="text-left w-50">
                            Free Updates:
                          </div>
                          <div class="text-left w-75">
                            <span class="pl-2 text-sm ml-3">12 months</span>
                          </div>
                        </div>
                      </li>
                    </ul>
                    <base-button
                      tag="a"
                      href="https://www.creative-tim.com/product/vue-argon-design-system-pro"
                      type="primary"
                      class="mt-4"
                    >
                      Buy Now!
                    </base-button>
                  </template>
                </card>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import { Glide, GlideSlide } from "vue-glide-js";
import "vue-glide-js/dist/vue-glide.css";
import GithubButton from "vue-github-button";
import { BCarousel } from "bootstrap-vue/esm/components/carousel/carousel";
import { BCarouselSlide } from "bootstrap-vue/esm/components/carousel/carousel-slide";

export default {
  bodyClass: "presentation-page",
  name: "presentation",
  components: {
    [Glide.name]: Glide,
    [GlideSlide.name]: GlideSlide,
    GithubButton,
    BCarousel,
    BCarouselSlide
  },
  methods: {
    prev() {
      this.$refs.carousel.prev();
    },
    next() {
      this.$refs.carousel.next();
    }
  }
};
</script>
<style>
.presentation-page .carousel-item {
  height: 70vh;
}

.presentation-page .github-button .widget {
  margin-top: 8px;
}

.presentation-page .section-patterns .row .pattern-4 {
  left: 0em;
  top: -7em;
}

.presentation-page .section-patterns .row .pattern-2 {
  left: 1em;
  top: -4em;
}

.presentation-page .section-patterns .row .pattern-3 {
  left: -5em;
  top: -1em;
}
</style>
