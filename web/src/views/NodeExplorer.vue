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
      <div class="container" style="padding-top: 100px;">
        <div class="card card-profile shadow mt--300">
          <div class="px-4">
            <div class="mt-3 py-3 text-center">
              <div class="row justify-content-center">
                <div class="col-lg-12">
                  <el-table
                    class="table table-striped table-flush"
                    header-row-class-name="thead-light"
                    v-if="
                      responseData.results && responseData.results.length > 0
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
                    <!-- <el-table-column
                      label="Node ID"
                      prop="name"
                      sortable
                      scope="row"
                    >
                      <template v-slot="{ row }">
                        <div>{{ row.node_id }}</div>
                      </template>
                    </el-table-column> -->
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
                      label="Version"
                      prop="name"
                      sortable
                      scope="row"
                    >
                      <template v-slot="{ row }">
                        <div>{{ row.version }}</div>
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
                  </el-table>
                  <div
                    class="row justify-content-center"
                    style="padding: 20px 0 20px 0;"
                    v-if="
                      responseData.results && responseData.results.length > 0
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
    </section>
  </div>
</template>

<script>
import { Table, TableColumn } from "element-ui";
import APIClient from "../plugins/apiClient";

export default {
  bodyClass: "profile-page",
  components: {
    [Table.name]: Table,
    [TableColumn.name]: TableColumn
  },
  created() {
    this.getNodes();
  },
  data() {
    return {
      responseData: {},
      pageURI: "?page=1&limit=25&order=id",
      pageSize: 25
    };
  },
  beforeRouteUpdate(to, from, next) {
    next();
  },
  watch: {
    pageURI: function() {
      this.getNodes();
    }
  },
  computed: {},
  methods: {
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
</style>
