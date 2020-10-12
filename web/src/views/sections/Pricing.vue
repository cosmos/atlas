<template>
  <div class="cd-section" id="pricing">
    <!--     *********     PRICING 1      *********      -->
    <Pricing1></Pricing1>
    <!--     *********     PRICING 2      *********      -->
    <Pricing2></Pricing2>
    <!--     *********     PRICING 3     *********      -->
    <Pricing3></Pricing3>
    <!--     *********     PRICING 4     *********      -->
    <Pricing4></Pricing4>
    <!--     *********     PRICING 5     *********      -->
    <Pricing5></Pricing5>
  </div>
</template>
<script>
import Pricing1 from "./Pricing/Pricing1";
import Pricing2 from "./Pricing/Pricing2";
import Pricing3 from "./Pricing/Pricing3";
import Pricing4 from "./Pricing/Pricing4";
import Pricing5 from "./Pricing/Pricing5";

export default {
  components: {
    Pricing1,
    Pricing2,
    Pricing3,
    Pricing4,
    Pricing5
  }
};
</script>
<style></style>
