#!/bin/sh

token=$1
path_manifest=$2
dry_run=$3

atlas login $token

if [[ "$dry_run" == "true" ]]; then
	atlas publish -m $path_manifest --dry-run
  echo "dry-run"
else
	atlas publish -m $path_manifest
fi
