package main

import (
	dockerRun "github.com/gianarb/slimmer/runner/docker"
	"github.com/mitchellh/cli"
	"github.com/gianarb/slimmer/command"
	"github.com/gianarb/slimmer/logger"
	"github.com/gianarb/slimmer/runner/stream"
	"github.com/fsouza/go-dockerclient"
	"log"
	"os"
)

func main() {
	logger := log.New(&logger.Console{}, "", 1)
	c := cli.NewCLI("slimmer", "0.0.0")
    c.Args = os.Args[1:]

	client, err := docker.NewClientFromEnv()

	if err != nil {
		logger.Fatal(err)
	}

	s := stream.ConsoleStream{}
	dockerRunner := dockerRun.DockerRunner{client, s}

    c.Commands = map[string]cli.CommandFactory{
        "build": func() (cli.Command, error) {
			return &command.BuildCommand{&dockerRunner, logger}, nil;
		},
    }

    exitStatus, _ := c.Run()

    os.Exit(exitStatus)
}
