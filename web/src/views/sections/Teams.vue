<template>
  <div class="cd-section" id="teams">
    <!--     *********     TEAM 1      *********      -->
    <Team1></Team1>
    <!--     *********     TEAM 2      *********      -->
    <Team2></Team2>
    <!--     *********     TEAM 3      *********      -->
    <Team3></Team3>
    <!--     *********     TEAM 4      *********      -->
    <Team4></Team4>
    <!--     *********     TEAM 5      *********      -->
    <Team5></Team5>
  </div>
</template>
<script>
import Team1 from "./Teams/Team1";
import Team2 from "./Teams/Team2";
import Team3 from "./Teams/Team3";
import Team4 from "./Teams/Team4";
import Team5 from "./Teams/Team5";

export default {
  components: {
    Team1,
    Team2,
    Team3,
    Team4,
    Team5
  }
};
</script>
<style>
#teams .carousel-item {
  height: 100vh;
}
</style>
