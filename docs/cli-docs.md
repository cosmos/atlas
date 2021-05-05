# CLI Binary

The atlas CLI binary is used to publish [manifest](./manifest.md) files that describe [Cosmos SDK Modules](https://docs.cosmos.network/v0.42/building-modules/intro.html) in order to make them discoverable at https://atlas.cosmos.network.

## Installation

To install the Atlas CLI locally just download the repository and install as follows:

```sh
git clone git@github.com:cosmos/atlas.git
cd atlas
make install
```


> Note: If you'd prefer to use docker to run commands take a look [here](#docker).


Once the atlas binary is installed you can get an overview of what it does by using `atlas --help`:

```sh
> atlas --help
NAME:
   Atlas CLI - A Cosmos SDK module registry framework

USAGE:
   atlas [global options] command [command options] [arguments...]

VERSION:
   0.0.2-16-ge7fede6-e7fede65720e9292f391276de0cb02ef09ecb7f1

COMMANDS:
   server  Start the atlas server
   login   Save an API token from the Atlas registry locally. It a token is not specified, it will read from stdin.
   publish  Publish a Cosmos SDK module to the Atlas registry.
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help (default: false)
   --version, -v  print the version (default: false)
```

## Manifest

In order for atlas to register your module, it needs to access a manifest `.toml` file that describes the module. This file needs to be accessible from the atlas binary and it's recommended to keep it within the module itself.

View a full outline of the expected contents of a manifest `.toml` file [here](./manifest.md).

For example: The [auth](https://github.com/cosmos/cosmos-sdk/blob/master/x/auth] module from the Cosmos SDK contains a manifest called `atlas.toml` located at `/x/auth/atlas/atlas.toml`. It looks like this:

```toml
[module]
description = "The auth module is responsible for specifying the base transaction and account types for an application, as well as AnteHandler and authentication logic."
homepage = "https://github.com/cosmos/cosmos-sdk"
keywords = [
  "authentication",
  "signatures",
  "ante",
  "transactions",
  "accounts",
]

name = "x/auth"

[bug_tracker]
url = "https://github.com/cosmos/cosmos-sdk/issues"

[[authors]]
name = "alexanderbez"

[[authors]]
name = "fedekunze"

[[authors]]
name = "cwgoes"

[[authors]]
name = "alessio"

[version]
documentation = "https://raw.githubusercontent.com/cosmos/cosmos-sdk/master/x/auth/atlas/atlas-v0.39.1.md"
repo = "https://github.com/cosmos/cosmos-sdk/releases/tag/v0.39.2"
sdk_compat = "v0.39.x"
version = "v0.39.2"
```

> Note: To automate the publication of modules to Atlas try using github actions like those seen [here](https://github.com/cosmos/cosmos-sdk/blob/master/.github/workflows/atlas.yml) in the Cosmos SDK repository.

## Publication

Once you've created a Manifest that describes your module you can use the atlas CLI to publish it to the atlas server at https://atlas.cosmos.network.

If it's the first time you've used atlas you'll need to login first. To do this you need to first obtain an API token by logging into https://atlas.cosmos.network using your github account. 

<img width="454" alt="Screen Shot 2021-05-05 at 11 29 36" src="https://user-images.githubusercontent.com/964052/117123113-e8f4f780-ad96-11eb-9926-48e2e42ae62a.png">


Once logged in go to your account page and create a new API token. Give it a name that you can remember and then copy the token that is generated afterwards.

<img width="702" alt="Screen Shot 2021-05-05 at 11 30 03" src="https://user-images.githubusercontent.com/964052/117123197-05912f80-ad97-11eb-87ad-959110e61be4.png">

Once you have your API token in hand, use it to login from the CLI using the following command:

```sh
atlas login [token]
```

Now you're able to publish your manifest by using the following command:

```sh
atlas publish -m [path/to/manifest]
```

> Note, you can execute a `--dry-run` first to ensure the manifest is valid

Congratulations! Your module will now show up on the atlas directory website 🎉

### Docker

To do the entire publication flow in one command using Docker try the following:

```sh
docker run -v $(shell pwd):/workspace --workdir /workspace interchainio/atlas:latest [APIkey] [path/to/manifest]] [dry-run, default false]
```
