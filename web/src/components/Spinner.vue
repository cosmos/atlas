<template>
  <component :is="tag" :class="[`spinner-${type}`]">
    <span class="sr-only">Loading...</span>
  </component>
</template>
<script>
export default {
  name: "spinner",
  props: {
    tag: {
      type: String,
      default: "div",
      description: "HTML tag for spinner"
    },
    type: {
      type: String,
      default: "border",
      description: "Spinner type (border|grow)"
    }
  }
};
</script>
<style></style>
