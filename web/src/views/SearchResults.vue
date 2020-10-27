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
      style="position: relative; padding-top: 20vh; min-height: 100vh;"
    >
      <div class="container mb-0">
        <div
          class="row"
          v-if="responseData.results && responseData.results.length > 0"
        >
          <div
            class="col-lg-4 col-md-6"
            v-for="mod in responseData.results"
            v-bind:key="mod.name"
          >
            <card class="card-blog">
              <template slot="body">
                <h6 class="card-category">
                  <i class="ni ni-badge"></i> {{ mod.team }}
                </h6>
                <h5 class="card-title">
                  <router-link
                    :to="{ name: 'modules', params: { id: mod.id } }"
                  >
                    {{ mod.name }}
                  </router-link>
                </h5>
                <p class="card-description">
                  {{ mod.description }}
                </p>
              </template>
              <template slot="footer">
                <div class="avatar-group author">
                  <router-link
                    v-for="author in mod.authors.slice(0, 3)"
                    v-bind:key="author.name"
                    :to="{ name: 'profile', params: { name: author.name } }"
                    class="avatar avatar-sm rounded-circle"
                    :title="author.name"
                  >
                    <img alt="Image placeholder" :src="avatarPicture(author)" />
                  </router-link>
                  <p class="extra-authors" v-if="mod.authors.length > 3">
                    ...
                  </p>
                </div>
                <div class="stats stats-right">
                  <i class="fa fa-star"></i> 0 ·
                  <i class="ni ni-archive-2"></i>
                  {{ latestVersion(mod.versions) }}
                </div>
              </template>
            </card>
          </div>
        </div>
        <div class="row justify-content-center" v-if="noMatch">
          <h1 class="title text-white">
            Sorry, no modules match your search criteria :(
          </h1>
        </div>
        <div
          class="row justify-content-center"
          v-if="responseData.results && responseData.results.length > 0"
        >
          <base-button
            class="align-self-center"
            nativeType="submit"
            type="neutral"
            :disabled="!this.responseData.prev_cursor"
            v-on:click="prevModules"
          >
            <i class="ni ni-bold-left"></i>
          </base-button>
          <base-button
            class="align-self-center"
            nativeType="submit"
            type="neutral"
            :disabled="!this.responseData.next_cursor"
            v-on:click="nextModules"
          >
            <i class="ni ni-bold-right"></i>
          </base-button>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import "bootstrap-vue/dist/bootstrap-vue.min.css";
import APIClient from "../plugins/apiClient";

export default {
  bodyClass: "search-results-page",
  components: {},
  watch: {
    cursor: function() {
      this.searchModules();
    }
  },
  methods: {
    filteredModules(x, y) {
      return this.responseData.results.slice(x, y);
    },

    prevModules() {
      this.page = "prev";
      this.cursor = this.responseData.prev_cursor;
    },

    nextModules() {
      this.page = "next";
      this.cursor = this.responseData.next_cursor;
    },

    searchModules() {
      APIClient.searchModules(
        this.$route.query.q,
        this.cursor,
        this.pageSize,
        this.page
      )
        .then(resp => {
          this.noMatch = resp.results.length === 0;
          this.responseData = resp;
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
    }
  },
  created() {
    this.searchModules();
  },
  data() {
    return {
      cursor: 0,
      page: "next",
      responseData: {},
      pageSize: 9,
      noMatch: false
    };
  },
  beforeRouteUpdate(to, from, next) {
    next();
    this.$Progress.start();
    this.cursor = 0;
    this.page = "next";
    this.searchModules();
    this.$Progress.finish();
  }
};
</script>

<style>
.stats i {
  top: 0;
}

.search-results-page .main {
  margin-top: 0;
}

.card-category {
  color: #ba3fd9;
}

.avatar-group .avatar {
  background-color: white;
}

.extra-authors {
  float: right;
  padding-left: 5px;
  margin-bottom: 0%;
  font-size: 1.2rem;
}
</style>
