<template>
  <div class="cd-section" id="accordion">
    <!--     *********     ACCORDION 1      *********      -->
    <div class="container accordion-1">
      <div class="row">
        <div class="col-md-6 mx-auto text-center">
          <h2 class="title mb-3 mt-5">Frequently asked question</h2>
        </div>
      </div>
      <div class="row">
        <div class="col-md-12">
          <collapse :multiple-active="false">
            <collapse-item name="1">
              <h5
                slot="title"
                class="mb-0 btn btn-link w-100 text-primary text-left"
              >
                How do I order? <i class="ni ni-bold-down float-right pt-1"></i>
              </h5>
              <div>
                Anim pariatur cliche reprehenderit, enim eiusmod high life
                accusamus terry richardson ad squid. 3 wolf moon officia aute,
                non cupidatat skateboard dolor brunch. Food truck quinoa
                nesciunt laborum eiusmod. Brunch 3 wolf moon tempor, sunt aliqua
                put a bird on it squid single-origin coffee nulla assumenda
                shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore
                wes anderson cred nesciunt sapiente ea proident. Ad vegan
                excepteur butcher vice lomo. Leggings occaecat craft beer
                farm-to-table, raw denim aesthetic synth nesciunt you probably
                haven't heard of them accusamus labore sustainable VHS.
              </div>
            </collapse-item>
            <collapse-item name="2">
              <h5
                slot="title"
                class="mb-0 btn btn-link w-100 text-primary text-left"
              >
                How can I make the payment?
                <i class="ni ni-bold-down float-right pt-1"></i>
              </h5>
              <div>
                Anim pariatur cliche reprehenderit, enim eiusmod high life
                accusamus terry richardson ad squid. 3 wolf moon officia aute,
                non cupidatat skateboard dolor brunch. Food truck quinoa
                nesciunt laborum eiusmod. Brunch 3 wolf moon tempor, sunt aliqua
                put a bird on it squid single-origin coffee nulla assumenda
                shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore
                wes anderson cred nesciunt sapiente ea proident. Ad vegan
                excepteur butcher vice lomo. Leggings occaecat craft beer
                farm-to-table, raw denim aesthetic synth nesciunt you probably
                haven't heard of them accusamus labore sustainable VHS.
              </div>
            </collapse-item>
            <collapse-item name="3">
              <h5
                slot="title"
                class="mb-0 btn btn-link w-100 text-primary text-left"
              >
                How much time does it take to receive the order?
                <i class="ni ni-bold-down float-right pt-1"></i>
              </h5>
              <div>
                Anim pariatur cliche reprehenderit, enim eiusmod high life
                accusamus terry richardson ad squid. 3 wolf moon officia aute,
                non cupidatat skateboard dolor brunch. Food truck quinoa
                nesciunt laborum eiusmod. Brunch 3 wolf moon tempor, sunt aliqua
                put a bird on it squid single-origin coffee nulla assumenda
                shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore
                wes anderson cred nesciunt sapiente ea proident. Ad vegan
                excepteur butcher vice lomo. Leggings occaecat craft beer
                farm-to-table, raw denim aesthetic synth nesciunt you probably
                haven't heard of them accusamus labore sustainable VHS.
              </div>
            </collapse-item>
            <collapse-item name="4">
              <h5
                slot="title"
                class="mb-0 btn btn-link w-100 text-primary text-left"
              >
                Can I resell the products?
                <i class="ni ni-bold-down float-right pt-1"></i>
              </h5>
              <div>
                Anim pariatur cliche reprehenderit, enim eiusmod high life
                accusamus terry richardson ad squid. 3 wolf moon officia aute,
                non cupidatat skateboard dolor brunch. Food truck quinoa
                nesciunt laborum eiusmod. Brunch 3 wolf moon tempor, sunt aliqua
                put a bird on it squid single-origin coffee nulla assumenda
                shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore
                wes anderson cred nesciunt sapiente ea proident. Ad vegan
                excepteur butcher vice lomo. Leggings occaecat craft beer
                farm-to-table, raw denim aesthetic synth nesciunt you probably
                haven't heard of them accusamus labore sustainable VHS.
              </div>
            </collapse-item>
            <collapse-item name="5">
              <h5
                slot="title"
                class="mb-0 btn btn-link w-100 text-primary text-left"
              >
                Where do I find the shipping details?
                <i class="ni ni-bold-down float-right pt-1"></i>
              </h5>
              <div>
                Anim pariatur cliche reprehenderit, enim eiusmod high life
                accusamus terry richardson ad squid. 3 wolf moon officia aute,
                non cupidatat skateboard dolor brunch. Food truck quinoa
                nesciunt laborum eiusmod. Brunch 3 wolf moon tempor, sunt aliqua
                put a bird on it squid single-origin coffee nulla assumenda
                shoreditch et. Nihil anim keffiyeh helvetica, craft beer labore
                wes anderson cred nesciunt sapiente ea proident. Ad vegan
                excepteur butcher vice lomo. Leggings occaecat craft beer
                farm-to-table, raw denim aesthetic synth nesciunt you probably
                haven't heard of them accusamus labore sustainable VHS.
              </div>
            </collapse-item>
          </collapse>
        </div>
      </div>
    </div>
  </div>
</template>
<script>
import Collapse from "@/components/Collapse/Collapse";
import CollapseItem from "@/components/Collapse/CollapseItem";

export default {
  components: {
    Collapse,
    CollapseItem
  }
};
</script>
<style></style>
