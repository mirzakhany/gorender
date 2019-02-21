#!/bin/bash

set -e

# Get rid of existing binaries
rm -f dist/$PROJECT_NAME*

echo "Building default binary"

# Set GOPATH and GOBIN
export GOPATH="$HOME/go"
export GOBIN="$GOPATH/src/$PKG_SRC/dist"

cd $GOPATH/src/$PKG_SRC

# get git hash, short and long and commit date
LONG_HASH=$(git log -n1 --pretty="format:%H" | cat)
SHORT_HASH=$(git log -n1 --pretty="format:%h"| cat)
BUILD_DATE=$(date "+%D/%H/%I/%S"| sed -e "s/\//-/g")
COMMIT_COUNT=$(git rev-list HEAD --count| cat)
COMMIT_DATE=$(git log -n1 --date="format:%D-%H-%I-%S" --pretty="format:%cd"| sed -e "s/\//-/g")

# ldflags to set version hash and commit hash in version pkg
FLAGS="-X $PKG_SRC/internal/app/$PROJECT_NAME.LongHash=$LONG_HASH 
-X $PKG_SRC/internal/app/$PROJECT_NAME.ShortHash=$SHORT_HASH 
-X $PKG_SRC/internal/app/$PROJECT_NAME.CommitCount=$COMMIT_COUNT 
-X $PKG_SRC/internal/app/$PROJECT_NAME.CommitDate=$COMMIT_DATE 
-X $PKG_SRC/internal/app/$PROJECT_NAME.BuildDate=$BUILD_DATE"

# Build the default application
CGO_ENABLED=0 go install -v -ldflags "$FLAGS" ./...