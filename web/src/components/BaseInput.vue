<template>
  <div
    class="form-group"
    :class="[
      { 'input-group': hasIcon },
      { 'has-danger': error },
      { focused: focused },
      { 'input-group-alternative': alternative },
      { 'has-label': label || $slots.label },
      { 'has-success': valid === true },
      { 'has-danger': valid === false }
    ]"
  >
    <slot name="label">
      <label v-if="label" :class="labelClasses">
        {{ label }}
      </label>
      <span class="text-danger" v-if="required">*</span>
    </slot>

    <div v-if="addonLeftIcon || $slots.addonLeft" class="input-group-prepend">
      <span class="input-group-text">
        <slot name="addonLeft">
          <i :class="addonLeftIcon"></i>
        </slot>
      </span>
    </div>

    <slot v-bind="slotData">
      <input
        :value="value"
        v-on="listeners"
        v-bind="$attrs"
        class="form-control"
        :class="[
          { 'is-valid': valid === true },
          { 'is-invalid': valid === false },
          inputClasses
        ]"
        aria-describedby="addon-right addon-left"
      />
    </slot>

    <div v-if="addonRightIcon || $slots.addonRight" class="input-group-append">
      <span class="input-group-text">
        <slot name="addonRight">
          <i :class="addonRightIcon"></i>
        </slot>
      </span>
    </div>

    <slot name="infoBlock"></slot>

    <slot name="helpBlock">
      <div
        class="text-danger invalid-feedback"
        style="display: block;"
        :class="{ 'mt-2': hasIcon }"
        v-if="error"
      >
        {{ error }}
      </div>
    </slot>
  </div>
</template>
<script>
export default {
  name: "base-input",
  props: {
    required: {
      type: Boolean,
      description: "Whether input is required"
    },
    valid: {
      type: Boolean,
      default: undefined,
      description: "Whether is valid"
    },
    alternative: {
      type: Boolean,
      description: "Whether input is type alternative"
    },
    label: {
      type: String,
      description: "Input label"
    },
    error: {
      type: String,
      description: "Input error"
    },
    labelClasses: {
      type: String,
      description: "Input label classes"
    },
    inputClasses: {
      type: String,
      description: "Input classes"
    },
    addonLeftIcon: {
      type: String,
      description: "Addon Left Icon"
    },
    addonRightIcon: {
      type: String,
      description: "Addon Right Icon"
    },
    value: {
      type: [String, Number],
      description: "Input value"
    }
  },
  data() {
    return {
      focused: false
    };
  },
  computed: {
    listeners() {
      return {
        ...this.$listeners,
        input: this.updateValue,
        focus: this.onFocus,
        blur: this.onBlur
      };
    },
    slotData() {
      return {
        focused: this.focused,
        ...this.listeners
      };
    },
    hasIcon() {
      const { addonRight, addonLeft } = this.$slots;
      return (
        addonRight !== undefined ||
        addonLeft !== undefined ||
        this.addonRightIcon !== undefined ||
        this.addonLeftIcon !== undefined
      );
    }
  },
  methods: {
    updateValue(evt) {
      let value = evt.target.value;
      this.$emit("input", value);
    },
    onFocus(value) {
      this.focused = true;
      this.$emit("focus", value);
    },
    onBlur(value) {
      this.focused = false;
      this.$emit("blur", value);
    }
  }
};
</script>
<style></style>
