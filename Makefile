.PHONY: build
build: clean dist/linkydink

.PHONY: build-linux-amd64
build-linux-amd64:
	./script/build-linux-amd64.sh

.PHONY: deploy
deploy: build-linux-amd64
	scp ./dist/linkydink andy@`doctl compute droplet get reading-dot-andydayton-dot-com --template {{.PublicIPv4}}`:/home/andy/bin/linkydink
	make clean-dist

# .PHONY: docker-test
# # note that this is for linux environments
# docker-test:
# 	docker build --tag andyinabox/linkydink:test .
# 	docker run --rm -p 8080:8080 -v database:/db andyinabox/linkydink:test

# .PHONY: docker-save-image
# docker-build-for-deploy:
# 	docker buildx build --platform linux/amd64/v3 --tag andyinabox/linkydink:deploy .
# 	docker save -o dist/andyinabox-linkydink-deploy.tar andyinabox/linkydink:deploy

.PHONY: clean-dist
clean-dist:
	rm -rf dist

.PHONY: clean-res
clean-res:
	rm -rf res

.PHONY: clean
clean: clean-dist clean-res

.PHONY: run
run: clean dist/linkydink db
	./dist/linkydink

.PHONY: build-docker-test
build-docker-test:
	docker build -t andyinabox/linkydink:test .

.PHONY: run-docker-test
run-docker-test:
	docker run -it -p 127.0.0.1:8080:8080 --rm --name linkydink-test andyinabox/linkydink:test

.PHONY: watch
watch:
	reflex -d fancy -G 'dist' -G 'res' -G 'db/*' -s make run

.PHONY: test
test:
	go test ./app/...

.PHONY: test-verbose
test-verbose:
	go test -v ./app/...

db:
	mkdir -p db

res/tmpl:
	go run ./cmd/copy -g='assets/**/*.tmpl' -o=res/tmpl

res/static:
	go run ./cmd/copy -g 'assets/static/*' -o=res/static

res/static/main.js:
	go run cmd/esbuild/main.go assets/main.js --bundle --minify --outfile=res/static/main.js

res/static/main.css:
	go run ./cmd/copy -g 'assets/main.css' -o=res/static

res: res/tmpl res/static res/static/main.js res/static/main.css

dist/linkydink: res
	go build -o dist/linkydink main.go




