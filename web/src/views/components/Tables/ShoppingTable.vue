<template>
  <div class="table table-shopping">
    <el-table style="width: 100%" :data="productsTable">
      <el-table-column min-width="150" align="left">
        <div slot-scope="{ row }" class="img-container">
          <img :src="row.image" alt="Agenda" />
        </div>
      </el-table-column>
      <el-table-column min-width="220" label="Product" align="left">
        <div class="td-name" slot-scope="{ row }">
          <a href="#jacket">{{ row.title }}</a>
          <br />
          <small>{{ row.description }}</small>
        </div>
      </el-table-column>
      <el-table-column
        min-width="80"
        label="Color"
        prop="color"
        align="left"
      ></el-table-column>
      <el-table-column
        min-width="60"
        label="Size"
        prop="size"
        align="left"
      ></el-table-column>
      <el-table-column min-width="180" label="Price" header-align="right">
        <div slot-scope="{ row }" class="td-number">
          <small>€</small>
          {{ row.price }}
        </div>
      </el-table-column>
      <el-table-column min-width="180" label="Quantity" header-align="right">
        <div slot-scope="{ row }" class="td-number">
          {{ row.quantity }}
          <div class="btn-group">
            <base-button
              type="info"
              size="sm"
              @click.native="decreaseQuantity(row)"
            >
              <i class="ni ni-fat-delete"></i>
            </base-button>
            <base-button
              type="info"
              size="sm"
              @click.native="increaseQuantity(row)"
            >
              <i class="ni ni-fat-add"></i>
            </base-button>
          </div>
        </div>
      </el-table-column>
      <el-table-column min-width="170" label="Amount" header-align="right">
        <div slot-scope="{ row }" class="td-number">
          {{ row.amount }}
        </div>
      </el-table-column>
      <el-table-column min-width="100" label="">
        <div class="td-actions">
          <base-button type="" link class="text-danger">
            <i class="ni ni-fat-remove"></i>
          </base-button>
        </div>
      </el-table-column>
    </el-table>
    <div class="table table-stats">
      <div class="td-total">
        Total
      </div>
      <div class="td-price">
        <small>€</small>
        {{ shoppingTotal }}
      </div>
      <div class="text-right">
        <button
          type="button"
          rel="tooltip"
          class="btn btn-info btn-round "
          data-original-title=""
          title=""
        >
          Complete Purchase
        </button>
      </div>
    </div>
  </div>
</template>
<script>
import { Table, TableColumn } from "element-ui";

export default {
  components: {
    [Table.name]: Table,
    [TableColumn.name]: TableColumn
  },
  data() {
    return {
      productsTable: [
        {
          image: "img/jacket.png",
          title: "Monaco bees natté jacket",
          description: "by Gucci",
          color: "Black",
          size: "M",
          price: 3390,
          quantity: 1,
          amount: 3390
        },
        {
          image: "img/boots.png",
          title: "Patent-leather ankle boots",
          description: "by Prada",
          color: "Black",
          size: "41",
          price: 499,
          quantity: 2,
          amount: 998
        },
        {
          image: "img/sweater.png",
          title: "Ophidia GG",
          description: "by Saint Laurent",
          color: "Red",
          size: "M",
          price: 200,
          quantity: 1,
          amount: 200
        }
      ]
    };
  },
  computed: {
    shoppingTotal() {
      return this.productsTable.reduce((accumulator, current) => {
        return accumulator + current.amount;
      }, 0);
    }
  },
  methods: {
    increaseQuantity(row) {
      row.quantity++;
      this.computeAmount(row);
    },
    decreaseQuantity(row) {
      if (row.quantity > 1) {
        row.quantity--;
        this.computeAmount(row);
      }
    },
    computeAmount(row) {
      row.amount = row.quantity * row.price;
    }
  }
};
</script>
<style>
.table-stats {
  display: flex;
  align-items: center;
  justify-content: space-between;
  flex-wrap: wrap;
}
</style>
