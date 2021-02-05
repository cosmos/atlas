# Publishing

There are two ways you can publish a module to atlas.

1. By downloading atlas locally:

   1. Create or update an existing manifest [manifest](./manifest.md).
   2. Ensure you have the atlas binary installed.
   3. If you haven't set your credentials locally:
        1. Obtain an API token by logging into the web app and either creating or use an existing API token.
        2. Set your credentials via `atlas login [token]`.
   4. Publish your module manifest via `atlas publish -m [path/to/manifest]`. Note,
      you can execute a `--dry-run` first to ensure the manifest is valid.

2. By using docker:

   1. Run: `docker run -v $(shell pwd):/workspace --workdir /workspace interchainio/atlas:latest [APIkey] [path/to/manifest]] [dry-run, default false]`
