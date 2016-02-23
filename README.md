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
slimmer -i golang:1.5 -v /Users/gianarb/go/slimmer:/slimmer -w /slimmer
```
In this case you can omit `-w` because the default value is `/slimmer`

# Notify 
Trigger a notification when the build fisished

## on mac
```bash
$ go run main.go api \
    -v /Users/gianarb/go/src/github.com/gianarb/slimmer:/tmp \
    -i golang:1.5 && osascript -e 'display notification "Build Successed" with title "Slimmer"' \
    || osascript -e 'display notification "Builld Failed" with title "Slimmer"'
```
