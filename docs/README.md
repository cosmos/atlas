# Welcome to Atlas!

Atlas is a registry of Cosmos SDK Modules that can be viewed at https://atlas.cosmos.network. Atlas is also a CLI compiled from golang that helps with publishing modules to the registry by authorizing users via Github and utilizing a Manifest file to describe the module. The CLI is also used to help in the operation of the actual Atlas server although most users will never need to worry about running the server themselves.

Alltogether Atlas consists of 3 parts:
1) The **CLI Binary** which is used to publish and update Cosmos SDK Modules.
  - To learn how to use the CLI Binary to publish modules, see the [CLI Docs](./cli-docs.md).
1) The **Node Explorer** which is an API and crawler that indexes publicly available information about nodes operating the Cosmos Hub blockchain.
  - To learn more about the Node Explorer, see the [Node Explorer Docs](./node-explorer.md).
1) The **Atlas Server** is used to store and serve all the information about modules and nodes as collected and configured by the CLI Binary and the Node Explorer.
  - To learn more about the Atlas server, see the [Server Docs](./server-docs.md).





