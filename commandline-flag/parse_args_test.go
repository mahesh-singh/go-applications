package main

import (
	"bytes"
	"errors"
	"testing"
)

type testConfig struct {
	args   []string
	err    error
	outupt string
	config
}

func TestParseArgs(t *testing.T) {

	tests := []testConfig{
		{
			args: []string{"-h"},
			outupt: `
			A greeter application which prints the name you entered a specified number of times.
			Usage of greeter: <options> [name]
			Options: 
			  -n int
					Number of times to greet
			`,
			err:    errors.New("flag: help requested"),
			config: config{numTime: 0},
		},
		{
			args:   []string{"-n", "10"},
			err:    nil,
			config: config{numTime: 10},
		},
		{
			args:   []string{"-n", "abc"},
			err:    errors.New("invalid value \"abc\" for flag -n: parse error"),
			config: config{numTime: 0},
		},
		{
			args:   []string{"-n", "1", "Mahesh"},
			err:    nil,
			config: config{numTime: 1, name: "mahesh"},
		},
		{
			args:   []string{"-n", "1", "foo", "bar"},
			err:    errors.New("more than one positional argument specified"),
			config: config{numTime: 0},
		},
	}
	byteBuff := new(bytes.Buffer)
	for _, tc := range tests {
		c, err := parseArgs(byteBuff, tc.args)
		if tc.err == nil && err != nil {
			t.Fatalf("Expected error to be: %v, got: %v\n", tc.err, err)
		}
		if tc.err != nil && tc.err.Error() != err.Error() {
			t.Fatalf("Expected error to be: %v, got : %v\n", tc.err, err)
		}

		if c.numTime != tc.numTime {
			t.Fatalf("Expected numtime to be: %v, got: %v", tc.numTime, c.numTime)
		}

		gotMsg := byteBuff.String()
		if len(tc.outupt) != 0 && tc.outupt != gotMsg {
			t.Fatalf("Expected stgout message to be: %v, got: %v", tc.outupt, gotMsg)
		}
		byteBuff.Reset()
	}
}
