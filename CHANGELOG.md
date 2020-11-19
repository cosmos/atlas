<!--
Guiding Principles:

Changelogs are for humans, not machines. There should be an entry for every
single version. The same types of changes should be grouped. Versions and
sections should be linkable. The latest version comes first. The release date
of each version is displayed. Mention whether you follow Semantic Versioning.

Ref: https://keepachangelog.com/en/1.0.0/

Usage:

Changelog entries are to be added to the Unreleased section under the
appropriate stanza (see below). Each entry should ideally include a tag and
the Github issue reference in the following format:

* [tag] [#issue-or-pr-number](link) message

Types of changes (Stanzas):

"Features" for new features.
"Improvements" for changes in existing functionality.
"Deprecated" for soon-to-be removed features.
"Fixed" for any bug fixes.
"Security" for any security issues or vulnerabilities.
-->

# Changelog

## [Unreleased]

## [0.0.2] - 2020-11-16

### Features

- [models] [#42](https://github.com/cosmos/atlas/issues/42) Move `Documentation`
  and `Repo` from the `modules` table to the `module_versions` table as to allow
  docs and version locations to be specific to the respective module version.

### Fixed

- [CLI] [#47](https://github.com/cosmos/atlas/issues/47) Fix the default value for the `--registry | -r` flag.
- [webapp] Fix profile page where the logged in user information was being displayed
  instead of the respective profile user.
- [router] [#44](https://github.com/cosmos/atlas/pull/44) Changes to Cosmos SDK publishing:
  - `Team` is no longer explicitly provided in the manifest. The `team` is now
  automatically inferred by the GitHub repository `owner`.
  - The publisher must be a contributor of the GitHub repository.

### Improvements

- [server] Stream logs to a syslog sink in addition to the standard writer when
  not in `dev` mode.
- [webapp] Change HTML and CSS for the astronaut used for error pages.

## [0.0.1] - 2020-11-11

- Initial release!

[Unreleased]: https://github.com/cosmos/atlas/compare/v0.0.2...HEAD
[0.0.2]: https://github.com/cosmos/atlas/releases/tag/v0.0.2
[0.0.1]: https://github.com/cosmos/atlas/releases/tag/v0.0.1
