<template>
  <div class="row">
    <div class="col-md-5 col-sm-8">
      <small class="d-block text-uppercase font-weight-bold mb-3"
        >Regular Image</small
      >
      <image-upload @change="onImageChange" select-text="Select Image" />
    </div>
    <div class="col-md-5 col-sm-8">
      <small class="d-block text-uppercase font-weight-bold mb-3">Avatar</small>
      <image-upload
        type="avatar"
        select-text="Add photo"
        @change="onAvatarChange"
      />
    </div>
  </div>
</template>
<script>
import ImageUpload from "@/components/ImageUpload";
export default {
  components: {
    ImageUpload
  },
  methods: {
    onImageChange(file) {
      this.images.regular = file;
    },
    onAvatarChange(file) {
      this.images.avatar = file;
    }
  }
};
</script>
<style></style>
