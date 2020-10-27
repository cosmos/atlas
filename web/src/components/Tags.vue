<template>
  <div>
    <input ref="select" :value="value" />
  </div>
</template>
<script>
import Choices from "choices.js";
import "choices.js/public/assets/styles/choices.min.css";

export default {
  name: "selects",
  props: ["options", "value"],
  mounted: function() {
    this.choicesInstance = new Choices(this.$refs.select, {
      delimiter: ",",
      editItems: true,
      removeItemButton: true,
      placeholder: true,
      paste: false,
      placeholderValue: "+ Add",
      addItems: true,
      defaultValue: "da, da"
    });
    this.$refs.select.addEventListener("addItem", this.handleSelectChange);
  },
  methods: {
    handleSelectChange(e) {
      this.$emit("input", e.target.value);
    }
  },
  destroyed: function() {
    this.choicesInstance.destroy();
  }
};
</script>
<style></style>
