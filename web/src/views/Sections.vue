<template>
  <div class="wrapper sections-page">
    <headers></headers>
    <features></features>
    <blogs></blogs>
    <teams></teams>
    <projects></projects>
    <pricing></pricing>
    <testimonials></testimonials>
    <contact></contact>
    <tables></tables>
    <accordion></accordion>
  </div>
</template>
<script>
import Headers from "./sections/Headers";
import Features from "./sections/Features";
import Blogs from "./sections/Blogs";
import Teams from "./sections/Teams";
import Projects from "./sections/Projects";
import Pricing from "./sections/Pricing";
import Testimonials from "./sections/Testimonials";
import Contact from "./sections/Contact";
import Tables from "./sections/Tables";
import Accordion from "./sections/Accordion";

export default {
  name: "sections",
  components: {
    Headers,
    Features,
    Blogs,
    Teams,
    Projects,
    Pricing,
    Testimonials,
    Contact,
    Tables,
    Accordion
  }
};
</script>
