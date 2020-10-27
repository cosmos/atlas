<template>
  <div class="cd-section" id="headers">
    <!--     *********     HEADER 1      *********      -->
    <Header1></Header1>
    <!--     *********     HEADER 2      *********      -->
    <Header2></Header2>
    <!--     *********     HEADER 3      *********      -->
    <Header3></Header3>
    <!--     *********     HEADER 4  w/ VIDEO    *********      -->
    <Header4></Header4>
    <!--     *********    HEADER 5      *********      -->
    <Header5></Header5>
  </div>
</template>
<script>
import Header1 from "./Headers/Header1.vue";
import Header2 from "./Headers/Header2.vue";
import Header3 from "./Headers/Header3.vue";
import Header4 from "./Headers/Header4.vue";
import Header5 from "./Headers/Header5.vue";

export default {
  components: {
    Header1,
    Header2,
    Header3,
    Header4,
    Header5
  }
};
</script>
<style></style>
