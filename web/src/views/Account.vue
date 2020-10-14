<template>
  <div class="wrapper">
    <div class="section-shaped my-0 skew-separator skew-mini">
      <div class="page-header page-header-small header-filter">
        <div
          class="page-header-image"
          style="background-image: url('img/pages/georgie.jpg');"
        ></div>
        <div class="container">
          <div class="header-body text-center mb-7">
            <div class="row justify-content-center"></div>
          </div>
        </div>
      </div>
    </div>
    <div class="bg-secondary">
      <div class="container bg-white card mb-0">
        <div class="row">
          <div class="col-md-3">
            <div class="section">
              <section class="text-center">
                <img
                  class="img img-raised shadow rounded-circle"
                  style="max-width: 180px;"
                  :src="user.avatar_url"
                />
                <h3 class="title mt-4">{{ user.name }}</h3>
                <p class="title mt-4" v-if="user.email">{{ user.email }}</p>
              </section>
            </div>
          </div>
          <div class="col-md-8 ml-auto">
            <div class="section">
              <div class="tab-content">
                <div v-if="accountTab === 'General'" class="tab-pane active">
                  <div>
                    <header>
                      <h2 class="text-uppercase">Account</h2>
                    </header>
                    <hr class="line-primary" />
                    <br />

                    <div class="row">
                      <div class="col-md-3 align-self-center">
                        <label class="labels" for="#email">Email</label>
                      </div>
                      <div class="col-md-9 align-self-center">
                        <base-input
                          id="email"
                          name="email"
                          type="email"
                          v-model="userEmail"
                          :placeholder="user.email"
                        ></base-input>
                      </div>
                    </div>

                    <div class="row mt-5">
                      <div class="col-md-6">
                        <base-button
                          nativeType="submit"
                          type="primary"
                          v-on:click="updateUserEmail"
                          :disabled="isEmailUpdateDisable"
                          >Save Changes</base-button
                        >
                      </div>
                    </div>
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
export default {
  bodyClass: "account-settings",
  components: {},
  created() {
    this.$Progress.start();
    this.$store.dispatch("getUser");
  },
  mounted() {
    this.$Progress.finish();
  },
  data() {
    return {
      query: "",
      accountTab: "General",
      userEmail: ""
    };
  },
  computed: {
    user() {
      return this.$store.getters.userRecord;
    },
    isEmailUpdateDisable() {
      return this.userEmail.length === 0 || !this.validEmail(this.userEmail);
    }
  },
  methods: {
    updateUserEmail() {
      this.$Progress.start();
      this.$store
        .dispatch("updateUser", { email: this.userEmail })
        .then(() => {
          this.$Progress.finish();
        })
        .catch(() => {
          this.$Progress.fail();
        });
    },
    validEmail: function(email) {
      var re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return re.test(email);
    }
  }
};
</script>
<style>
.account-settings .nav {
  text-align: left;
}

.account-settings .nav .nav-item {
  padding: 1rem 0;
}

.account-settings .nav .nav-item:not(:last-child) {
  border-bottom: 1px solid #5e72e4;
}
</style>
