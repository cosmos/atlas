package v1

import (
	"testing"

	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
)

func TestModuleFromManifest(t *testing.T) {
	manifest := Manifest{
		Module: ModuleManifest{
			Name:          "x/test",
			Team:          "test_team",
			Repo:          "https://github.com/test/test-repo",
			Keywords:      []string{"tokens", "transfer"},
			Description:   "A test description about a test module.",
			Homepage:      "https://testmodule.com",
			Documentation: "https://github.com/test/test-repo/blob/master/x/test/readme.md",
		},
		BugTracker: BugTackerManifest{
			URL:     "https://testmodule.com/bugs",
			Contact: "bugs@testmodule.com",
		},
		Authors: []AuthorsManifest{
			{Name: "test_author1", Email: "testauthor1@testmodule.com"},
			{Name: "test_author2", Email: "testauthor2@testmodule.com"},
		},
		Version: VersionManifest{
			Version:   "v1.0.0",
			SDKCompat: "v0.40.x",
		},
	}

	module := ModuleFromManifest(manifest)
	require.Equal(t, manifest.Module.Name, module.Name)
	require.Equal(t, manifest.Module.Team, module.Team)
	require.Equal(t, manifest.Module.Repo, module.Repo)
	require.Equal(t, manifest.Module.Description, module.Description)
	require.Equal(t, manifest.Module.Homepage, module.Homepage)
	require.Equal(t, manifest.Module.Documentation, module.Documentation)
	require.Len(t, module.Keywords, 2)
	require.Equal(t, manifest.Module.Keywords[0], module.Keywords[0].Name)
	require.Equal(t, manifest.Module.Keywords[1], module.Keywords[1].Name)
	require.Equal(t, manifest.BugTracker.URL, module.BugTracker.URL.String)
	require.Equal(t, manifest.BugTracker.Contact, module.BugTracker.Contact.String)
	require.Len(t, module.Authors, 2)
	require.Equal(t, manifest.Authors[0].Name, module.Authors[0].Name)
	require.Equal(t, manifest.Authors[0].Email, module.Authors[0].Email.String)
	require.Equal(t, manifest.Authors[1].Name, module.Authors[1].Name)
	require.Equal(t, manifest.Authors[1].Email, module.Authors[1].Email.String)
	require.Equal(t, manifest.Version.Version, module.Version.Version)
	require.Equal(t, manifest.Version.SDKCompat, module.Version.SDKCompat.String)
}

func TestValidateManifest(t *testing.T) {
	validate := validator.New()

	testCases := []struct {
		name      string
		manifest  Manifest
		expectErr bool
	}{
		{
			"valid manifest all fields",
			Manifest{
				Module: ModuleManifest{
					Name:          "x/test",
					Team:          "test_team",
					Repo:          "https://github.com/test/test-repo",
					Keywords:      []string{"tokens", "transfer"},
					Description:   "A test description about a test module.",
					Homepage:      "https://testmodule.com",
					Documentation: "https://github.com/test/test-repo/blob/master/x/test/readme.md",
				},
				BugTracker: BugTackerManifest{
					URL:     "https://testmodule.com/bugs",
					Contact: "bugs@testmodule.com",
				},
				Authors: []AuthorsManifest{
					{Name: "test_author1", Email: "testauthor1@testmodule.com"},
					{Name: "test_author2", Email: "testauthor2@testmodule.com"},
				},
				Version: VersionManifest{
					Version:   "v1.0.0",
					SDKCompat: "v0.40.x",
				},
			},
			false,
		},
		{
			"valid manifest no bug tracker",
			Manifest{
				Module: ModuleManifest{
					Name:          "x/test",
					Team:          "test_team",
					Repo:          "https://github.com/test/test-repo",
					Keywords:      []string{"tokens", "transfer"},
					Description:   "A test description about a test module.",
					Homepage:      "https://testmodule.com",
					Documentation: "https://github.com/test/test-repo/blob/master/x/test/readme.md",
				},
				Authors: []AuthorsManifest{
					{Name: "test_author1", Email: "testauthor1@testmodule.com"},
					{Name: "test_author2", Email: "testauthor2@testmodule.com"},
				},
				Version: VersionManifest{
					Version:   "v1.0.0",
					SDKCompat: "v0.40.x",
				},
			},
			false,
		},
		{
			"invalid bug tracker",
			Manifest{
				Module: ModuleManifest{
					Name:          "x/test",
					Team:          "test_team",
					Repo:          "https://github.com/test/test-repo",
					Keywords:      []string{"tokens", "transfer"},
					Description:   "A test description about a test module.",
					Homepage:      "https://testmodule.com",
					Documentation: "https://github.com/test/test-repo/blob/master/x/test/readme.md",
				},
				BugTracker: BugTackerManifest{
					URL:     "https://testmodule",
					Contact: "testmodule.com",
				},
				Authors: []AuthorsManifest{
					{Name: "test_author1", Email: "testauthor1@testmodule.com"},
					{Name: "test_author2", Email: "testauthor2@testmodule.com"},
				},
				Version: VersionManifest{
					Version:   "v1.0.0",
					SDKCompat: "v0.40.x",
				},
			},
			true,
		},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			err := validate.Struct(tc.manifest)
			require.Equal(t, tc.expectErr, err != nil, err)
		})
	}
}
