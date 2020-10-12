<template>
  <div class="project-1">
    <div class="container">
      <div class="row">
        <div class="col-lg-8 mr-auto ml-auto text-center mb-5">
          <h3 class="diaplay-3">Some of Our Awesome Projects</h3>
        </div>
      </div>
    </div>
    <nav class="navbar navbar-expand-lg bg-transparent mb-5">
      <div class="container">
        <div class="navbar-translate">
          <p>52 projects found</p>
        </div>
      </div>
    </nav>
    <div class="container">
      <div class="row">
        <div class="col-lg-4 col-md-6">
          <card headerClasses="mt-2" bodyClasses="text-center p-4">
            <template slot="header">
              <div class="float-left">
                <i class="ni ni-basket d-inline text-info"></i>
                <p class="d-inline">H-23</p>
              </div>
              <div class="float-right">
                <base-dropdown menuClasses="dropdown-menu-right">
                  <div
                    slot="title"
                    class="nav-link dropdown-toggle text-primary"
                    data-toggle="dropdown"
                  >
                    <i class="ni ni-settings-gear-65"></i>
                  </div>
                  <a class="dropdown-item" href="javascript:;">Action</a>
                  <a class="dropdown-item" href="javascript:;"
                    >Another action</a
                  >
                  <a class="dropdown-item" href="javascript:;"
                    >Something else</a
                  >
                </base-dropdown>
              </div>
            </template>
            <template slot="body">
              <a href="javascript:;">
                <img
                  src="img/slack.png"
                  alt="Circle image"
                  class="img-fluid rounded-circle shadow"
                  style="width: 100px;"
                />
              </a>
              <h4 class="card-title mt-3 mb-0">Slack</h4>
              <p class="card-description">
                We are happy to work at such a great project.
              </p>
              <h5 class="card-title mt-4">Members</h5>
              <div class="avatar-group">
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Ryan Tompson"
                >
                  <img alt="Image placeholder" src="img/faces/team-1.jpg" />
                </a>
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Romina Hadid"
                >
                  <img alt="Image placeholder" src="img/faces/team-2.jpg" />
                </a>
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Alexander Smith"
                >
                  <img alt="Image placeholder" src="img/faces/team-3.jpg" />
                </a>
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Jessica Doe"
                >
                  <img alt="Image placeholder" src="img//faces/team-4.jpg" />
                </a>
              </div>
            </template>
          </card>
        </div>

        <div class="col-lg-4 col-md-6">
          <card headerClasses="mt-2" bodyClasses="text-center p-4">
            <template slot="header">
              <div class="float-left">
                <i class="ni ni-button-power d-inline text-success"></i>
                <p class="d-inline">F-43</p>
              </div>
              <div class="float-right">
                <base-dropdown menuClasses="dropdown-menu-right">
                  <div
                    slot="title"
                    class="nav-link dropdown-toggle text-primary"
                    data-toggle="dropdown"
                  >
                    <i class="ni ni-settings-gear-65"></i>
                  </div>
                  <a class="dropdown-item" href="javascript:;">Action</a>
                  <a class="dropdown-item" href="javascript:;"
                    >Another action</a
                  >
                  <a class="dropdown-item" href="javascript:;"
                    >Something else</a
                  >
                </base-dropdown>
              </div>
            </template>
            <template slot="body">
              <a href="javascript:;">
                <img
                  src="img/spotify.jpeg"
                  alt="Circle image"
                  class="img-fluid rounded-circle shadow"
                  style="width: 100px;"
                />
              </a>
              <h4 class="card-title mt-3 mb-0">Spotify</h4>
              <p class="card-description">
                We strive to embrace and drive change in our industry.
              </p>
              <h5 class="card-title mt-4">Members</h5>
              <div class="avatar-group">
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Ryan Tompson"
                >
                  <img alt="Image placeholder" src="img/faces/team-1.jpg" />
                </a>
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Romina Hadid"
                >
                  <img alt="Image placeholder" src="img/faces/team-2.jpg" />
                </a>
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Alexander Smith"
                >
                  <img alt="Image placeholder" src="img/faces/team-3.jpg" />
                </a>
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Jessica Doe"
                >
                  <img alt="Image placeholder" src="img//faces/team-4.jpg" />
                </a>
              </div>
            </template>
          </card>
        </div>

        <div class="col-lg-4 col-md-6">
          <card headerClasses="mt-2" bodyClasses="text-center p-4">
            <template slot="header">
              <div class="float-left">
                <i class="ni ni-building d-inline text-danger"></i>
                <p class="d-inline ml-1">J-11</p>
              </div>
              <div class="float-right">
                <base-dropdown menuClasses="dropdown-menu-right">
                  <div
                    slot="title"
                    class="nav-link dropdown-toggle text-primary"
                    data-toggle="dropdown"
                  >
                    <i class="ni ni-settings-gear-65"></i>
                  </div>
                  <a class="dropdown-item" href="javascript:;">Action</a>
                  <a class="dropdown-item" href="javascript:;"
                    >Another action</a
                  >
                  <a class="dropdown-item" href="javascript:;"
                    >Something else</a
                  >
                </base-dropdown>
              </div>
            </template>
            <template slot="body">
              <a href="javascript:;">
                <img
                  src="img/dribbble.png"
                  alt="Circle image"
                  class="img-fluid rounded-circle shadow"
                  style="width: 100px;"
                />
              </a>
              <h4 class="card-title mt-3 mb-0">Dribbble</h4>
              <p class="card-description">
                The time has come to bring our plans and ideas to life.
              </p>
              <h5 class="card-title mt-4">Members</h5>
              <div class="avatar-group">
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Ryan Tompson"
                >
                  <img alt="Image placeholder" src="img/faces/team-1.jpg" />
                </a>
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Romina Hadid"
                >
                  <img alt="Image placeholder" src="img/faces/team-2.jpg" />
                </a>
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Alexander Smith"
                >
                  <img alt="Image placeholder" src="img/faces/team-3.jpg" />
                </a>
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Jessica Doe"
                >
                  <img alt="Image placeholder" src="img//faces/team-4.jpg" />
                </a>
              </div>
            </template>
          </card>
        </div>

        <div class="col-lg-4 col-md-6">
          <card headerClasses="mt-2" bodyClasses="text-center p-4">
            <template slot="header">
              <div class="float-left">
                <i class="ni ni-spaceship d-inline text-warning"></i>
                <p class="d-inline">A-11</p>
              </div>
              <div class="float-right">
                <base-dropdown menuClasses="dropdown-menu-right">
                  <div
                    slot="title"
                    class="nav-link dropdown-toggle text-primary"
                    data-toggle="dropdown"
                  >
                    <i class="ni ni-settings-gear-65"></i>
                  </div>
                  <a class="dropdown-item" href="javascript:;">Action</a>
                  <a class="dropdown-item" href="javascript:;"
                    >Another action</a
                  >
                  <a class="dropdown-item" href="javascript:;"
                    >Something else</a
                  >
                </base-dropdown>
              </div>
            </template>
            <template slot="body">
              <a href="javascript:;">
                <img
                  src="img/tim.png"
                  alt="Circle image"
                  class="img-fluid rounded-circle shadow"
                  style="width: 100px;"
                />
              </a>
              <h4 class="card-title mt-3 mb-0">CreativeTim</h4>
              <p class="card-description">
                We are developing the best design projects.
              </p>
              <h5 class="card-title mt-4">Members</h5>
              <div class="avatar-group">
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Ryan Tompson"
                >
                  <img alt="Image placeholder" src="img/faces/team-1.jpg" />
                </a>
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Romina Hadid"
                >
                  <img alt="Image placeholder" src="img/faces/team-2.jpg" />
                </a>
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Alexander Smith"
                >
                  <img alt="Image placeholder" src="img/faces/team-3.jpg" />
                </a>
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Jessica Doe"
                >
                  <img alt="Image placeholder" src="img//faces/team-4.jpg" />
                </a>
              </div>
            </template>
          </card>
        </div>

        <div class="col-lg-4 col-md-6">
          <card headerClasses="mt-2" bodyClasses="text-center p-4">
            <template slot="header">
              <div class="float-left">
                <i class="ni ni-umbrella-13 d-inline text-primary"></i>
                <p class="d-inline">A-11</p>
              </div>
              <div class="float-right">
                <base-dropdown menuClasses="dropdown-menu-right">
                  <div
                    slot="title"
                    class="nav-link dropdown-toggle text-primary"
                    data-toggle="dropdown"
                  >
                    <i class="ni ni-settings-gear-65"></i>
                  </div>
                  <a class="dropdown-item" href="javascript:;">Action</a>
                  <a class="dropdown-item" href="javascript:;"
                    >Another action</a
                  >
                  <a class="dropdown-item" href="javascript:;"
                    >Something else</a
                  >
                </base-dropdown>
              </div>
            </template>
            <template slot="body">
              <a href="javascript:;">
                <img
                  src="img/dropbox.png"
                  alt="Circle image"
                  class="img-fluid rounded-circle shadow"
                  style="width: 100px;"
                />
              </a>
              <h4 class="card-title mt-3 mb-0">DropBox</h4>
              <p class="card-description">
                It is important to save every project safely with our app.
              </p>
              <h5 class="card-title mt-4">Members</h5>
              <div class="avatar-group">
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Ryan Tompson"
                >
                  <img alt="Image placeholder" src="img/faces/team-1.jpg" />
                </a>
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Romina Hadid"
                >
                  <img alt="Image placeholder" src="img/faces/team-2.jpg" />
                </a>
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Alexander Smith"
                >
                  <img alt="Image placeholder" src="img/faces/team-3.jpg" />
                </a>
                <a
                  href="javascript:;"
                  class="avatar avatar-lg rounded-circle"
                  v-b-tooltip.hover.top
                  title="Jessica Doe"
                >
                  <img alt="Image placeholder" src="img//faces/team-4.jpg" />
                </a>
              </div>
            </template>
          </card>
        </div>

        <div class="col-lg-4 col-md-6">
          <card headerClasses="mt-2" bodyClasses="text-center p-4">
            <template slot="header">
              <div class="float-left">
                <i class="ni ni-bus-front-12 d-inline"></i>
                <p class="d-inline">E-28</p>
              </div>
              <div class="float-right">
                <base-dropdown menuClasses="dropdown-menu-right">
                  <div
                    slot="title"
                    class="nav-link dropdown-toggle text-primary"
                    data-toggle="dropdown"
                  >
                    <i class="ni ni-settings-gear-65"></i>
                  </div>
                  <a class="dropdown-item" href="javascript:;">Action</a>
                  <a class="dropdown-item" href="javascript:;"
                    >Another action</a
                  >
                  <a class="dropdown-item" href="javascript:;"
                    >Something else</a
                  >
                </base-dropdown>
              </div>
            </template>
            <template slot="body">
              <a href="javascript:;">
                <img
                  src="img/unass.jpg"
                  alt="Circle image"
                  class="img-fluid rounded-circle shadow"
                  style="width: 100px;"
                />
              </a>
              <h4 class="card-title mt-3 mb-0">Unassigned</h4>
              <p class="card-description">
                Here you can add your description and bellow add your members.
              </p>
              <h5 class="card-title mt-4">Members</h5>
              <div class="avatar-group">
                <a
                  href="javascript:;"
                  class="avatar avatar-sm rounded-circle"
                  v-b-tooltip.hover.top
                  title="Add members"
                >
                  <i class="ni ni-key-25"></i>
                </a>
              </div>
            </template>
          </card>
        </div>
      </div>
      <footer class="mt-5">
        <base-pagination
          v-model="currentPage"
          class="float-left"
          :total="50"
        ></base-pagination>

        <div class="text-right">
          <p>Showing 6 out of 36</p>
          <div></div>
        </div>
      </footer>
    </div>
  </div>
</template>
<script>
import BaseDropdown from "@/components/BaseDropdown";
import { VBTooltip } from "bootstrap-vue/esm/directives/tooltip/tooltip";

export default {
  components: {
    BaseDropdown
  },
  directives: {
    BTooltip: VBTooltip
  },
  data() {
    return {
      currentPage: 1
    };
  }
};
</script>
<style></style>
