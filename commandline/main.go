package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
)

type config struct {
	numTime   int
	printUses bool
}

var usesString = fmt.Sprintf(`Uses: %s <integer> [-h | --help]
A greeter application which prints the name you entered <integer> number of times
`, os.Args[0])

func parseArgs(args []string) (config, error) {
	var numTimes int
	var err error
	c := config{}
	if len(args) != 1 {
		c.printUses = true
		return c, errors.New("Invalid number of arguments")
	}

	if args[0] == "-h" || args[0] == "--help" {
		c.printUses = true
		return c, nil
	}

	numTimes, err = strconv.Atoi(args[0])
	if err != nil {
		c.printUses = true
		return c, err
	}
	c.numTime = numTimes
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
		return "", errors.New("No name found")
	}
	return name, nil
}

func greetUser(c config, name string, w io.Writer) {
	for i := 0; i < c.numTime; i++ {
		fmt.Fprintf(w, "Nice to meet you %s \n", name)
	}
}

func runCmd(r io.Reader, w io.Writer, c config) error {
	if c.printUses {
		printUses(w)
		return nil
	}

	name, err := getName(r, w)
	if err != nil {
		return err
	}

	greetUser(c, name, w)
	return nil
}

func printUses(w io.Writer) {

	fmt.Fprintf(w, usesString)
}

func validateArgs(c config) error {
	if c.numTime == 0 {
		return errors.New("Specify number greater than 0 to greet")
	}
	return nil
}

func main() {
	c, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		printUses(os.Stdout)
		os.Exit(1)
	}

	err = validateArgs(c)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		printUses(os.Stdout)
		os.Exit(1)
	}

	err = runCmd(os.Stdin, os.Stdout, c)
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
}
