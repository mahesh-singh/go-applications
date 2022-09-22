package main

import (
	"bytes"
	"errors"
	"strings"
	"testing"
)

func TestRunCmd(t *testing.T) {
	tests := []struct {
		c      config
		input  string
		outupt string
		err    error
	}{
		{
			c:      config{printUses: true},
			outupt: usesString,
		},
		{
			c:      config{numTime: 3},
			input:  "",
			outupt: strings.Repeat("Type your name and press enter/return key when done. \n", 1),
			err:    errors.New("No name found"),
		},
		{
			c:      config{numTime: 5},
			input:  "Mahesh",
			outupt: "Type your name and press enter/return key when done. \n" + strings.Repeat("Nice to meet you Mahesh\n", 5),
		},
	}
	byteBuff := new(bytes.Buffer)
	for _, tc := range tests {
		r := strings.NewReader(tc.input)
		err := runCmd(r, byteBuff, tc.c)
		if err != nil && tc.err == nil {
			t.Fatalf("Expected error: %v, Got error: %v", tc.err.Error(), err.Error())
		}
		if tc.err != nil && tc.err.Error() != err.Error() {
			t.Fatalf("Expected: %v, got %v", tc.err.Error(), err.Error())
		}
		gotMsg := byteBuff.String()
		if gotMsg != tc.outupt {
			t.Fatalf("Expected: %s, got: %s", tc.outupt, gotMsg)
		}
		byteBuff.Reset()
	}
}
