<template>
  <div class="cd-section" id="projects">
    <!--     *********     PROJECT 1      *********      -->
    <Project1></Project1>
    <!--     *********     PROJECT 2      *********      -->
    <Project2></Project2>
    <!--     *********     PROJECT 3      *********      -->
    <Project3></Project3>
  </div>
</template>
<script>
import Project1 from "./Projects/Project1";
import Project2 from "./Projects/Project2";
import Project3 from "./Projects/Project3";

export default {
  components: {
    Project1,
    Project2,
    Project3
  }
};
</script>
<style></style>
