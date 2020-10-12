<template>
  <div class="section section-components">
    <div class="container">
      <h3 class="h4 text-success font-weight-bold mb-4">Tabs</h3>
      <div class="row jutify-content-center">
        <div class="col-lg-6">
          <div class="mb-3">
            <small class="text-uppercase font-weight-bold">With icons</small>
          </div>
          <tabs pills fill>
            <tab-pane active>
              <template slot="label">
                <i class="ni ni-cloud-upload-96 mr-2"></i> Home
              </template>
              <p>
                I think that’s a responsibility that I have, to push
                possibilities, to show people, this is the level that things
                could be at. So when you get something that has the name Kanye
                West on it, it’s supposed to be pushing the furthest
                possibilities. I will be the leader of a company that ends up
                being worth billions of dollars, because I got the answers. I
                understand culture. I am the nucleus.
              </p>
            </tab-pane>
            <tab-pane>
              <template slot="label">
                <i class="ni ni-bell-55 mr-2"></i> Profile
              </template>
              <p>
                I will be the leader of a company that ends up being worth
                billions of dollars, because I got the answers. I understand
                culture. I am the nucleus. I think that’s a responsibility that
                I have, to push possibilities, to show people, this is the level
                that things could be at. I think that’s a responsibility that I
                have, to push possibilities, to show people, this is the level
                that things could be at.
              </p>
            </tab-pane>
            <tab-pane>
              <template slot="label">
                <i class="ni ni-calendar-grid-58 mr-2"></i> Messages
              </template>
              <p>
                I think that’s a responsibility that I have, to push
                possibilities, to show people, this is the level that things
                could be at. So when you get something that has the name Kanye
                West on it, it’s supposed to be pushing the furthest
                possibilities. I will be the leader of a company that ends up
                being worth billions of dollars, because I got the answers. I
                understand culture. I am the nucleus.
              </p>
            </tab-pane>
          </tabs>
        </div>
        <div class="col-lg-6 mt-5 mt-lg-0">
          <div class="mb-3">
            <small class="text-uppercase font-weight-bold">With text</small>
          </div>
          <tabs pills fill>
            <tab-pane active>
              <template slot="label">
                Home
              </template>
              <p>
                I think that’s a responsibility that I have, to push
                possibilities, to show people, this is the level that things
                could be at. So when you get something that has the name Kanye
                West on it, it’s supposed to be pushing the furthest
                possibilities. I will be the leader of a company that ends up
                being worth billions of dollars, because I got the answers. I
                understand culture. I am the nucleus.
              </p>
            </tab-pane>
            <tab-pane>
              <template slot="label">
                Profile
              </template>
              <p>
                I will be the leader of a company that ends up being worth
                billions of dollars, because I got the answers. I understand
                culture. I am the nucleus. I think that’s a responsibility that
                I have, to push possibilities, to show people, this is the level
                that things could be at. I think that’s a responsibility that I
                have, to push possibilities, to show people, this is the level
                that things could be at.
              </p>
            </tab-pane>
            <tab-pane>
              <template slot="label">
                Messages
              </template>
              <p>
                I think that’s a responsibility that I have, to push
                possibilities, to show people, this is the level that things
                could be at. So when you get something that has the name Kanye
                West on it, it’s supposed to be pushing the furthest
                possibilities. I will be the leader of a company that ends up
                being worth billions of dollars, because I got the answers. I
                understand culture. I am the nucleus.
              </p>
            </tab-pane>
          </tabs>
        </div>
      </div>
      <div class="section section-pills">
        <div class="container">
          <div id="navigation-pills">
            <div class="title">
              <h4 class="h4 text-success font-weight-bold mb-4">
                Navigation Pills
              </h4>
            </div>
            <div class="row">
              <div class="col-md-6">
                <div class="mb-3">
                  <small class="text-uppercase font-weight-bold"
                    >Horizontal tabs</small
                  >
                </div>
                <tabs pills type="primary">
                  <tab-pane label="Profile">
                    <p>
                      Collaboratively administrate empowered markets via
                      plug-and-play networks. Dynamically procrastinate B2C
                      users after installed base benefits.
                    </p>
                    <br />
                    <p>
                      Dramatically visualize customer directed convergence
                      without revolutionary ROI.
                    </p>
                  </tab-pane>
                  <tab-pane label="Settings">
                    <p>
                      Efficiently unleash cross-media information without
                      cross-media value. Quickly maximize timely deliverables
                      for real-time schemas.
                    </p>
                    <br />
                    <p>
                      Dramatically maintain clicks-and-mortar solutions without
                      functional solutions.
                    </p>
                  </tab-pane>
                  <tab-pane label="Options">
                    <p>
                      Completely synergize resource taxing relationships via
                      premier niche markets. Professionally cultivate one-to-one
                      customer service with robust ideas.
                    </p>
                    <br />
                    <p>
                      Dynamically innovate resource-leveling customer service
                      for state of the art customer service.
                    </p>
                  </tab-pane>
                </tabs>
              </div>
              <div class="col-md-6">
                <div class="mb-3">
                  <small class="text-uppercase font-weight-bold"
                    >Vertical tabs</small
                  >
                </div>
                <tabs fill pills type="primary" vertical class="row">
                  <tab-pane label="Profile">
                    <p>
                      Collaboratively administrate empowered markets via
                      plug-and-play networks. Dynamically procrastinate B2C
                      users after installed base benefits.
                    </p>
                    <br />
                    <p>
                      Dramatically visualize customer directed convergence
                      without revolutionary ROI.
                    </p>
                  </tab-pane>
                  <tab-pane label="Settings">
                    <p>
                      Efficiently unleash cross-media information without
                      cross-media value. Quickly maximize timely deliverables
                      for real-time schemas.
                    </p>
                    <br />
                    <p>
                      Dramatically maintain clicks-and-mortar solutions without
                      functional solutions.
                    </p>
                  </tab-pane>
                  <tab-pane label="Options">
                    <p>
                      Completely synergize resource taxing relationships via
                      premier niche markets. Professionally cultivate one-to-one
                      customer service with robust ideas.
                    </p>

                    <br />
                    <p>
                      Dynamically innovate resource-leveling customer service
                      for state of the art customer service.
                    </p>
                  </tab-pane>
                </tabs>
              </div>
            </div>

            <div class="row">
              <div class="col-md-12">
                <div class="mb-3">
                  <small class="text-uppercase font-weight-bold"
                    >With icons</small
                  >
                </div>
              </div>
              <div class="col-md-6">
                <tabs pills fill type="info">
                  <tab-pane>
                    <template slot="label">
                      <i class="ni ni-cloud-upload-96 mr-2"></i> Home
                    </template>
                    <p>
                      Collaboratively administrate empowered markets via
                      plug-and-play networks. Dynamically procrastinate B2C
                      users after installed base benefits.
                    </p>
                    <br />
                    <p>
                      Dramatically visualize customer directed convergence
                      without revolutionary ROI.
                    </p>
                  </tab-pane>
                  <tab-pane>
                    <template slot="label">
                      <i class="ni ni-bell-55 mr-2"></i> Profile
                    </template>
                    <p>
                      Efficiently unleash cross-media information without
                      cross-media value. Quickly maximize timely deliverables
                      for real-time schemas.
                    </p>
                    <br />
                    <p>
                      Dramatically maintain clicks-and-mortar solutions without
                      functional solutions.
                    </p>
                  </tab-pane>
                  <tab-pane>
                    <template slot="label">
                      <i class="ni ni-calendar-grid-58 mr-2"></i> Messages
                    </template>
                    <p>
                      Completely synergize resource taxing relationships via
                      premier niche markets. Professionally cultivate one-to-one
                      customer service with robust ideas.
                    </p>
                    <br />
                    <p>
                      Dynamically innovate resource-leveling customer service
                      for state of the art customer service.
                    </p>
                  </tab-pane>
                </tabs>
              </div>
              <div class="col-md-6">
                <tabs fill pills icons type="danger" vertical class="row">
                  <tab-pane>
                    <template slot="label">
                      <i class="ni ni-calendar-grid-58 mr-2"></i>
                      Messages
                    </template>
                    <p>
                      Collaboratively administrate empowered markets via
                      plug-and-play networks. Dynamically procrastinate B2C
                      users after installed base benefits.
                    </p>
                    <br />
                    <p>
                      Dramatically visualize customer directed convergence
                      without revolutionary ROI.
                    </p>
                  </tab-pane>
                  <tab-pane label="Settings">
                    <template slot="label">
                      <i class="ni ni-settings-gear-65"></i>
                      Settings
                    </template>
                    <p>
                      Efficiently unleash cross-media information without
                      cross-media value. Quickly maximize timely deliverables
                      for real-time schemas.
                    </p>
                    <br />
                    <p>
                      Dramatically maintain clicks-and-mortar solutions without
                      functional solutions.
                    </p>
                  </tab-pane>
                </tabs>
              </div>
            </div>

            <div class="row">
              <div class="col-md-12">
                <div class="mb-3">
                  <small class="text-uppercase font-weight-bold"
                    >Just icons</small
                  >
                </div>
              </div>
              <div class="col-md-6">
                <tabs pills circle type="warning" vertical class="row">
                  <tab-pane>
                    <template slot="label">
                      <span class="nav-link-icon d-block"
                        ><i class="ni ni-atom"></i
                      ></span>
                    </template>
                    <p>
                      Collaboratively administrate empowered markets via
                      plug-and-play networks. Dynamically procrastinate B2C
                      users after installed base benefits.
                    </p>
                    <br />
                    <p>
                      Dramatically visualize customer directed convergence
                      without revolutionary ROI.
                    </p>
                  </tab-pane>
                  <tab-pane label="Settings">
                    <template slot="label">
                      <i class="ni ni-settings-gear-65"></i>
                    </template>
                    <p>
                      Efficiently unleash cross-media information without
                      cross-media value. Quickly maximize timely deliverables
                      for real-time schemas.
                    </p>
                    <br />
                    <p>
                      Dramatically maintain clicks-and-mortar solutions without
                      functional solutions.
                    </p>
                  </tab-pane>
                </tabs>
              </div>
              <div class="col-md-6">
                <tabs pills circle type="success">
                  <tab-pane>
                    <template slot="label">
                      <i class="ni ni-cloud-upload-96"></i>
                    </template>
                    <p>
                      Collaboratively administrate empowered markets via
                      plug-and-play networks. Dynamically procrastinate B2C
                      users after installed base benefits.
                    </p>
                    <br />
                    <p>
                      Dramatically visualize customer directed convergence
                      without revolutionary ROI.
                    </p>
                  </tab-pane>
                  <tab-pane>
                    <template slot="label">
                      <i class="ni ni-bell-55"></i>
                    </template>
                    <p>
                      Efficiently unleash cross-media information without
                      cross-media value. Quickly maximize timely deliverables
                      for real-time schemas.
                    </p>
                    <br />
                    <p>
                      Dramatically maintain clicks-and-mortar solutions without
                      functional solutions.
                    </p>
                  </tab-pane>
                  <tab-pane>
                    <template slot="label">
                      <i class="ni ni-calendar-grid-58"></i>
                    </template>
                    <p>
                      Completely synergize resource taxing relationships via
                      premier niche markets. Professionally cultivate one-to-one
                      customer service with robust ideas.
                    </p>
                    <br />
                    <p>
                      Dynamically innovate resource-leveling customer service
                      for state of the art customer service.
                    </p>
                  </tab-pane>
                </tabs>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import Tabs from "@/components/Tabs/Tabs.vue";
import TabPane from "@/components/Tabs/TabPane.vue";
export default {
  components: {
    Tabs,
    TabPane
  }
};
</script>
<style></style>
