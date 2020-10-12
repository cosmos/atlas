const GlobalMixins = {
  install(Vue) {
    Vue.mixin({
      mounted() {
        let {bodyClass} = this.$options;
        if (bodyClass) {
          document.body.classList.add(bodyClass);
        }

        if (!this.$store.getters.isAuthenticated) {
          this.$store.dispatch('getUser');
        }
      },
      beforeDestroy() {
        let {bodyClass} = this.$options;
        if (bodyClass) {
          document.body.classList.remove(bodyClass);
        }
      },
      methods: {
        queryModules: function() {
          this.$router.push({path: 'search', query: {q: this.searchCriteria}});
        },
        logout: function() {
          if (this.$store.getters.isAuthenticated) {
            this.$store.dispatch('logoutUser', this.$router);
          }
        }
      }
    });
  }
};

export default GlobalMixins;
