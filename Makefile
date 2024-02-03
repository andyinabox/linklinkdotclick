.PHONY: build
build: clean dist/linkydink

.PHONY: deploy
deploy: clean-dist dist/linkydink-linux-amd64
	./script/deploy.sh

.PHONY: clean-dist
clean-dist:
	go run ./cmd/clean -g 'dist/*'

.PHONY: clean-res
clean-res:
	go run ./cmd/clean -g 'res/**/*'

.PHONY: clean
clean: clean-dist clean-res

.PHONY: run
run: clean resources
	go run .

.PHONY: watch
watch:
	reflex -d fancy -G 'dist/**/*' -G 'res/**/*' -G 'db/**/*' -s make run

.PHONY: test
test:
	go test ./app/...

.PHONY: test-verbose
test-verbose:
	go test -v ./app/...

.PHONY: resources
resources: res/static/main.js res/static/main.css res/tmpl

# .PHONY: resources
# resources:
# 	go run cmd/esbuild/main.go assets/main.js --bundle --minify --outfile=res/static/main.js
# 	cp assets/main.css res/static/main.css
# 	go run ./cmd/copy -g='assets/**/*.tmpl' -o=res/tmpl

dist:
	mkdir dist

res:
	mkdir res

res/tmpl: res
	go run ./cmd/copy -g='assets/**/*.tmpl' -o=res/tmpl

# res/static: res
# 	go run ./cmd/copy -g 'assets/static/*' -o=res/static

res/static/main.js: res
	go run cmd/esbuild/main.go assets/main.js --bundle --minify --outfile=res/static/main.js

res/static/main.css: res
	go run ./cmd/copy -g 'assets/main.css' -o=res/static

dist/linkydink: resources
	go build -o dist/linkydink main.go

dist/linkydink-linux-amd64:
	./script/build-linux-amd64.sh


