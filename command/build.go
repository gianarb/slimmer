package command

import (
	"strings"
	"flag"
	"github.com/gianarb/slimmer/runner"
	"log"
)

type BuildCommand struct {
	Runner runner.Runner
	Logger *log.Logger
}

func (c *BuildCommand) Run(args []string) int {
	cmdFlags := flag.NewFlagSet("event", flag.ContinueOnError)
	if err := cmdFlags.Parse(args); err != nil {
		c.Logger.Fatal(err)
	}

	containerId, err := c.Runner.BuildContainer("cf/php:7", []string{})
	if err != nil {
		c.Runner.RemoveContainer(containerId)
		c.Logger.Fatal(err)
	}
	c.Logger.Printf("Build %s started - source %s", containerId[0:12], "cf/php:7")

	c.Runner.Exec(containerId, []string{"bash", "./build.sh"})
	c.Runner.RemoveContainer(containerId)

	return 0
}

func (c *BuildCommand) Help() string {
	helpText := `
Usage: start slimmer build.
Options:
`
	return strings.TrimSpace(helpText)
}

func (r *BuildCommand) Synopsis() string {
	return "Run local build"
}
