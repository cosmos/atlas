<template>
  <div class="section section-select">
    <div class="container">
      <div class="row">
        <div class="col-md-6">
          <h3 class="h4 text-success font-weight-bold mb-4">
            Customizable Select
          </h3>
          <div class="row">
            <div class="col-md-6">
              <div class="mb-3">
                <label
                  class="text-uppercase font-weight-bold"
                  for="choices-single-default"
                  >Single Option</label
                >
                <choices-single :options="options" v-model="selected">
                  <option disabled value="0">Select one</option>
                </choices-single>
              </div>
            </div>
            <div class="col-md-6">
              <div class="mb-3">
                <label
                  class="text-uppercase font-weight-bold"
                  for="choices-multiple-default"
                  >Multiple Options</label
                >
                <choices-multiple :options="optionsMultiple" v-model="selected">
                  <option disabled value="0">Select one</option>
                </choices-multiple>
              </div>
            </div>
          </div>
        </div>
        <div class="col-lg-6 col-md-12 mt-md-5 mt-lg-0">
          <h3 class="h4 text-success font-weight-bold mb-4">Dropdowns</h3>
          <div class="row">
            <div class="col-md-4">
              <div class="mb-3">
                <small class="text-uppercase font-weight-bold"
                  >Multilevel</small
                >
              </div>
              <base-dropdown multiLevel>
                <a
                  role="button"
                  class="btn btn-primary text-white"
                  slot="title"
                  data-toggle="dropdown"
                >
                  <span class="nav-link-inner--text">Click here</span>
                </a>
                <a class="dropdown-item" href="javascript:;">Action</a>
                <a class="dropdown-item" href="javascript:;">Another action</a>
                <a
                  class="dropdown-item dropdown-toggle"
                  :class="{ open: multiLevel }"
                  @click="toggleMultiLevel"
                >
                  Submenu
                </a>
                <ul class="dropdown-menu" :class="{ show: multiLevel }">
                  <a class="dropdown-item" href="javascript:void(0)"
                    >Submenu action</a
                  >
                  <a class="dropdown-item" href="javascript:void(0)"
                    >Submenu action</a
                  >
                  <a
                    class="dropdown-item dropdown-toggle"
                    :class="{ open: multiLevel2 }"
                    @click="toggleMultiLevel2()"
                  >
                    Subsubmenu
                  </a>
                  <ul class="dropdown-menu" :class="{ show: multiLevel2 }">
                    <a class="dropdown-item" href="javascript:void(0)"
                      >Subsubmenu action 1</a
                    >
                    <a class="dropdown-item" href="javascript:void(0)"
                      >Subsubmenu action 2</a
                    >
                  </ul>
                  <a
                    class="dropdown-item dropdown-toggle"
                    :class="{ open: multiLevel3 }"
                    @click="toggleMultiLevel3()"
                  >
                    Second Subsubmenu
                  </a>
                  <ul class="dropdown-menu" :class="{ show: multiLevel3 }">
                    <a class="dropdown-item" href="javascript:void(0)"
                      >Subsubmenu action 1</a
                    >
                    <a class="dropdown-item" href="javascript:void(0)"
                      >Subsubmenu action 2</a
                    >
                  </ul>
                </ul>
              </base-dropdown>
            </div>
            <div class="col-md-4">
              <div class="mb-3">
                <small class="text-uppercase font-weight-bold">Dropdown</small>
              </div>
              <base-dropdown class="btn btn-primary">
                <a role="button" slot="title" data-toggle="dropdown">
                  <span class="nav-link-inner--text">Click me</span>
                </a>
                <a href="javascript:;" class="dropdown-item">
                  <i class="ni ni-single-02"></i>
                  <span>My profile</span>
                </a>
                <a href="javascript:;" class="dropdown-item">
                  <i class="ni ni-settings-gear-65"></i>
                  <span>Settings</span>
                </a>
                <a href="javascript:;" class="dropdown-item">
                  <i class="ni ni-calendar-grid-58"></i>
                  <span>Activity</span>
                </a>
                <a href="javascript:;" class="dropdown-item">
                  <i class="ni ni-support-16"></i>
                  <span>Support</span>
                </a>
                <div class="dropdown-divider"></div>
                <a href="javascript:;" class="dropdown-item">
                  <i class="ni ni-user-run"></i>
                  <span>Logout</span>
                </a>
              </base-dropdown>
            </div>
            <div class="col-md-4">
              <div class="mb-3">
                <small class="text-uppercase font-weight-bold">Dropup</small>
              </div>
              <base-dropdown class="btn btn-primary" direction="up">
                <a role="button" slot="title" data-toggle="dropdown">
                  <span class="nav-link-inner--text">Click me</span>
                </a>
                <a href="javascript:;" class="dropdown-item">
                  <i class="ni ni-single-02"></i>
                  <span>My profile</span>
                </a>
                <a href="javascript:;" class="dropdown-item">
                  <i class="ni ni-settings-gear-65"></i>
                  <span>Settings</span>
                </a>
                <a href="javascript:;" class="dropdown-item">
                  <i class="ni ni-calendar-grid-58"></i>
                  <span>Activity</span>
                </a>
                <a href="javascript:;" class="dropdown-item">
                  <i class="ni ni-support-16"></i>
                  <span>Support</span>
                </a>
                <div class="dropdown-divider"></div>
                <a href="javascript:;" class="dropdown-item">
                  <i class="ni ni-user-run"></i>
                  <span>Logout</span>
                </a>
              </base-dropdown>
            </div>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="col-md-6">
          <div class="form-group">
            <div class="mb-3 mt-5">
              <small class="text-uppercase font-weight-bold">Textarea</small>
            </div>
            <textarea
              class="form-control"
              id="exampleFormControlTextarea1"
              rows="3"
              spellcheck="false"
              placeholder="You can write your text here..."
            ></textarea>
          </div>
        </div>
        <div class="col-md-6">
          <div class="form-group">
            <div class="mb-3 mt-5">
              <small class="text-uppercase font-weight-bold">Tags</small>
              <tags value="Amsterdam, Sydney, Rome, Singapore"></tags>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import BaseDropdown from "@/components/BaseDropdown";
import ChoicesSingle from "@/components/SingleSelect";
import ChoicesMultiple from "@/components/MultipleSelect";
import Tags from "@/components/Tags";

export default {
  components: {
    BaseDropdown,
    ChoicesSingle,
    ChoicesMultiple,
    Tags
  },
  data() {
    return {
      multiLevel: false,
      multiLevel2: false,
      multiLevel3: false,
      selected: 2,
      options: [
        { id: 1, text: "Foobar" },
        { id: 2, text: "Is Great" },
        { id: 3, text: "Pick One" }
      ],
      optionsMultiple: [
        { id: 1, text: "Bucharest" },
        { id: 2, text: "Los Santos" },
        { id: 3, text: "Miami" },
        { id: 4, text: "New York" },
        { id: 5, text: "Paris" },
        { id: 6, text: "Piatra Neamt" },
        { id: 7, text: "Rome" },
        { id: 8, text: "Sydney" }
      ],
      tags: [
        { id: 1, text: "Amsterdam" },
        { id: 2, text: "Sydney" },
        { id: 3, text: "Rome" },
        { id: 4, text: "Singapore" }
      ]
    };
  },
  methods: {
    toggleMultiLevel() {
      this.multiLevel = !this.multiLevel;
    },
    toggleMultiLevel2() {
      this.multiLevel2 = !this.multiLevel2;
      this.multiLevel3 = false;
    },
    toggleMultiLevel3() {
      this.multiLevel3 = !this.multiLevel3;
      this.multiLevel2 = false;
    }
  }
};
</script>
<style>
.choices__list--single {
  padding-top: 0;
}

li.dropdown {
  z-index: 2;
}
</style>
