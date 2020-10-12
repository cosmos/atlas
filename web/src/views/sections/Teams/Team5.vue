<template>
  <div class="team-5">
    <div class="container">
      <div class="row">
        <div class="col-md-8 mx-auto text-center mb-5">
          <h3 class="display-3">Meet our awesome team</h3>
          <p class="lead">
            People are so scared to lose that they don't even try. Like, one
            thing people can't say is that I'm not trying, and I'm not trying my
            hardest, and I'm not trying to do the best way I know how.
          </p>
        </div>
      </div>
      <div class="row">
        <div class="col-lg-4 col-md-6">
          <card class="card-profile" headerClasses="bg-info">
            <template slot="header">
              <img
                src="img/ill/inn.svg"
                class="position-absolute top-0 left-0"
              />
              <div class="card-avatar">
                <a href="javascript:;">
                  <img
                    class="img img-raised rounded-circle"
                    src="img/faces/team-2.jpg"
                  />
                </a>
              </div>
            </template>
            <template slot="body">
              <div class="d-flex justify-content-between">
                <a href="javascript:;" class="btn btn-sm btn-info mr-4 mt-3"
                  >Connect</a
                >
                <a
                  href="javascript:;"
                  class="btn btn-sm btn-default float-right mt-3"
                  >Message</a
                >
              </div>
              <div class="row">
                <div class="col">
                  <div class="card-profile-stats d-flex justify-content-center">
                    <div>
                      <span class="heading">22</span>
                      <span class="description">Friends</span>
                    </div>
                    <div>
                      <span class="heading">10</span>
                      <span class="description">Photos</span>
                    </div>
                    <div>
                      <span class="heading">89</span>
                      <span class="description">Comments</span>
                    </div>
                  </div>
                </div>
              </div>
              <div class="text-center">
                <h5 class="h4">
                  Melanie Paisley<span class="font-weight-light">, 31</span>
                </h5>
                <div class="font-weight-300">
                  <i class="ni location_pin mr-2"></i>Beijing, China
                </div>
              </div>
            </template>
          </card>
        </div>
        <div class="col-lg-4 col-md-6">
          <card class="card-profile" data-background="cover">
            <img
              class="img pattern pattern-top"
              src="img/ill/linth3.svg"
              alt="cover"
            />
            <template slot="header">
              <div class="card-avatar">
                <a href="javascript:;">
                  <img
                    class="img img-raised rounded-circle"
                    src="img/faces/michael.jpg"
                  />
                </a>
              </div>
            </template>
            <template slot="body">
              <div class="text-center">
                <h5 class="h5 mt-3">
                  Byron Reese<span class="font-weight-light">, 29</span>
                </h5>
                <div class="font-weight-300">
                  <i class="ni location_pin mr-2"></i>California
                </div>
              </div>
              <div class="card-profile-stats d-flex justify-content-center">
                <div>
                  <span class="heading">23K</span>
                  <span class="description">Followers</span>
                </div>
                <div>
                  <span class="heading">47</span>
                  <span class="description">Posts</span>
                </div>
              </div>
              <base-button type="success" class="btn-block">Like!</base-button>
            </template>
          </card>
        </div>
        <div class="col-lg-4 col-md-6">
          <card class="card-profile" headerClasses="bg-info">
            <template slot="header">
              <img
                src="img/ill/inn.svg"
                class="position-absolute top-0 left-0"
              />
              <div class="card-avatar">
                <a href="javascript:;">
                  <img
                    class="img img-raised rounded-circle"
                    src="img/faces/team-3.jpg"
                  />
                </a>
              </div>
            </template>
            <template slot="body">
              <div class="d-flex justify-content-between">
                <a href="javascript:;" class="btn btn-sm btn-info mr-4 mt-3"
                  >Connect</a
                >
                <a
                  href="javascript:;"
                  class="btn btn-sm btn-default float-right mt-3"
                  >Message</a
                >
              </div>
              <div class="row">
                <div class="col">
                  <div class="card-profile-stats d-flex justify-content-center">
                    <div>
                      <span class="heading">22</span>
                      <span class="description">Friends</span>
                    </div>
                    <div>
                      <span class="heading">10</span>
                      <span class="description">Photos</span>
                    </div>
                    <div>
                      <span class="heading">89</span>
                      <span class="description">Comments</span>
                    </div>
                  </div>
                </div>
              </div>
              <div class="text-center">
                <h5 class="h4">
                  Johana Doe<span class="font-weight-light">, 29</span>
                </h5>
                <div class="font-weight-300">
                  <i class="ni location_pin mr-2"></i>California
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
