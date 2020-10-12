<template>
  <div class="container">
    <div class="row">
      <div class="col-md-6">
        <h3 class="h4 text-success font-weight-bold mb-4">Progress bars</h3>
        <base-progress :value="60" label="Default"></base-progress>
        <base-progress
          :value="60"
          type="primary"
          label="primary"
        ></base-progress>
        <base-progress
          :value="60"
          type="secondary"
          label="secondary"
        ></base-progress>
        <base-progress :value="60" type="info" label="info"></base-progress>
        <base-progress
          :value="60"
          type="success"
          label="success"
        ></base-progress>
        <base-progress :value="60" type="danger" label="danger"></base-progress>
        <base-progress
          :value="60"
          type="warning"
          label="warning"
        ></base-progress>
      </div>
      <div class="col-md-5 pagination-area">
        <h4 class="h4 text-success font-weight-bold mb-4">Pagination</h4>
        <base-pagination
          :page-count="10"
          v-model="pagination.default"
          type="primary"
          :showArrows="false"
        ></base-pagination>
        <base-pagination
          :page-count="10"
          v-model="pagination.info"
          iconLeft="ni ni-bold-left"
          iconRight="ni ni-bold-right"
          type="info"
        ></base-pagination>
        <base-pagination
          :page-count="10"
          v-model="pagination.success"
          iconLeft="fa fa-angle-double-left"
          iconRight="fa fa-angle-double-right"
          type="success"
        ></base-pagination>
        <base-pagination
          :page-count="10"
          v-model="pagination.warning"
          type="warning"
        ></base-pagination>
        <base-pagination
          :page-count="10"
          v-model="pagination.danger"
          type="danger"
        ></base-pagination>
        <base-pagination
          :page-count="10"
          v-model="pagination.secondary"
          type="secondary"
        ></base-pagination>
        <base-pagination
          :page-count="10"
          v-model="pagination.default"
          type="default"
        ></base-pagination>
      </div>
    </div>
  </div>
</template>
<script>
export default {
  data() {
    return {
      pagination: {
        default: 3,
        info: 2,
        success: 3,
        warning: 1,
        secondary: 5,
        danger: 4
      }
    };
  }
};
</script>
<style></style>
