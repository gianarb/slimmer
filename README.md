# Slimmer

Run your testsuite in a container with only 1 command and with a simple bash
script, you are ready to create your perfect build.

Dead simple build executor with superpowers.
The only thing you have to to is create an executable file
with the proper [shebang line](http://unix.stackexchange.com/questions/87560/does-the-shebang-determine-the-shell-which-runs-the-script) named `build.slimmer` under
the project root and enjoy the achievement.

# Install
This project is in written go, so it's distributed as a self contained binary.
There are two ways to obtain the `slimmer` binary:

## Compile From source

```bash
go get github.com/gianarb/slimmer
```

## Download distribution binaries
**coming soon**

# Run Test
We have a WIP test suite

## with go test
```bash
go test -v ./...
```

## with slimmer

```bash
slimmer build -i golang:1.6 -v /Users/gianarb/go/slimmer:/slimmer -w /slimmer
```

In this case you can omit the `-w` parameter because the default value for the working dir is `/slimmer`

# Environment Variables

Environment variables can be injected passed with `-e` 

```bash
slimmer build -i golang:1.5 -v /Users/gianarb/go/slimmer:/tmp -w /tmp -e SECRET_TOKEN=43t3gse4ts4st4ts4s
```

# Notify 
Trigger a notification when the build fisished

## On OS X

```bash
$ go run main.go build \
    -v /Users/gianarb/go/src/github.com/gianarb/slimmer:/tmp \
    -i golang:1.5 && osascript -e 'display notification "Build Successed" with title "Slimmer"' \
    || osascript -e 'display notification "Builld Failed" with title "Slimmer"'
```
