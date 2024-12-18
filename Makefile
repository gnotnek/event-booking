.PHONY: tidy
tidy:
	go fmt ./...
	go mod tidy -v

.PHONY: build
build:
	mkdir -p tmp/bin
	go build -o tmp/bin/app.exe main.go

.PHONY: run
run: 
	./tmp/bin/app api

.PHONY: run/live
run/live:
	go run github.com/cosmtrek/air@v1.43.0 \
		--build.cmd "make build" --build.bin "./tmp/bin/app $(bin)" --build.delay "100" \
		--build.exclude_dir "" \
		--build.include_ext "go, tpl, tmpl, html, css, scss, js, ts, sql, jpeg, jpg, gif, png, bmp, svg, webp, ico" \
		--misc.clean_on_exit "true"

.PHONY: test
test:
	go test -race -coverprofile="coverage.out" ./...

.PHONY: generate
generate:
	go generate ./...
