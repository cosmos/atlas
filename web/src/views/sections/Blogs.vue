<template>
  <div class="cd-section" id="blogs">
    <!--     *********     BLOG 1      *********      -->
    <Blog1></Blog1>
    <!--     *********     BLOG 2      *********      -->
    <Blog2></Blog2>
    <!--     *********     BLOG 3      *********      -->
    <Blog3></Blog3>
    <!--     *********     BLOG 4      *********      -->
    <Blog4></Blog4>
    <!--     *********     BLOG 5      *********      -->
    <Blog5></Blog5>
    <!--     *********     BLOG 6      *********      -->
    <Blog6></Blog6>
    <!--     *********     BLOG 7      *********      -->
    <Blog7></Blog7>
  </div>
</template>
<script>
import Blog1 from "./Blogs/Blog1.vue";
import Blog2 from "./Blogs/Blog2.vue";
import Blog3 from "./Blogs/Blog3.vue";
import Blog4 from "./Blogs/Blog4.vue";
import Blog5 from "./Blogs/Blog5.vue";
import Blog6 from "./Blogs/Blog6.vue";
import Blog7 from "./Blogs/Blog7.vue";

export default {
  components: {
    Blog1,
    Blog2,
    Blog3,
    Blog4,
    Blog5,
    Blog6,
    Blog7
  }
};
</script>
<style></style>
