package cli

import (
	"fmt"
	"os"

	"github.com/raygervais/xavier/client/pkg/api"
)

type CommandLineInterface struct {
	commands map[string]func() func(string) error
	api      api.API
}

// Init creates the CLI instance which handles all user input
func Init(api api.API) CommandLineInterface {
	cli := CommandLineInterface{
		api: api,
	}

	cli.commands = map[string]func() func(string) error{
		"health": cli.health,
	}

	return cli
}

// Handler is the primary entry point after initialization,
// validating the provided commands and then calling their respective
// callbacks.
func (cli CommandLineInterface) Handler() error {
	// Validate that we aren't calling the application haphazardly.
	if len(os.Args) == 1 {
		return fmt.Errorf("no command provided:  %s", os.Args)
	}

	action := os.Args[1]
	cmd, ok := cli.commands[action]
	if !ok {
		return fmt.Errorf("invalid command provided: '%s'", action)
	}

	return cmd()(action)
}

// Help displays the appropriate help message per function.
func (cli CommandLineInterface) Help() {
	var help string

	for name := range cli.commands {
		help += name + "\t --help\n"
	}

	fmt.Printf("Usage of %s:\n <command> [<args>]\n%s", os.Args[0], help)
}

func (cli CommandLineInterface) health() func(string) error {
	return func(cmd string) error {
		return cli.api.HealthCheck()
	}
}
