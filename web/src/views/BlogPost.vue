<template>
  <div class="wrapper">
    <div class="page-header header-filter">
      <div
        class="page-header-image"
        style="background-image: url('img/sections/photo-1.jpg');"
      ></div>
      <div class="container">
        <div class="row">
          <div class="col-lg-10 mx-auto text-center">
            <h2 class="display-2 text-white">Check our latest post!</h2>
          </div>
        </div>
        <div class="row">
          <div class="floating-box bg-default">
            <div class="box text-center">
              <div
                class="icon icon-shape bg-primary icon-xl rounded-circle text-white"
              >
                <i class="ni ni-spaceship"></i>
              </div>
            </div>
            <blockquote class="blockquote text-center mt-4">
              <p class="mb-0 text-white">posted 13 Sept 2019</p>
              <footer class="blockquote-footer text-white">
                Someone famous in
                <cite title="Source Title">Source Title</cite>
              </footer>
            </blockquote>
            <h2 class="lead text-white p-5">
              As a result of growing greenhouse gas emissions, climate models
              predict that our planet will get significantly warmer, that
              ecosystems will be changed or destroyed, and that enormous human
              and economic costs will be incurred. These scenarios aren’t
              guaranteed, but avoiding them will be very hard. We’re trying to
              take small steps to mitigate our impact.
            </h2>
          </div>
        </div>
      </div>
    </div>
    <section class="section">
      <div class="container">
        <div class="row">
          <div class="col-md-8 ml-auto mr-auto">
            <h6 class="category">March 1, 2019</h6>
            <h3 class="title">The Castle Looks Different at Night...</h3>
            <p>
              This is the paragraph where you can write more details about your
              product. Keep you user engaged by providing meaningful
              information. Remember that by this time, the user is curious,
              otherwise he wouldn't scroll to get here. Add a button if you want
              the user to see more. We are here to make life better. <br />
              <br />
              And now I look and look around and there’s so many Kanyes I've
              been trying to figure out the bed design for the master bedroom at
              our Hidden Hills compound... and thank you for turning my personal
              jean jacket into a couture piece.
            </p>
            <br />
            <br />
            <p class="blockquote">
              “And thank you for turning my personal jean jacket into a couture
              piece.”
              <br />
              <br />
              <small>
                Kanye West, Producer.
              </small>
            </p>
            <h3 class="title">Using Video Games To Generate Clicks</h3>
            <br />
            <p>
              This is the paragraph where you can write more details about your
              product. Keep you user engaged by providing meaningful
              information. Remember that by this time, the user is curious,
              otherwise he wouldn't scroll to get here. Add a button if you want
              the user to see more. We are here to make life better.
            </p>
          </div>
        </div>
      </div>
    </section>
    <section>
      <vue-glide
        :startAt="0"
        :perView="4"
        :perTouch="1"
        :focutAt="2"
        type="carousel"
      >
        <vue-glide-slide>
          <img src="img/theme/victor-garcia.jpg" height="500" width="450" />
        </vue-glide-slide>
        <vue-glide-slide>
          <img src="img/theme/sacha-styles.jpg" height="500" width="450" />
        </vue-glide-slide>
        <vue-glide-slide>
          <img src="img/theme/doyoun-seo.jpg" height="500" width="450" />
        </vue-glide-slide>
        <vue-glide-slide>
          <img src="img/theme/ayo-ogunseinde.jpg" height="500" width="450" />
        </vue-glide-slide>
        <vue-glide-slide>
          <img src="img/theme/sofia-kuniakina.jpg" height="500" width="450" />
        </vue-glide-slide>
        <template slot="control">
          <button class="glide__arrow glide__arrow--left" data-glide-dir="<">
            <i class="ni ni-bold-left"></i>
          </button>
          <button class="glide__arrow glide__arrow--right" data-glide-dir=">">
            <i class="ni ni-bold-right"></i>
          </button>
        </template>
      </vue-glide>
    </section>
    <section class="section mt-5">
      <div class="container">
        <div class="row">
          <div class="col-md-8 mx-auto">
            <h3 class="title">Rest of the Story:</h3>
            <p>
              We are here to make life better. And now I look and look around
              and there’s so many Kanyes I've been trying to figure out the bed
              design for the master bedroom at our Hidden Hills compound... and
              thank you for turning my personal jean jacket into a couture
              piece. <br />
              I speak yell scream directly at the old guard on behalf of the
              future. daytime All respect prayers and love to Phife’s family
              Thank you for so much inspiration.
            </p>
            <br />
            <p>
              Thank you Anna for the invite thank you to the whole Vogue team
              And I love you like Kanye loves Kanye Pand Pand Panda I've been
              trying to figure out the bed design for the master bedroom at our
              Hidden Hills compound...The Pablo pop up was almost a pop up of
              influence. All respect prayers and love to Phife’s family Thank
              you for so much inspiration daytime I love this new Ferg album!
              The Life of Pablo is now available for purchase I have a dream.
              Thank you to everybody who made The Life of Pablo the number 1
              album in the world! I'm so proud of the nr #1 song in the country.
              Panda! Good music 2016!
            </p>
            <br />
            <p>
              I love this new Ferg album! The Life of Pablo is now available for
              purchase I have a dream. Thank you to everybody who made The Life
              of Pablo the number 1 album in the world! I'm so proud of the nr
              #1 song in the country. Panda! Good music 2016!
            </p>
          </div>
        </div>
      </div>
    </section>
    <section class="section section-blog-info">
      <div class="container">
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
                  <base-button type="primary" size="sm">
                    <template slot="icon">
                      <span class="btn-inner--icon icon-big"
                        ><i class="ni ni-fat-add"></i
                      ></span>
                    </template>
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
    </section>
    <!-- Blog Section -->
    <Blog></Blog>
  </div>
</template>
<script>
import { Glide, GlideSlide } from "vue-glide-js";
import "vue-glide-js/dist/vue-glide.css";
import Blog from "./sections/Blogs/Blog3.vue";

export default {
  bodyClass: "blog-post",
  components: {
    [Glide.name]: Glide,
    [GlideSlide.name]: GlideSlide,
    Blog
  }
};
</script>
<style></style>
