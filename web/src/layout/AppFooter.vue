<template>
  <footer class="footer">
    <div class="container">
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
                href="https://github.com/cosmos/atlas"
                class="nav-link"
                target="_blank"
                >GitHub</a
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
