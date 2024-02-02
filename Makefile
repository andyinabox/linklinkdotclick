.PHONY: build
build: clean dist/linkydink

.PHONY: clean-dist
clean-dist:
	rm -rf dist

.PHONY: clean-res
clean-res:
	rm -rf res

.PHONY: clean
clean: clean-dist clean-res

.PHONY: run
run: dist/linkydink db
	./dist/linkydink

.PHONY: build-docker-test
build-docker-test:
	docker build -t andyinabox/linkydink:test .

.PHONY: run-docker-test
run-docker-test:
	docker run -it -p 127.0.0.1:8080:8080 --rm --name linkydink-test andyinabox/linkydink:test

.PHONY: watch
watch:
	reflex -G 'dist' -G 'res' -G 'db/*' -s make clean run

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
	go run cmd/esbuild/main.go assets/main.js --bundle --outfile=res/static/main.js

res: res/tmpl res/static res/static/main.js

dist/linkydink: res
	go build -o dist/linkydink main.go




