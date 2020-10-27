<template>
  <div>
    <div class="row">
      <div class="col-md-4">
        <base-button
          block
          type="primary"
          class=" mb-3"
          @click="modals.modal1 = true"
        >
          Default
        </base-button>
        <modal :show.sync="modals.modal1">
          <h6 slot="header" class="modal-title" id="modal-title-default">
            Type your modal title
          </h6>

          <p>
            Far far away, behind the word mountains, far from the countries
            Vokalia and Consonantia, there live the blind texts. Separated they
            live in Bookmarksgrove right at the coast of the Semantics, a large
            language ocean.
          </p>
          <p>
            A small river named Duden flows by their place and supplies it with
            the necessary regelialia. It is a paradisematic country, in which
            roasted parts of sentences fly into your mouth.
          </p>

          <template slot="footer">
            <base-button type="primary">Save changes</base-button>
            <base-button
              type="link"
              class="ml-auto"
              @click="modals.modal1 = false"
              >Close
            </base-button>
          </template>
        </modal>
      </div>
      <div class="col-md-4">
        <base-button
          block
          type="warning"
          class=" mb-3"
          @click="modals.modal2 = true"
        >
          Notification
        </base-button>

        <modal
          :show.sync="modals.modal2"
          gradient="danger"
          modal-classes="modal-danger modal-dialog-centered"
        >
          <h6 slot="header" class="modal-title" id="modal-title-notification">
            Your attention is required
          </h6>

          <div class="py-3 text-center">
            <i class="ni ni-bell-55 ni-3x"></i>
            <h4 class="heading mt-4">You should read this!</h4>
            <p>
              A small river named Duden flows by their place and supplies it
              with the necessary regelialia.
            </p>
          </div>

          <template slot="footer">
            <base-button type="white">Ok, Got it</base-button>
            <base-button
              type="link"
              text-color="white"
              class="ml-auto"
              @click="modals.modal2 = false"
            >
              Close
            </base-button>
          </template>
        </modal>
      </div>
      <div class="col-md-4">
        <base-button
          block
          type="default"
          class=" mb-3"
          @click="modals.modal3 = true"
        >
          Form
        </base-button>

        <modal
          :show.sync="modals.modal3"
          body-classes="p-0"
          modal-classes="modal-dialog-centered modal-sm"
        >
          <card
            type="secondary"
            shadow
            header-classes="bg-white pb-5"
            body-classes="px-lg-5 py-lg-5"
            class="border-0 mb-0"
          >
            <template slot="header">
              <div class="text-muted text-center mb-3">
                <small>Sign in with</small>
              </div>
              <div class="btn-wrapper text-center">
                <base-button type="neutral">
                  <img slot="icon" src="img/icons/common/github.svg" />
                  Github
                </base-button>

                <base-button type="neutral">
                  <img slot="icon" src="img/icons/common/google.svg" />
                  Google
                </base-button>
              </div>
            </template>
            <template slot="body">
              <div class="text-center text-muted mb-4">
                <small>Or sign in with credentials</small>
              </div>
              <form role="form">
                <base-input
                  alternative
                  class="mb-3"
                  placeholder="Email"
                  addon-left-icon="ni ni-email-83"
                >
                </base-input>
                <base-input
                  alternative
                  type="password"
                  placeholder="Password"
                  addon-left-icon="ni ni-lock-circle-open"
                >
                </base-input>
                <base-checkbox>
                  Remember me
                </base-checkbox>
                <div class="text-center">
                  <base-button type="primary" class="my-4">Sign In</base-button>
                </div>
              </form>
            </template>
          </card>
        </modal>
      </div>
      <div class="col-md-4">
        <base-button
          block
          type="danger"
          class=" mb-3"
          @click="modals.modal4 = true"
        >
          Long Modal
        </base-button>

        <modal :show.sync="modals.modal4" modal-classes="modal-dialog-centered">
          <h5 slot="header" class="modal-title" id="modal-title-default">
            Modal title
          </h5>

          I always felt like I could do anything. That’s the main thing people
          are controlled by! Thoughts- their perception of themselves! They're
          slowed down by their perception of themselves. If you're taught you
          can’t do anything, you won’t do anything. I was taught I could do
          everything.
          <br /><br />
          As we live, our hearts turn colder. Cause pain is what we go through
          as we become older. We get insulted by others, lose trust for those
          others. We get back stabbed by friends. It becomes harder for us to
          give others a hand. We get our heart broken by people we love, even
          that we give them all we have. Then we lose family over time. What
          else could rust the heart more over time? Blackgold.
          <br /><br />
          We’re not always in the position that we want to be at. We’re
          constantly growing. We’re constantly making mistakes. We’re constantly
          trying to express ourselves and actualize our dreams. If you have the
          opportunity to play this game of life you need to appreciate every
          moment. A lot of people don’t appreciate the moment until it’s passed.
          <br /><br />
          There’s nothing I really wanted to do in life that I wasn’t able to
          get good at. That’s my skill. I’m not really specifically talented at
          anything except for the ability to learn. That’s what I do. That’s
          what I’m here for. Don’t be afraid to be wrong because you can’t learn
          anything from a compliment.
          <br /><br />
          It really matters and then like it really doesn’t matter. What matters
          is the people who are sparked by it. And the people who are like
          offended by it, it doesn’t matter. Because it's about motivating the
          doers. Because I’m here to follow my dreams and inspire other people
          to follow their dreams, too.
          <br /><br />
          The time is now for it to be okay to be great. People in this world
          shun people for being great. For being a bright color. For standing
          out. But the time is now to be okay to be the greatest you. Would you
          believe in what you believe in, if you were the only one who believed
          it?
        </modal>
      </div>
      <div class="col-md-4">
        <base-button
          block
          type="success"
          class=" mb-3"
          @click="modals.modal5 = true"
        >
          Message Modal
        </base-button>

        <modal
          :show.sync="modals.modal5"
          modal-classes="modal-dialog-centered"
          body-classes="bg-secondary"
          footer-classes="bg-secondary"
        >
          <h5 slot="header" class="modal-title" id="modal-title-default">
            New message to CT
          </h5>

          <form>
            <div class="form-group">
              <label for="recipient-name" class="col-form-label"
                >Recipient:</label
              >
              <base-input alternative id="recipient-name"></base-input>
            </div>
            <div class="form-group">
              <label for="message-text" class="col-form-label">Message:</label>
              <textarea
                class="form-control form-control-alternative"
                id="message-text"
              ></textarea>
            </div>
          </form>

          <template slot="footer">
            <base-button
              type="secondary"
              class="ml-auto"
              @click="modals.modal5 = false"
              >Close
            </base-button>
            <base-button type="primary">Send message</base-button>
          </template>
        </modal>
      </div>
      <div class="col-md-4">
        <base-button
          block
          type="info"
          class=" mb-3"
          @click="modals.modal6 = true"
        >
          Signup Modal
        </base-button>

        <modal
          :show.sync="modals.modal6"
          body-classes="p-0"
          modal-classes="modal-dialog-centered modal-sm"
        >
          <card
            type="secondary"
            shadow
            header-classes="bg-white pb-5"
            body-classes="px-lg-5 py-lg-5"
            class="border-0 mb-0"
          >
            <template slot="header">
              <div class="text-muted text-center mb-3">
                <small>Sign Up with</small>
              </div>
              <div class="btn-wrapper text-center">
                <base-button type="neutral">
                  <img slot="icon" src="img/icons/common/google.svg" />
                  Google
                </base-button>
                <base-button type="neutral">
                  <img slot="icon" src="img/icons/common/github.svg" />
                  Github
                </base-button>
              </div>
            </template>
            <template slot="body">
              <div class="text-center text-muted mb-4">
                <small>Or sign up here:</small>
              </div>
              <form role="form">
                <base-input
                  alternative
                  class="mb-3"
                  placeholder="Email"
                  addon-left-icon="ni ni-email-83"
                >
                </base-input>
                <base-input
                  alternative
                  type="password"
                  placeholder="Password"
                  addon-left-icon="ni ni-lock-circle-open"
                >
                </base-input>
                <base-input
                  alternative
                  type="password"
                  placeholder="Confirm Password"
                  addon-left-icon="ni ni-lock-circle-open"
                >
                </base-input>
                <div class="text-center">
                  <base-button type="primary" class="my-4">Sign up</base-button>
                </div>
              </form>
            </template>
          </card>
        </modal>
      </div>
    </div>
    <h5 class="h5 text-default font-weight-bold mb-5 mt-5">Sizes</h5>
    <base-button type="primary" class=" mb-3" @click="modals.modal7 = true">
      Extra Large Modal
    </base-button>

    <modal
      :show.sync="modals.modal7"
      modal-classes="modal-dialog-centered modal-xl"
    >
      <h5 slot="header" class="modal-title" id="modal-title-default">
        Modal title
      </h5>

      Society has put up so many boundaries, so many limitations on what’s right
      and wrong that it’s almost impossible to get a pure thought out. It’s like
      a little kid, a little boy, looking at colors, and no one told him what
      colors are good, before somebody tells you you shouldn’t like pink because
      that’s for girls, or you’d instantly become a gay two-year-old. Why would
      anyone pick blue over pink? Pink is obviously a better color. Everyone’s
      born confident, and everything’s taken away from you
      <br />
      <br />

      As we live, our hearts turn colder. Cause pain is what we go through as we
      become older. We get insulted by others, lose trust for those others. We
      get back stabbed by friends. It becomes harder for us to give others a
      hand. We get our heart broken by people we love, even that we give them
      all we have. Then we lose family over time. What else could rust the heart
      more over time? Blackgold.

      <template slot="footer">
        <base-button
          type="secondary"
          class="ml-auto"
          @click="modals.modal7 = false"
          >Close
        </base-button>
        <base-button type="primary">Save changes</base-button>
      </template>
    </modal>

    <base-button type="primary" class=" mb-3" @click="modals.modal8 = true">
      Large Modal
    </base-button>

    <modal
      :show.sync="modals.modal8"
      modal-classes="modal-dialog-centered modal-lg"
    >
      <h5 slot="header" class="modal-title" id="modal-title-default">
        Modal title
      </h5>

      The time is now for it to be okay to be great. People in this world shun
      people for being great. For being a bright color. For standing out. But
      the time is now to be okay to be the greatest you. Would you believe in
      what you believe in, if you were the only one who believed it?
      <br />
      <br />
      I always felt like I could do anything. That’s the main thing people are
      controlled by! Thoughts- their perception of themselves! They're slowed
      down by their perception of themselves. If you're taught you can’t do
      anything, you won’t do anything. I was taught I could do everything.

      <template slot="footer">
        <base-button
          type="secondary"
          class="ml-auto"
          @click="modals.modal7 = false"
          >Close
        </base-button>
        <base-button type="primary">Save changes</base-button>
      </template>
    </modal>

    <base-button type="primary" class=" mb-3" @click="modals.modal9 = true">
      Small Modal
    </base-button>

    <modal
      :show.sync="modals.modal9"
      modal-classes="modal-dialog-centered modal-sm"
    >
      <h5 slot="header" class="modal-title" id="modal-title-default">
        Modal title
      </h5>

      I always felt like I could do anything. That’s the main thing people are
      controlled by! Thoughts- their perception of themselves! They're slowed
      down by their perception of themselves. If you're taught you can’t do
      anything, you won’t do anything. I was taught I could do everything.

      <template slot="footer">
        <base-button
          type="secondary"
          class="ml-auto"
          @click="modals.modal7 = false"
          >Close
        </base-button>
        <base-button type="primary">Save changes</base-button>
      </template>
    </modal>
  </div>
</template>
<script>
import Modal from "@/components/Modal.vue";

export default {
  components: {
    Modal
  },
  data() {
    return {
      modals: {
        modal1: false,
        modal2: false,
        modal3: false,
        modal4: false,
        modal5: false,
        modal6: false,
        modal7: false,
        modal8: false,
        modal9: false
      }
    };
  }
};
</script>
<style></style>
