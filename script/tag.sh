#!/bin/bash

set -e

ALLOWED_BRANCH=develop

if [[ "$(git branch --show-current)" == "$ALLOWED_BRANCH" ]]; then
  echo " -> ✅ tagging from '$ALLOWED_BRANCH' branch"
else
  echo " -> 🚫 tagging should happen in the '$ALLOWED_BRANCH' branch"
  exit 1
fi

if [[ -z $(git status -s) ]]; then
  echo " -> ✅ no uncommitted changes in this branch"
else
  echo " -> 🚫 won't tag with uncommitted changes on this branch"
  exit 1
fi

VERSION=$(cat ./VERSION)
echo " -> 🏷 tagging $VERSION for release"
git tag $VERSION