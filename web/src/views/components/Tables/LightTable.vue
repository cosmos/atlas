<template>
  <div>
    <el-table
      class="table table-responsive table-flush align-items-center"
      header-row-class-name="thead-light"
      :data="projects"
    >
      <el-table-column
        label="Project"
        min-width="310px"
        prop="name"
        sortable
        scope="row"
      >
        <template v-slot="{ row }">
          <div class="media align-items-center">
            <a href="javascript:void(0)" class="avatar rounded-circle mr-3">
              <img alt="Image placeholder" :src="row.img" />
            </a>
            <div class="media-body">
              <span class="name mb-0 text-sm">{{ row.title }}</span>
            </div>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="Budget" prop="budget" min-width="140px" sortable>
      </el-table-column>
      <el-table-column label="Status" min-width="170px" prop="status" sortable>
        <template v-slot="{ row }">
          <badge class="badge-dot mr-4" type="">
            <i :class="`bg-${row.statusType}`"></i>
            <span class="status">{{ row.status }}</span>
          </badge>
        </template>
      </el-table-column>
      <el-table-column label="Users" min-width="190px">
        <div class="avatar-group">
          <a
            href="javascript:void(0)"
            class="avatar avatar-sm rounded-circle"
            data-toggle="tooltip"
            data-original-title="Ryan Tompson"
          >
            <img alt="Image placeholder" src="img/faces/team-1.jpg" />
          </a>
          <a
            href="javascript:void(0)"
            class="avatar avatar-sm rounded-circle"
            data-toggle="tooltip"
            data-original-title="Romina Hadid"
          >
            <img alt="Image placeholder" src="img/faces/team-2.jpg" />
          </a>
          <a
            href="javascript:void(0)"
            class="avatar avatar-sm rounded-circle"
            data-toggle="tooltip"
            data-original-title="Alexander Smith"
          >
            <img alt="Image placeholder" src="img/faces/team-3.jpg" />
          </a>
          <a
            href="javascript:void(0)"
            class="avatar avatar-sm rounded-circle"
            data-toggle="tooltip"
            data-original-title="Jessica Doe"
          >
            <img alt="Image placeholder" src="img/faces/team-4.jpg" />
          </a>
        </div>
      </el-table-column>
      <el-table-column
        label="Completion"
        prop="completion"
        min-width="240px"
        sortable
      >
        <template v-slot="{ row }">
          <div class="d-flex align-items-center">
            <span class="completion mr-2">{{ row.completion }}%</span>
            <div>
              <base-progress
                :type="row.statusType"
                :value="row.completion"
              ></base-progress>
            </div>
          </div>
        </template>
      </el-table-column>
      <el-table-column min-width="180px">
        <template>
          <el-dropdown trigger="click" class="dropdown">
            <span class="btn btn-sm btn-icon-only text-light">
              <i class="fa fa-ellipsis-v mt-2"></i>
            </span>
            <el-dropdown-menu
              class="dropdown-menu dropdown-menu-arrow show"
              slot="dropdown"
            >
              <a class="dropdown-item" href="javascript:void(0)">Action</a>
              <a class="dropdown-item" href="javascript:void(0)"
                >Another action</a
              >
              <a class="dropdown-item" href="javascript:void(0)"
                >Something else here</a
              >
            </el-dropdown-menu>
          </el-dropdown>
        </template>
      </el-table-column>
    </el-table>

    <div class="card-footer bg-white py-4 d-flex justify-content-end">
      <base-pagination v-model="currentPage" :total="50"></base-pagination>
    </div>
  </div>
</template>
<script>
import {
  Table,
  TableColumn,
  DropdownMenu,
  DropdownItem,
  Dropdown
} from "element-ui";

export default {
  name: "light-table",
  components: {
    [Table.name]: Table,
    [TableColumn.name]: TableColumn,
    [Dropdown.name]: Dropdown,
    [DropdownItem.name]: DropdownItem,
    [DropdownMenu.name]: DropdownMenu
  },
  data() {
    return {
      projects: [
        {
          img: "img/theme/bootstrap.jpg",
          title: "Argon Design System",
          budget: "$2500 USD",
          status: "pending",
          statusType: "warning",
          completion: 60
        },
        {
          img: "img/theme/angular.jpg",
          title: "Angular Now UI Kit PRO",
          budget: "$1800 USD",
          status: "completed",
          statusType: "success",
          completion: 100
        },
        {
          img: "img/theme/sketch.jpg",
          title: "Black Dashboard",
          budget: "$3150 USD",
          status: "delayed",
          statusType: "danger",
          completion: 72
        },
        {
          img: "img/theme/react.jpg",
          title: "React Material Dashboard",
          budget: "$4400 USD",
          status: "on schedule",
          statusType: "info",
          completion: 90
        },
        {
          img: "img/theme/vue.jpg",
          title: "Vue Paper UI Kit PRO",
          budget: "$2200 USD",
          status: "completed",
          statusType: "success",
          completion: 100
        },
        {
          img: "img/theme/bootstrap.jpg",
          title: "Argon Design System",
          budget: "$2500 USD",
          status: "pending",
          statusType: "warning",
          completion: 60
        },
        {
          img: "img/theme/angular.jpg",
          title: "Angular Now UI Kit PRO",
          budget: "$1800 USD",
          status: "completed",
          statusType: "success",
          completion: 100
        },
        {
          img: "img/theme/sketch.jpg",
          title: "Black Dashboard",
          budget: "$3150 USD",
          status: "delayed",
          statusType: "danger",
          completion: 72
        },
        {
          img: "img/theme/vue.jpg",
          title: "Vue Paper UI Kit PRO",
          budget: "$2200 USD",
          status: "completed",
          statusType: "success",
          completion: 100
        }
      ],
      currentPage: 1
    };
  }
};
</script>
<style>
.index-page .table .progress-wrapper {
  padding-top: 0;
}
.table.align-items-center td {
  vertical-align: middle;
}

.el-table .hidden-columns {
  visibility: hidden;
  position: absolute;
  z-index: -1;
}
</style>
