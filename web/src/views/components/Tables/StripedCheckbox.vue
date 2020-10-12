<template>
  <div>
    <div class="tools float-right">
      <base-dropdown menuClasses="dropdown-menu-right">
        <div
          slot="title"
          class="nav-link dropdown-toggle text-primary"
          data-toggle="dropdown"
        >
          <i class="ni ni-settings-gear-65"></i>
        </div>
        <a class="dropdown-item" href="javascript:;">Edit Profile</a>
        <a class="dropdown-item" href="javascript:;">Settings</a>
        <a class="dropdown-item" href="javascript:;">Log out</a>
        <a class="dropdown-item text-danger" href="javascript:;">Remove data</a>
      </base-dropdown>
    </div>
    <el-table
      class="table table-striped table-flush align-items-center mb-0"
      header-row-class-name="text-primary"
      :data="projects"
    >
      <el-table-column
        label="#"
        min-width="20px"
        prop="active"
        sortable
        scope="row"
      >
        <template>
          <div class="media align-items-center">
            <base-checkbox></base-checkbox>
          </div>
        </template>
      </el-table-column>
      <el-table-column
        label="Product name"
        min-width="60px"
        prop="name"
        sortable
      >
        <template v-slot="{ row }">
          <div class="media align-items-center">
            {{ row.name }}
          </div>
        </template>
      </el-table-column>
      <el-table-column label="Type" min-width="60px" prop="type" sortable>
        <template v-slot="{ row }">
          <div class="media align-items-center">
            {{ row.type }}
          </div>
        </template>
      </el-table-column>
      <el-table-column
        label="Milestone"
        min-width="60px"
        prop="communication"
        sortable
      >
        <template v-slot="{ row }">
          <div class="media align-items-center">
            <div class="progress-container">
              <div class="progress-badge">
                {{ row.version }}
              </div>
              <base-progress
                type="primary"
                :value="row.completion"
              ></base-progress>
            </div>
          </div>
        </template>
      </el-table-column>
      <el-table-column label="Qty" min-width="60px" prop="quantity" sortable>
        <template v-slot="{ row }">
          <div class="media align-items-center">
            {{ row.qty }}
          </div>
        </template>
      </el-table-column>
      <el-table-column label="Price" min-width="60px" prop="price" sortable>
        <template v-slot="{ row }">
          <div class="media align-items-center">
            {{ row.price }}
          </div>
        </template>
      </el-table-column>
      <el-table-column label="Amount" min-width="60px" prop="amount" sortable>
        <template v-slot="{ row }">
          <div class="media align-items-center">
            {{ row.amount }}
          </div>
        </template>
      </el-table-column>
    </el-table>
    <div class="table">
      <td width="300px"></td>
      <td width="300px"></td>
      <td width="300px"></td>
      <td class="td-total">Total</td>
      <td class="td-price">€ 35,999</td>
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
import BaseDropdown from "@/components/BaseDropdown";

export default {
  name: "light-table",
  components: {
    [Table.name]: Table,
    [TableColumn.name]: TableColumn,
    [Dropdown.name]: Dropdown,
    [DropdownItem.name]: DropdownItem,
    [DropdownMenu.name]: DropdownMenu,
    BaseDropdown
  },
  data() {
    return {
      projects: [
        {
          name: "Andrew Mike",
          type: "Office",
          salary: "€ 99,225",
          active: false,
          completion: 20,
          version: "v1.2.0",
          qty: 25,
          price: "€ 49",
          amount: "€ 1,225"
        },
        {
          name: "John Doe",
          type: "Office",
          salary: "€ 89,241",
          active: false,
          completion: 30,
          version: "v1.4.0",
          qty: 30,
          price: "€ 10",
          amount: "€ 300"
        },
        {
          name: "Alex Mike",
          type: "Office",
          salary: "€ 92,144",
          active: false,
          completion: 10,
          version: "v2.0.0",
          qty: 50,
          price: "€ 109",
          amount: "€1125"
        },
        {
          name: "Mike Monday",
          type: "Meeting",
          salary: "€ 49,990",
          active: true,
          completion: 80,
          version: "v1.5.0",
          qty: 10,
          price: "€ 4,990",
          amount: "€1125"
        },
        {
          name: "Paul dickens",
          type: "Communication",
          salary: "€ 69,201",
          active: true,
          completion: 20,
          version: "v1.0.0",
          qty: 10,
          price: "€ 5,999",
          amount: "€1125"
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
