package main

import (
	"bufio"
	"errors"
	"flag"
	"fmt"
	"io"
	"os"
)

type config struct {
	numTime int
	name    string
}

func parseArgs(w io.Writer, args []string) (config, error) {
	c := config{}
	fs := flag.NewFlagSet("greeter", flag.ContinueOnError)
	fs.SetOutput(w)
	fs.Usage = func() {
		var usesString = `A greeter application which prints the name you entered a apecified number of times.
		
		Uses of %s: <options> [name]`
		fmt.Fprintf(w, usesString, fs.Name())
		fmt.Fprintln(w)
		fmt.Fprintln(w, "Options: ")
		fs.PrintDefaults()
	}
	fs.IntVar(&c.numTime, "n", 0, "Number of time to greet")
	err := fs.Parse(args)
	if err != nil {
		return c, err
	}
	if fs.NArg() > 1 {
		return c, errors.New("more than one positional argument specified")
	}

	if fs.NArg() == 1 {
		c.name = fs.Arg(0)
	}
	return c, nil
}

func getName(r io.Reader, w io.Writer) (string, error) {
	msg := "Type your name and press enter/return key when done. \n"
	fmt.Fprint(w, msg)
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		return "", err
	}
	name := scanner.Text()
	if len(name) == 0 {
		return "", errors.New("no name found")
	}
	return name, nil
}

func greetUser(c config, w io.Writer) {
	for i := 0; i < c.numTime; i++ {
		fmt.Fprintf(w, "Nice to meet you %s \n", c.name)
	}
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	var err error
	if len(c.name) == 0 {
		c.name, err = getName(r, w)
	}
	if err != nil {
		return err
	}

	greetUser(c, w)
	return nil
}

func validateArgs(c config) error {
	if c.numTime == 0 {
		return errors.New("specify number greater than 0 to greet")
	}
	return nil
}

func main() {
	c, err := parseArgs(os.Stdout, os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}

	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
