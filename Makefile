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
	reflex -G 'dist' -G 'res' -s make run

res/tmpl:
	go run ./cmd/copy -g='assets/**/*.tmpl' -o=res/tmpl

res/static:
	go run ./cmd/copy -g 'assets/static/*' -o=res/static

res/static/main.js:
	go run cmd/esbuild/main.go assets/main.js --bundle --outfile=res/static/main.js

res: clean-res res/tmpl res/static res/static/main.js res/static/data.json


dist/server: res
	go build -o dist/server main.go

dist: clean-dist dist/server



