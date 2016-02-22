package command

import (
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/gianarb/slimmer/runner"
)

type BuildCommand struct {
	Runner runner.Runner
	Logger *log.Logger
}

type envVars []string

func (e *envVars) Set(value string) error {
	*e = append(*e, value)
	return nil
}

func (e *envVars) String() string {
	return fmt.Sprintf("%d", *e)
}

func (c *BuildCommand) Run(args []string) int {
	var image string
	var envVar envVars
	cmdFlags := flag.NewFlagSet("event", flag.ContinueOnError)
	cmdFlags.StringVar(&image, "i", "", "Start docker image")
	cmdFlags.Var(&envVar, "e", "List envaironment variables")

	if err := cmdFlags.Parse(args); err != nil {
		c.Logger.Fatal(err)
	}

	containerId, err := c.Runner.BuildContainer(image, envVar)
	if err != nil {
		c.Runner.RemoveContainer(containerId)
		c.Logger.Fatal(err)
	}
	c.Logger.Printf("Build %s started - source %s", containerId[0:12], image)

	exitCode, err := c.Runner.Exec(containerId, []string{"bash", "./build.sh"})
	if err != nil {
		c.Runner.RemoveContainer(containerId)
		c.Logger.Fatal(err)
	}
	c.Runner.RemoveContainer(containerId)

	return exitCode
}

func (c *BuildCommand) Help() string {
	helpText := `
Usage: start slimmer build.
Options:
	-i=			Start docker image
	-e=			Start docker image
`
	return strings.TrimSpace(helpText)
}

func (r *BuildCommand) Synopsis() string {
	return "Run local build"
}
