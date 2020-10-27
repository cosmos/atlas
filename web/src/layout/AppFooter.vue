<template>
  <footer class="footer">
    <div class="container">
      <!-- <div class="row row-grid align-items-center mb-5">
        <div class="col-lg-6">
          <h3 class="text-primary font-weight-light mb-2">
            Thank you for supporting us!
          </h3>
          <h4 class="mb-0 font-weight-light">
            Let's get in touch on any of these platforms.
          </h4>
        </div>
        <div class="col-lg-6 text-lg-center btn-wrapper">
          <button
            target="_blank"
            href="https://twitter.com/creativetim"
            rel="nofollow"
            class="btn btn-icon-only btn-twitter rounded-circle"
            data-toggle="tooltip"
            data-original-title="Follow us"
          >
            <span class="btn-inner--icon"><i class="fa fa-twitter"></i></span>
          </button>
          <button
            target="_blank"
            href="https://www.facebook.com/CreativeTim/"
            rel="nofollow"
            class="btn-icon-only rounded-circle btn btn-facebook"
            data-toggle="tooltip"
            data-original-title="Like us"
          >
            <span class="btn-inner--icon"><i class="fa fa-facebook"></i></span>
          </button>
          <button
            target="_blank"
            href="https://dribbble.com/creativetim"
            rel="nofollow"
            class="btn btn-icon-only btn-dribbble rounded-circle"
            data-toggle="tooltip"
            data-original-title="Follow us"
          >
            <span class="btn-inner--icon"><i class="fa fa-dribbble"></i></span>
          </button>
          <button
            target="_blank"
            href="https://github.com/creativetimofficial"
            rel="nofollow"
            class="btn btn-icon-only btn-github rounded-circle"
            data-toggle="tooltip"
            data-original-title="Star on Github"
          >
            <span class="btn-inner--icon"><i class="fa fa-github"></i></span>
          </button>
        </div>
      </div> -->
      <div class="row align-items-center justify-content-md-between">
        <div class="col-md-auto">
          <div class="copyright">
            &copy; {{ year }}
            <a href="" target="_self">Cosmos Atlas</a>
          </div>
        </div>
        <div class="col-md-auto">
          <ul class="nav nav-footer justify-content-end">
            <li class="nav-item">
              <a
                href="https://cosmos.network/sdk"
                class="nav-link"
                target="_blank"
                >Cosmos SDK</a
              >
            </li>
            <li class="nav-item">
              <a
                href="https://hub.cosmos.network/"
                class="nav-link"
                target="_blank"
                >Cosmos Hub</a
              >
            </li>
            <li class="nav-item">
              <a href="https://tendermint.com/" class="nav-link" target="_blank"
                >Tendermint Core</a
              >
            </li>
            <li class="nav-item">
              <a
                href="https://github.com/cosmos/ics"
                class="nav-link"
                target="_blank"
                >Interchain Standards</a
              >
            </li>
          </ul>
        </div>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  name: "app-footer",
  data() {
    return {
      year: new Date().getFullYear()
    };
  }
};
</script>

<style></style>
