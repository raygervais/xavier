package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/raygervais/xavier/client/pkg/api"
	"github.com/raygervais/xavier/client/pkg/cli"
	"github.com/raygervais/xavier/client/pkg/conf"
)

var (
	helpFlag = flag.Bool("help", false, "Display application usage documentation")
)

func main() {
	// Determine and configure the application for the user's OS.
	configPath, err := conf.DetermineStorageLocation()
	if err != nil {
		fmt.Printf("Could not determine configuration location: %s\n", err)
	}

	if err := conf.InitializeConfigurationLocation(configPath); err != nil {
		fmt.Printf("Could not initialize configuration location: %s", err)
	}

	cli := cli.Init(api.Init(conf.ApplicationConfiguration))

	// Flag for --help
	flag.Parse()

	if *helpFlag {
		cli.Help()
		os.Exit(0)
	}

	if err := cli.Handler(); err != nil {
		fmt.Printf("Command error: %s\n", err)
	}
}
