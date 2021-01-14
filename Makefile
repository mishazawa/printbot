include .env
export

build:
	go build -o bin/$(BASE) .

run:
	go run .

clean:
	go clean -testcache

test: clean
	go test -v ./...

commit: fmt
	git add .
	git commit --allow-empty

fmt:
	gofmt -s -w .

critic:
	~/go/bin/golangci-lint run
