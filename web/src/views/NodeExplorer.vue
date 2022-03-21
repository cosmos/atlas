<template>
  <div class="wrapper">
    <div
      class="page-header page-header-small header-filter skew-separator skew-mini"
      style="position: absolute;"
    >
      <div class="page-header-image"></div>
    </div>

    <section
      class="section bg-secondary"
      style="position: relative; padding-top: 40vh; min-height: 100vh;"
    >
      <div class="row">
        <div class="container">
          <modal :show.sync="copyTextModal">
            <h4 style="text-align: center;">
              Copied!
            </h4>
          </modal>
          <div class="mt--300">
            <div class="px-4">
              <div class="mt-3 py-3 text-center">
                <div class="row justify-content-center">
                  <div class="col-lg-12">
                    <div class="row">
                      <div class="text-lg-left align-self-lg-left">
                        <base-button
                          nativeType="submit"
                          type="primary"
                          :disabled="displayMode === 'list'"
                          v-on:click="switchMode"
                          >List</base-button
                        >
                      </div>
                      <div class="col-lg-2 text-lg-left align-self-lg-left">
                        <base-button
                          nativeType="submit"
                          type="primary"
                          :disabled="displayMode === 'map'"
                          v-on:click="switchMode"
                          >Map</base-button
                        >
                      </div>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
      <div class="row">
        <div class="container" style="padding-top: 100px;">
          <div
            class="card card-profile shadow mt--300"
            v-bind:class="{ chart_bg: displayMode === 'map' }"
          >
            <div class="px-4">
              <div class="mt-3 py-3 text-center">
                <div class="row justify-content-center">
                  <div class="col-lg-12">
                    <div
                      class="chart"
                      ref="chartdiv"
                      v-show="displayMode === 'map'"
                    ></div>
                    <el-table
                      class="table table-striped table-flush"
                      header-row-class-name="thead-light"
                      v-if="
                        responseData.results &&
                          responseData.results.length > 0 &&
                          displayMode === 'list'
                      "
                      :data="responseData.results"
                    >
                      <el-table-column
                        label="moniker"
                        prop="name"
                        sortable
                        scope="row"
                      >
                        <template v-slot="{ row }">
                          <div>{{ row.moniker }}</div>
                        </template>
                      </el-table-column>
                      <el-table-column
                        label="address"
                        prop="name"
                        sortable
                        scope="row"
                      >
                        <template v-slot="{ row }">
                          <div>{{ row.address }}</div>
                        </template>
                      </el-table-column>
                      <el-table-column
                        label="Node ID"
                        prop="name"
                        sortable
                        scope="row"
                      >
                        <template v-slot="{ row }">
                          <div v-on:dblclick="copyToClipboard(row.node_id)">
                            {{ row.node_id }}
                          </div>
                        </template>
                      </el-table-column>
                      <el-table-column
                        label="Network"
                        prop="name"
                        sortable
                        scope="row"
                      >
                        <template v-slot="{ row }">
                          <div>{{ row.network }}</div>
                        </template>
                      </el-table-column>
                      <el-table-column
                        label="Tx Indexing"
                        prop="name"
                        sortable
                        scope="row"
                      >
                        <template v-slot="{ row }">
                          <div>{{ row.tx_index }}</div>
                        </template>
                      </el-table-column>
                      <el-table-column
                        label="Location"
                        prop="name"
                        sortable
                        scope="row"
                      >
                        <template v-slot="{ row }">
                          <div>
                            {{ row.location.city }}, {{ row.location.country }}
                          </div>
                        </template>
                      </el-table-column>
                      <el-table-column
                        label="Last Sync"
                        prop="name"
                        sortable
                        scope="row"
                      >
                        <template v-slot="{ row }">
                          <div>{{ formatDate(row.updated_at) }}</div>
                        </template>
                      </el-table-column>
                    </el-table>
                    <div
                      class="row justify-content-center"
                      style="padding: 20px 0 20px 0;"
                      v-if="
                        responseData.results &&
                          responseData.results.length > 0 &&
                          displayMode === 'list'
                      "
                    >
                      <base-button
                        class="align-self-center"
                        nativeType="submit"
                        type="primary"
                        :disabled="!this.responseData.prev_uri"
                        v-on:click="prevNodes"
                      >
                        <i class="ni ni-bold-left"></i>
                      </base-button>
                      <base-button
                        class="align-self-center"
                        nativeType="submit"
                        type="primary"
                        :disabled="!this.responseData.next_uri"
                        v-on:click="nextNodes"
                      >
                        <i class="ni ni-bold-right"></i>
                      </base-button>
                    </div>
                  </div>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </section>
  </div>
</template>

<script>
import { Table, TableColumn } from "element-ui";
import APIClient from "../plugins/apiClient";
import * as am4core from "@amcharts/amcharts4/core";
import * as am4maps from "@amcharts/amcharts4/maps";
import am4geodata_worldLow from "@amcharts/amcharts4-geodata/worldLow";
import am4themes_dark from "@amcharts/amcharts4/themes/dark";
import am4themes_animated from "@amcharts/amcharts4/themes/animated";
import Modal from "@/components/Modal.vue";

export default {
  bodyClass: "profile-page",
  components: {
    [Table.name]: Table,
    [TableColumn.name]: TableColumn,
    Modal
  },
  beforeDestroy() {
    if (this.chart) {
      this.chart.dispose();
    }
  },

  created() {
    this.getNodes();
  },

  mounted() {
    this.createChart();
    this.populateMapData("?page=1&limit=100&order=id");
  },

  data() {
    return {
      displayMode: "list",
      responseData: {},
      firstPageURI: "?page=1&limit=25&order=moniker,id&reverse=true",
      pageURI: "?page=1&limit=25&order=moniker,id&reverse=true",
      pageSize: 25,
      copyTextModal: false
    };
  },
  beforeRouteUpdate(to, from, next) {
    next();
  },
  watch: {
    pageURI: function() {
      this.getNodes();
    },

    displayMode: function() {
      this.pageURI = this.firstPageURI;
    }
  },
  computed: {},
  methods: {
    copyToClipboard(value) {
      this.$copyText(value).then(() => {
        this.copyTextModal = true;
        setTimeout(() => {
          this.copyTextModal = false;
        }, 1000);
      });
    },

    prevNodes() {
      this.pageURI = this.responseData.prev_uri;
    },

    nextNodes() {
      this.pageURI = this.responseData.next_uri;
    },

    getNodes() {
      APIClient.getNodes(this.pageURI)
        .then(resp => {
          this.responseData = resp;
        })
        .catch(err => {
          console.log(err);
          this.$notify({
            group: "errors",
            type: "error",
            duration: 3000,
            title: "Error",
            text: this.getResponseError(err)
          });
        });
    },

    populateMapData(nextURI) {
      APIClient.getNodes(nextURI)
        .then(resp => {
          if (resp.results && resp.results.length > 0) {
            resp.results.forEach(node => {
              this.chartSeries.addData({
                value: 10,
                color: "#5064fb",
                name: node.location.country,
                moniker: node.moniker,
                address: node.address,
                nodeID: node.node_id,
                latitude: parseFloat(node.location.latitude),
                longitude: parseFloat(node.location.longitude)
              });
            });

            if (resp.next_uri !== "") {
              this.populateMapData(resp.next_uri);
            }
          }
        })
        .catch(err => {
          console.log(err);
          this.$notify({
            group: "errors",
            type: "error",
            duration: 3000,
            title: "Error",
            text: this.getResponseError(err)
          });
        });
    },

    nodeIDLimited(nodeID) {
      if (nodeID.length > 0) {
        return nodeID.substring(0, 9) + "...";
      } else {
        return nodeID;
      }
    },

    switchMode() {
      if (this.displayMode === "list") {
        this.displayMode = "map";

        if (this.chart) {
          this.chartSeries.removeData(this.chartSeries.data.length);
          this.populateMapData("?page=1&limit=100&order=id");
          this.chart.show();
        }
      } else {
        this.displayMode = "list";

        if (this.chart) {
          this.chart.hide();
        }
      }
    },

    createChart() {
      am4core.useTheme(am4themes_dark);
      am4core.useTheme(am4themes_animated);

      let chart = am4core.create(this.$refs.chartdiv, am4maps.MapChart);

      chart.geodata = am4geodata_worldLow;
      chart.projection = new am4maps.projections.Miller();

      let polygonSeries = chart.series.push(new am4maps.MapPolygonSeries());

      polygonSeries.exclude = ["AQ"];
      polygonSeries.useGeodata = true;
      polygonSeries.nonScalingStroke = true;
      polygonSeries.strokeWidth = 0.5;
      polygonSeries.calculateVisualCenter = true;

      let polygonTemplate = polygonSeries.mapPolygons.template;
      polygonTemplate.tooltipText = "{name}";

      let hs = polygonTemplate.states.create("hover");
      hs.properties.fill = chart.colors.getIndex(0);

      let imageSeries = chart.series.push(new am4maps.MapImageSeries());
      imageSeries.dataFields.value = "value";

      let imageTemplate = imageSeries.mapImages.template;
      imageTemplate.nonScaling = true;
      imageTemplate.propertyFields.latitude = "latitude";
      imageTemplate.propertyFields.longitude = "longitude";

      let circle = imageTemplate.createChild(am4core.Circle);
      circle.fillOpacity = 0.7;
      circle.propertyFields.fill = "color";
      circle.tooltipText =
        "[bold]Address: {address}[/]\n[bold]Moniker: {moniker}[/]\n[bold]Node ID: {nodeID}[/]";

      imageSeries.heatRules.push({
        target: circle,
        property: "radius",
        min: 1,
        max: 15,
        dataField: "value"
      });

      this.chartSeries = imageSeries;
      this.chart = chart;
    }
  }
};
</script>

<style>
section.bg-secondary {
  background: url(/img/stars.d8924548.d8924548.svg) repeat top,
    linear-gradient(145.11deg, #202854 9.49%, #171b39 91.06%);
}

.table.align-items-center td {
  vertical-align: middle;
}

.table th,
.table td {
  border-top: none;
}

.table thead th {
  border-bottom: none;
}

.el-table .hidden-columns {
  visibility: hidden;
  position: absolute;
  z-index: -1;
}

.chart_bg {
  background: #1a1a1f;
}

.chart {
  height: 600px;
  width: 100%;
}
</style>
