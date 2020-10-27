<template>
  <div class="wrapper">
    <div class="page-header page-header-small header-filter">
      <div
        class="page-header-image"
        style="background-image: url('img/pages/nicolas-prieto.jpg');"
      ></div>
      <div class="content-center">
        <div class="row">
          <div class="col-md-6 mx-auto text-center">
            <h1 class="title text-white">
              A Place for Entrepreneurs to Share and Discover New Stories
            </h1>
            <base-button
              type="warning"
              icon="fa fa-twitter"
              rounded
              icon-only
            ></base-button>
            <base-button
              type="warning"
              icon="fa fa-instagram"
              rounded
              icon-only
            ></base-button>
          </div>
        </div>
      </div>
    </div>
    <div class="main main-raised">
      <div class="container">
        <div class="row">
          <div class="col-lg-4 col-md-6">
            <card type="default" class="card-blog">
              <template slot="image">
                <img class="img rounded mt-n7" src="img/ill/p7.svg" />
              </template>
              <template slot="body">
                <h4 class="display-4 text-white">Bussiness</h4>
                <p class="lead text-white mt-0">
                  Society has put up so many boundaries and limitations on
                  what’s right and wrong that it’s almost impossible...
                </p>
                <a href="javascript:;" class="btn btn-link text-white px-0"
                  >Read more</a
                >
              </template>
            </card>
          </div>
          <div class="col-lg-4 col-md-6">
            <card type="default" class="card-blog">
              <template slot="image">
                <img class="img rounded mt-n7" src="img/ill/p22.png" />
              </template>
              <template slot="body">
                <h4 class="display-4 text-white">Premium</h4>
                <p class="lead text-white mt-0">
                  Colors are good, before somebody tells you you shouldn’t like
                  pink because that’s for girls.
                </p>
                <a href="javascript:;" class="btn btn-link text-white px-0"
                  >Read more</a
                >
              </template>
            </card>
          </div>
          <div class="col-lg-4 col-md-6">
            <card type="default" class="card-blog">
              <template slot="image">
                <img class="img rounded mt-n7" src="img/ill/p7.svg" />
              </template>
              <template slot="body">
                <h4 class="display-4 text-white">Focus</h4>
                <p class="lead text-white mt-0">
                  Because it's about motivating the doers. Because I’m here to
                  follow my dreams and inspire other people.
                </p>
                <a href="javascript:;" class="btn btn-link text-white px-0"
                  >Read more</a
                >
              </template>
            </card>
          </div>
        </div>
      </div>
      <br />
      <!-- Related Stories -->
      <Blog></Blog>
      <br />
      <br />
      <!-- Latest Blogposts -->
      <Blog2></Blog2>
      <div class="section team-4 bg-default">
        <div class="team-2">
          <div class="container">
            <div class="row">
              <div class="col-md-8 mx-auto text-center mb-5">
                <h3 class="display-3 text-white">The Executive Team</h3>
                <h4 class="lead text-white">
                  This is the paragraph where you can write more details about
                  your team. Keep you user engaged by providing meaningful
                  information.
                </h4>
              </div>
            </div>
            <div class="row">
              <div class="col-lg-4 col-md-6">
                <card class="card-profile" data-image="profile-image">
                  <template slot="header">
                    <div class="card-image">
                      <base-dropdown menuClasses="dropdown-menu-right">
                        <div
                          slot="title"
                          class="nav-link dropdown-toggle text-white"
                          data-toggle="dropdown"
                        >
                          <i class="ni ni-settings-gear-65"></i>
                        </div>
                        <a class="dropdown-item" href="javascript:;"
                          >Edit Profile</a
                        >
                        <a class="dropdown-item" href="javascript:;"
                          >Settings</a
                        >
                        <a class="dropdown-item" href="javascript:;">Log out</a>
                      </base-dropdown>
                      <a href="javascript:;">
                        <img
                          class="img rounded"
                          src="img/faces/ali-pazani.jpg"
                        />
                      </a>
                    </div>
                  </template>
                  <template slot="body">
                    <h4 class="display-4 mb-0">Sofia Scarlett</h4>
                    <p class="lead">UX Designer</p>
                    <div class="table-responsive">
                      <ul class="list-unstyled ">
                        <li class="py-1">
                          <div class="d-flex align-items-center">
                            <div>
                              <div class="badge badge-circle badge-info mr-3">
                                <i class="ni ni-atom"></i>
                              </div>
                            </div>
                            <div>
                              <h6 class="mb-1">Dedicated entrepreneur</h6>
                            </div>
                          </div>
                        </li>
                        <li class="py-1">
                          <div class="d-flex align-items-center">
                            <div>
                              <div
                                class="badge badge-circle badge-success mr-3"
                              >
                                <i class="ni ni-user-run"></i>
                              </div>
                            </div>
                            <div>
                              <h6 class="mb-1">Outdors lover</h6>
                            </div>
                          </div>
                        </li>
                        <li class="py-1">
                          <div class="d-flex align-items-center">
                            <div>
                              <div class="badge badge-circle badge-danger mr-3">
                                <i class="ni ni-chart-bar-32"></i>
                              </div>
                            </div>
                            <div>
                              <h6 class="mb-1">Super friendly support team</h6>
                            </div>
                          </div>
                        </li>
                      </ul>
                    </div>
                  </template>
                </card>
              </div>
              <div class="col-lg-4 col-md-6">
                <card class="card-profile" data-image="profile-image">
                  <template slot="header">
                    <div class="card-image">
                      <base-dropdown menuClasses="dropdown-menu-right">
                        <div
                          slot="title"
                          class="nav-link dropdown-toggle text-white"
                          data-toggle="dropdown"
                        >
                          <i class="ni ni-settings-gear-65"></i>
                        </div>
                        <a class="dropdown-item" href="javascript:;"
                          >Edit Profile</a
                        >
                        <a class="dropdown-item" href="javascript:;"
                          >Settings</a
                        >
                        <a class="dropdown-item" href="javascript:;">Log out</a>
                      </base-dropdown>
                      <a href="javascript:;">
                        <img class="img rounded" src="img/faces/team-5.jpg" />
                      </a>
                    </div>
                  </template>
                  <template slot="body">
                    <h4 class="display-4 mb-0">Dylan Wyatt</h4>
                    <p class="lead">Team Lead</p>
                    <div class="table-responsive">
                      <ul class="list-unstyled ">
                        <li class="py-1">
                          <div class="d-flex align-items-center">
                            <div>
                              <div class="badge badge-circle badge-info mr-3">
                                <i class="ni ni-atom"></i>
                              </div>
                            </div>
                            <div>
                              <h6 class="mb-1">Dedicated entrepreneur</h6>
                            </div>
                          </div>
                        </li>
                        <li class="py-1">
                          <div class="d-flex align-items-center">
                            <div>
                              <div
                                class="badge badge-circle badge-success mr-3"
                              >
                                <i class="ni ni-user-run"></i>
                              </div>
                            </div>
                            <div>
                              <h6 class="mb-1">Urban exploration</h6>
                            </div>
                          </div>
                        </li>
                        <li class="py-1">
                          <div class="d-flex align-items-center">
                            <div>
                              <div class="badge badge-circle badge-danger mr-3">
                                <i class="ni ni-chart-bar-32"></i>
                              </div>
                            </div>
                            <div>
                              <h6 class="mb-1">
                                Able to get good at everything
                              </h6>
                            </div>
                          </div>
                        </li>
                      </ul>
                    </div>
                  </template>
                </card>
              </div>
              <div class="col-lg-4 col-md-6">
                <card class="card-profile" data-image="profile-image">
                  <template slot="header">
                    <div class="card-image">
                      <base-dropdown menuClasses="dropdown-menu-right">
                        <div
                          slot="title"
                          class="nav-link dropdown-toggle text-white"
                          data-toggle="dropdown"
                        >
                          <i class="ni ni-settings-gear-65"></i>
                        </div>
                        <a class="dropdown-item" href="javascript:;"
                          >Edit Profile</a
                        >
                        <a class="dropdown-item" href="javascript:;"
                          >Settings</a
                        >
                        <a class="dropdown-item" href="javascript:;">Log out</a>
                      </base-dropdown>
                      <a href="javascript:;">
                        <img class="img rounded" src="img/faces/atikh.jpg" />
                      </a>
                    </div>
                  </template>
                  <template slot="body">
                    <h4 class="display-4 mb-0">Mila Skylar</h4>
                    <p class="lead">Content Creator</p>
                    <div class="table-responsive">
                      <ul class="list-unstyled ">
                        <li class="py-1">
                          <div class="d-flex align-items-center">
                            <div>
                              <div class="badge badge-circle badge-info mr-3">
                                <i class="ni ni-atom"></i>
                              </div>
                            </div>
                            <div>
                              <h6 class="mb-1">High quality publication</h6>
                            </div>
                          </div>
                        </li>
                        <li class="py-1">
                          <div class="d-flex align-items-center">
                            <div>
                              <div
                                class="badge badge-circle badge-success mr-3"
                              >
                                <i class="ni ni-user-run"></i>
                              </div>
                            </div>
                            <div>
                              <h6 class="mb-1">Storytelling</h6>
                            </div>
                          </div>
                        </li>
                        <li class="py-1">
                          <div class="d-flex align-items-center">
                            <div>
                              <div class="badge badge-circle badge-danger mr-3">
                                <i class="ni ni-chart-bar-32"></i>
                              </div>
                            </div>
                            <div>
                              <h6 class="mb-1">
                                Master of words qualification
                              </h6>
                            </div>
                          </div>
                        </li>
                      </ul>
                    </div>
                  </template>
                </card>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="title text-center bg-default">
        <h3 class="text-white mb-0 pb-3">Show us some love</h3>
      </div>
      <div class="social-line social-line-big-icons bg-default">
        <div class="container">
          <div class="row">
            <div class="col-lg-3 col-md-6 mb-md-4 mb-lg-0">
              <base-button
                tag="a"
                href="javascript:void(0)"
                class="btn-icon btn-gradient-twitter"
                type=""
              >
                <i class="fa fa-twitter"></i>
                <p class="title">twitter</p>
                <p class="subtitle">@creativetim</p>
              </base-button>
            </div>
            <div class="col-lg-3 col-md-6 mb-md-4 mb-lg-0">
              <base-button
                tag="a"
                href="javascript:void(0)"
                class="btn-icon btn-gradient-facebook"
                type=""
              >
                <i class="fa fa-facebook"></i>
                <p class="title">facebook</p>
                <p class="subtitle">@creativetim</p>
              </base-button>
            </div>
            <div class="col-lg-3 col-md-6">
              <base-button
                tag="a"
                href="javascript:void(0)"
                class="btn-icon btn-gradient-slack"
                type=""
              >
                <i class="fa fa-slack"></i>
                <p class="title">slack</p>
                <p class="subtitle">@creativetim</p>
              </base-button>
            </div>
            <div class="col-lg-3 col-md-6">
              <base-button
                tag="a"
                href="javascript:void(0)"
                class="btn-icon btn-gradient-instagram"
                type=""
              >
                <i class="fa fa-instagram"></i>
                <p class="title">instagram</p>
                <p class="subtitle">@creativetim</p>
              </base-button>
            </div>
          </div>
        </div>
      </div>
      <div class="subscribe-line bg-default">
        <div class="container">
          <div class="row">
            <div class="col-lg-8 col-10 mx-auto">
              <div class="text-center">
                <h4 class="title text-white">Subscribe to our Newsletter</h4>
                <p class="description text-white">
                  Join our newsletter and get news in your inbox every week! We
                  hate spam too, so no worries about this.
                </p>
              </div>
              <div class="card card-raised card-form-horizontal">
                <div class="card-body">
                  <form method="" action="">
                    <div class="row">
                      <div class="col-sm-8">
                        <base-input
                          placeholder="Your Email..."
                          addonLeftIcon="ni ni-email-83"
                        ></base-input>
                      </div>
                      <div class="col-sm-4">
                        <base-button type="primary" class="btn-block"
                          >Subscribe</base-button
                        >
                      </div>
                    </div>
                  </form>
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
import Blog from "./sections/Blogs/Blog3.vue";
import Blog2 from "./sections/Blogs/Blog4.vue";
import BaseDropdown from "@/components/BaseDropdown";

export default {
  components: {
    Blog,
    Blog2,
    BaseDropdown
  }
};
</script>
<style></style>
