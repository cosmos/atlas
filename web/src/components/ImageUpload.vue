<template>
  <div class="fileinput text-center">
    <div
      class="fileinput-new thumbnail"
      :class="{ 'img-circle': type === 'avatar' }"
    >
      <img :src="image" alt="preview" />
    </div>
    <div>
      <span class="btn btn-file" :class="btnClasses">
        <span class="fileinput-new">{{
          fileExists ? changeText : selectText
        }}</span>
        <input type="hidden" value="" name="" />
        <input
          accept="image/*"
          @change="handlePreview"
          type="file"
          name="..."
          class="valid"
          :multiple="false"
          aria-invalid="false"
        />
      </span>
      <base-button v-if="fileExists" @click="removeFile" round type="danger">
        <i class="fa fa-times"></i> {{ removeText }}
      </base-button>
    </div>
  </div>
</template>
<script>
export default {
  name: "image-upload",
  props: {
    type: {
      type: String,
      default: "",
      description: 'Image upload type (""|avatar)'
    },
    btnClasses: {
      type: String,
      default: "",
      description: "Add photo button classes"
    },
    src: {
      type: String,
      default: "",
      description: "Initial image to display"
    },
    selectText: {
      type: String,
      default: "Select image"
    },
    changeText: {
      type: String,
      default: "Change"
    },
    removeText: {
      type: String,
      default: "Remove"
    }
  },
  data() {
    let avatarPlaceholder = "img/placeholder.jpg";
    let imgPlaceholder = "img/image_placeholder.jpg";
    return {
      placeholder: this.type === "avatar" ? avatarPlaceholder : imgPlaceholder,
      imagePreview: null
    };
  },
  computed: {
    fileExists() {
      return this.imagePreview !== null;
    },
    image() {
      return this.imagePreview || this.src || this.placeholder;
    }
  },
  methods: {
    handlePreview(event) {
      let file = event.target.files[0];
      this.imagePreview = URL.createObjectURL(file);
      this.$emit("change", file);
    },
    removeFile() {
      this.imagePreview = null;
      this.$emit("change", null);
    }
  }
};
</script>
<style></style>
