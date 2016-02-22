# Slimmer

Run your testsuite in a container with only 1 command and with a simple bash
script, you are ready to create your perfect build.

# Install
This project is in go, you can use this repo or my compiled binary.

## From source
```bash
go get github.com/gianarb/slimmer
```

## From compiled
WIP

# Run Test
We have a WIP test suite

## with go test
```bash
go test -v ./...
```

## with slimmer
```bash
slimmer -i golang:1.5 -v /Users/gianarb/go/slimmer:/tmp -w /tmp
```
In this case you can omit `-w` because the default value is `/tmp`
