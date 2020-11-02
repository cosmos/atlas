<template>
  <div id="app">
    <notifications
      group="errors"
      position="top right"
      class="error-notification"
    />
    <router-view name="header"></router-view>
    <vue-progress-bar></vue-progress-bar>
    <main>
      <fade-transition origin="center" mode="out-in" :duration="250">
        <router-view />
      </fade-transition>
    </main>
    <router-view name="footer"></router-view>
  </div>
</template>

<script>
import { FadeTransition } from "vue2-transitions";

export default {
  components: {
    FadeTransition
  }
};
</script>

<style>
.error-notification {
  width: 500px;
  top: 30px !important;
  right: 20px !important;
  position: fixed;
  z-index: 10000;
}
</style>
