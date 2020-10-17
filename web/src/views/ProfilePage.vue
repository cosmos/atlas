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
      style="position: relative; padding-top: 50vh; min-height: 100vh;"
    >
      <div class="container">
        <div class="card card-profile shadow mt--300">
          <div class="px-4">
            <div class="row justify-content-center">
              <div class="col-lg-3 order-lg-2">
                <div class="card-profile-image">
                  <a>
                    <img :src="avatarPicture" class="rounded-circle" />
                  </a>
                </div>
              </div>
              <div
                class="col-lg-4 order-lg-3 text-lg-right align-self-lg-center"
              >
                <div class="card-profile-actions py-4 mt-lg-0">
                  <base-button
                    tag="a"
                    target="_blank"
                    :href="`https://github.com/${user.name}`"
                    type="primary"
                    size="sm"
                    class="mr-4"
                    >GitHub</base-button
                  >
                  <base-button
                    tag="a"
                    v-show="user.email != null"
                    :href="`mailto:${user.email}`"
                    type="default"
                    size="sm"
                    class="float-right"
                    >Message</base-button
                  >
                </div>
              </div>
              <div class="col-lg-4 order-lg-1">
                <div class="card-profile-stats d-flex justify-content-center">
                  <div>
                    <span class="heading">{{ userModules.length }}</span>
                    <span class="description">Published Modules</span>
                  </div>
                  <div>
                    <span class="heading">0</span>
                    <span class="description">Followers</span>
                  </div>
                </div>
              </div>
            </div>
            <div class="text-center mt-5">
              <h3>{{ user.full_name }}</h3>
              <div class="h6 font-weight-300">
                <i class="ni location_pin mr-2"></i>{{ user.name }}
              </div>
            </div>
            <div class="mt-5 py-5 border-top text-center">
              <div class="row justify-content-center">
                <div class="col-lg-9">
                  <h3>
                    Modules
                  </h3>
                  <!--  -->

                  <el-table
                    class="table table-striped table-flush"
                    header-row-class-name="thead-light"
                    :data="paginatedUserModules"
                  >
                    <el-table-column prop="name" sortable scope="row">
                      <template v-slot="{ row }">
                        <div>
                          {{ row.name }}
                        </div>
                      </template>
                    </el-table-column>
                    <el-table-column prop="name" sortable scope="row">
                      <template v-slot="{ row }">
                        <div>
                          <a :href="row.documentation">Documentation</a>
                        </div>
                      </template>
                    </el-table-column>
                    <el-table-column prop="name" sortable scope="row">
                      <template v-slot="{ row }">
                        <div>
                          <a :href="row.repo">Repository</a>
                        </div>
                      </template>
                    </el-table-column>
                    <el-table-column prop="name" sortable scope="row">
                      <template v-slot="{ row }">
                        <div>Updated: {{ formatDate(row.updated_at) }}</div>
                      </template>
                    </el-table-column>
                    <el-table-column prop="name" sortable scope="row">
                      <template v-slot="{ row }">
                        <div>{{ latestVersion(row.versions) }}</div>
                      </template>
                    </el-table-column>
                  </el-table>
                  <div class="row justify-content-center">
                    <div class="col-md-5">
                      <base-pagination
                        class="justify-content-center"
                        style="margin-top: revert;"
                        v-if="userModules.length > 0"
                        type="primary"
                        v-model="currentPage"
                        :perPage="pageSize"
                        :total="userModules.length"
                      ></base-pagination>
                    </div>
                  </div>
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
import APIClient from "../plugins/apiClient";
import moment from "moment";

export default {
  bodyClass: "profile-page",
  components: {
    [Table.name]: Table,
    [TableColumn.name]: TableColumn
  },
  created() {
    this.getUserByName();
    this.getUserModules();
  },
  data() {
    return {
      user: {},
      userModules: [],
      currentPage: 1,
      pageSize: 10
    };
  },
  beforeRouteUpdate(to, from, next) {
    next();
    this.$Progress.start();
    this.getUserByName();
    this.getUserModules();
    this.$Progress.finish();
  },
  computed: {
    avatarPicture() {
      return this.user.avatar_url != ""
        ? this.user.avatar_url
        : "img/generic-avatar.png";
    },

    paginatedUserModules() {
      return this.userModules.filter((row, index) => {
        let start = (this.currentPage - 1) * this.pageSize;
        let end = this.currentPage * this.pageSize;

        if (index >= start && index < end) {
          return true;
        }
      });
    }
  },
  methods: {
    latestVersion(versions) {
      return versions.reduce((a, b) => {
        let aUpdated = new Date(a.updated_at);
        let bUpdated = new Date(b.updated_at);

        return aUpdated > bUpdated ? a : b;
      }).version;
    },

    formatDate(timestamp) {
      return moment(timestamp).fromNow();
    },

    getUserByName() {
      APIClient.getUserByName(this.$route.params.name)
        .then(resp => {
          this.user = resp;
        })
        .catch(err => {
          console.log(err);
          this.$router.push("/");
        });
    },

    getUserModules() {
      APIClient.getUserModules(this.$route.params.name)
        .then(resp => {
          this.userModules = resp;
        })
        .catch(err => {
          console.log(err);
        });
    }
  }
};
</script>

<style>
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
</style>
