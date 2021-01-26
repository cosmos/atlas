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
			var buffer bytes.Buffer

			// Get current working directory. This is avoid generating files elsewhere
			manifestPath, err := os.Getwd()
			if err != nil {
				return err
			}

			if _, err = buffer.WriteString(defaultManifestTemplate); err != nil {
				return err
			}

			err = ioutil.WriteFile(filepath.Join(manifestPath, filepath.Base("atlas.toml")), buffer.Bytes(), 0644)
			if err != nil {
				fmt.Printf("MustWriteFile failed: %v", err)
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
