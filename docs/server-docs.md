
## Server

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
contained in the `web` directory. The web application is executed as a separate
process external from the Atlas API server. The webapp requires `VUE_APP_ATLAS_API_ADDR`
to be populated in order to know how to speak with the Atlas API. This can be
set as an explicit environment variable or populated in a `.env` file at the root
of the `web` directory.

To run the webapp locally and watch for lives changes:

```shell
$ cd web && yarn serve
```

To build for production:

```shell
$ cd web && yarn build
```

### Publishing

To publish a Cosmos SDK module, please see the publishing [doc](https://github.com/cosmos/atlas/blob/main/docs/publishing.md).

### Action

To publish a Cosmos SDK module with github actions, please see the action [doc](https://github.com/cosmos/atlas/blob/main/docs/action.md).

## Migrations

Atlas performs migrations through the [migrate](https://github.com/golang-migrate/migrate)
tool. The migrations are defined in `db/migrations`. In order to run migrations,
you must provide a `ATLAS_DATABASE_URL` environment variable.

```shell
$ ATLAS_DATABASE_URL=... make migrate
```

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
   ATLAS_DATABASE_URL=postgres://postgres:postgres@localhost:6432/postgres?sslmode=disable
   ATLAS_LOG_FORMAT=debug
   ATLAS_DEV=true
   ATLAS_GH_CLIENT_ID=...
   ATLAS_GH_CLIENT_SECRET=...
   ATLAS_ALLOWED_ORIGINS=http://localhost:8081

   # Testing session cookie (e.g. securecookie.GenerateRandomKey(32))
   ATLAS_SESSION_KEY=UIla7DSIVXzhvd9yHxexEExel9HQpSCQ+Rsn3y+e2Rs=
   ```

4. Start Atlas:

   ```shell
   $ atlas server
   ```

5. Start the webapp:

   ```shell
   $ cd web && yarn serve
   ```

Note, if you choose to run Atlas at a different listen address, be sure to populate
`VUE_APP_ATLAS_API_ADDR` and `ATLAS_ALLOWED_ORIGINS` accordingly. Where the former
is the listen address of the Atlas API server and the later is the address of
the webapp (yarn will automatically allocate a free port).

## Tests

Atlas performs all database relevant tests through a Docker Postgres instance.
Executing the `$ make test` target will automatically start a Postgres Docker
instance and populate all relevant environment variables. If you'd like to execute
tests on a different Postgres instance, you must provide the `ATLAS_MIGRATIONS_DIR`
and `ATLAS_TEST_DATABASE_URL` environment variables.

```shell
$ ATLAS_TEST_DATABASE_URL=... ATLAS_MIGRATIONS_DIR=... make test
```
