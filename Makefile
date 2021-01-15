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

commit: fmt critic
	git add .
	git commit --allow-empty

fmt:
	goimports -l -w .
	gofmt -s -w .

critic:
	~/go/bin/golangci-lint run
