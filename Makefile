.PHONY: build
build: clean dist

.PHONY: clean
clean:
	rm -rf dist
	rm -rf res

.PHONY: run
run: build
	./dist/server

.PHONY: watch
	reflex -s make run

res/public/main.js:
	go run cmd/esbuild/main.go src/main.js --bundle --outfile=res/public/main.js

res/tmpl:
	cp -r src/tmpl res/tmpl

res/data/data.json:
	mkdir -p res/data
	cp db/data.json res/data/data.json

res: res/public/main.js res/data/data.json res/tmpl

dist/server: res
	go build -o dist/server ./cmd/server

dist: dist/server



