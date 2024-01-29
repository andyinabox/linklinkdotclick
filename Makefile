.PHONY: build
build: clean dist

res/public/main.js:
	go run cmd/esbuild/main.go src/main.js --bundle --outfile=res/public/main.js

res/tmpl:
	cp -r src/tmpl res/tmpl

res: res/public/main.js res/tmpl

dist/server: res
	go build -o dist/server ./cmd/server

dist: dist/server

.PHONY: clean
clean:
	rm -rf dist
	rm -rf res