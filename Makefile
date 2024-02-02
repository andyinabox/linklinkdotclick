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
run: clean dist/linkydink db
	./dist/linkydink

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




