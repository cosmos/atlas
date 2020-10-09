package cmd_test

import (
	"context"
	"os"
	"path"
	"testing"

	"github.com/BurntSushi/toml"
	"github.com/stretchr/testify/require"

	"github.com/cosmos/atlas/cmd"
	v1 "github.com/cosmos/atlas/server/router/v1"
)

func TestPublishCommand_DryRun_Valid(t *testing.T) {
	app := cmd.NewApp()
	mockIn, mockOut := cmd.ApplyMockIO(app)
	ctx := cmd.ContextWithReader(context.Background(), mockIn)

	manifest := v1.Manifest{
		Module: v1.ModuleManifest{
			Name:          "x/test",
			Team:          "test_team",
			Repo:          "https://github.com/test/test-repo",
			Keywords:      []string{"tokens", "transfer"},
			Description:   "A test description about a test module.",
			Homepage:      "https://testmodule.com",
			Documentation: "https://github.com/test/test-repo/blob/master/x/test/readme.md",
		},
		BugTracker: v1.BugTackerManifest{
			URL:     "https://testmodule.com/bugs",
			Contact: "bugs@testmodule.com",
		},
		Authors: []v1.AuthorsManifest{
			{Name: "test_author1", Email: "testauthor1@testmodule.com"},
			{Name: "test_author2", Email: "testauthor2@testmodule.com"},
		},
		Version: v1.VersionManifest{
			Version:   "v1.0.0",
			SDKCompat: "v0.40.x",
		},
	}

	// create temp directory and write a valid manifest to it
	tmpDir := t.TempDir()
	manifestPath := path.Join(tmpDir, "manifest.toml")

	file, err := os.Create(manifestPath)
	require.NoError(t, err)

	defer func() {
		_ = file.Close()
	}()

	encoder := toml.NewEncoder(file)
	require.NoError(t, encoder.Encode(manifest))

	// execute command and verify output
	require.NoError(t, cmd.ExecTestCmd(ctx, app, []string{"atlas", "publish", "-m", manifestPath, "--dry-run"}))
	require.Contains(t, mockOut.String(), "manifest successfully verified!", mockOut.String())
}

func TestPublishCommand_DryRun_Invalid(t *testing.T) {
	app := cmd.NewApp()
	mockIn, mockOut := cmd.ApplyMockIO(app)
	ctx := cmd.ContextWithReader(context.Background(), mockIn)

	manifest := v1.Manifest{
		Module: v1.ModuleManifest{
			Team:          "test_team",
			Repo:          "https://github.com/test/test-repo",
			Keywords:      []string{"tokens", "transfer"},
			Description:   "A test description about a test module.",
			Homepage:      "https://testmodule.com",
			Documentation: "https://github.com/test/test-repo/blob/master/x/test/readme.md",
		},
		BugTracker: v1.BugTackerManifest{
			URL:     "bad_url",
			Contact: "testmodule.com",
		},
		Authors: []v1.AuthorsManifest{
			{Email: "testauthor1@testmodule.com"},
			{Name: "test_author2", Email: "testauthor2@testmodule.com"},
		},
		Version: v1.VersionManifest{
			Version:   "v1.0.0",
			SDKCompat: "v0.40.x",
		},
	}

	// create temp directory and write an invalid manifest to it
	tmpDir := t.TempDir()
	manifestPath := path.Join(tmpDir, "manifest.toml")

	file, err := os.Create(manifestPath)
	require.NoError(t, err)

	defer func() {
		_ = file.Close()
	}()

	encoder := toml.NewEncoder(file)
	require.NoError(t, encoder.Encode(manifest))

	// execute command and verify output
	require.Error(t, cmd.ExecTestCmd(ctx, app, []string{"atlas", "publish", "-m", manifestPath, "--dry-run"}))
	require.Contains(t, mockOut.String(), "failed to verify manifest", mockOut.String())
}
