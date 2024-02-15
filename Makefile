# build the main app
.PHONY: build
build: bin resources
	go build -o bin/linkydink main.go

# build a linux version for release
.PHONY: build-release
build-release: dist
	GOOS=linux GOARCH=amd64 go build -o bin/linkydink-linux-amd64 main.go
	cd ./bin && tar -czvf ../dist/linkydink-linux-amd64.tar.gz linkydink-linux-amd64 

# deploy to staging
.PHONY: deploy
deploy: dist/linkydink-linux-amd64.tar.gz
	./script/deploy.sh

# tag 
.PHONY: tag
tag:
	./script/tag.sh

# prepare release and deploy to production
.PHONY: release
release: tag build-release
	DEPLOY_ENV=production ./script/deploy.sh

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
	go run ./cmd/copy/main.go -g='assets/static/**/*' -o res/static
	go run ./cmd/esbuild/main.go assets/main.css --bundle --outfile=res/static/main.css
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
	./script/cert.sh

dist/linkydink-linux-amd64.tar.gz: build-release