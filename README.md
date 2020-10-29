# Atlas

![GitHub Logo](./images/atlas_logo.png)

![Build, Test and Cover](https://github.com/cosmos/atlas/workflows/Build,%20Test%20and%20Cover/badge.svg?branch=bez%2F13-client-cli-commands)
[![Go Report Card](https://goreportcard.com/badge/github.com/cosmos/atlas)](https://goreportcard.com/report/github.com/cosmos/atlas)
[![codecov](https://codecov.io/gh/cosmos/atlas/branch/main/graph/badge.svg)](https://codecov.io/gh/cosmos/atlas)
[![GoDoc](https://godoc.org/github.com/cosmos/atlas?status.png)](https://pkg.go.dev/github.com/cosmos/atlas)
[![license](https://img.shields.io/github/license/cosmos/atlas.svg)](https://github.com/cosmos/atlas/blob/main/LICENSE)

> Source code for the default [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) module
registry, viewable online at [atlas.cosmos.network](https://atlas.cosmos.network).

## Table of Contents

- [Atlas](#atlas)
  - [Table of Contents](#table-of-contents)
  - [Background](#background)
  - [Usage](#usage)
    - [Server](#server)
    - [Web App](#web-app)
  - [Migrations](#migrations)
  - [Tests](#tests)
  - [License](#license)

## Background

Atlas implements a [Cosmos SDK](https://github.com/cosmos/cosmos-sdk) module registry,
where developers are able to publish and update modules. The registry provides a singular and
holistic interface for application developers to discover [Cosmos SDK](https://github.com/cosmos/cosmos-sdk)
modules when building their blockchain applications.

More information about the architecture, publishing and module configuration can
be found under [docs](./docs/README.md).

## Usage

Atlas is composed of two primary components, the server and the web application.
The server is responsible for providing a RESTful API, handling user authentication
via Github OAuth and persisting modules and relevant data to PostgreSQL.

### Server

In order to start the Atlas server, you must provide a series of configuration
values that may be defined in environment variables, a (TOML) configuration file
or via CLI flags (in order of precedence). See the sample [env](./.env.sample) or
[config](./config.sample.toml) files for all possible configurations.

```shel
$ atlas server --config=/path/to/atlas/config.toml
```

Note:

1. Atlas will look for environment variables defined in a `.env` file in the
root directory. Any explicit environment variables defined will override those
defined in this file.
2. Certain configuration values are not exposed or able to be provided via CLI flags.
3. All environment variables must be prefixed with `ATLAS_*`.

See `--help` for further documentation.

### API Documentation

Atlas leverages [Swagger](https://swagger.io/) to document its API. The documentation
is compiled automatically via [swag](https://github.com/swaggo/swag/) through
annotated REST handlers. The compiled documentation resides in the `docs/api`
directory and is served at `/api/docs/`.

The [Swagger](https://swagger.io/) documentation can be recompiled as follows:

```shell
$ make update-swagger-docs
```

### Web App

The Atlas web application is built using [Vue.js](https://vuejs.org/) and is
contained in the `web` directory. The web application is served as a static
resource from the same service as the API. It contains a `.env` file at the root
which must contain a `VUE_APP_ATLAS_API_ADDR` environment variable that describes
how to reach the Atlas API.

To build locally and watch for lives changes:

```shell
$ cd web && npm run build-watch
```

To build for production:

```shell
$ cd web && npm run build
```

## Migrations

Atlas performs migrations through the [migrate](https://github.com/golang-migrate/migrate)
tool. The migrations are defined in `db/migrations`. In order to run migrations,
you must provide a `ATLAS_DATABASE_URL` environment variable.

```shell
$ ATLAS_DATABASE_URL=... make migrate
```

To install [migrate](https://github.com/golang-migrate/migrate), please visit [here](https://github.com/golang-migrate/migrate/tree/master/cmd/migrate)
for instructions.

## Local Development

To run, test and experiment with Atlas in a local development environment, execute
the following:

1. Start a postgres database using Docker:

   ```shell
   $ docker-compose up -d
   ```

2. Run migrations:

   ```shell
   $ ATLAS_DATABASE_URL="postgres://postgres:postgres@localhost:6432/postgres?sslmode=disable" make migrate
   ```

3. Populate your Atlas server config or root `.env`:

   ```env
   # .env

   # Database and Atlas server options
   ATLAS_DATABASE_URL=postgres://postgres:postgres@localhost:6432/postgres?sslmode=disable
   ATLAS_LOG_FORMAT=text

   # GitHub OAuth
   ATLAS_GH_CLIENT_ID=...
   ATLAS_GH_CLIENT_SECRET=...
   ATLAS_GH_REDIRECT_URL=http://localhost:8080/api/v1/session/authorize

   # Testing session cookie (e.g. securecookie.GenerateRandomKey(32))
   ATLAS_SESSION_KEY=UIla7DSIVXzhvd9yHxexEExel9HQpSCQ+Rsn3y+e2Rs=
   ```

4. Start Atlas:

   ```shell
   $ atlas server --dev=true --log.level=debug
   ```

Note, if you run Atlas with a custom listening address, be sure to update the
`VUE_APP_ATLAS_API_ADDR` environment variable in `web/.env` and the `ATLAS_GH_REDIRECT_URL`.

## Tests

Atlas performs all database relevant tests through a Docker Postgres instance.
Executing the `$ make test` target will automatically start a Postgres Docker
instance and populate all relevant environment variables. If you'd like to execute
tests on a different Postgres instance, you must provide the `ATLAS_MIGRATIONS_DIR`
and `ATLAS_TEST_DATABASE_URL` environment variables.

```shell
$ ATLAS_TEST_DATABASE_URL=... ATLAS_MIGRATIONS_DIR=... make test
```

## License

- [Apache License, Version 2.0](https://www.apache.org/licenses/LICENSE-2.0)
