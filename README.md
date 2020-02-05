# govalues

- [govalues](#govalues)
  - [Building](#building)
  - [Testing](#testing)
  - [Running](#running)

## Building

- Vet the package, make sure the tests run:

```bash
$ go vet ./... && go clean -testcache && go test ./...
?       github.com/TheDonDope/govalues/cmd/govalues     [no test files]
?       github.com/TheDonDope/govalues/pkg/politics     [no test files]
?       github.com/TheDonDope/govalues/pkg/simulation   [no test files]
```

- Build the main command:

```bash
$ go build ./...
# No output on sucess
```

**Attention:** The built binary is not (yet) ignored by git, so be careful to not commit your binaries to the repository.

## Testing

- Run the testsuite with coverage enabled:

```bash
$ go test ./... -coverprofile=coverage.out
?       github.com/TheDonDope/govalues/cmd/govalues     [no test files]
?       github.com/TheDonDope/govalues/pkg/politics     [no test files]
?       github.com/TheDonDope/govalues/pkg/simulation   [no test files]
```

- Open the results in the browser:

```bash
$ go tool cover -html=coverage.out
# Opens Browser
```

Output files are ignored by git, so don't you even worry :)

## Running

- Run the govalues command:

```bash
$ go run ./cmd/govalues/main.go
# Log output
```

- Run the govalues binary:

```bash
$ ./govalues
# Log output
```
