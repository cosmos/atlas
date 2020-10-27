<template>
  <div class="wrapper">
    <div
      class="page-header page-header-small header-filter skew-separator skew-mini"
      style="position: absolute;"
    >
      <div class="page-header-image"></div>
    </div>

    <div
      class="main"
      style="position: relative; padding-top: 40vh; min-height: 100vh;"
    >
      <div class="container bg-white card">
        <div class="row" style="padding-left: 25px; padding-right: 25px;">
          <div class="col-md-3">
            <div class="section">
              <section class="text-center">
                <img
                  class="img img-raised shadow rounded-circle"
                  style="max-width: 180px;"
                  :src="user.avatar_url"
                />
                <h3 class="title mt-4">{{ user.name }}</h3>
                <p v-if="user.email">{{ user.email }}</p>
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

                    <div class="row mt-5">
                      <div class="col-md-12 align-self-center">
                        <header>
                          <h5 class="text-uppercase">Tokens</h5>
                        </header>
                        <br />

                        <div>
                          <el-table
                            class="table table-striped table-flush"
                            :data="paginatedUserTokens"
                          >
                            <el-table-column
                              label="token"
                              prop="active"
                              sortable
                              scope="row"
                            >
                              <template v-slot="{ row }">
                                <div class="media align-items-center">
                                  {{ row.token }}
                                </div>
                              </template>
                            </el-table-column>

                            <div class="last-col">
                              <el-table-column
                                label="Revoke"
                                prop="revoke"
                                class="foo"
                              >
                                <template v-slot="{ row }">
                                  <base-button
                                    size="sm"
                                    icon="ni ni-fat-remove pt-1"
                                    type="danger"
                                    v-on:click="revokeUserToken(row)"
                                  ></base-button>
                                </template>
                              </el-table-column>
                            </div>
                          </el-table>
                          <div class="row">
                            <div class="col-md-5 align-self-center">
                              <base-button
                                nativeType="submit"
                                type="primary"
                                v-on:click="createUserToken"
                                >New Token</base-button
                              >
                            </div>

                            <div class="col-md-5">
                              <base-pagination
                                style="margin-top: revert;"
                                v-if="userTokens.length > 0"
                                type="primary"
                                v-model="currentPage"
                                :perPage="pageSize"
                                :total="userTokens.length"
                              ></base-pagination>
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
      </div>
    </div>
  </div>
</template>
<script>
import { Table, TableColumn } from "element-ui";

export default {
  bodyClass: "account-settings",
  components: {
    [Table.name]: Table,
    [TableColumn.name]: TableColumn
  },
  created() {
    this.$store.dispatch("getUser");
    this.$store.dispatch("getUserTokens");
  },
  data() {
    return {
      query: "",
      accountTab: "General",
      userEmail: "",
      currentPage: 1,
      pageSize: 5
    };
  },
  computed: {
    user() {
      return this.$store.getters.userRecord;
    },

    userTokens() {
      return this.$store.getters.userTokens
        .filter(token => !token.revoked)
        .sort((a, b) => {
          let aUpdated = new Date(a.updated_at);
          let bUpdated = new Date(b.updated_at);

          if (aUpdated > bUpdated) {
            return -1;
          }
          if (aUpdated < bUpdated) {
            return 1;
          }

          return 0;
        });
    },

    paginatedUserTokens() {
      return this.userTokens.filter((row, index) => {
        let start = (this.currentPage - 1) * this.pageSize;
        let end = this.currentPage * this.pageSize;

        if (index >= start && index < end) {
          return true;
        }
      });
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
    },

    createUserToken() {
      this.$Progress.start();
      this.$store
        .dispatch("createUserToken")
        .then(() => {
          this.$Progress.finish();
        })
        .catch(err => {
          this.$Progress.fail();
          this.$notify({
            group: "errors",
            type: "error",
            duration: 3000,
            title: "Error",
            text: err
          });
        });
    },

    revokeUserToken(token) {
      this.$Progress.start();
      this.$store
        .dispatch("revokeUserToken", token)
        .then(() => {
          this.$Progress.finish();
        })
        .catch(err => {
          this.$Progress.fail();
          this.$notify({
            group: "errors",
            type: "error",
            duration: 3000,
            title: "Error",
            text: err
          });
        });
    }
  }
};
</script>

<style>
div.main {
  background: url(/img/stars.d8924548.d8924548.svg) repeat top,
    linear-gradient(145.11deg, #202854 9.49%, #171b39 91.06%);
}

.account-settings .nav {
  text-align: left;
}

.account-settings .nav .nav-item {
  padding: 1rem 0;
}

.el-table .hidden-columns {
  visibility: hidden;
  position: absolute;
  z-index: -1;
  width: 100%;
}

.table th,
.table td {
  border-top: none;
}

.el-table__empty-text {
  display: none;
}
</style>
