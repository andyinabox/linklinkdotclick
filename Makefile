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

res/tmpl:
	go run ./cmd/copy -g='assets/**/*.tmpl' -o=res/tmpl

res/public:
	go run ./cmd/copy -g 'assets/public/*' -o=res/public

res/data/data.json:
	go run ./cmd/copy -g 'db/data.json' -o=res/data

res/public/main.js:
	go run cmd/esbuild/main.go assets/main.js --bundle --outfile=res/public/main.js

res: clean-res res/tmpl res/public res/public/main.js res/data/data.json


dist/server: res
	go build -o dist/server main.go

dist: clean-dist dist/server



