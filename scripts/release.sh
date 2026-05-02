#!/usr/bin/env bash
# scripts/release.sh — tag a new release after sanity checks.
#
# Usage: scripts/release.sh v0.1.0
#
# Verifies: clean tree, on main, pushed to origin, version format,
# tag does not yet exist. Then creates an annotated tag and pushes it,
# which triggers .github/workflows/release.yml (goreleaser).
set -euo pipefail

if [ "$#" -ne 1 ]; then
  echo "usage: $0 <vX.Y.Z>" >&2
  exit 2
fi

TAG=$1

if ! [[ "$TAG" =~ ^v[0-9]+\.[0-9]+\.[0-9]+(-[a-zA-Z0-9.-]+)?$ ]]; then
  echo "error: tag must look like vX.Y.Z (optional -prerelease), got '$TAG'" >&2
  exit 1
fi

if [ -n "$(git status --porcelain)" ]; then
  echo "error: working tree not clean — commit or stash first" >&2
  exit 1
fi

BRANCH=$(git rev-parse --abbrev-ref HEAD)
if [ "$BRANCH" != "main" ]; then
  echo "error: must be on main, currently on '$BRANCH'" >&2
  exit 1
fi

if ! git diff --quiet "@{u}"; then
  echo "error: local main differs from origin/main — push first" >&2
  exit 1
fi

if git rev-parse "$TAG" >/dev/null 2>&1; then
  echo "error: tag '$TAG' already exists" >&2
  exit 1
fi

echo "tagging $TAG ..."
git tag -a "$TAG" -m "Release $TAG"
git push origin "$TAG"
echo "pushed $TAG — release workflow will run on GitHub."
