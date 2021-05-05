# Atlas

![GitHub Logo](../images/atlas_logo.png)

![Build, Test and Cover](https://github.com/cosmos/atlas/workflows/Build,%20Test%20and%20Cover/badge.svg?branch=bez%2F13-client-cli-commands)
[![Go Report Card](https://goreportcard.com/badge/github.com/cosmos/atlas)](https://goreportcard.com/report/github.com/cosmos/atlas)
[![codecov](https://codecov.io/gh/cosmos/atlas/branch/main/graph/badge.svg)](https://codecov.io/gh/cosmos/atlas)
[![GoDoc](https://godoc.org/github.com/cosmos/atlas?status.png)](https://pkg.go.dev/github.com/cosmos/atlas)
[![license](https://img.shields.io/github/license/cosmos/atlas.svg)](https://github.com/cosmos/atlas/blob/main/LICENSE)
[![Netlify Status](https://api.netlify.com/api/v1/badges/76c69961-2403-433d-a115-061ce17148af/deploy-status)](https://app.netlify.com/sites/cosmos-atlas/deploys)

> Source code for the default [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) module
registry, viewable online at [atlas.cosmos.network](https://atlas.cosmos.network).

## Table of Contents

- [Atlas](#atlas)
  - [Table of Contents](#table-of-contents)
  - [Summary](#sumary)
  - [Usage](#usage)
    - [CLI Binary](#cli-binary)
    - [Node Explorer](#node-explorer)
    - [Atlas Server](#atlas-server)
  - [License](#license)


## Summary

Atlas implements a [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) module registry,
where developers are able to publish and update modules. The registry provides a singular and
holistic interface for application developers to discover [Cosmos SDK](https://github.com/cosmos/cosmos-sdk)
modules when building their blockchain applications.

## Usage

Atlas is composed of two primary components, the server and the web application.
The server is responsible for providing a RESTful API, handling user authentication
via Github OAuth and persisting modules and relevant data to PostgreSQL.

Alltogether Atlas consists of 3 parts: the **CLI Binary**, the **Node Explorer** and the **Atlas Server**.

### **[CLI Binary](./cli-docs.md)** 
The CLI Binary is used to publish and update Cosmos SDK Modules.

To learn how to use the CLI Binary to publish modules, see the [CLI Docs](./cli-docs.md).  

### **[Node Explorer](./node-explorer.md)**
The Node Explorer is an API and crawler that indexes publicly available information about nodes operating the Cosmos Hub blockchain.

To learn more about the Node Explorer, see the [Node Explorer Docs](./node-explorer.md).

### **[Atlas Server](./server-docs.md)**
The Atlas Server is responsible for providing a RESTful API, handling user authentication
via Github OAuth and persisting modules and relevant data to PostgreSQL and serving all information
using a web interface.

To learn more about the Atlas server, see the [Server Docs](./server-docs.md).

## License

- [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)
