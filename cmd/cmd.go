package cmd

import "github.com/alecthomas/kong"

// Build information set at link time.
// See Makefile to know the assigned values.
var (
	BuildName        string
	BuildDescription string
	BuildVersion     string
	BuildDate        string
)

// A Root defines the base structure of CLI commands.
type Root struct {
	Version kong.VersionFlag `help:"Show the current version."`
}

// Parse parses and executes the command line interface.
func Parse() {
	var root Root
	cli := kong.Parse(&root,
		kong.Name(BuildName),
		kong.Description(BuildDescription),
		kong.Vars{"version": BuildVersion},
		kong.UsageOnError(),
	)
	cli.FatalIfErrorf(cli.Run())
}
