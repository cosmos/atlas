# Architecture

- [Architecture](#architecture)
  - [Authentication](#authentication)
  - [Users](#users)
  - [Publishing](#publishing)
  - [Router](#router)

## Authentication

Atlas utilizes [GitHub OAuth](https://docs.github.com/en/free-pro-team@latest/developers/apps/authorizing-oauth-apps)
and [server-side session cookies](https://github.com/gorilla/sessions) for user
authentication and authorization. Through the webapp, Users may gran or reject
access to the Atlas GitHub OAuth application.

Upon granting access, the user is taken through an [implicit flow](https://tools.ietf.org/html/rfc6749#section-1.3.2)
where the resource server (Atlas) completes authorization and grants an access
token. Atlas does not use the access token itself, but rather the server-side
session cookie to authenticate.

### API Tokens

Through the webapp, users are only authenticated through the presence of a valid
session cookie. However, users may also utilize authenticated API endpoints through
the use of API tokens.

API tokens can be created and revoked through the user portal of the webapp only.
In addition to being used for direct API access to authenticated routes, API
tokens are also the primary means in which owners can publish Cosmos SDK modules
through the Atlas CLI.

## Users

The data model of Atlas describes a single user model, where any given user can
be a Cosmos SDK module owner and/or author. However, only owners can publish and
update existing published Cosmos SDK modules.

The initial publisher of a Cosmos SDK module is the first and only owner of that
respective module. That owner may then invite other users of Atlas to become owners.
On the other hand, authors are arbitrary and do not need to be invited. The
owners of a Cosmos SDK module define who the authors are.

## Publishing

Publishing Cosmos SDK modules requires a user to login via the CLI with an API
token. Once logged in, the owner can publish a new or existing module which they
have access to. Publishing modules requires a valid manifest. See [here](./manifest.md)
for more information.

Publishers can publish a Cosmos SDK module only if they are an owner of the
respective module and a contributor to the GitHub repository where the module
resides in.

## Router

All Atlas API routes are versioned via with a path prefix of `/api/<version>`.
Regardless of the API version, all requests come bundled with CORS and request
logging middleware.

In addition, Atlas documents it's API via [Swagger](https://swagger.io/). Note,
currently only the latest API version is documented.
