import moment from "moment";

const GlobalMixins = {
  install(Vue) {
    Vue.mixin({
      created() {
        this.$Progress.start();
      },
      mounted() {
        let {bodyClass} = this.$options;
        if (bodyClass) {
          document.body.classList.add(bodyClass);
        }

        if (this.$store != null && !this.$store.getters.isAuthenticated) {
          this.$store.dispatch('getUser');
        }

        this.$Progress.finish();
      },
      beforeDestroy() {
        let {bodyClass} = this.$options;
        if (bodyClass) {
          document.body.classList.remove(bodyClass);
        }
      },
      methods: {
        avatarPicture(author) {
          return author.avatar_url != ""
            ? author.avatar_url
            : "img/generic-avatar.png";
        },
        
        formatDate(timestamp) {
          return moment(timestamp).fromNow();
        },

        queryModules: function() {
          if (this.$route.name === 'search-results' && this.$route.query.q === this.searchCriteria) {
            // prevent routing when we're on the results page with the same query
            return
          }

          this.$router.push({name: 'search-results', query: {q: this.searchCriteria}});
          this.searchCriteria = "";
        },

        logout: function() {
          if (this.$store.getters.isAuthenticated) {
            this.$store.dispatch('logoutUser', this.$router);
          }
        },

        latestVersion(versions) {
          return versions.reduce((a, b) => {
            let aUpdated = new Date(a.updated_at);
            let bUpdated = new Date(b.updated_at);
    
            return aUpdated > bUpdated ? a : b;
          }).version;
        },
      }
    });
  }
};

export default GlobalMixins;
