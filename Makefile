run: build
	@./bin/govalues

install:
	go install golang.org/x/tools/cmd/godoc@latest

build:
	go build \
	  -v \
	  -o ./bin/govalues \
	  ./cmd/govalues/main.go

clean:
	rm -f ./bin/govalues
	rm -f coverage.html
	rm -f coverage.out
	rm -rf tmp
	rm -rf vendor

doc:
	godoc

test:
	go test -race -v ./... -coverprofile coverage.out

test-ci:
	go test -race -v ./... -coverprofile coverage.out -covermode=atomic
	# bash -c "bash <(curl -s https://codecov.io/bash)"

cover: test
	go tool cover -html coverage.out -o coverage.html

show-cover: cover
	open coverage.html

vet:
	go vet ./...


