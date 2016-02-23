# Slimmer

Run your testsuite in a container with only 1 command and with a simple bash
script, you are ready to create your perfect build.

It use a `build.slimmer` an executable file to specify your build.

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

# Environemt Variables
We support env vars, you can inject it:
```bash
slimmer -i golang:1.5 -v /Users/gianarb/go/slimmer:/tmp -w /tmp -e SECRET_TOKEN=43t3gse4ts4st4ts4s
```

# Notify 
Trigger a notification when the build fisished

## on mac
```bash
$ go run main.go api \
    -v /Users/gianarb/go/src/github.com/gianarb/slimmer:/tmp \
    -i golang:1.5 && osascript -e 'display notification "Build Successed" with title "Slimmer"' \
    || osascript -e 'display notification "Builld Failed" with title "Slimmer"'
```
