.PHONY: build
build: clean dist

.PHONY: clean-dist
clean-dist:
	rm -rf dist

.PHONY: clean-res
clean-res:
	rm -rf cmd/server/res

.PHONY: clean
clean: clean-dist clean-res

.PHONY: run
run: build
	./dist/server

.PHONY: watch
watch:
	reflex -g '*.go' -s make run

cmd/server/res/main.js:
	go run cmd/esbuild/main.go src/main.js --bundle --outfile=cmd/server/res/main.js

cmd/server/res: clean-res cmd/server/res/main.js
	cp db/data.json cmd/server/res/data.json
	cp src/tmpl/* cmd/server/res/

dist/server: cmd/server/res
	go build -o dist/server ./cmd/server

dist: clean-dist dist/server



