<template>
  <div>
    <div class="section section-prefooter">
      <div class="container">
        <h4 class="h4 text-success font-weight-bold mb-4">Pre-footer Areas</h4>
      </div>
      <div class="container-fluid">
        <!--     *********    SIMPLE SOCIAL LINE     *********      -->
        <div class="social-line social-line-blue text-center">
          <div class="container">
            <div class="row justify-content-center">
              <div class="col-md-12">
                <h4 class="title mb-4">Thank you for your support!</h4>
              </div>
              <base-button class="btn-icon btn-twitter btn-round" type="">
                <i class="fa fa-twitter"></i>
                Twitter · 2.5k
              </base-button>
              <base-button class="btn-icon btn-facebook btn-round" type="">
                <i class="fa fa-facebook"></i>
                Facebook · 3.2k
              </base-button>
              <base-button class="btn-icon btn-tumblr btn-round" type="">
                <i class="fa fa-tumblr"></i>
                Tumblr · 1.2k
              </base-button>
              <base-button class="btn-icon btn-dribbble btn-round" type="">
                <i class="fa fa-dribbble"></i>
                Dribbble · 1.8k
              </base-button>
            </div>
          </div>
        </div>
        <br />
        <br />
        <!--     *********   SIMPLE SOCIAL LINE     *********      -->
        <div class="social-line social-line-big-icons">
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
        <br />
        <br />
        <!--     *********    SIMPLE BLACK SOCIAL LINE     *********      -->
        <div class="social-line social-line-big-icons bg-default">
          <div class="container">
            <div class="row">
              <div class="col-md-2">
                <base-button simple>
                  <i class="fa fa-twitter text-white"></i>
                </base-button>
              </div>
              <div class="col-md-2">
                <base-button simple>
                  <i class="fa fa-facebook text-white"></i>
                </base-button>
              </div>
              <div class="col-md-2">
                <base-button simple>
                  <i class="fa fa-google-plus text-white"></i>
                </base-button>
              </div>
              <div class="col-md-2">
                <base-button simple>
                  <i class="fa fa-dribbble text-white"></i>
                </base-button>
              </div>
              <div class="col-md-2">
                <base-button simple>
                  <i class="fa fa-youtube text-white"></i>
                </base-button>
              </div>
              <div class="col-md-2">
                <base-button simple>
                  <i class="fa fa-instagram text-white"></i>
                </base-button>
              </div>
            </div>
          </div>
        </div>
        <br />
        <br />
        <!--     *********    SIMPLE SUBSCRIBE LINE     *********      -->
        <div class="subscribe-line subscribe-line-white">
          <div class="container">
            <div class="row">
              <div class="col-lg-5 mr-auto">
                <h4 class="title">Get Tips &amp; Tricks every Week!</h4>
                <p class="description">
                  Join our newsletter and get news in your inbox every week! We
                  hate spam too, so no worries about this.
                </p>
              </div>
              <div
                class="col-lg-6 d-flex justify-content-center flex-column ml-auto"
              >
                <form method="" action="">
                  <div class="row">
                    <div class="col-sm-8">
                      <base-input
                        placeholder="Your Email..."
                        addonLeftIcon="ni ni-email-83"
                      ></base-input>
                    </div>
                    <div class="col-sm-4">
                      <base-button type="success" class="btn-round btn-block"
                        >SUBSCRIBE</base-button
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
    <div class="section section-footer">
      <div class="container">
        <h4 class="h4 text-success font-weight-bold mb-4">Footer Areas</h4>
      </div>
      <div class="container-fluid">
        <div id="footer-areas">
          <div class="section section-shaped no-tilt">
            <div class="shape shape-style-1 shape-default">
              <span></span>
              <span></span>
              <span></span>
              <span></span>
              <span></span>
              <span></span>
            </div>
            <div class="subscribe-line">
              <div class="container">
                <div class="row">
                  <div class="col-lg-6 col-sm-10 ml-auto mr-auto">
                    <div class="text-center">
                      <h4 class="title text-white">
                        Subscribe to our Newsletter
                      </h4>
                      <p class="description text-white">
                        Join our newsletter and get news in your inbox every
                        week! We hate spam too, so no worries about this.
                      </p>
                    </div>
                    <form class="mt-4" method="" action="">
                      <div class="row">
                        <div class="col-sm-8">
                          <base-input
                            placeholder="Search"
                            addonLeftIcon="ni ni-email-83"
                          ></base-input>
                        </div>
                        <div class="col-sm-4">
                          <base-button type="info" class="btn-block"
                            >SUBSCRIBE</base-button
                          >
                        </div>
                      </div>
                    </form>
                  </div>
                </div>
              </div>
            </div>
            <!--     *********    SIMPLE FOOTER WITH SOCIAL AND BRAND     *********      -->
            <footer class="footer bg-transparent">
              <div class="container">
                <div class="row">
                  <div class="col-md-3">
                    <base-dropdown>
                      <div
                        slot="title"
                        class="nav-link dropdown-toggle btn btn-link text-white"
                        data-toggle="dropdown"
                      >
                        <i class="ni ni-pin-3"></i>
                        <span>United States</span>
                      </div>
                      <a href="javascript:;" class="dropdown-item">
                        <span>Australia</span>
                      </a>
                      <a href="javascript:;" class="dropdown-item">
                        <span>Belgium</span>
                      </a>
                      <a href="javascript:;" class="dropdown-item">
                        <span>Canada</span>
                      </a>
                      <a href="javascript:;" class="dropdown-item">
                        <span>France</span>
                      </a>
                      <a href="javascript:;" class="dropdown-item">
                        <span>Italy</span>
                      </a>
                      <a href="javascript:;" class="dropdown-item">
                        <span>Portugal</span>
                      </a>
                      <a href="javascript:;" class="dropdown-item">
                        <span>Spain</span>
                      </a>
                      <a href="javascript:;" class="dropdown-item">
                        <span>United States</span>
                      </a>
                    </base-dropdown>
                    <br />
                    <base-dropdown>
                      <div
                        slot="title"
                        class="nav-link dropdown-toggle btn btn-link text-white"
                        data-toggle="dropdown"
                      >
                        <i class="ni ni-chat-round"></i>
                        <span>English</span>
                      </div>
                      <a href="javascript:;" class="dropdown-item">
                        <span>Deutch</span>
                      </a>
                      <a href="javascript:;" class="dropdown-item">
                        <span>English</span>
                      </a>
                      <a href="javascript:;" class="dropdown-item">
                        <span>Español</span>
                      </a>
                      <a href="javascript:;" class="dropdown-item">
                        <span>Français</span>
                      </a>
                      <a href="javascript:;" class="dropdown-item">
                        <span>Italiano</span>
                      </a>
                    </base-dropdown>
                    <li class="copyrights">© Argon Design System</li>
                  </div>
                  <div class="col-md-2 col-6">
                    <div class="column">
                      <h4 class="mt-3">Products</h4>
                      <ul>
                        <li>
                          <a href="javascript:;"><span>Payments</span></a>
                        </li>
                        <li>
                          <a href="javascript:;"><span>Billing</span></a>
                        </li>
                        <li>
                          <a href="javascript:;"><span>Connect</span></a>
                        </li>
                        <li>
                          <a href="javascript:;"><span>Sigma</span></a>
                        </li>
                        <li>
                          <a href="javascript:;"><span>Issuing</span></a>
                        </li>
                        <li>
                          <a href="javascript:;"><span>Terminal</span></a>
                        </li>
                      </ul>
                    </div>
                  </div>
                  <div class="col-md-2 col-6">
                    <div class="column">
                      <h4 class="mt-3">Resources</h4>
                      <ul>
                        <li class="nav-item">
                          <a href="javascript:void(0)">
                            Contact Us
                          </a>
                        </li>
                        <li class="nav-item">
                          <a href="javascript:void(0)">
                            About Us
                          </a>
                        </li>
                        <li class="nav-item">
                          <a href="javascript:void(0)">
                            Blog
                          </a>
                        </li>
                        <li class="nav-item">
                          <a href="javascript:void(0)">
                            License
                          </a>
                        </li>
                      </ul>
                    </div>
                  </div>
                  <div class="col-md-2 col-6">
                    <div class="column">
                      <h4 class="mt-3">Company</h4>
                      <ul>
                        <li class="nav-item">
                          <a href="javascript:;">
                            Support
                          </a>
                        </li>
                        <li class="nav-item">
                          <a href="javascript:;">
                            Jobs
                          </a>
                        </li>
                        <li class="nav-item">
                          <a href="javascript:;">
                            Privacy
                          </a>
                        </li>
                        <li class="nav-item">
                          <a href="javascript:;">
                            Guides
                          </a>
                        </li>
                        <li class="nav-item">
                          <a href="javascript:;">
                            Pricing
                          </a>
                        </li>
                      </ul>
                    </div>
                  </div>
                  <div class="col-md-3 col-6">
                    <div class="column">
                      <h4 class="mt-3">Follow us:</h4>
                      <div class="btn-wrapper profile text-left">
                        <base-button
                          tag="a"
                          size="sm"
                          class="btn-twitter"
                          href="javascript:void(0);"
                        >
                          <i class="fa fa-twitter"></i>
                        </base-button>
                        <base-button
                          tag="a"
                          size="sm"
                          class="btn-facebook"
                          href="javascript:void(0);"
                        >
                          <i class="fa fa-facebook"></i>
                        </base-button>
                        <base-button
                          tag="a"
                          size="sm"
                          class="btn-slack"
                          href="javascript:void(0);"
                        >
                          <i class="fa fa-slack"></i>
                        </base-button>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </footer>
          </div>
          <!--     *********    SIMPLE FOOTER     *********      -->
          <footer class="footer footer-simple">
            <div class="container">
              <nav>
                <ul>
                  <li>
                    <a href="javascript:void(0)" class="nav-link">
                      Creative Tim
                    </a>
                  </li>
                  <li>
                    <a href="javascript:void(0)" class="nav-link">
                      About Us
                    </a>
                  </li>
                  <li>
                    <a href="javascript:void(0)" class="nav-link">
                      Blog
                    </a>
                  </li>
                  <li>
                    <a
                      href="javascript:void(0)"
                      target="_blank"
                      class="nav-link"
                    >
                      License
                    </a>
                  </li>
                </ul>
              </nav>
              <div class="copyright text-right">
                © 2020, Proudly Coded by
                <a
                  href="javascript:void(0)"
                  target="_blank"
                  class="copyright-link"
                  >Creative Tim</a
                >.
              </div>
            </div>
          </footer>
          <!--     *********    BLACK SIMPLE FOOTER WITH SOCIAL AND BRAND     *********      -->
          <footer class="footer footer-simple bg-gradient-default">
            <div class="container">
              <div class="row">
                <div class="col-md-3">
                  <a class="footer-brand" href="javascript:;">Argon Design</a>
                </div>
                <div class="col-md-6">
                  <div class="column">
                    <ul class="inline-menu">
                      <li>
                        <a class="nav-link" href="javascript:;">
                          Blog
                        </a>
                      </li>
                      <li>
                        <a class="nav-link" href="javascript:;">
                          Presentation
                        </a>
                      </li>
                      <li>
                        <a class="nav-link" href="javascript:;">
                          Discover
                        </a>
                      </li>
                      <li>
                        <a class="nav-link" href="javascript:;">
                          Payment
                        </a>
                      </li>
                      <li>
                        <a class="nav-link" href="javascript:;">
                          Contact Us
                        </a>
                      </li>
                    </ul>
                  </div>
                </div>
                <div class="col-md-3">
                  <ul class="social-buttons pull-right">
                    <li>
                      <base-button
                        tag="a"
                        size="sm"
                        class="btn-twitter mr-2"
                        href="javascript:void(0);"
                      >
                        <i class="fa fa-twitter"></i>
                      </base-button>
                    </li>
                    <li>
                      <base-button
                        tag="a"
                        size="sm"
                        class="btn-facebook mr-2"
                        href="javascript:void(0);"
                      >
                        <i class="fa fa-facebook"></i>
                      </base-button>
                    </li>
                    <li>
                      <base-button
                        tag="a"
                        size="sm"
                        class="btn-instagram mr-2"
                        href="javascript:void(0);"
                      >
                        <i class="fa fa-instagram"></i>
                      </base-button>
                    </li>
                  </ul>
                </div>
              </div>
            </div>
          </footer>
          <!--     *********    BIG FOOTER     *********      -->
          <footer class="footer footer-big">
            <div class="container">
              <div class="content">
                <div class="row">
                  <div class="col-md-4">
                    <div class="column">
                      <img src="img/brand/blue.png" class="logo" alt="" />
                    </div>
                  </div>
                  <div class="col-md-2 col-6">
                    <div class="column">
                      <h4 class="mt-3">Company</h4>
                      <ul>
                        <li>
                          <a href="javascript:;">
                            About Us
                          </a>
                        </li>
                        <li>
                          <a href="javascript:;">
                            Careers
                          </a>
                        </li>
                        <li>
                          <a href="javascript:;">
                            Press
                          </a>
                        </li>
                      </ul>
                    </div>
                  </div>
                  <div class="col-md-2 col-6">
                    <div class="column">
                      <h4 class="mt-3">Product</h4>
                      <ul>
                        <li>
                          <a href="javascript:;">
                            How it works
                          </a>
                        </li>
                        <li>
                          <a href="javascript:;">
                            Pricing
                          </a>
                        </li>
                      </ul>
                    </div>
                  </div>
                  <div class="col-md-2 col-6">
                    <div class="column">
                      <h4 class="mt-3">Resources</h4>
                      <ul>
                        <li>
                          <a href="javascript:;">
                            Blog
                          </a>
                        </li>
                        <li>
                          <a href="javascript:;">
                            Case study
                          </a>
                        </li>
                        <li>
                          <a href="javascript:;">
                            Connect
                          </a>
                        </li>
                        <li>
                          <a href="javascript:;">
                            Help Center
                          </a>
                        </li>
                      </ul>
                    </div>
                  </div>
                  <div class="col-md-2 col-6">
                    <div class="column">
                      <h4 class="mt-3">Contact</h4>
                      <ul>
                        <li>
                          <a href="javascript:;">
                            Book a demo
                          </a>
                        </li>
                        <li>
                          <a href="javascript:;">
                            Mail us
                          </a>
                        </li>
                      </ul>
                      <div class="btn-wrapper profile text-left mt-3">
                        <base-button
                          tag="a"
                          size="sm"
                          class="btn-link"
                          type=""
                          href="javascript:void(0);"
                        >
                          <i class="fa fa-twitter"></i>
                        </base-button>
                        <base-button
                          tag="a"
                          size="sm"
                          class="btn-link"
                          type=""
                          href="javascript:void(0);"
                        >
                          <i class="fa fa-facebook"></i>
                        </base-button>
                        <base-button
                          tag="a"
                          size="sm"
                          class="btn-link"
                          type=""
                          href="javascript:void(0);"
                        >
                          <i class="fa fa-slack"></i>
                        </base-button>
                      </div>
                    </div>
                  </div>
                </div>
                <hr />
                <div class="row">
                  <div class="col-md-9">
                    <div class="column">
                      <nav>
                        <ul>
                          <li class="d-inline-block">
                            <a href="javascript:void(0)" class="nav-link">
                              @Argon
                            </a>
                          </li>
                          <li class="d-inline-block">
                            <a href="javascript:void(0)" class="nav-link">
                              Terms &amp; Conditions
                            </a>
                          </li>
                          <li class="d-inline-block">
                            <a href="javascript:void(0)" class="nav-link">
                              Privacy Policy
                            </a>
                          </li>
                          <li class="d-inline-block">
                            <a
                              href="javascript:void(0)"
                              target="_blank"
                              class="nav-link"
                            >
                              Legal Notice
                            </a>
                          </li>
                        </ul>
                      </nav>
                    </div>
                  </div>
                  <div class="col-md-3">
                    <div class="dropdown">
                      Language:
                      <base-dropdown>
                        <div
                          slot="title"
                          class="nav-link dropdown-toggle btn btn-link text-primary"
                          data-toggle="dropdown"
                        >
                          <i class="ni ni-chat-round"></i>
                          <span>English</span>
                        </div>
                        <a href="javascript:;" class="dropdown-item">
                          <span>Deutch</span>
                        </a>
                        <a href="javascript:;" class="dropdown-item">
                          <span>English</span>
                        </a>
                        <a href="javascript:;" class="dropdown-item">
                          <span>Español</span>
                        </a>
                        <a href="javascript:;" class="dropdown-item">
                          <span>Français</span>
                        </a>
                        <a href="javascript:;" class="dropdown-item">
                          <span>Italiano</span>
                        </a>
                      </base-dropdown>
                    </div>
                  </div>
                </div>
                <div class="row">
                  <div class="col-md-12">
                    <div class="footer-description">
                      Mastercard is a registered trademark and the circles
                      design is a trademark of Mastercard International
                      Incorporated. The Spendesk Prepaid Card/This card is
                      issued by Transact Payments Limited pursuant to licence by
                      Mastercard International Incorporated. The Spendesk
                      accounts are issued by Transact Payments Limited. Transact
                      Payments Limited is authorised and regulated by the
                      Gibraltar Financial Services Commission.
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </footer>
          <!--     *********    BIG FOOTER     *********      -->
          <footer class="footer footer-big">
            <div class="container">
              <div class="content">
                <div class="row">
                  <div class="col-md-2 col-6">
                    <div class="column">
                      <h4 class="mt-3">About Us</h4>
                      <ul>
                        <li class="nav-item">
                          <a href="javascript:void(0)">
                            Contact Us
                          </a>
                        </li>
                        <li class="nav-item">
                          <a href="javascript:void(0)">
                            About Us
                          </a>
                        </li>
                        <li class="nav-item">
                          <a href="javascript:void(0)">
                            Blog
                          </a>
                        </li>
                        <li class="nav-item">
                          <a href="javascript:void(0)">
                            License
                          </a>
                        </li>
                      </ul>
                    </div>
                  </div>
                  <div class="col-md-2 col-6">
                    <div class="column">
                      <h4 class="mt-3">Market</h4>
                      <ul>
                        <li>
                          <a href="javascript:;">
                            Sales FAQ
                          </a>
                        </li>
                        <li>
                          <a href="javascript:;">
                            How to Register
                          </a>
                        </li>
                        <li>
                          <a href="javascript:;">
                            Sell Goods
                          </a>
                        </li>
                        <li>
                          <a href="javascript:;">
                            Receive Payment
                          </a>
                        </li>
                        <li>
                          <a href="javascript:;">
                            Transactions Issues
                          </a>
                        </li>
                      </ul>
                    </div>
                  </div>
                  <div class="col-md-4 col-6">
                    <div class="column">
                      <h4 class="mt-3">Social Feed</h4>
                      <div class="social-feed">
                        <div class="feed-line">
                          <i class="fa fa-instagram"></i>
                          <p>
                            How to handle ethical disagreements with your
                            clients.
                          </p>
                        </div>
                        <div class="feed-line">
                          <i class="fa fa-twitter"></i>
                          <p>
                            The tangible benefits of designing at 1x pixel
                            density.
                          </p>
                        </div>
                        <div class="feed-line">
                          <i class="fa fa-facebook-square"></i>
                          <p>
                            A collection of 25 stunning sites that you can use
                            for inspiration.
                          </p>
                        </div>
                      </div>
                    </div>
                  </div>
                  <div class="col-md-4 col-6">
                    <div class="column">
                      <h4 class="mt-3">Numbers Don't Lie</h4>
                      <h5>
                        14.521 <small class="text-muted">Freelancers</small>
                      </h5>
                      <h5>
                        1.423.183 <small class="text-muted">Transactions</small>
                      </h5>
                      <h4 class="mt-4 mb-4">Follow Us</h4>
                      <ul class="social-buttons">
                        <li>
                          <base-button
                            tag="a"
                            class="btn-twitter"
                            type=""
                            href="javascript:void(0);"
                          >
                            <i class="fa fa-twitter"></i>
                          </base-button>
                        </li>
                        <li>
                          <base-button
                            tag="a"
                            class="btn-facebook"
                            type=""
                            href="javascript:void(0);"
                          >
                            <i class="fa fa-facebook"></i>
                          </base-button>
                        </li>
                        <li>
                          <base-button
                            tag="a"
                            class="btn-dribbble"
                            type=""
                            href="javascript:void(0);"
                          >
                            <i class="fa fa-dribbble"></i>
                          </base-button>
                        </li>
                      </ul>
                    </div>
                  </div>
                </div>
              </div>
              <hr />
              <div class="copyright d-flex flex-row-reverse">
                Copyright © 2020 Creative Tim All Rights Reserved.
              </div>
            </div>
          </footer>
          <!--     *********    BIG BLACK FOOTER V2     *********      -->
          <footer class="footer footer-big bg-gradient-default">
            <div class="container">
              <div class="content">
                <div class="row mb-5">
                  <div class="column mx-auto">
                    <img src="img/brand/blue.png" class="logo logo-sm" alt="" />
                  </div>
                </div>
                <div class="row">
                  <div class="col-md-2 col-6">
                    <div class="column">
                      <h4 class="mb-4">About Us</h4>
                      <ul class="links-vertical">
                        <li>
                          <a href="javascript:;">
                            Our mission
                          </a>
                        </li>
                        <li>
                          <a href="javascript:;">
                            Media Kit
                          </a>
                        </li>
                        <li>
                          <a href="javascript:;">
                            Contact
                          </a>
                        </li>
                      </ul>
                    </div>
                  </div>
                  <div class="col-md-2 col-6">
                    <div class="column">
                      <h4 class="mb-4">Social</h4>
                      <ul class="links-vertical">
                        <li>
                          <a href="javascript:;">
                            Community
                          </a>
                        </li>
                        <li>
                          <a href="javascript:;">
                            Dribbble
                          </a>
                        </li>
                        <li>
                          <a href="javascript:;">
                            Twitter
                          </a>
                        </li>
                        <li>
                          <a href="javascript:;">
                            Facebook
                          </a>
                        </li>
                      </ul>
                    </div>
                  </div>
                  <div class="col-md-8">
                    <div class="column">
                      <h4 class="mb-4">Newsletter</h4>
                      <ul class="links-vertical">
                        <p>
                          We handpicked the most interesting content for you. No
                          spam. It takes one click to unsubscribe. Give it a
                          try!
                        </p>
                        <form class="form form-newsletter" method="" action="">
                          <base-input placeholder="Your Email"></base-input>

                          <base-button type="primary">SUBSCRIBE</base-button>
                        </form>
                      </ul>
                    </div>
                  </div>
                </div>
              </div>
              <hr class="bg-white opacity-3" />
              <div class="column">
                <ul class="d-flex justify-content-center">
                  <li>
                    <a href="javascript:;" class="nav-link">
                      Terms &amp; conditions
                    </a>
                  </li>
                  <li>
                    <a href="javascript:;" class="nav-link">
                      Privacy policy
                    </a>
                  </li>
                  <li>
                    <a href="javascript:;" class="nav-link">
                      Legal notice
                    </a>
                  </li>
                </ul>
              </div>
            </div>
          </footer>
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
