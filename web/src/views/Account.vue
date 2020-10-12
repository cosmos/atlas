<template>
  <div class="wrapper">
    <div class="section-shaped my-0 skew-separator skew-mini">
      <div class="page-header page-header-small header-filter">
        <div
          class="page-header-image"
          style="background-image: url('img/pages/georgie.jpg');"
        ></div>
        <div class="container">
          <div class="header-body text-center mb-7">
            <div class="row justify-content-center">
              <div class="col-xl-5 col-lg-6 col-md-8 px-5">
                <h1 class="text-white">Your account</h1>
                <p class="text-lead text-white">
                  That’s the main thing people are controlled by! Thoughts -
                  their perception of themselves!
                </p>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="bg-secondary">
      <div class="container bg-white card mb-0">
        <div class="row">
          <div class="col-md-3">
            <div class="section">
              <!-- User Information -->
              <section class="text-center">
                <div
                  class="fileinput fileinput-new text-center"
                  data-provides="fileinput"
                >
                  <image-upload
                    type="avatar"
                    select-text="Add photo"
                    @change="onAvatarChange"
                    btnClasses="btn-primary btn-sm btn-round mt-3"
                  />
                </div>
                <h3 class="title mt-4">{{ query }}</h3>
              </section>
              <!-- User Information -->
              <!-- Profile Sidebar -->
              <section>
                <br />
                <tabs
                  v-model="accountTab"
                  vertical
                  tabNavWrapperClasses="col-md-12 text-center"
                  no-content-space
                >
                  <tab-pane label="General"></tab-pane>
                  <tab-pane label="Billing"></tab-pane>
                  <tab-pane label="Security"></tab-pane>
                  <tab-pane label="Notifications"></tab-pane>
                </tabs>
              </section>
              <!-- End Profile Sidebar -->
              <!-- Profile Completion -->
              <br /><br /><br />
              <div>
                <div class="progress-container progress-primary">
                  <span class="progress-badge">Profile Completion</span>
                  <base-progress :value="80" type="primary" showLabel>
                  </base-progress>
                </div>
              </div>
              <!-- End Profile Completion -->
            </div>
          </div>
          <div class="col-md-8 ml-auto">
            <div class="section">
              <div class="tab-content">
                <div v-if="accountTab === 'General'" class="tab-pane active">
                  <div>
                    <header>
                      <h2 class="text-uppercase">General information</h2>
                    </header>
                    <hr class="line-primary" />
                    <br />
                    <div class="row">
                      <div class="col-md-3 align-self-center">
                        <label class="labels" for="#firstName"
                          >First Name</label
                        >
                      </div>
                      <div class="col-md-9 align-self-center">
                        <base-input
                          id="firstName"
                          name="firstName"
                          type="text"
                          value="Charlie"
                        ></base-input>
                      </div>
                    </div>
                    <div class="row">
                      <div class="col-md-3 align-self-center">
                        <label class="labels" for="#lastName">Last Name</label>
                      </div>
                      <div class="col-md-9 align-self-center">
                        <base-input
                          id="lasttName"
                          name="lastName"
                          type="text"
                          value="Bailey"
                        ></base-input>
                      </div>
                    </div>
                    <div class="row">
                      <div class="col-md-3 align-self-center">
                        <label class="labels">I’m</label>
                      </div>
                      <div class="col-md-4 align-self-center">
                        <div class="form-group">
                          <choices-single :options="gender" v-model="selected">
                            <option disabled value="0">Select one</option>
                          </choices-single>
                        </div>
                      </div>
                    </div>
                    <div class="row">
                      <div class="col-md-3 align-self-center">
                        <label class="labels">Birth Date</label>
                      </div>
                      <div class="col-md-9 align-self-center">
                        <div class="row">
                          <div class="col-md-4 align-self-center">
                            <div class="form-group">
                              <choices-single
                                :options="months"
                                v-model="selected"
                              >
                                <option disabled value="0">Select one</option>
                              </choices-single>
                            </div>
                          </div>
                          <div class="col-md-4">
                            <div class="form-group">
                              <choices-single
                                :options="days"
                                v-model="selected"
                              >
                                <option disabled value="0">Select one</option>
                              </choices-single>
                            </div>
                          </div>
                          <div class="col-md-4">
                            <div class="form-group">
                              <choices-single
                                :options="years"
                                v-model="selected"
                              >
                                <option disabled value="0">Select one</option>
                              </choices-single>
                            </div>
                          </div>
                        </div>
                      </div>
                    </div>
                    <div class="row">
                      <div class="col-md-3 align-self-center">
                        <label class="labels" for="#email">Email</label>
                      </div>
                      <div class="col-md-9 align-self-center">
                        <base-input
                          id="email"
                          name="email"
                          type="text"
                          value="charlie.bailey@example.com"
                        ></base-input>
                      </div>
                    </div>
                    <div class="row">
                      <div class="col-md-3 align-self-center">
                        <label class="labels" for="#confirmEmail"
                          >Confirm Email</label
                        >
                      </div>
                      <div class="col-md-9 align-self-center">
                        <base-input
                          id="confirmEmail"
                          name="confirmEmail"
                          type="email"
                          value="charlie.bailey@example.com"
                        ></base-input>
                      </div>
                    </div>
                    <div class="row">
                      <div class="col-md-3 align-self-center">
                        <label class="labels" for="#location"
                          >Your Location</label
                        >
                      </div>
                      <div class="col-md-9 align-self-center">
                        <base-input
                          id="location"
                          name="location"
                          type="text"
                          value="Sydney, A"
                        ></base-input>
                      </div>
                    </div>
                    <div class="row">
                      <div class="col-md-3 align-self-center">
                        <label class="labels" for="#phone">Phone Number</label>
                      </div>
                      <div class="col-md-4 align-self-center">
                        <base-input
                          id="phone"
                          name="phone"
                          type="tel"
                          value="+40 745 031 200"
                        ></base-input>
                      </div>
                    </div>
                    <div class="row">
                      <div class="col-md-3 align-self-center">
                        <label class="labels">Language</label>
                      </div>
                      <div class="col-md-4 align-self-center">
                        <div class="form-group">
                          <choices-single
                            :options="languages"
                            v-model="selected"
                          >
                            <option disabled value="0">Select one</option>
                          </choices-single>
                        </div>
                      </div>
                    </div>
                    <div class="row">
                      <div class="col-md-3 align-self-center">
                        <label class="labels">Skills</label>
                      </div>
                      <div class="col-md-9 align-self-center">
                        <tags value="VueJS, Angular, Laravel, React"></tags>
                      </div>
                    </div>
                    <div class="row mt-5">
                      <div class="col-md-6">
                        <base-button nativeType="submit" type="primary"
                          >Save Changes</base-button
                        >
                        <base-button nativeType="reset" outline type="primary"
                          >Cancel</base-button
                        >
                      </div>
                    </div>
                  </div>
                </div>

                <div v-if="accountTab === 'Billing'" class="tab-pane active">
                  <header>
                    <h2 class="text-uppercase">Billing method</h2>
                  </header>
                  <hr class="line-primary" />
                  <br />
                  <table class="table align-items-center">
                    <thead>
                      <tr>
                        <th scope="col">Card Type</th>
                        <th scope="col">Card Number</th>
                        <th scope="col">Payment Method</th>
                        <th scope="col">Action</th>
                      </tr>
                    </thead>
                    <tbody>
                      <tr>
                        <th scope="row">
                          <img alt="Image" src="img/visas.png" class="avatar" />
                        </th>
                        <td>
                          <span class="d-block">•••• •••• •••• 8372</span>
                          <small class="text-muted">Exp: 06/22</small>
                        </td>
                        <td class="text-center">
                          <base-radio id="Radios1" checked></base-radio>
                        </td>
                        <td>
                          <base-button type="danger" simple size="sm">
                            Remove card
                          </base-button>
                        </td>
                      </tr>
                      <tr>
                        <th scope="row">
                          <img
                            alt="Image"
                            src="img/mastercard.png"
                            class="avatar mastercard"
                          />
                        </th>
                        <td>
                          <span class="d-block">•••• •••• •••• 1225</span>
                          <small class="text-muted">Exp: 07/21</small>
                        </td>
                        <td class="text-center">
                          <base-radio id="Radios1"></base-radio>
                        </td>
                        <td>
                          <base-button type="danger" simple size="sm">
                            Remove card
                          </base-button>
                        </td>
                      </tr>
                    </tbody>
                  </table>
                  <base-button type="primary" size="sm">
                    Add card
                  </base-button>
                </div>

                <div v-if="accountTab === 'Security'" class="tab-pane active">
                  <header>
                    <h2 class="text-uppercase mb-0">Security Questions</h2>
                  </header>
                  <hr class="line-primary" />
                  <form>
                    <div class="row">
                      <div class="col-md-6">
                        <label>Security Question</label>
                        <div class="form-group">
                          <choices-single
                            :options="questions"
                            v-model="selected"
                          >
                            <option disabled value="0">Select one</option>
                          </choices-single>
                        </div>
                      </div>
                      <div class="col-md-6">
                        <label>Your Answer</label>
                        <base-input
                          placeholder="Enter your answer"
                          type="text"
                        ></base-input>
                      </div>
                    </div>
                    <br />
                    <br />
                    <header>
                      <h2 class="text-uppercase">Security Settings</h2>
                    </header>
                    <hr class="line-primary" />
                    <div
                      class="form-group d-flex align-items-center justify-content-between"
                    >
                      <span>Notify me via email when logging in</span>
                      <base-switch
                        v-model="switches.off"
                        type="primary"
                      ></base-switch>
                    </div>
                    <div
                      class="form-group d-flex align-items-center justify-content-between"
                    >
                      <span>Send SMS confirmation for all online payments</span>
                      <base-switch
                        v-model="switches.on"
                        type="primary"
                      ></base-switch>
                    </div>
                    <div
                      class="form-group d-flex align-items-center justify-content-between"
                    >
                      <span>Check which devices accessed your account</span>
                      <base-switch
                        v-model="switches.off"
                        type="primary"
                      ></base-switch>
                    </div>
                    <div
                      class="form-group d-flex align-items-center justify-content-between"
                    >
                      <span
                        >Find My Device, make sure your device can be found if
                        it gets lost</span
                      >
                      <base-switch
                        v-model="switches.on"
                        type="primary"
                      ></base-switch>
                    </div>
                    <div
                      class="form-group d-flex align-items-center justify-content-between"
                    >
                      <span
                        >Lock your device with a PIN, pattern, or password</span
                      >
                      <base-switch
                        v-model="switches.on"
                        type="primary"
                      ></base-switch>
                    </div>
                    <div
                      class="form-group d-flex align-items-center justify-content-between"
                    >
                      <span
                        >Manage what apps have access to app-usage data on your
                        device</span
                      >
                      <base-switch
                        v-model="switches.off"
                        type="primary"
                      ></base-switch>
                    </div>
                    <div class="row mt-5 justify-content-end">
                      <div class="col-md-4 actions">
                        <base-button
                          type="primary"
                          outline
                          size="sm"
                          nativeType="reset"
                          >Cancel</base-button
                        >
                        <base-button type="primary" size="sm"
                          >Save Changes</base-button
                        >
                      </div>
                    </div>
                  </form>
                </div>

                <div
                  v-if="accountTab === 'Notifications'"
                  class="tab-pane active"
                >
                  <div class="container">
                    <div class="row">
                      <div class="col-12">
                        <base-alert type="primary">
                          We will never distribute your email address to third
                          parties. Read about email communication in our privacy
                          policy.
                        </base-alert>
                      </div>
                    </div>
                    <hr />
                    <div class="row">
                      <div class="col-12">
                        <form>
                          <h5 class="mb-4">Notification Preferences</h5>
                          <base-checkbox class="mb-3" checked
                            >Someone mentions me</base-checkbox
                          >
                          <base-checkbox class="mb-3" checked
                            >Someone follows me</base-checkbox
                          >
                          <base-checkbox class="mb-3"
                            >Someone shares my activty</base-checkbox
                          >
                          <base-checkbox class="mb-3"
                            >Someone messages me</base-checkbox
                          >
                          <base-checkbox class="mb-3"
                            >Someone adds me to a project</base-checkbox
                          >
                          <base-checkbox class="mb-3"
                            >Sales and promotions</base-checkbox
                          >
                          <base-button
                            nativeType="submit"
                            type="primary"
                            size="sm"
                            class="mt-4 mb-5"
                            >Update preferences</base-button
                          >
                        </form>
                      </div>
                    </div>
                    <hr />
                    <div class="row">
                      <div class="col-12">
                        <form>
                          <h5>Notification Frequency</h5>
                          <base-radio
                            name="radio0"
                            class="mb-3"
                            v-model="radio.radio1"
                            >Daily</base-radio
                          >
                          <base-radio
                            name="radio1"
                            class="mb-3"
                            v-model="radio.radio1"
                            >Weekly</base-radio
                          >
                          <base-radio
                            name="radio2"
                            class="mb-3"
                            v-model="radio.radio1"
                            >Monthly</base-radio
                          >
                          <base-radio
                            name="radio3"
                            class="mb-3"
                            v-model="radio.radio1"
                            >Never</base-radio
                          >
                        </form>
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import ImageUpload from "@/components/ImageUpload";
import Tabs from "@/components/Tabs/Tabs.vue";
import TabPane from "@/components/Tabs/TabPane.vue";
import ChoicesSingle from "@/components/SingleSelect";
import Tags from "@/components/Tags";
import BaseAlert from "@/components/BaseAlert";

export default {
  bodyClass: "account-settings",
  components: {
    ImageUpload,
    Tabs,
    TabPane,
    ChoicesSingle,
    Tags,
    BaseAlert
  },
  methods: {
    onImageChange(file) {
      this.images.regular = file;
    },
    onAvatarChange(file) {
      this.images.avatar = file;
    }
  },
  created() {
    this.query = this.$route.query.q;
  },
  data() {
    return {
      query: "",
      accountTab: "General",
      selected: 2,
      gender: [
        { id: 1, text: "Male" },
        { id: 2, text: "Female" }
      ],
      months: [
        { id: 1, text: "January" },
        { id: 2, text: "February" },
        { id: 3, text: "March" },
        { id: 4, text: "April" },
        { id: 5, text: "May" },
        { id: 6, text: "June" },
        { id: 7, text: "July" },
        { id: 8, text: "August" },
        { id: 9, text: "September" },
        { id: 10, text: "October" },
        { id: 11, text: "November" },
        { id: 12, text: "December" }
      ],
      days: [
        { id: 1, text: "1" },
        { id: 2, text: "2" },
        { id: 3, text: "3" },
        { id: 4, text: "4" },
        { id: 5, text: "5" },
        { id: 6, text: "6" },
        { id: 7, text: "7" },
        { id: 8, text: "8" },
        { id: 9, text: "9" },
        { id: 10, text: "10" },
        { id: 11, text: "11" },
        { id: 12, text: "12" },
        { id: 13, text: "13" },
        { id: 14, text: "14" },
        { id: 15, text: "15" },
        { id: 16, text: "16" },
        { id: 17, text: "17" },
        { id: 18, text: "18" },
        { id: 19, text: "19" },
        { id: 20, text: "20" },
        { id: 21, text: "21" },
        { id: 22, text: "22" },
        { id: 23, text: "23" },
        { id: 24, text: "24" },
        { id: 25, text: "25" },
        { id: 26, text: "26" },
        { id: 27, text: "27" },
        { id: 28, text: "28" },
        { id: 29, text: "29" },
        { id: 30, text: "30" },
        { id: 31, text: "31" }
      ],
      years: [
        { id: 1, text: "1986" },
        { id: 2, text: "1987" },
        { id: 3, text: "1988" },
        { id: 4, text: "1989" },
        { id: 5, text: "1990" },
        { id: 6, text: "1991" },
        { id: 7, text: "1992" },
        { id: 8, text: "1993" },
        { id: 9, text: "1994" },
        { id: 10, text: "1995" },
        { id: 11, text: "1996" },
        { id: 12, text: "1997" },
        { id: 13, text: "1998" },
        { id: 14, text: "1999" },
        { id: 15, text: "2000" },
        { id: 16, text: "2001" },
        { id: 17, text: "2002" },
        { id: 18, text: "2003" },
        { id: 19, text: "2004" },
        { id: 20, text: "2005" },
        { id: 21, text: "2006" },
        { id: 22, text: "2007" },
        { id: 23, text: "2008" },
        { id: 24, text: "2009" },
        { id: 25, text: "2010" },
        { id: 26, text: "2011" },
        { id: 27, text: "2012" },
        { id: 28, text: "2013" },
        { id: 29, text: "2014" },
        { id: 30, text: "2015" },
        { id: 31, text: "2016" },
        { id: 32, text: "2017" },
        { id: 33, text: "2018" },
        { id: 34, text: "2019" },
        { id: 35, text: "2020" }
      ],
      languages: [
        { id: 1, text: "Deutsche" },
        { id: 2, text: "English" },
        { id: 3, text: "French" },
        { id: 4, text: "Russian" },
        { id: 5, text: "Spanish" }
      ],
      questions: [
        { id: 1, text: "Question 1" },
        { id: 2, text: "Question 2" },
        { id: 3, text: "Question 3" }
      ],
      switches: {
        off: false,
        on: true
      },
      radio: {
        radio1: "radio1"
      }
    };
  }
};
</script>
<style>
.account-settings .nav {
  text-align: left;
}

.account-settings .nav .nav-item {
  padding: 1rem 0;
}

.account-settings .nav .nav-item:not(:last-child) {
  border-bottom: 1px solid #5e72e4;
}
</style>
