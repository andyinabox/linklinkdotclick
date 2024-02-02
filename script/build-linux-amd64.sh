#!/bin/bash
# this command fails in the Makefile for some reason
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=linux -e GOARCH=amd64 golang:1.18 make
