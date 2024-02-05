#!/bin/bash

set -e

# cross-compile using docker container
echo "building app in docker container for linux/amd64"
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=linux -e GOARCH=amd64 golang:1.18 make

# rename file
echo "renaming file"
mv bin/linkydink bin/linkydink-linux-amd64

# compress release ind add to dist folder
echo "compressing dist file"
# temporarily change working dir so we can compress the executable without parent dir
pushd ./bin
tar -czvf ../dist/linkydink-linux-amd64.tar.gz linkydink-linux-amd64 
popd