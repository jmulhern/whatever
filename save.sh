#!/usr/bin/env bash

version=${1}
if [ "${version}" == "" ]; then
  git describe --abbrev=0 --tags
  exit 1
fi

message=${2:-"nothing important"}

git add .
git commit -m "${version} - ${message}"
git tag "v${version}"
git push origin main --tags

