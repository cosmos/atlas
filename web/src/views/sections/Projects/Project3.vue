<template>
  <div class="projects-3">
    <div class="container">
      <div class="row">
        <div class="col-md-8 mr-auto ml-auto text-center">
          <h6 class="category text-muted">Our work</h6>
          <h2 class="title mb-5">Some of Our Awesome Projects - 3</h2>
        </div>
      </div>
      <div class="row">
        <div class="col-lg-4">
          <card
            class="card-background"
            bodyClasses="text-left"
            style="background-image: url('img/theme/masha-rostovskaya.jpg')"
          >
            <template slot="body">
              <div
                class="icon icon-shape bg-gradient-white shadow rounded-circle mb-3"
              >
                <i class="ni ni-atom text-warning"></i>
              </div>
              <a href="javascript:;">
                <h2 class="card-title">Project Alpha</h2>
              </a>
              <h6 class="desc text-white opacity-8">Java App</h6>
              <base-button tag="a" class="text-white" type="warning" size="sm">
                Watch Later
              </base-button>
            </template>
          </card>
        </div>
        <div class="col-lg-4">
          <card
            class="card-background"
            style="background-image: url('img/theme/ali-pazani.jpg')"
          >
            <template slot="body">
              <div
                class="icon icon-shape bg-gradient-white shadow rounded-circle mb-3"
              >
                <i class="ni ni-controller text-danger"></i>
              </div>
              <a href="javascript:;">
                <h2 class="card-title">Project Beta</h2>
              </a>
              <h6 class="desc text-white opacity-8">College project</h6>
              <base-button tag="a" class="text-white" type="danger" size="sm">
                Watch Later
              </base-button>
            </template>
          </card>
        </div>
        <div class="col-lg-4">
          <card
            class="card-background"
            bodyClasses="text-right"
            style="background-image: url('img/theme/willy-dade.jpg')"
          >
            <template slot="body">
              <div
                class="icon icon-shape bg-gradient-white shadow rounded-circle mb-3"
              >
                <i class="ni ni-html5 text-success"></i>
              </div>
              <a href="javascript:;">
                <h2 class="card-title">Project Gama</h2>
              </a>
              <h6 class="desc text-white opacity-8">HTML code</h6>
              <base-button tag="a" class="text-white" type="success" size="sm">
                Watch Later
              </base-button>
            </template>
          </card>
        </div>
      </div>
    </div>
  </div>
</template>
<script></script>
<style></style>
