package cmd_test

import (
	"context"
	"path"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/atlas/cmd"
)

func TestLoginCommand_Stdin(t *testing.T) {
	app := cmd.NewApp()
	mockIn, mockOut := cmd.ApplyMockIO(app)

	// write a fake test token to stdin
	fakeToken := "fake_token"
	mockIn.Reset(fakeToken + "\n")

	// create a temporary directory to write credentials to
	tmpDir := t.TempDir()
	credsPath := path.Join(tmpDir, ".atlas", "credentials")

	ctx := cmd.ContextWithReader(context.Background(), mockIn)
	require.NoError(t, cmd.ExecCmd(ctx, app, []string{"atlas", "login", "-d", tmpDir}))
	require.Contains(t, mockOut.String(), "login token successfully saved", mockOut.String())

	var creds cmd.Credentials

	_, err := toml.DecodeFile(credsPath, &creds)
	require.NoError(t, err)
	require.Equal(t, fakeToken, creds.Registry.Token)
}

func TestLoginCommand_Arg(t *testing.T) {
	app := cmd.NewApp()
	mockIn, mockOut := cmd.ApplyMockIO(app)

	// write a fake test token to stdin
	fakeToken := "fake_token"

	// create a temporary directory to write credentials to
	tmpDir := t.TempDir()
	credsPath := path.Join(tmpDir, ".atlas", "credentials")

	ctx := cmd.ContextWithReader(context.Background(), mockIn)
	require.NoError(t, cmd.ExecCmd(ctx, app, []string{"atlas", "login", "-d", tmpDir, fakeToken}))
	require.Contains(t, mockOut.String(), "login token successfully saved", mockOut.String())

	var creds cmd.Credentials

	_, err := toml.DecodeFile(credsPath, &creds)
	require.NoError(t, err)
	require.Equal(t, fakeToken, creds.Registry.Token)
}
