// Sample command-line application.
//
// Usage:
//
//	go-cli-sample <action> [parameters...]
//
// Each action defines its own parameters (flags and/or positional arguments).
package main

import (
	"fmt"
	"os"
	"sort"
)

// command describes an action available in the CLI.
type command struct {
	description string
	run         func(args []string) error
}

var commands = map[string]command{
	"greet":   {"Prints a greeting", cmdGreet},
	"sum":     {"Sums a list of numbers", cmdSum},
	"reverse": {"Reverses a text", cmdReverse},
}

func main() {
	if len(os.Args) < 2 {
		usage()
		os.Exit(2)
	}

	action, args := os.Args[1], os.Args[2:]

	if action == "help" || action == "-h" || action == "--help" {
		usage()
		return
	}

	cmd, ok := commands[action]
	if !ok {
		fmt.Fprintf(os.Stderr, "error: unknown action %q\n\n", action)
		usage()
		os.Exit(2)
	}

	if err := cmd.run(args); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage: go-cli-sample <action> [parameters...]")
	fmt.Fprintln(os.Stderr, "\nAvailable actions:")

	names := make([]string, 0, len(commands))
	for name := range commands {
		names = append(names, name)
	}
	sort.Strings(names)

	for _, name := range names {
		fmt.Fprintf(os.Stderr, "  %-10s %s\n", name, commands[name].description)
	}
	fmt.Fprintln(os.Stderr, "  help       Shows this message")
	fmt.Fprintln(os.Stderr, "\nUse \"go-cli-sample <action> -h\" to see the parameters of each action.")
}
