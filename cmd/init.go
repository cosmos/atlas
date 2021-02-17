package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
)

func InitCommand() *cli.Command {
	return &cli.Command{
		Name:  "init",
		Usage: `Initialize a manifest.`,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "dir",
				Aliases: []string{"d"},
				Usage:   "The directory to generate the manifest in",
			},
		},
		Action: func(ctx *cli.Context) error {
			var manifestBuffer bytes.Buffer
			var docBuffer bytes.Buffer

			// Get current working directory. This is avoid generating files elsewhere
			manifestPath := ctx.String("dir")

			if _, err := manifestBuffer.WriteString(defaultManifestTemplate); err != nil {
				return err
			}
			if _, err := docBuffer.WriteString(defaultDocsTemplate); err != nil {
				return err
			}

			if err := ioutil.WriteFile(filepath.Join(manifestPath, filepath.Base("manifest.toml")), manifestBuffer.Bytes(), 0644); err != nil {
				fmt.Printf("Creation of manifest.toml failed: %v", err)
				os.Exit(1)
			}

			if err := ioutil.WriteFile(filepath.Join(manifestPath, filepath.Base("atlas.md")), docBuffer.Bytes(), 0644); err != nil {
				fmt.Printf("Creation of manifest.toml failed: %v", err)
				os.Exit(1)
			}
			return nil
		},
	}
}

const defaultManifestTemplate = `[module]
# Name of the module. (Required)
name = ""

# Description of the module. (Optional)
description = ""

# Link to where the module is located, it can also be a link to your project. (Optional)
homepage = ""

#List of key words describing your module (Optional)
keywords = []


[bug_tracker]
# A URL to a site that provides information or guidance on how to submit or deal
# with security vulnerabilities and bug reports.
url = ""

# An email address to submit bug reports and security vulnerabilities to.
contact = ""

# To list multiple authors, multiple [[authors]] need to be created
[[authors]]
# Name of one of the authors. Typically their Github name. (Required)
name = ""

# Email of the author mentioned. (Optional)
email = ""

[version]
# The repository field should be a URL to the source repository for your module.
# Typically, this will point to the specific GitHub repository release/tag for the
# module, although this is not enforced or required. (Required)
repo = ""

# The documentation field specifies a URL to a website hosting the module's documentation. (Optional)
documentation = ""

# The module version to be published. (Required)
version = ""

# An optional Cosmos SDK version compatibility may be provided. (Optional)
sdk_compat = ""
`

const defaultDocsTemplate = `
# <module_name>

<!-- Short description of the module --> 

## Usage

1. Import the module.

<!--
Show how the module should be imported and what other imports are needed

Example:
   import (
    distr "github.com/cosmos/cosmos-sdk/x/distribution"
    distrclient "github.com/cosmos/cosmos-sdk/x/distribution/client"
    distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
    distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
   )
-->
2. Add AppModuleBasic to your ModuleBasics.

<!--
Example:
    var (
      ModuleBasics = module.NewBasicManager(
        // ...
        distr.AppModuleBasic{},
      }
    )
-->

3. Give distribution module account permissions.

<!--
If account permissions are needed show an example.
Example:
  	// module account permissions
    var maccPerms = map[string][]string{
      distrtypes.ModuleName:          nil,
    }
-->

4. Allow the <module_name> module to receive funds.

<!--
If the module receives funds it should be regestered to do so. 
Example:
      allowedReceivingModAcc = map[string]bool{
        distrtypes.ModuleName: true,
      }
-->
5. Add the <module_name> keeper to your apps struct.

<!--
Example:
      type app struct {
        // ...
        DistrKeeper      distrkeeper.Keeper
        // ...
      }
-->
6. Add the <module_name> store key to the group of store keys.

<!--
   func NewApp(...) *App {
     // ...
      keys := sdk.NewKVStoreKeys(
       distrtypes.StoreKey,
      )
     // ...
   }
-->
7. Create the keeper. 

<!--
Example:
   func NewApp(...) *App {
      // ...
      // create capability keeper with router
      app.DistrKeeper = distrkeeper.NewKeeper(
		    appCodec, keys[distrtypes.StoreKey], app.GetSubspace(distrtypes.ModuleName), app.AccountKeeper, app.BankKeeper,
		    &stakingKeeper, authtypes.FeeCollectorName, app.ModuleAccountAddrs(),
			)
   }
-->

8. Add the <module_name> module to the app's ModuleManager.

<!--
Example:
   func NewApp(...) *App {
     // ...
     app.mm = module.NewManager(
       // ...
       distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
       // ...
     )
   }
-->

9. Set the <module_name> module begin blocker order.

<!--
Example:
    func NewApp(...) *App {
     // ...
      app.mm.SetOrderBeginBlockers(
        // ...
        distrtypes.ModuleName,
        // ...
      )
    }
    -->

10.  Set the <module_name> module genesis order.

<!--
Example:
   func NewApp(...) *App {
     // ...
     app.mm.SetOrderInitGenesis(distrtypes.ModuleName,, ...)
   }
-->

11. Add the <module_name> module to the simulation manager (if you have one set).

<!--
Example:
   func NewApp(...) *App {
     // ...
     app.sm = module.NewSimulationManager(
       // ...
       distr.NewAppModule(appCodec, app.DistrKeeper, app.AccountKeeper, app.BankKeeper, app.StakingKeeper),
       // ...
     )
   }
-->

<!-- These examples only cover part of what a module require to use. Please add enough information so that a developer can quickly add your module to their chain.-->

## Genesis

## Messages

<!-- Todo: add a short description about client interactions -->

### CLI
<!-- Todo: add a short description about client interactions -->

#### Queries
<!-- Todo: add a short description about cli query interactions -->

#### Transactions
<!-- Todo: add a short description about cli transaction interactions -->


### REST
<!-- Todo: add a short description about REST interactions -->

#### Query
<!-- Todo: add a short description about REST query interactions -->

#### Tx
<!-- Todo: add a short description about REST transaction interactions -->

### gRPC
<!-- Todo: add a short description about gRPC interactions -->

#### Query
<!-- Todo: add a short description about gRPC query interactions -->

#### Tx
<!-- Todo: add a short description about gRPC transactions interactions -->
`
