# Action

This action is used to publish updates or new modules to [Atlas](https://atlas.cosmos.network/). The cosmos registry of modules

## Inputs

### `token`

**Required** The token generated on atlas.cosmos.network that links to your account.

### `path`

**Required** The path to your manifest file.

### `dry-run`

**Optional** If you would like to test the validity of your manifest file.
**Default**: False

## Example usage

```yaml
uses: marbar3778/atlas@v1
with:
  token: "testKey"
  path: ./example/bank/atlas.toml
  dry-run: ${{ github.event_name != 'pull_request' }}
```
