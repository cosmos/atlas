package cmd_test

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/urfave/cli/v2"

	"github.com/cosmos/atlas/cmd"
	"github.com/cosmos/atlas/config"
)

func TestParseServerConfig(t *testing.T) {
	svrCmd := cmd.StartServerCommand()
	fs, err := cmd.FlagSetFromCmdFlags("test", svrCmd.Flags)
	require.NoError(t, err)

	ctx := cli.NewContext(nil, fs, nil)
	ctx.Command = svrCmd

	tmpDir := t.TempDir()
	configFile, err := os.Create(path.Join(tmpDir, "config.toml"))
	require.NoError(t, err)

	defer func() {
		_ = configFile.Close()
	}()

	configBz := []byte(fmt.Sprintf(`%s="1"
%s="2"
%s="3"
`, config.FlagListenAddr, config.FlagGHClientSecret, config.FlagGHRedirectURL))
	_, err = configFile.Write(configBz)
	require.NoError(t, err)

	// set config file path flag
	require.NoError(t, ctx.Set(config.FlagConfig, configFile.Name()))

	// parse config and ensure all values are read from file
	cfg, err := cmd.ParseServerConfig(ctx)
	require.NoError(t, err)
	require.Equal(t, "1", cfg.String(config.FlagListenAddr))
	require.Equal(t, "2", cfg.String(config.FlagGHClientSecret))
	require.Equal(t, "3", cfg.String(config.FlagGHRedirectURL))

	// set OS environment variables and ensure environment variables take precedence
	os.Setenv("ATLAS_"+strings.ReplaceAll(strings.ToUpper(config.FlagListenAddr), ".", "_"), "11")
	os.Setenv("ATLAS_"+strings.ReplaceAll(strings.ToUpper(config.FlagGHClientSecret), ".", "_"), "12")

	cfg, err = cmd.ParseServerConfig(ctx)
	require.NoError(t, err)
	require.Equal(t, "11", cfg.String(config.FlagListenAddr))
	require.Equal(t, "12", cfg.String(config.FlagGHClientSecret))
	require.Equal(t, "3", cfg.String(config.FlagGHRedirectURL))

	// set a flag and ensure it takes precedence over env and file
	require.NoError(t, ctx.Set(config.FlagListenAddr, "21"))
	cfg, err = cmd.ParseServerConfig(ctx)
	require.NoError(t, err)
	require.Equal(t, "21", cfg.String(config.FlagListenAddr))
	require.Equal(t, "12", cfg.String(config.FlagGHClientSecret))
	require.Equal(t, "3", cfg.String(config.FlagGHRedirectURL))
}
