#!/bin/bash

set -e

ALLOWED_BRANCH=main

if [[ -z $(git status -s) ]]; then
  echo " -> ✅ no uncommitted changes in this branch"
else
  echo " -> 🚫 won't tag with uncommitted changes on this branch"
  exit 1
fi

if [[ "$(git branch --show-current)" == "$ALLOWED_BRANCH" ]]; then
  echo " -> ✅ tagging from '$ALLOWED_BRANCH' branch"
else
  echo " -> 🚫 tagging should happen in the '$ALLOWED_BRANCH' branch"
  exit 1
fi

git pull origin $ALLOWED_BRANCH
git fetch --tags

VERSION=$(cat ./VERSION)
echo " -> tagging 🏷 $VERSION for release"
git tag $VERSION

git push --tags