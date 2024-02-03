.PHONY: build
build: clean bin/linkydink

.PHONY: dist
dist: bin/linkydink-linux-amd64
	-mkdir dist
	tar -czvf dist/linkydink-linux-amd64.tar.gz bin/linkydink-linux-amd64

.PHONY: deploy
deploy: clean-bin bin/linkydink-linux-amd64
	./script/deploy.sh

.PHONY: clean-bin
clean-bin: bin/clean
	./bin/clean -g 'bin/linkydink*'

.PHONY: clean-res
clean-res: bin/clean
	./bin/clean -g 'res/**/*'

.PHONY: clean
clean: clean-bin clean-res

.PHONY: run
run: clean resources
	go run .

.PHONY: watch
watch:
	reflex -d fancy -G 'bin/**/*' -G 'res/**/*' -G 'db/**/*' -s make run

.PHONY: test
test:
	go test ./app/...

.PHONY: resources
resources: res/static/main.js res/static/main.css res/tmpl

bin:
	mkdir bin

res:
	mkdir res

res/tmpl: res bin/copy
	./bin/copy -g='assets/**/*.tmpl' -o=res/tmpl

res/static/main.js: res bin/esbuild
	./bin/esbuild assets/main.js --bundle --minify --outfile=res/static/main.js

res/static/main.css: res bin/copy
	./bin/copy -g 'assets/main.css' -o=res/static

bin/clean: bin
	go build -o bin/clean cmd/clean/main.go

bin/copy: bin
	go build -o bin/copy cmd/copy/main.go

bin/esbuild: bin
	go build -o bin/esbuild cmd/esbuild/main.go

bin/linkydink: bin resources
	go build -o bin/linkydink main.go

# this actually just runs `make build` inside a docker container
bin/linkydink-linux-amd64:
	./script/build-linux-amd64.sh
