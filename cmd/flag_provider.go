package cmd

import (
	"errors"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/maps"
	"github.com/urfave/cli/v2"
)

var _ koanf.Provider = (*CLIFlagProvider)(nil)

// CLIFlagProvider implements a simple Koanf CLI flag provider.
type CLIFlagProvider struct {
	delim  string
	ctx    *cli.Context
	konfig *koanf.Koanf
}

func NewCLIFlagProvider(ctx *cli.Context, delim string, konfig *koanf.Koanf) *CLIFlagProvider {
	return &CLIFlagProvider{
		ctx:    ctx,
		delim:  delim,
		konfig: konfig,
	}
}

// Read reads the flag variables and returns a nested conf map.
func (p *CLIFlagProvider) Read() (map[string]interface{}, error) {
	mp := make(map[string]interface{})

	setFlags := make(map[string]bool)
	for _, fn := range p.ctx.FlagNames() {
		setFlags[fn] = true
	}

	for _, flag := range p.ctx.Command.VisibleFlags() {
		// first name is the primary name
		name := flag.Names()[0]

		// set the flag value iff one of the two criteria are met:
		//
		// 1. The flag/config does not exist already in the config object.
		// 2. The flag/config does exist but was overridden by the CLI flag.
		if !p.konfig.Exists(name) || (p.konfig.Exists(name) && setFlags[name]) {
			mp[name] = p.ctx.Value(name)
		}
	}

	return maps.Unflatten(mp, p.delim), nil
}

// ReadBytes is not supported by the env koanf.
func (p *CLIFlagProvider) ReadBytes() ([]byte, error) {
	return nil, errors.New("cli flag provider does not support this method")
}

// Watch is not supported.
func (p *CLIFlagProvider) Watch(cb func(event interface{}, err error)) error {
	return errors.New("cli flag provider does not support this method")
}
