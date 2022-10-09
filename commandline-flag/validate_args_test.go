package main

import (
	"errors"
	"testing"
)

func TestVlidateArgs(t *testing.T) {
	tests := []struct {
		c   config
		err error
	}{
		{c: config{},
			err: errors.New("Specify number greater than 0 to greet")},
		{c: config{numTime: -1},
			err: errors.New("Specify number greater than 0 to greet")},
		{c: config{numTime: 10},
			err: nil},
	}

	for _, tc := range tests {
		err := validateArgs(tc.c)
		if err != nil && tc.err.Error() != err.Error() {
			t.Fatalf("Expected to be: %v, got: %v", tc.err, err)
		}
		if tc.err == nil && err != nil {
			t.Fatalf("Exected to be: %v, got: %v", tc.err, err)
		}
	}
}
