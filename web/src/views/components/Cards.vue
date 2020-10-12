<template>
  <div class="section section-cards bg-secondary">
    <div class="container">
      <h2 class="mb-5">
        <span>Cards</span>
      </h2>
      <h3 class="h4 text-success font-weight-bold mb-4">Blog Cards</h3>
      <div class="row">
        <div class="col-md-6 col-lg-4">
          <card gradient="danger">
            <template slot="body">
              <h6 class="category-social text-white text-uppercase">
                <i class="fa fa-fire"></i> Focus
              </h6>
              <p class="lead text-white my-0">
                Our brains are finely attuned to distraction
              </p>
            </template>
            <a href="javascript:;">
              <img class="img pattern rounded" src="img/ill/p51.svg" />
            </a>
          </card>

          <card type="primary" class="card-blog">
            <template slot="image">
              <img class="img rounded" src="img/ill/p22.png" />
            </template>
            <template slot="body">
              <h6 class="category text-white">
                <i class="ni ni-atom"></i>
                Focus
              </h6>
              <h5 class="card-title">
                <a class="text-white" href="javascript:;">Stay focused</a>
              </h5>
              <p class="card-description text-white">
                Our brains are finely attuned to distraction, so today's digital
                environment makes it especially hard to focus...
              </p>
            </template>
          </card>

          <card class="card-blog" data-header="pattern">
            <a href="javascript:;">
              <img class="img pattern rounded" src="img/ill/p2.svg" />
            </a>
            <template slot="body">
              <h6 class="card-category text-danger">
                <i class="ni ni-badge"></i> Features
              </h6>
              <h5 class="card-title">
                <a href="javascript:;">FiftyThree Files For Paper</a>
              </h5>
              <p class="card-description">
                Yesterday, as Facebook launched its news reader app Paper,
                design-focused startup FiftyThree called out Facebook...
              </p>
            </template>
            <template slot="footer">
              <div class="author">
                <img
                  src="img/faces/team-3.jpg"
                  alt="..."
                  class="avatar img-raised"
                />
                <span>Johana Doe</span>
              </div>
              <div class="stats stats-right">
                <i class="ni ni-favourite-28"></i> 10.4K ·
                <i class="ni ni-archive-2"></i> 425
              </div>
            </template>
          </card>
        </div>
        <div class="col-md-6 col-lg-4">
          <card class="text-white bg-default" data-background="pattern">
            <a href="javascript:;">
              <img class="img pattern rounded" src="img/ill/inn.svg" />
            </a>
            <template slot="body">
              <div class="content">
                <h5 class="h2 card-title text-white mb-2">Card Title</h5>
                <p class="card-description">
                  This is a wider card with supporting text below as a natural
                  lead-in to additional content.
                </p>
                <p class="card-text text-sm font-weight-bold">
                  Last updated 3 mins ago
                </p>
              </div>
            </template>
          </card>

          <card class="card-blog" data-header="skew" headerClasses="bg-yellow">
            <template slot="header">
              <div class="card-image">
                <img class="card-img-top" src="img/ill/inn.svg" />
              </div>
            </template>
            <template slot="body">
              <h6 class="card-category">
                <i class="ni ni-paper-diploma"></i>
                Business
              </h6>
              <h5 class="card-title">
                <a href="javascript:;"
                  >Springer Spends $343M To Buy Business Insider</a
                >
              </h5>
              <p class="card-description">
                German media giant Axel Springer has announced it’s acquiring
                online business news publication Business Inside...
              </p>
            </template>
          </card>

          <card class="bg-gradient-warning">
            <template slot="body">
              <h6 class="category text-white">
                <i class="ni ni-money-coins"></i> Features
              </h6>
              <h5 class="card-title">
                <a class="text-white" href="javascript:;"
                  >That’s One Way To Ditch Your Passenger</a
                >
              </h5>
              <p class="card-description text-white">
                As near as we can tell, this guy must have thought he was going
                over backwards break to bring the nose down...
              </p>
              <div class="card-footer">
                <div class="author">
                  <img
                    src="img/faces/team-3.jpg"
                    alt="..."
                    class="avatar img-raised"
                  />
                  <span>Johana Doe</span>
                </div>
                <div class="stats stats-right">
                  <i class="ni ni-favourite-28"></i> 10.4K ·
                  <i class="ni ni-archive-2"></i> 425
                </div>
              </div>
            </template>
          </card>
        </div>
        <div class="col-md-6 col-lg-4">
          <card class="card-blog bg-info">
            <template slot="image">
              <img class="img pattern rounded" src="img/ill/p7.svg" />
            </template>
            <template slot="body">
              <h4 class="display-4 text-white">Here Be Dragons</h4>
              <p class="lead text-white mt-0">
                Our brains are finely attuned to distraction, so today's digital
                environment makes it especially hard to focus...
              </p>
            </template>
          </card>

          <card class="card-blog" data-background="pattern">
            <a href="javascript:;">
              <img class="img pattern rounded" src="img/ill/p8.svg" />
            </a>
            <template slot="body">
              <h6 class="card-category text-danger">
                <i class="ni ni-badge"></i> News
              </h6>
              <h5 class="card-title">
                <a href="javascript:;">Focus on Your Employees</a>
              </h5>
              <p class="card-description">
                Don't be scared of the truth because we need to restart the
                human foundation in truth And I love you like Kanye
              </p>
            </template>
            <template slot="footer">
              <div class="author">
                <img
                  src="img/faces/team-3.jpg"
                  alt="..."
                  class="avatar img-raised"
                />
                <span>Dylan Wyatt</span>
              </div>
              <div class="stats stats-right">
                <i class="ni ni-favourite-28"></i> 10.4K ·
                <i class="ni ni-archive-2"></i> 425
              </div>
            </template>
          </card>

          <card class="card-blog">
            <template slot="image">
              <img
                class="card-img-top"
                src="img/ill/linth3.svg"
                alt="Image placeholder"
              />
            </template>
            <template slot="body">
              <h5 class="h5 card-title mb-0">Get started with Argon</h5>
              <small class="text-muted"
                >by John Snow on Oct 29th at 10:23 AM</small
              >
              <p class="card-text mt-4">
                Argon is a great free UI package based on Bootstrap 4 that
                includes the most important components and features.
              </p>
              <a href="#!" class="btn btn-link px-0">View article</a>
            </template>
          </card>
        </div>
      </div>
      <h3 class="h4 text-success font-weight-bold mb-5 mt-5">Profile Cards</h3>
      <div class="row">
        <div class="col-md-6 col-lg-4">
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
                  <a class="dropdown-item" href="javascript:;">Edit Profile</a>
                  <a class="dropdown-item" href="javascript:;">Settings</a>
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
                        <div class="badge badge-circle badge-success mr-3">
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
                        <h6 class="mb-1">Able to get good at everything</h6>
                      </div>
                    </div>
                  </li>
                </ul>
              </div>
            </template>
          </card>
        </div>
        <div class="col-md-6 col-lg-4">
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
                    src="img/faces/team-4.jpg"
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
                  Daisy Paisley<span class="font-weight-light">, 28</span>
                </h5>
                <div class="font-weight-300">
                  <i class="ni location_pin mr-2"></i>Bucharest, Romania
                </div>
              </div>
            </template>
          </card>
        </div>
        <div class="col-md-6 col-lg-4">
          <card
            class="card-profile"
            data-image="img-raised"
            footerClasses="text-center"
          >
            <div slot="headerImage">
              <a href="javascript:;">
                <img class="img" src="img/faces/alejandro-escamilla.jpg" />
              </a>
              <div class="card-title text-white">
                Austin Johnson
              </div>
            </div>
            <template slot="body">
              <h6 class="card-category text-info text-center">
                Project Manager
              </h6>
              <p class="card-description">
                Don't be scared of the truth because we need to restart the
                human foundation in truth And I love you...
              </p>
            </template>
            <template slot="footer">
              <base-button
                link
                type=""
                class="btn-instagram btn-icon-only btn-simple btn btn-default"
              >
                <span class="btn-inner--icon"
                  ><i class="fa fa-instagram"></i
                ></span>
              </base-button>
              <base-button
                link
                type=""
                class="btn-twitter btn-icon-only btn-simple btn btn-default"
              >
                <span class="btn-inner--icon"
                  ><i class="fa fa-twitter"></i
                ></span>
              </base-button>
              <base-button
                link
                type=""
                class="btn-dribbble btn-icon-only btn-simple btn btn-default"
              >
                <span class="btn-inner--icon"
                  ><i class="fa fa-dribbble"></i
                ></span>
              </base-button>
            </template>
          </card>
        </div>
        <div class="col-md-6 col-lg-4">
          <card class="card-profile" data-image="img-rounded">
            <template slot="avatar">
              <a href="javascript:;">
                <img class="img" src="img/faces/team-1.jpg" />
              </a>
            </template>
            <template slot="body">
              <h6 class="card-category text-gray">Project Manager</h6>
              <h4 class="card-title">Austin Johnson</h4>
              <p class="card-description">
                Don't be scared of the truth because we need to restart the
                human foundation in truth And I love you...
              </p>
              <a href="javascript:;" class="btn btn-info btn-round">Follow</a>
            </template>
          </card>
        </div>
        <div class="col-md-6 col-lg-4">
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
                    src="img/faces/team-3.jpg"
                  />
                </a>
              </div>
            </template>
            <template slot="body">
              <div class="text-center">
                <h5 class="h5 mt-3">
                  Johana Doe<span class="font-weight-light">, 37</span>
                </h5>
                <div class="font-weight-300">
                  <i class="ni location_pin mr-2"></i>Chicago, USA
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
        <div class="col-md-6 col-lg-4">
          <card class="card-profile" data-background="full">
            <a href="javascript:;">
              <img class="img img-raised rounded" src="img/theme/lucy.jpg" />
            </a>
            <template slot="image">
              <base-dropdown menuClasses="dropdown-menu-right">
                <div
                  slot="title"
                  class="nav-link dropdown-toggle text-white"
                  data-toggle="dropdown"
                >
                  <i class="ni ni-settings-gear-65"></i>
                </div>
                <a class="dropdown-item" href="javascript:;">Edit Profile</a>
                <a class="dropdown-item" href="javascript:;">Settings</a>
                <a class="dropdown-item" href="javascript:;">Log out</a>
              </base-dropdown>
            </template>
            <template slot="body">
              <h4 class="display-4">Jane Doe</h4>
              <p class="lead mt-0 mb-1">Team Lead</p>
              <div class="table-responsive">
                <table class="table tablesorter " id="plain-table">
                  <tbody>
                    <tr>
                      <td class="text-left pl-0">
                        <div class="badge badge-circle badge-info mr-2">
                          <i class="ni ni-hat-3"></i>
                        </div>
                        Skills
                      </td>
                      <td class="text-right">
                        Leadership
                      </td>
                    </tr>
                    <tr>
                      <td class="text-left pl-0">
                        <div class="badge badge-circle badge-info mr-2">
                          <i class="ni ni-satisfied"></i>
                        </div>
                        Hobbies
                      </td>
                      <td class="text-right">
                        Skiing, Chess
                      </td>
                    </tr>
                    <tr>
                      <td class="text-left pl-0">
                        <div class="badge badge-circle badge-info mr-2">
                          <i class="ni ni-bullet-list-67"></i>
                        </div>
                        Level
                      </td>
                      <td class="text-right">
                        • • • • •
                      </td>
                    </tr>
                  </tbody>
                </table>
              </div>
            </template>
          </card>
        </div>
      </div>
      <h3 class="h4 text-success font-weight-bold mb-5 mt-5">
        Full Background Cards
      </h3>
      <div class="row">
        <div class="col-md-6 col-lg-3">
          <card class="card-blog card-background" data-animation="zooming">
            <div
              class="full-background"
              style="background-image: url('img/theme/josh-appel.jpg')"
            ></div>
            <template slot="body">
              <div class="content-bottom">
                <h6 class="card-category text-white opacity-8">
                  New Challenges
                </h6>
                <h5 class="card-title">Touch on a trend</h5>
              </div>
            </template>
          </card>
        </div>
        <div class="col-md-6 col-lg-3">
          <card class="card-blog card-background" data-animation="zooming">
            <div
              class="full-background"
              style="background-image: url('img/theme/john-hoang.jpg')"
            ></div>
            <template slot="body">
              <div class="content-bottom">
                <h6 class="card-category text-white opacity-8">
                  New Opportunities
                </h6>
                <h5 class="card-title">Constantly growing</h5>
              </div>
            </template>
          </card>
        </div>
        <div class="col-md-6">
          <card class="card-blog card-background" data-animation="zooming">
            <div
              class="full-background"
              style="background-image: url('img/theme/kit-suman.jpg')"
            ></div>
            <template slot="body">
              <div class="content-bottom">
                <h6 class="card-category text-white opacity-8">Sales Leads</h6>
                <h5 class="card-title">Configure Blockchain Technology</h5>
              </div>
            </template>
          </card>
        </div>
      </div>
      <h3 class="h4 text-success font-weight-bold mb-5 mt-5">Pricing Cards</h3>
      <div class="row">
        <div class="col-md-6 col-lg-4">
          <card
            class="card-pricing text-center"
            type="success"
            headerClasses="bg-transparent"
            footerClasses="bg-transparent"
            style="background-image: url('img/ill/pattern_pricing1.svg"
          >
            <template slot="header">
              <h6 class="text-uppercase ls-1 text-white py-3 mb-0">
                Bravo pack
              </h6>
            </template>
            <template slot="body">
              <div class="display-2 text-white">$49</div>
              <span class=" text-white">per application</span>
              <ul class="list-unstyled my-4">
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape bg-white shadow rounded-circle text-success"
                      >
                        <i class="ni ni-book-bookmark"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm text-white"
                        >Complete documentation</span
                      >
                    </div>
                  </div>
                </li>
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape bg-white shadow rounded-circle text-success"
                      >
                        <i class="ni ni-diamond"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm text-white"
                        >Working materials in Sketch</span
                      >
                    </div>
                  </div>
                </li>
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape bg-white shadow rounded-circle text-success"
                      >
                        <i class="ni ni-chart-pie-35"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm text-white"
                        >2GB cloud storage</span
                      >
                    </div>
                  </div>
                </li>
              </ul>
              <base-button link class="text-white" type=""
                >FREE TRIAL</base-button
              >
            </template>
            <template slot="footer">
              <a href="#!" class=" text-white">Request a demo</a>
            </template>
          </card>
        </div>
        <div class="col-md-6 col-lg-4">
          <card
            class="card-pricing text-center"
            headerClasses="bg-transparent"
            footerClasses="bg-transparent"
            style="background-image: url('img/ill/pattern_pricing1.svg"
          >
            <template slot="header">
              <h6 class="text-uppercase ls-1 py-3 mb-0">Premium pack</h6>
            </template>
            <template slot="body">
              <div class="display-2">$89</div>
              <span>per month</span>
              <ul class="list-unstyled my-4">
                <li class="align-items-center">
                  <b class="text-primary">20GB</b> <span>File Storage</span>
                </li>
                <li class="align-items-center">
                  <b class="text-primary">Sketch</b>
                  <span>Working materials in</span>
                </li>
                <li class="align-items-center">
                  <b class="text-primary">Premium</b> <span>Support</span>
                </li>
              </ul>
            </template>
            <template slot="footer">
              <a href="#!" class="text-primary">Request a demo</a>
            </template>
          </card>
        </div>
        <div class="col-md-6 col-lg-4">
          <card
            class="card-pricing card-background text-center"
            headerClasses="bg-transparent"
            footerClasses="bg-transparent"
            style="background-image: url('img/ill/pattern_pricing6.svg"
          >
            <template slot="body">
              <h6 class="card-category text-danger text-uppercase">Standard</h6>
              <h1 class="card-title text-default"><small>$</small>69</h1>
              <ul>
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape icon-shape-info shadow rounded-circle"
                      >
                        <i class="ni ni-check-bold"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm text-whit">Sharing Tools</span>
                    </div>
                  </div>
                </li>
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape icon-shape-info shadow rounded-circle"
                      >
                        <i class="ni ni-check-bold"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm text-whit">Design Tools</span>
                    </div>
                  </div>
                </li>
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape icon-shape-danger shadow rounded-circle"
                      >
                        <i class="ni ni-fat-remove"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm text-whit"
                        >Private Messages</span
                      >
                    </div>
                  </div>
                </li>
                <li>
                  <div class="d-flex align-items-center">
                    <div>
                      <div
                        class="icon icon-xs icon-shape icon-shape-danger shadow rounded-circle"
                      >
                        <i class="ni ni-fat-remove"></i>
                      </div>
                    </div>
                    <div>
                      <span class="pl-2 text-sm text-whit">Private Brand</span>
                    </div>
                  </div>
                </li>
              </ul>
              <a href="#pablo" class="btn btn-danger mt-4">
                Choose plan
              </a>
            </template>
          </card>
        </div>
        <div class="col-md-6 col-lg-4">
          <card
            class="card-pricing card-background text-center"
            type="default"
            headerClasses="bg-transparent"
            footerClasses="bg-transparent"
            style="background-image: url('img/ill/pricing_bg.svg"
            data-background="image"
          >
            <template slot="body">
              <h6 class="card-category text-white text-uppercase">Premium</h6>
              <h1 class="card-title"><small class="text-white">$</small>109</h1>
              <ul>
                <li><b>100</b> Projects</li>
                <li><b>5</b> Team Members</li>
                <li><b>55</b> Personal Contacts</li>
                <li><b>5K</b> Messages</li>
              </ul>
              <a href="javascript:;" class="btn btn-white btn-sm mt-3">
                Get Started
              </a>
            </template>
          </card>
        </div>
        <div class="col-md-6 col-lg-4">
          <card
            class="card-pricing card-coin"
            footerClasses="text-center bg-transparent"
            style="background-image: url('img/ill/pattern_pricing5.svg"
          >
            <template slot="header">
              <img src="img/dropbox.png" class="img-center shadow" />
            </template>
            <template slot="body">
              <div class="row">
                <div class="col-md-12 text-center">
                  <h5 class="text-uppercase">Platinum Pack</h5>
                  <span>£ 349</span>
                  <hr class="bg-info" />
                </div>
              </div>
              <div class="row">
                <ul class="list-group">
                  <li class="list-group-item">350 messages</li>
                  <li class="list-group-item">10K emails</li>
                  <li class="list-group-item">24/7 Support</li>
                </ul>
              </div>
            </template>
            <template slot="footer">
              <base-button type="primary" class="mb-3">Get plan</base-button>
            </template>
          </card>
        </div>
        <div class="col-md-6 col-lg-4">
          <card
            class="card-pricing text-center"
            gradient="warning"
            footerClasses="text-center bg-transparent"
            headerClasses="bg-transparent"
            style="background-image: url('img/ill/pattern_pricing5.svg"
            data-icon="big-icon"
          >
            <template slot="header">
              <div class="icon text-white">
                <i class="ni ni-camera-compact"></i>
              </div>
              <p class="text-white">Premium pack</p>
            </template>
            <template slot="body">
              <div class="price d-inline">
                <span class="currency text-white">$</span>
                <span class="amount text-white">89</span>
                <span class="frequency text-white">/mo</span>
              </div>
              <ul class="list-unstyled align-items-center">
                <li>
                  <p class="text-sm text-white">10 Templates + Updates</p>
                </li>
                <li>
                  <p class="text-sm text-white">+ Adobe Sketch</p>
                </li>
              </ul>
            </template>
            <template slot="footer">
              <base-button type="white" outline class="mb-3"
                >Upgrade</base-button
              >
            </template>
          </card>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import BaseDropdown from "@/components/BaseDropdown";

export default {
  components: {
    BaseDropdown
  }
};
</script>
<style></style>
