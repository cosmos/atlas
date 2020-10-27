<template>
  <div
    class="toast"
    role="alert"
    aria-live="assertive"
    artia-atomic="true"
    data-animation="auto-hide"
  >
    <div class="toast-header">
      <slot name="header-content"></slot>
      <button
        type="button"
        class="ml-2 close"
        data-dismiss="toast"
        aria-label="Close"
      >
        <span aria-hidden="true">×</span>
      </button>
    </div>
    <div class="toast-body">
      <slot name="text"></slot>
    </div>
  </div>
</template>
<script></script>
<style></style>
