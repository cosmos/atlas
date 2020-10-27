<template>
  <div class="wrapper">
    <div class="section-shaped my-0 skew-separator skew-mini">
      <div class="page-header page-header-small header-filter">
        <div
          class="page-header-image"
          style="background-image: url('img/sections/broofoac.jpg');"
        ></div>
        <div class="container">
          <div class="header-body text-center mb-7">
            <div class="row justify-content-center">
              <div class="col-xl-5 col-lg-6 col-md-8 px-5">
                <h1 class="text-white">Chat / Messages</h1>
                <p class="text-lead text-white">
                  Don’t be afraid to be wrong because you can’t learn anything
                  from a compliment.
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <section>
      <div class="container pt-5 mb-5 upper">
        <div class="row flex-row chat">
          <div class="col-lg-4">
            <div class="card bg-secondary">
              <form class="card-header mb-3 w-100">
                <div class="form-group w-100 mb-0">
                  <base-input
                    class="input-group-alternative mb-0"
                    placeholder="Search contact"
                    addonRightIcon="ni ni-zoom-split-in"
                  >
                  </base-input>
                </div>
              </form>
              <div class="list-group list-group-chat list-group-flush">
                <a
                  href="javascript:;"
                  class="list-group-item active bg-gradient-primary"
                >
                  <div class="media">
                    <img
                      alt="Image"
                      src="img/faces/christian.jpg"
                      class="avatar"
                    />
                    <div class="media-body ml-2">
                      <div class="justify-content-between align-items-center">
                        <h6 class="mb-0 text-white">
                          Charlie Watson
                          <span class="badge badge-success"></span>
                        </h6>
                        <div>
                          <small>Typing...</small>
                        </div>
                      </div>
                    </div>
                  </div>
                </a>
                <a href="javascript:;" class="list-group-item">
                  <div class="media">
                    <img
                      alt="Image"
                      src="img/faces/team-2.jpg"
                      class="avatar shadow"
                    />
                    <div class="media-body ml-2">
                      <div class="justify-content-between align-items-center">
                        <h6 class="mb-0">Jane Doe</h6>
                        <div>
                          <small class="text-muted">1 hour ago</small>
                        </div>
                      </div>
                      <span
                        class="text-muted text-small col-10 p-0 text-truncate d-block"
                        >Computer users and programmers</span
                      >
                    </div>
                  </div>
                </a>
                <a href="javascript:;" class="list-group-item">
                  <div class="media">
                    <img
                      alt="Image"
                      src="img/faces/team-3.jpg"
                      class="avatar shadow"
                    />
                    <div class="media-body ml-2">
                      <div class="justify-content-between align-items-center">
                        <h6 class="mb-0">Mila Skylar</h6>
                        <div>
                          <small class="text-muted">24 min ago</small>
                        </div>
                      </div>
                      <span
                        class="text-muted text-small col-10 p-0 text-truncate d-block"
                        >You can subscribe to receive weekly...</span
                      >
                    </div>
                  </div>
                </a>
                <a href="javascript:;" class="list-group-item">
                  <div class="media">
                    <img
                      alt="Image"
                      src="img/faces/team-4.jpg"
                      class="avatar shadow"
                    />
                    <div class="media-body ml-2">
                      <div class="justify-content-between align-items-center">
                        <h6 class="mb-0">Sofia Scarlett</h6>
                        <div>
                          <small class="text-muted">7 hours ago</small>
                        </div>
                      </div>
                      <span
                        class="text-muted text-small col-10 p-0 text-truncate d-block"
                        >It’s an effective resource regardless..</span
                      >
                    </div>
                  </div>
                </a>
                <a href="javascript:;" class="list-group-item">
                  <div class="media">
                    <img
                      alt="Image"
                      src="img/faces/team-5.jpg"
                      class="avatar shadow"
                    />
                    <div class="media-body ml-2">
                      <div class="justify-content-between align-items-center">
                        <h6 class="mb-0">Tom Klein</h6>
                        <div>
                          <small class="text-muted">1 day ago</small>
                        </div>
                      </div>
                      <span
                        class="text-muted text-small col-10 p-0 text-truncate d-block"
                        >Be sure to check it out if your dev pro...</span
                      >
                    </div>
                  </div>
                </a>
              </div>
            </div>
          </div>
          <div class="col-lg-8">
            <div class="card">
              <div class="card-header d-inline-block">
                <div class="row">
                  <div class="col-md-9">
                    <div class="media align-items-center">
                      <img
                        alt="Image"
                        src="img/faces/christian.jpg"
                        class="avatar shadow"
                      />
                      <div class="media-body">
                        <h6 class="mb-0 d-block">Charlie Watson</h6>
                        <span class="text-muted text-small"
                          >last seen today at 1:53am</span
                        >
                      </div>
                    </div>
                  </div>
                  <div class="col-md-1 col-3">
                    <base-dropdown menuClasses="dropdown-menu-right">
                      <div
                        slot="title"
                        class="nav-link dropdown-toggle text-primary text-sm mt-1"
                        data-toggle="dropdown"
                      >
                        <i class="ni ni-settings-gear-65"></i>
                      </div>
                      <a class="dropdown-item" href="javascript:;">
                        <i class="ni ni-single-02"></i> Profile
                      </a>
                      <a class="dropdown-item" href="javascript:;">
                        <i class="ni ni-notification-70"></i> Mute conversation
                      </a>
                      <a class="dropdown-item" href="javascript:;">
                        <i class="ni ni-key-25"></i> Block
                      </a>
                      <a class="dropdown-item" href="javascript:;">
                        <i class="ni ni-button-power"></i> Clear chat
                      </a>
                      <div class="dropdown-divider"></div>
                      <a class="dropdown-item" href="javascript:;">
                        <i class="ni ni-fat-remove"></i> Delete chat
                      </a>
                    </base-dropdown>
                  </div>
                  <div class="col-md-2 col-3 text-right">
                    <base-button icon type="" class="text-primary"
                      ><i class="ni ni-book-bookmark"></i
                    ></base-button>
                  </div>
                </div>
              </div>
              <div class="card-body">
                <div class="row justify-content-start">
                  <div class="col-auto">
                    <div class="card">
                      <div class="card-body p-2">
                        <p class="mb-1">
                          It contains a lot of good lessons about effective
                          practices
                        </p>
                        <div>
                          <small class="opacity-60"
                            ><i class="far fa-clock"></i> 3:14am</small
                          >
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="row justify-content-end text-right">
                  <div class="col-auto">
                    <div class="card bg-gradient-primary text-white">
                      <div class="card-body p-2">
                        <p class="mb-1">
                          Can it generate daily design links that include essays
                          and data visualizations ?<br />
                        </p>
                        <div>
                          <small class="opacity-60">3:30am</small>
                          <i class="ni ni-check-bold"></i>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="row mt-4">
                  <div class="col-md-12 text-center">
                    <span class="badge text-default">Wed, 3:27pm</span>
                  </div>
                </div>
                <div class="row justify-content-start">
                  <div class="col-auto">
                    <div class="card">
                      <div class="card-body p-2">
                        <p class="mb-1">
                          Yeah! Responsive Design is geared towards those trying
                          to build web apps
                        </p>
                        <div>
                          <small class="opacity-60"
                            ><i class="far fa-clock"></i> 4:31pm</small
                          >
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="row justify-content-end text-right">
                  <div class="col-auto">
                    <div class="card bg-gradient-primary text-white">
                      <div class="card-body p-2">
                        <p class="mb-1">
                          Excellent, I want it now !
                        </p>
                        <div>
                          <small class="opacity-60">4:40pm</small>
                          <i class="ni ni-check-bold"></i>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="row justify-content-start">
                  <div class="col-auto">
                    <div class="card">
                      <div class="card-body p-2">
                        <p class="mb-1">
                          You can easily get it; The content here is all free
                        </p>
                        <div>
                          <small class="opacity-60"
                            ><i class="far fa-clock"></i> 4:42pm</small
                          >
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="row justify-content-end text-right">
                  <div class="col-auto">
                    <div class="card bg-gradient-primary text-white">
                      <div class="card-body p-2">
                        <p class="mb-1">
                          Awesome, blog is important source material for anyone
                          who creates apps? <br />
                          beacuse these blogs offer a lot of information about
                          website development.
                        </p>
                        <div>
                          <small class="opacity-60">4:46pm</small>
                          <i class="ni ni-check-bold"></i>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="row justify-content-start">
                  <div class="col-5">
                    <div class="card">
                      <div class="card-body p-2">
                        <div class="col-12 p-0">
                          <img
                            src="img/theme/img-1-1200x1000.jpg"
                            alt="Rounded image"
                            class="img-fluid rounded mb-1"
                          />
                        </div>
                        <div>
                          <small class="opacity-60"
                            ><i class="far fa-clock"></i> 4:47pm</small
                          >
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="row justify-content-end text-right">
                  <div class="col-auto">
                    <div class="card bg-gradient-primary text-white">
                      <div class="card-body p-2">
                        <p class="mb-0">
                          At the end of the day … the native dev apps is where
                          users are
                        </p>
                        <div>
                          <small class="opacity-60">4:47pm</small>
                          <i class="ni ni-check-bold"></i>
                        </div>
                      </div>
                    </div>
                  </div>
                </div>
                <div class="row justify-content-start">
                  <div class="col-auto">
                    <div class="card ">
                      <div class="card-body p-2">
                        <div class="spinner">
                          <div class="bounce1"></div>
                          <div class="bounce2"></div>
                          <div class="bounce3"></div>
                        </div>
                        <p class="d-inline-block mr-2 mb-1 mt-1">
                          Typing...
                        </p>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
              <div class="card-footer d-block">
                <div class="form-group">
                  <base-input
                    placeholder="Your message"
                    addonRightIcon="ni ni-send"
                  >
                  </base-input>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>
<script>
import BaseDropdown from "@/components/BaseDropdown.vue";

export default {
  components: {
    BaseDropdown
  },
  bodyClass: "chat-page"
};
</script>
<style></style>
