<template>
  <div class="cd-section" id="testimonials">
    <!--     *********     TESTIMONIALS 1      *********      -->
    <Testimonial1></Testimonial1>
    <!--     *********     TESTIMONIALS 2      *********      -->
    <Testimonial2></Testimonial2>
    <!--     *********     TESTIMONIALS 3      *********      -->
    <Testimonial3></Testimonial3>
  </div>
</template>
<script>
import Testimonial1 from "./Testimonial/Testimonial1";
import Testimonial2 from "./Testimonial/Testimonial2";
import Testimonial3 from "./Testimonial/Testimonial3";

export default {
  components: {
    Testimonial1,
    Testimonial2,
    Testimonial3
  }
};
</script>
<style></style>
