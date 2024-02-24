#!/bin/bash

VERSION=$(cat ./VERSION)

CSS_INFILE=assets/main.css
CSS_OUTFILE="res/static/main.$VERSION.css"
JS_INFILE=assets/main.js
JS_OUTFILE="res/static/main.$VERSION.js"

go run ./cmd/copy/main.go -g='assets/**/*.tmpl' -o=res/tmpl
go run ./cmd/copy/main.go -g='assets/static/**/*' -o res/static
go run ./cmd/copy/main.go -g='assets/main.css' -o res/static

if [[ "${DEPLOY_ENV}" == "production" ]]; then
	go run ./cmd/esbuild/main.go $CSS_INFILE --bundle --outfile=$CSS_OUTFILE
	go run ./cmd/esbuild/main.go $JS_INFILE --bundle --minify --outfile=$JS_OUTFILE
else
	go run ./cmd/esbuild/main.go $CSS_INFILE --bundle --outfile=$CSS_OUTFILE
	go run ./cmd/esbuild/main.go $JS_INFILE --bundle --outfile=$JS_OUTFILE
fi
