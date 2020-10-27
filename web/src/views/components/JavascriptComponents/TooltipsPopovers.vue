<template>
  <div class="row tooltips-popovers">
    <div class="col-lg-6">
      <h3 class="h4 text-success font-weight-bold mt-md mb-4">Tooltips</h3>
      <small class="d-block text-uppercase font-weight-bold mb-3">Colors</small>
      <base-button
        size="sm"
        type="primary"
        class="btn-tooltip"
        v-b-tooltip.hover.top
        title="Tooltip on top"
      >
        Primary
      </base-button>

      <base-button
        size="sm"
        type="info"
        class="btn-tooltip"
        v-b-tooltip.hover.top
        title="Tooltip on top"
      >
        Info
      </base-button>

      <base-button
        size="sm"
        type="success"
        class="btn-tooltip"
        v-b-tooltip.hover.top
        title="Tooltip on top"
      >
        Success
      </base-button>

      <base-button
        size="sm"
        type="warning"
        class="btn-tooltip"
        v-b-tooltip.hover.top
        title="Tooltip on top"
      >
        Warning
      </base-button>

      <base-button
        size="sm"
        type="danger"
        class="btn-tooltip"
        v-b-tooltip.hover.top
        title="Tooltip on top"
      >
        Danger
      </base-button>

      <base-button
        size="sm"
        type="default"
        class="btn-tooltip"
        v-b-tooltip.hover.top
        title="Tooltip on top"
      >
        Default
      </base-button>

      <base-button
        size="sm"
        type="secondary"
        class="btn-tooltip"
        v-b-tooltip.hover.top
        title="Tooltip on top"
      >
        Secondary
      </base-button>
      <br />
      <br />
      <small class="d-block text-uppercase font-weight-bold mb-3"
        >Position</small
      >
      <base-button
        size="sm"
        type="primary"
        class="btn-tooltip"
        v-b-tooltip.hover.left
        title="Tooltip on left"
        >On left
      </base-button>
      <base-button
        size="sm"
        type="primary"
        class="btn-tooltip"
        v-b-tooltip.hover.top
        title="Tooltip on top"
        >On top
      </base-button>
      <base-button
        size="sm"
        type="primary"
        class="btn-tooltip"
        v-b-tooltip.hover.bottom
        title="Tooltip on bottom"
        >On bottom
      </base-button>
      <base-button
        size="sm"
        type="primary"
        class="btn-tooltip"
        v-b-tooltip.hover.right
        title="Tooltip on right"
        >On right
      </base-button>
    </div>
    <div class="col-lg-6">
      <h3 class="h4 text-success font-weight-bold mt-md mb-4">Popovers</h3>
      <small class="d-block text-uppercase font-weight-bold mb-3">Colors</small>
      <base-button
        size="sm"
        type="primary"
        v-b-popover.hover.top="
          'Vivamus sagittis lacus vel augue laoreet rutrum faucibus.'
        "
        title="Popover On Top"
        >Primary
      </base-button>
      <base-button
        size="sm"
        type="info"
        v-b-popover.hover.top="
          'Vivamus sagittis lacus vel augue laoreet rutrum faucibus.'
        "
        title="Popover On Top"
        >Info
      </base-button>
      <base-button
        size="sm"
        type="success"
        v-b-popover.hover.top="
          'Vivamus sagittis lacus vel augue laoreet rutrum faucibus.'
        "
        title="Popover On Top"
        >Success
      </base-button>
      <base-button
        size="sm"
        type="warning"
        v-b-popover.hover.top="
          'Vivamus sagittis lacus vel augue laoreet rutrum faucibus.'
        "
        title="Popover On Top"
        >Warning
      </base-button>
      <base-button
        size="sm"
        type="danger"
        v-b-popover.hover.top="
          'Vivamus sagittis lacus vel augue laoreet rutrum faucibus.'
        "
        title="Popover On Top"
        >Danger
      </base-button>
      <base-button
        size="sm"
        type="default"
        v-b-popover.hover.top="
          'Vivamus sagittis lacus vel augue laoreet rutrum faucibus.'
        "
        title="Popover On Top"
        >Default
      </base-button>
      <base-button
        size="sm"
        type="secondary"
        v-b-popover.hover.top="
          'Vivamus sagittis lacus vel augue laoreet rutrum faucibus.'
        "
        title="Popover On Top"
        >Secondary
      </base-button>
      <br />
      <br />
      <small class="d-block text-uppercase font-weight-bold mb-3"
        >Position</small
      >

      <base-button
        size="sm"
        type="default"
        v-b-popover.hover.left="
          'Vivamus sagittis lacus vel augue laoreet rutrum faucibus.'
        "
        title="Popover On Left"
        >On left
      </base-button>

      <base-button
        size="sm"
        type="default"
        v-b-popover.hover.top="
          'Vivamus sagittis lacus vel augue laoreet rutrum faucibus.'
        "
        title="Popover On Top"
        >On top
      </base-button>
      <base-button
        size="sm"
        type="default"
        v-b-popover.hover.right="
          'Vivamus sagittis lacus vel augue laoreet rutrum faucibus.'
        "
        title="Popover On right"
        >On right
      </base-button>
      <base-button
        size="sm"
        type="default"
        v-b-popover.hover.bottom="
          'Vivamus sagittis lacus vel augue laoreet rutrum faucibus.'
        "
        title="Popover On bottom"
        >On bottom
      </base-button>
    </div>
  </div>
</template>
<script>
import { VBTooltip } from "bootstrap-vue/esm/directives/tooltip/tooltip";
import { VBPopover } from "bootstrap-vue/esm/directives/popover/popover";
import "bootstrap-vue/dist/bootstrap-vue.min.css";

export default {
  directives: {
    BTooltip: VBTooltip,
    BPopover: VBPopover
  }
};
</script>
<style></style>
