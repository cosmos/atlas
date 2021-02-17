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

Below is an example of the action being used. To see an example of the toml file, see [here](../example/bank/atlas.toml)

```yaml
    steps:
      uses: actions/checkout@v2
      uses: cosmos/atlas@v0.0.3
      with:
        token: "testKey"
        path: ./example/bank/atlas.toml
  dry-run: ${{ github.event_name != 'pull_request' }}
```
