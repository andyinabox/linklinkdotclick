.PHONY: build
build: clean dist

.PHONY: clean-dist
clean-dist:
	rm -rf dist

.PHONY: clean-res
clean-res:
	rm -rf res

.PHONY: clean
clean: clean-dist clean-res

.PHONY: run
run: build
	./dist/server

.PHONY: watch
watch:
	reflex -g '*.go' -s make run

res/main.js:
	go run cmd/esbuild/main.go assets/main.js --bundle --outfile=res/main.js

res: clean-res res/main.js
	cp db/data.json res/data.json
	cp assets/tmpl/* res/

dist/server: res
	go build -o dist/server main.go

dist: clean-dist dist/server



