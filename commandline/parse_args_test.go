package main

import (
	"errors"
	"testing"
)

type testConfig struct {
	args []string
	err  error
	config
}

func TestParseArgs(t *testing.T) {

	tests := []testConfig{
		{
			args:   []string{"-h"},
			err:    nil,
			config: config{printUses: true, numTime: 0},
		},
		{
			args:   []string{"10"},
			err:    nil,
			config: config{printUses: false, numTime: 10},
		},
		{
			args:   []string{"abc"},
			err:    errors.New("strconv.Atoi: parsing \"abc\": invalid syntax"),
			config: config{printUses: true, numTime: 0},
		},
		{
			args:   []string{"1", "foo"},
			err:    errors.New("Invalid number of arguments"),
			config: config{printUses: true, numTime: 0},
		},
	}

	for _, tc := range tests {
		c, err := parseArgs(tc.args)
		if tc.err != nil && tc.err.Error() != err.Error() {
			t.Fatalf("Expected error to be: %v, got : %v\n", tc.err, err)
		}
		if tc.err == nil && err != nil {
			t.Fatalf("Expected error to be: %v, got: %v\n", tc.err, err)
		}
		if c.printUses != tc.printUses {
			t.Fatalf("Expected printUses to be: %v, got: %v", tc.printUses, c.printUses)
		}
		if c.numTime != tc.numTime {
			t.Fatalf("Expected numtime to be: %v, got: %v", tc.numTime, c.numTime)
		}
	}
}
