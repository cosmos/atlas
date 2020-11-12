<template>
  <div
    class="section section-hero section-shaped"
    style="padding-bottom: 0; padding-top: 0;"
  >
    <div class="page-header">
      <div class="container shape-container d-flex align-items-center py-lg">
        <div class="col px-0">
          <div class="row align-items-center justify-content-center">
            <div class="col-lg-8 text-center">
              <div class="row">
                <div class="col-md-12 text-center">
                  <h1 style="color: white;" v-if="emailConfirmSuccess >= 1">
                    Email confirmed!
                  </h1>
                  <div v-if="emailConfirmSuccess <= -1">
                    <img
                      class="card-img"
                      src="/img/cosmosnaut-floating.svg"
                      style="width: 350px; height: 350px; padding-bottom: 30px;"
                    />

                    <h2 style="color: white;">
                      Invalid link. Please reconfirm.
                    </h2>
                  </div>
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
import APIClient from "../plugins/apiClient";

export default {
  created() {
    APIClient.confirmEmail(this.$route.params.token)
      .then(resp => {
        this.$store.commit("setUser", resp);
        this.emailConfirmSuccess = 1;

        this.$confetti.start({
          windSpeedMax: 0,
          particlesPerFrame: 4,
          particles: [
            {
              type: "rect",
              size: 5
            },
            {
              type: "circle",
              size: 5
            }
          ]
        });
      })
      .catch(err => {
        console.log(err);
        this.emailConfirmSuccess = -1;
      });
  },

  destroyed() {
    this.$confetti.stop();
  },

  data() {
    return {
      emailConfirmSuccess: 0
    };
  }
};
</script>

<style></style>
