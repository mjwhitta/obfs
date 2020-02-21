package main

import (
	"os"
	"strings"

	"gitlab.com/mjwhitta/cli"
	hl "gitlab.com/mjwhitta/hilighter"
	"gitlab.com/mjwhitta/where"
)

// Exit status
const (
	Good             int = 0
	InvalidOption    int = 1
	InvalidArgument  int = 2
	MissingArguments int = 3
	ExtraArguments   int = 4
	Exception        int = 5
)

// Flags
type cliFlags struct {
	nocolor bool
	verbose bool
	version bool
}

var flags cliFlags

func init() {
	// Configure cli package
	cli.Align = true
	cli.Authors = []string{"Miles Whittaker <mj@whitta.dev>"}
	cli.Banner = hl.Sprintf("%s [OPTIONS] <cmd>", os.Args[0])
	cli.BugEmail = "where.bugs@whitta.dev"
	cli.ExitStatus = strings.Join(
		[]string{
			"Normally the exit status is 0. In the event of an error",
			"the exit status will be one of the below:\n\n",
			"1: Invalid option\n",
			"2: Invalid argument\n",
			"3: Missing arguments\n",
			"4: Extra arguments\n",
			"5: Exception",
		},
		" ",
	)
	cli.Info = strings.Join(
		[]string{"Simple which-like example binary."},
		" ",
	)
	cli.SeeAlso = []string{"command", "which"}
	cli.Title = "Where"

	// Parse cli flags
	cli.Flag(
		&flags.nocolor,
		"no-color",
		false,
		"Disable colorized output.",
	)
	cli.Flag(
		&flags.verbose,
		"v",
		"verbose",
		false,
		"Show show stacktrace if error.",
	)
	cli.Flag(&flags.version, "V", "version", false, "Show version.")
	cli.Parse()
}

// Process cli flags and ensure no issues
func validate() {
	hl.Disable(flags.nocolor)

	// Short circuit if version was requested
	if flags.version {
		hl.Printf("where version %s\n", where.Version)
		os.Exit(Good)
	}

	// Validate cli flags
	if cli.NArg() < 1 {
		cli.Usage(MissingArguments)
	} else if cli.NArg() > 1 {
		cli.Usage(ExtraArguments)
	}
}
