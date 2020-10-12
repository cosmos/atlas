<template>
  <div class="container">
    <h2 class="mt-lg mb-5">
      <span>Javascript Components</span>
    </h2>
    <h3 class="h4 text-success font-weight-bold mb-4">Modals</h3>
    <modals></modals>
    <h3 class="h4 text-success font-weight-bold mt-md mb-4">Datepicker</h3>
    <datepicker></datepicker>
    <tooltips-popovers></tooltips-popovers>
    <h3 class="h4 text-success font-weight-bold mt-md mb-4 mt-5">
      File Upload
    </h3>
    <file-upload></file-upload>
  </div>
</template>
<script>
import Modals from "./JavascriptComponents/Modals";
import Datepicker from "./JavascriptComponents/Datepicker";
import TooltipsPopovers from "./JavascriptComponents/TooltipsPopovers";
import FileUpload from "./JavascriptComponents/FileUpload";

export default {
  components: {
    Modals,
    Datepicker,
    TooltipsPopovers,
    FileUpload
  }
};
</script>
<style></style>
