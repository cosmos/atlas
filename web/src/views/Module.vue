<template>
  <div class="wrapper">
    <div
      class="page-header page-header-small header-filter skew-separator skew-mini"
      style="position: absolute;"
    >
      <div class="page-header-image"></div>
    </div>

    <section
      class="section bg-secondary"
      style="position: relative; padding-top: 40vh; min-height: 100vh;"
    >
      <div class="container" style="padding-top: 100px;">
        <div class="card card-profile shadow mt--300">
          <modal :show.sync="showInviteModal">
            <div class="row">
              <div class="col-md-3 align-self-center">
                <label class="labels" for="#userToInvite">User</label>
              </div>
              <div class="col-md-9 align-self-center">
                <base-input
                  id="userToInvite"
                  name="userToInvite"
                  v-model="userToInvite"
                ></base-input>
              </div>
            </div>
            <template slot="footer">
              <base-button
                type="primary"
                v-on:click="inviteUser"
                :disabled="!userToInvite"
                >Invite</base-button
              >
              <base-button type="primary" @click="showInviteModal = false"
                >Close
              </base-button>
            </template>
          </modal>
          <div class="px-4">
            <div class="row justify-content-center">
              <div class="col-lg-4 order-lg-2">
                <div class="title text-center justify-content-center">
                  <h3>{{ module.name }}</h3>
                  <div class="h6 font-weight-300">
                    <i class="ni location_pin mr-2"></i>{{ module.team }}
                  </div>
                </div>
              </div>
              <div
                class="col-lg-4 order-lg-3 text-lg-right align-self-lg-center"
              >
                <div class="card-profile-actions py-4 mt-lg-0">
                  <base-button
                    tag="a"
                    target="_blank"
                    v-if="version.repo"
                    :href="version.repo"
                    type="primary"
                    size="sm"
                    >Repo</base-button
                  >
                  <base-button
                    tag="a"
                    target="_blank"
                    type="primary"
                    v-if="module.homepage"
                    :href="module.homepage"
                    size="sm"
                    >Homepage</base-button
                  >
                  <base-button
                    tag="a"
                    target="_blank"
                    type="primary"
                    v-if="module.bug_tracker && module.bug_tracker.url"
                    :href="module.bug_tracker.url"
                    size="sm"
                    >Bugs</base-button
                  >
                </div>
              </div>
              <div class="col-lg-4 order-lg-1">
                <div class="card-profile-stats d-flex justify-content-center">
                  <div class="title text-center justify-content-center">
                    <div class="module-header mb-2">
                      <base-button
                        type="primary"
                        size="sm"
                        :disabled="!isAuthenticated"
                        v-on:click="handleFavorite(starToggle)"
                        style="margin-right: 10px;"
                        ><i class="fa fa-star"></i>
                        {{ starToggle }}</base-button
                      >
                      {{ moduleStars }}
                    </div>
                    <div>
                      <span>Updated: {{ formatDate(module.updated_at) }}</span>
                    </div>
                  </div>
                </div>
              </div>
            </div>
            <div class="py-5 border-top text-center">
              <div class="row">
                <div
                  class="col-lg-8 text-left"
                  style="margin-right: 50px; padding-left: 30px;"
                >
                  <vue-markdown
                    :source="documentation"
                    :anchor-attributes="anchorAttrs"
                  ></vue-markdown>
                </div>
                <div class="col-lg-3 text-lg-right">
                  <h5 class="card-title mt-4">Owners</h5>
                  <div class="avatar-group">
                    <router-link
                      v-for="owner in module.owners"
                      v-bind:key="owner.name"
                      :to="{ name: 'profile', params: { name: owner.name } }"
                      class="avatar avatar-lg rounded-circle"
                      :title="owner.name"
                    >
                      <img
                        alt="Image placeholder"
                        :src="avatarPicture(owner)"
                      />
                    </router-link>
                    <div
                      class="avatar-group"
                      style="padding-top: 15px"
                      v-if="isOwner"
                    >
                      <base-button
                        type="primary"
                        size="sm"
                        @click="showInviteModal = true"
                        ><i class="fa fa-plus"></i> Invite</base-button
                      >
                    </div>
                  </div>
                  <h5 class="card-title mt-4">Authors</h5>
                  <div class="avatar-group">
                    <router-link
                      v-for="author in module.authors"
                      v-bind:key="author.name"
                      :to="{ name: 'profile', params: { name: author.name } }"
                      class="avatar avatar-lg rounded-circle"
                      :title="author.name"
                    >
                      <img
                        alt="Image placeholder"
                        :src="avatarPicture(author)"
                      />
                    </router-link>
                  </div>
                  <h5 class="card-title mt-4">Keywords</h5>
                  <div class="avatar-group">
                    <span
                      class="badge badge-pill badge-primary"
                      style="margin-right: 4px;"
                      v-for="keyword in module.keywords"
                      v-bind:key="keyword.name"
                    >
                      {{ keyword.name }}</span
                    >
                  </div>
                  <h5 class="card-title mt-4"></h5>
                  <el-table
                    class="table table-striped table-flush"
                    header-row-class-name="thead-light"
                    :data="sortedVersions"
                  >
                    <el-table-column
                      label="version"
                      prop="version"
                      sortable
                      scope="row"
                    >
                      <template v-slot="{ row }">
                        <router-link
                          :to="{
                            name: 'modulesVersioned',
                            params: { id: module.id, version: row.version }
                          }"
                          >{{ row.version }}</router-link
                        >
                      </template>
                    </el-table-column>
                    <el-table-column
                      label="SDK Compatibility"
                      prop="name"
                      sortable
                      scope="row"
                    >
                      <template v-slot="{ row }">
                        <div>
                          {{ row.sdk_compat }}
                        </div>
                      </template>
                    </el-table-column>
                  </el-table>
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
import { Table, TableColumn } from "element-ui";
import Modal from "@/components/Modal.vue";
import APIClient from "../plugins/apiClient";
import axios from "axios";
import VueMarkdown from "vue-markdown";
import xss from "xss";

export default {
  bodyClass: "profile-page",
  components: {
    VueMarkdown,
    Modal,
    [Table.name]: Table,
    [TableColumn.name]: TableColumn
  },

  created() {
    this.getModule();
  },

  data() {
    return {
      version: {},
      module: {},
      moduleStars: 0,
      documentation: "",
      userToInvite: "",
      anchorAttrs: {
        target: "_blank",
        rel: "noopener noreferrer nofollow"
      },
      showInviteModal: false
    };
  },

  computed: {
    isOwner() {
      if (!this.objectEmpty(this.user) && !this.objectEmpty(this.module)) {
        return this.module.owners.some(owner => owner.name === this.user.name);
      }

      return false;
    },

    starToggle() {
      if (!this.objectEmpty(this.user) && !this.objectEmpty(this.module)) {
        return this.user.stars.includes(this.module.id) ? "unstar" : "star";
      }

      return "";
    },

    sortedVersions() {
      if (this.objectEmpty(this.module)) {
        return [];
      }

      let versions = this.module.versions;
      versions.sort((a, b) => {
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

      return versions;
    }
  },

  watch: {
    $route() {
      this.$Progress.start();
      this.getModule();
      this.$Progress.finish();
    }
  },

  methods: {
    inviteUser() {
      this.$Progress.start();

      APIClient.inviteModuleOwner(this.userToInvite, this.module.id)
        .then(() => {
          this.showInviteModal = false;
          this.userToInvite = "";
          this.$Progress.finish();
        })
        .catch(err => {
          console.log(err);
          this.$Progress.fail();
          this.$notify({
            group: "errors",
            type: "error",
            duration: 3000,
            title: "Error",
            text: this.getResponseError(err)
          });
        });
    },

    handleFavorite(toggle) {
      if (toggle === "star") {
        APIClient.starModule(this.module.id)
          .then(resp => {
            this.moduleStars = resp.stars;
            this.$store.dispatch("getUser");
          })
          .catch(err => {
            console.log(err);
            this.$notify({
              group: "errors",
              type: "error",
              duration: 3000,
              title: "Error",
              text: this.getResponseError(err)
            });
          });
      } else if (toggle === "unstar") {
        APIClient.unstarModule(this.module.id)
          .then(resp => {
            this.moduleStars = resp.stars;
            this.$store.dispatch("getUser");
          })
          .catch(err => {
            console.log(err);
            this.$notify({
              group: "errors",
              type: "error",
              duration: 3000,
              title: "Error",
              text: this.getResponseError(err)
            });
          });
      }
    },

    getDocumentation(version) {
      axios
        .get(version.documentation)
        .then(resp => {
          this.documentation = xss(resp.data);
        })
        .catch(err => {
          console.log(err);
          this.$notify({
            group: "errors",
            type: "error",
            duration: 3000,
            title: "Error",
            text: err
          });
        });
    },

    getModule() {
      APIClient.getModule(this.$route.params.id)
        .then(resp => {
          if (!this.$route.params.version) {
            this.version = this.latestVersion(resp.versions);
          } else {
            this.version = resp.versions.find(v => {
              return v.version === this.$route.params.version;
            });
          }

          if (!this.version) {
            this.$router.push({ name: "error" });
            return;
          }

          this.module = resp;
          this.moduleStars = resp.stars;
          this.getDocumentation(this.version);
        })
        .catch(err => {
          console.log(err);
          this.$router.push({ name: "error" });
          return;
        });
    }
  }
};
</script>

<style>
.module-header {
  padding-top: 30px !important;
}

section.bg-secondary {
  background: url(/img/stars.d8924548.d8924548.svg) repeat top,
    linear-gradient(145.11deg, #202854 9.49%, #171b39 91.06%);
}

.table.align-items-center td {
  vertical-align: middle;
}

.table th,
.table td {
  border-top: none;
}

.table thead th {
  border-bottom: none;
}

.el-table .hidden-columns {
  visibility: hidden;
  position: absolute;
  z-index: -1;
}

.title {
  padding-top: 30px;
}
</style>
