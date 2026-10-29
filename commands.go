package main

import (
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

// cmdGreet: go-cli-sample greet -name Mary -times 2 -shout
func cmdGreet(args []string) error {
	fs := flag.NewFlagSet("greet", flag.ContinueOnError)
	name := fs.String("name", "world", "name of the person to greet")
	times := fs.Int("times", 1, "number of repetitions")
	shout := fs.Bool("shout", false, "prints the greeting in uppercase")
	fs.Usage = func() {
		fmt.Fprintln(fs.Output(), "Usage: go-cli-sample greet [-name NAME] [-times N] [-shout]")
		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		return ignoreHelp(err)
	}
	if *times < 1 {
		return errors.New("-times must be greater than zero")
	}

	msg := fmt.Sprintf("Hello, %s!", *name)
	if *shout {
		msg = strings.ToUpper(msg)
	}
	for i := 0; i < *times; i++ {
		fmt.Println(msg)
	}
	return nil
}

// cmdSum: go-cli-sample sum 1 2 3.5
func cmdSum(args []string) error {
	fs := flag.NewFlagSet("sum", flag.ContinueOnError)
	decimals := fs.Int("decimals", 2, "decimal places in the result")
	fs.Usage = func() {
		fmt.Fprintln(fs.Output(), "Usage: go-cli-sample sum [-decimals N] <number> [number...]")
		fs.PrintDefaults()
	}

	if err := fs.Parse(args); err != nil {
		return ignoreHelp(err)
	}
	if fs.NArg() == 0 {
		return errors.New("provide at least one number")
	}

	var total float64
	for _, a := range fs.Args() {
		n, err := strconv.ParseFloat(a, 64)
		if err != nil {
			return fmt.Errorf("%q is not a valid number", a)
		}
		total += n
	}

	fmt.Printf("%.*f\n", *decimals, total)
	return nil
}

// cmdReverse: go-cli-sample reverse "some text"
func cmdReverse(args []string) error {
	fs := flag.NewFlagSet("reverse", flag.ContinueOnError)
	fs.Usage = func() {
		fmt.Fprintln(fs.Output(), "Usage: go-cli-sample reverse <text...>")
	}

	if err := fs.Parse(args); err != nil {
		return ignoreHelp(err)
	}
	if fs.NArg() == 0 {
		return errors.New("provide the text to reverse")
	}

	runes := []rune(strings.Join(fs.Args(), " "))
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	fmt.Println(string(runes))
	return nil
}

// ignoreHelp treats -h/--help as success, since the FlagSet already printed the usage.
func ignoreHelp(err error) error {
	if errors.Is(err, flag.ErrHelp) {
		return nil
	}
	return err
}
