# build the main app
.PHONY: build
build: bin resources
	go build -o bin/linkydink main.go

# build a linux version for release
# under the hood this cross-compiles `make build` using a docker container 
.PHONY: build-release
build-release: dist
	./script/build/release.sh

# deploy 
.PHONY: deploy
deploy: dist/linkydink-linux-amd64.tar.gz
	./script/deploy/deploy.sh

# run the main application
.PHONY: run
run: build .cert/localhost.crt
	./bin/linkydink

# run the main application and reload when files are changed
.PHONY: watch
watch:
	reflex -d fancy -G 'bin/*' -G 'release/*' -G 'res/**/*' -G 'db/*' -s make run


# run tests
.PHONY: test
test:
	go test ./app/... ./pkg/...

.PHONY: resources
resources: res
	go run ./cmd/copy/main.go -g='assets/**/*.tmpl' -o=res/tmpl
	go run ./cmd/copy/main.go -g 'assets/main.css' -o=res/static
	go run ./cmd/esbuild/main.go assets/main.js --bundle --minify --outfile=res/static/main.js

.PHONY: clean
clean:
	rm -rf bin
	rm -rf res
	rm -rf dist

.PHONY: clobber-db
clobber-db:
	rm -rf db

.PHONY: smtp
smtp:
	mailpit

bin:
	mkdir bin

res:
	mkdir res

dist:
	mkdir dist

.cert/localhost.crt:
	./script/build/cert.sh

dist/linkydink-linux-amd64.tar.gz: build-release