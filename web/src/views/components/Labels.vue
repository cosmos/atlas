<template>
  <div class="container section-components">
    <h3 class="h4 text-success font-weight-bold mb-4 mt-4">Labels</h3>
    <div class="row">
      <div class="col-md-6">
        <div class="mb-3">
          <small class="text-uppercase font-weight-bold"
            >Contextual variations</small
          >
        </div>
        <badge type="default" class="mr-2">Default</badge>
        <badge type="primary" class="mr-2">Primary</badge>
        <badge type="secondary" class="mr-2">Secondary</badge>
        <badge type="info" class="mr-2">Info</badge>
        <badge type="success" class="mr-2">Success</badge>
        <badge type="danger" class="mr-2">Danger</badge>
        <badge type="warning" class="mr-2">Warning</badge>
        <div class="mb-2 mt-2">
          <small class="text-uppercase font-weight-bold">Pill Badges</small>
        </div>
        <badge type="default" rounded class="mr-2">Default</badge>
        <badge type="primary" rounded class="mr-2">Primary</badge>
        <badge type="secondary" rounded class="mr-2">Secondary</badge>
        <badge type="info" rounded class="mr-2">Info</badge>
        <badge type="success" rounded class="mr-2">Success</badge>
        <badge type="danger" rounded class="mr-2">Danger</badge>
        <badge type="warning" rounded class="mr-2">Warning</badge>
        <div class="mb-2 mt-2">
          <small class="text-uppercase font-weight-bold">Links</small>
        </div>
        <badge
          type="default"
          tag="a"
          href="javascript:void(0)"
          rounded
          class="mr-2"
          >Default</badge
        >
        <badge
          type="primary"
          tag="a"
          href="javascript:void(0)"
          rounded
          class="mr-2"
          >Primary</badge
        >
        <badge
          type="secondary"
          tag="a"
          href="javascript:void(0)"
          rounded
          class="mr-2"
          >Secondary</badge
        >
        <badge
          type="info"
          tag="a"
          href="javascript:void(0)"
          rounded
          class="mr-2"
          >Info</badge
        >
        <badge
          type="success"
          tag="a"
          href="javascript:void(0)"
          rounded
          class="mr-2"
          >Success</badge
        >
        <badge
          type="danger"
          tag="a"
          href="javascript:void(0)"
          rounded
          class="mr-2"
          >Danger</badge
        >
        <badge
          type="warning"
          tag="a"
          href="javascript:void(0)"
          rounded
          class="mr-2"
          >Warning</badge
        >
      </div>
      <div class="col-md-6">
        <div class="mb-3">
          <small class="text-uppercase font-weight-bold">Style</small>
        </div>
        <base-button type="default">
          Unread messages
          <badge type="primary" class="ml-2">24</badge>
        </base-button>

        <base-button type="primary">
          <span>Notifications</span>
          <badge
            type="danger"
            rounded
            class="badge-floating badge-md border-white badge-circle"
            >4</badge
          >
        </base-button>
        <div class="mb-3 mt-2">
          <small class="text-uppercase font-weight-bold">Sizes</small>
        </div>
        <badge type="primary" rounded class="ml-2">Default</badge>
        <badge type="warning" rounded class="ml-2 badge-md">Medium</badge>
        <badge type="success" rounded class="ml-2 badge-lg">Large badge</badge>
      </div>
    </div>
  </div>
</template>
<script></script>
<style></style>
