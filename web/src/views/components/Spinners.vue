<template>
  <div class="section pt-0">
    <div class="container">
      <h3 class="h4 text-success font-weight-bold mb-2">Spinners</h3>
      <div class="row">
        <div class="col-md-6">
          <div class="mb-3 mt-5">
            <small class="text-uppercase font-weight-bold"
              >Border Spinners</small
            >
          </div>
          <spinner></spinner>
          <base-button type="primary" class="btn-sm ml-2 mb-2" disabled="">
            <spinner class="spinner-border-sm"></spinner>
          </base-button>
          <base-button type="primary" class="btn-sm ml-2 mb-2" disabled="">
            <spinner class="spinner-border-sm"></spinner>
            Loading...
          </base-button>
          <br /><br />
          <small class="text-uppercase font-weight-bold">Colors</small>
          <br /><br />
          <spinner class="text-primary mr-2"></spinner>
          <spinner class="text-secondary mr-2"></spinner>
          <spinner class="text-success mr-2"></spinner>
          <spinner class="text-danger mr-2"></spinner>
          <spinner class="text-warning mr-2"></spinner>
          <spinner class="text-info mr-2"></spinner>
          <spinner class="text-light mr-2"></spinner>
          <spinner class="text-default mr-2"></spinner>
        </div>
        <div class="col-md-6">
          <div class="mb-3 mt-5">
            <small class="text-uppercase font-weight-bold"
              >Growing Spinners</small
            >
          </div>
          <spinner type="grow"></spinner>
          <base-button type="primary" class="btn-sm ml-2 mb-2" disabled="">
            <spinner type="grow" class="spinner-grow-sm"></spinner>
          </base-button>
          <base-button type="primary" class="btn-sm ml-2 mb-2" disabled="">
            <spinner type="grow" class="spinner-grow-sm"></spinner>
            Loading...
          </base-button>
          <br /><br />
          <small class="text-uppercase font-weight-bold">Colors</small>
          <br /><br />
          <spinner type="grow" class="text-primary mr-2"></spinner>
          <spinner type="grow" class="text-secondary mr-2"></spinner>
          <spinner type="grow" class="text-success mr-2"></spinner>
          <spinner type="grow" class="text-danger mr-2"></spinner>
          <spinner type="grow" class="text-warning mr-2"></spinner>
          <spinner type="grow" class="text-info mr-2"></spinner>
          <spinner type="grow" class="text-light mr-2"></spinner>
          <spinner type="grow" class="text-default mr-2"></spinner>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import Spinner from "@/components/Spinner";

export default {
  components: {
    Spinner
  }
};
</script>
<style></style>
