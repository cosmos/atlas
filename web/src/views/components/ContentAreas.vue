<template>
  <div class="section section-content-areas pt-0">
    <div class="container">
      <h2 class="mt-lg mb-5">
        <span>Content Areas</span>
      </h2>
      <h3 class="h4 text-success font-weight-bold mb-4">Tables</h3>
      <div class="mb-4 mt-4">
        <small class="text-uppercase font-weight-bold"
          >Simple With Actions</small
        >
      </div>
      <light-table></light-table>
      <div class="mb-4 mt-4">
        <small class="text-uppercase font-weight-bold"
          >Striped With Checkboxes</small
        >
      </div>
      <striped-checkbox></striped-checkbox>
      <h4><small>Shopping Cart Table</small></h4>
      <shopping-table></shopping-table>
      <div id="comments">
        <h3 class="h4 text-success font-weight-bold mb-5 mt-5">Comments</h3>
        <div class="row">
          <div class="col-md-8 mx-auto">
            <div class="card">
              <div class="card-header">
                <h5 class="h3 mb-0">Activity feed</h5>
              </div>
              <div class="card-header d-flex align-items-center">
                <div class="d-flex align-items-center">
                  <a href="javascript:;">
                    <img src="img/faces/team-1.jpg" class="avatar" />
                  </a>
                  <div class="mx-3">
                    <a
                      href="javascript:;"
                      class="text-dark font-weight-600 text-sm"
                      >John Snow</a
                    >
                    <small class="d-block text-muted">3 days ago</small>
                  </div>
                </div>
                <div class="text-right ml-auto">
                  <base-button type="primary" icon="ni ni-fat-add" size="sm">
                    Follow
                  </base-button>
                </div>
              </div>
              <div class="card-body">
                <p class="mb-4">
                  Personal profiles are the perfect way for you to grab their
                  attention and persuade recruiters to continue reading your CV
                  because you’re telling them from the off exactly why they
                  should hire you.
                </p>
                <img
                  alt="Image placeholder"
                  src="img/sections/mohamed.jpg"
                  class="img-fluid rounded"
                />
                <div class="row align-items-center my-3 pb-3 border-bottom">
                  <div class="col-sm-6">
                    <div class="icon-actions">
                      <a href="javascript:;" class="like active">
                        <i class="ni ni-like-2"></i>
                        <span class="text-muted">150</span>
                      </a>
                      <a href="javascript:;">
                        <i class="ni ni-chat-round"></i>
                        <span class="text-muted">36</span>
                      </a>
                      <a href="javascript:;">
                        <i class="ni ni-curved-next"></i>
                        <span class="text-muted">12</span>
                      </a>
                    </div>
                  </div>
                  <div class="col-sm-6 d-none d-sm-block">
                    <div
                      class="d-flex align-items-center justify-content-sm-end"
                    >
                      <div class="avatar-group">
                        <a
                          href="javascript:;"
                          class="avatar avatar-xs rounded-circle"
                          data-toggle="tooltip"
                          data-original-title="Jessica Rowland"
                        >
                          <img
                            alt="Image placeholder"
                            src="img/faces/team-1.jpg"
                            class=""
                          />
                        </a>
                        <a
                          href="javascript:;"
                          class="avatar avatar-xs rounded-circle"
                          data-toggle="tooltip"
                          data-original-title="Audrey Love"
                        >
                          <img
                            alt="Image placeholder"
                            src="img/faces/team-2.jpg"
                            class="rounded-circle"
                          />
                        </a>
                        <a
                          href="javascript:;"
                          class="avatar avatar-xs rounded-circle"
                          data-toggle="tooltip"
                          data-original-title="Michael Lewis"
                        >
                          <img
                            alt="Image placeholder"
                            src="img/faces/team-3.jpg"
                            class="rounded-circle"
                          />
                        </a>
                      </div>
                      <small class="pl-2 font-weight-bold">and 30+ more</small>
                    </div>
                  </div>
                </div>
                <!-- Comments -->
                <div class="mb-1">
                  <div class="media media-comment">
                    <img
                      alt="Image placeholder"
                      class="media-comment-avatar rounded-circle"
                      src="img/faces/team-1.jpg"
                    />
                    <div class="media-body">
                      <div class="media-comment-text">
                        <h6 class="h5 mt-0">Michael Lewis</h6>
                        <p class="text-sm lh-160">
                          You have the opportunity to play this game of life you
                          need to appreciate every moment. A lot of people don’t
                          appreciate the moment until it’s passed.
                        </p>
                        <div class="icon-actions">
                          <a href="javascript:;" class="like active">
                            <i class="ni ni-like-2"></i>
                            <span class="text-muted">3 likes</span>
                          </a>
                          <a href="javascript:;">
                            <i class="ni ni-curved-next"></i>
                            <span class="text-muted">2 shares</span>
                          </a>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="media media-comment">
                    <img
                      alt="Image placeholder"
                      class="media-comment-avatar rounded-circle"
                      src="img/faces/team-2.jpg"
                    />
                    <div class="media-body">
                      <div class="media-comment-text">
                        <h6 class="h5 mt-0">Jessica Stones</h6>
                        <p class="text-sm lh-160">
                          I always felt like I could do anything. That’s the
                          main thing people are controlled by! Thoughts- their
                          perception of themselves! They're slowed down.
                        </p>
                        <div class="icon-actions">
                          <a href="javascript:;" class="like active">
                            <i class="ni ni-like-2"></i>
                            <span class="text-muted">10 likes</span>
                          </a>
                          <a href="javascript:;">
                            <i class="ni ni-curved-next"></i>
                            <span class="text-muted">1 share</span>
                          </a>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="media align-items-center mt-5">
                    <img
                      alt="Image placeholder"
                      class="avatar avatar-lg rounded-circle mb-4"
                      src="img/faces/team-3.jpg"
                    />
                    <div class="media-body">
                      <form>
                        <textarea
                          class="form-control"
                          placeholder="Write your comment"
                          rows="1"
                        ></textarea>
                      </form>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <h3 class="h4 text-success font-weight-bold mb-5 mt-5">Chat</h3>
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
          <div class="card card-plain">
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
                  <base-button icon type="" class="text-primary"
                    ><i class="ni ni-book-bookmark"></i
                  ></base-button>
                </div>
                <div class="col-md-2 col-3 text-right">
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
  </div>
</template>
<script>
import LightTable from "./Tables/LightTable.vue";
import StripedCheckbox from "./Tables/StripedCheckbox.vue";
import ShoppingTable from "./Tables/ShoppingTable.vue";
import BaseDropdown from "@/components/BaseDropdown.vue";
export default {
  components: {
    LightTable,
    StripedCheckbox,
    ShoppingTable,
    BaseDropdown
  }
};
</script>
<style></style>
