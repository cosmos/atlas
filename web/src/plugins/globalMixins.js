import moment from "moment";

const GlobalMixins = {
  install(Vue) {
    Vue.mixin({
      created() {
        this.$Progress.start();
      },

      mounted() {
        let { bodyClass } = this.$options;
        if (bodyClass) {
          document.body.classList.add(bodyClass);
        }

        if (this.$store != null && !this.$store.getters.isAuthenticated) {
          this.$store.dispatch("getUser");
        }

        this.$Progress.finish();
      },

      beforeDestroy() {
        let { bodyClass } = this.$options;
        if (bodyClass) {
          document.body.classList.remove(bodyClass);
        }
      },

      computed: {
        isAuthenticated() {
          return this.$store.getters.isAuthenticated;
        },

        user() {
          return this.$store.getters.userRecord;
        }
      },

      methods: {
        objectEmpty(obj) {
          return Object.keys(obj).length === 0;
        },

        avatarPicture(author) {
          return author.avatar_url != ""
            ? author.avatar_url
            : "/img/generic-avatar.png";
        },

        formatDate(timestamp) {
          return moment(timestamp).fromNow();
        },

        queryModules: function(searchCriteria) {
          if (
            this.$route.name === "search" &&
            this.$route.query.q === searchCriteria
          ) {
            // prevent routing when we're on the results page with the same query
            return;
          }

          this.$router.push({
            name: "search",
            query: { q: searchCriteria }
          });
        },

        logout: function() {
          if (this.$store.getters.isAuthenticated) {
            this.$store.dispatch("logoutUser", this.$router);
          }
        },

        latestVersion(versions) {
          return versions.reduce((a, b) => {
            let aUpdated = new Date(a.updated_at);
            let bUpdated = new Date(b.updated_at);

            return aUpdated > bUpdated ? a : b;
          }).version;
        }
      }
    });
  }
};

export default GlobalMixins;
