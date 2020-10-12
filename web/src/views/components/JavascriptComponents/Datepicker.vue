<template>
  <div class="row">
    <div class="col-md-4">
      <small class="d-block text-uppercase font-weight-bold mb-3"
        >Single date</small
      >
      <base-input addon-left-icon="ni ni-calendar-grid-58">
        <flat-picker
          slot-scope="{ focus, blur }"
          @on-open="focus"
          @on-close="blur"
          :config="{ allowInput: true }"
          class="form-control datepicker"
          v-model="dates.simple"
        >
        </flat-picker>
      </base-input>
    </div>
    <div class="col-md-4">
      <small class="d-block text-uppercase font-weight-bold mb-3"
        >Date range</small
      >

      <div class="input-daterange datepicker row align-items-center">
        <div class="col">
          <base-input addon-left-icon="ni ni-calendar-grid-58">
            <flat-picker
              slot-scope="{ focus, blur }"
              @on-open="focus"
              @on-close="blur"
              :config="{ allowInput: true, mode: 'range' }"
              class="form-control datepicker"
              v-model="dates.range"
            >
            </flat-picker>
          </base-input>
        </div>
      </div>
    </div>
    <div class="col-md-4">
      <small class="d-block text-uppercase font-weight-bold mb-3"
        >Datetimepicker</small
      >

      <div class="input-datetime datepicker row align-items-center">
        <div class="col">
          <base-input addon-left-icon="ni ni-calendar-grid-58">
            <flat-picker
              slot-scope="{ focus, blur }"
              @on-open="focus"
              @on-close="blur"
              :config="{ allowInput: true, enableTime: true }"
              class="form-control datepicker"
              v-model="dates.datetime"
            >
            </flat-picker>
          </base-input>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import flatPicker from "vue-flatpickr-component";
import "flatpickr/dist/flatpickr.css";

export default {
  components: {
    flatPicker
  },
  data() {
    return {
      dates: {
        simple: "2020-07-17",
        range: "2020-07-17 to 2020-07-19",
        datetime: "2020-07-31 12:00"
      }
    };
  }
};
</script>
<style></style>
