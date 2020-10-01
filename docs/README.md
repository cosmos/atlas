# Atlas Documentation

## Manifest

Atlas reads a canonical TOML-based manifest when publishing Cosmos SDK modules.
The manifest schema is defined as follows

- [[module]](#module)
  - [name](#name-required)
  - [team](#team-required)
  - [description](#description)
  - [documentation](#documentation)
  - [homepage](#homepage)
  - [repo](#repo-required)

### [module]

#### `name` (required)

The name of the Cosmos SDK module. The name does not necessarily
have to be unique as module's can be forked from other teams and organizations.
Typically, a module will be named as `x/<name>`, where `<name>` is concise and
meaningful. Note, the combination of `name` and `team` (see below) must be unique.

```toml
[module]

name = "x/poa"
# ...
```

#### `team` (required)

The team or organization that primarily owns and maintains
the module. The team combined with the `name` must be globally unique.

```toml
[module]

team = "Interchain"
# ...
```

#### `description`

The description is a short blurb about the Cosmos SDK module. Atlas will display
this with the module. This should be plain text (not Markdown).

```toml
[module]

description = "A short description of my module."
# ...
```

#### `documentation`

The documentation field specifies a URL to a website hosting the module's documentation.
Typically, this is a Markdown file hosted in the module's root directory in Github.

```toml
[module]

documentation = "https://github.com/cosmos/cosmos-sdk/blob/master/x/slashing/readme.md"
# ...
```

#### `homepage`

The homepage field should be a URL to a site that is the home page for your module,
organization or team.

```toml
[module]

homepage = "https://interchain.io/"
# ...
```

#### `repo` (required)

The repository field should be a URL to the source repository for your module.

```toml
[module]

repo = "https://github.com/cosmos/cosmos-sdk"
# ...
```

<!-- 
	// one-to-one relationships
		BugTracker BugTracker `json:"bug_tracker" yaml:"bug_tracker" gorm:"foreignKey:module_id"`

		// many-to-many relationships
		Keywords []Keyword `gorm:"many2many:module_keywords" json:"keywords" yaml:"keywords"`
		Authors  []User    `gorm:"many2many:module_authors" json:"authors" yaml:"authors"`
		Owners   []User    `gorm:"many2many:module_owners" json:"owners" yaml:"owners"`

		// one-to-many relationships
		Version  ModuleVersion   `gorm:"-" json:"-" yaml:"-"` // current version in manifest
		Versions []ModuleVersion `gorm:"foreignKey:module_id" json:"versions" yaml:"versions"` -->
