package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		s    string
		err  error
		want int
	}{
		{"no error", nil, 0},
		{"one error", fmt.Errorf("Test-Error"), 1},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			errList := &errorList{}
			errList.add(tt.err)
			require.Lenf(t, errList.errors, tt.want, "add(%v) errorList size = %d, want %d", tt.err, len(errList.errors), tt.want)
		})
	}
}

func TestHasError(t *testing.T) {
	tests := []struct {
		s    string
		err  error
		want bool
	}{
		{"no error", nil, false},
		{"one Error", fmt.Errorf("Test-Error"), true},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			errList := &errorList{}
			errList.add(tt.err)
			require.Equalf(t, tt.want, errList.hasError(), "hasError() returns %v, want %v", errList.hasError(), tt.want)
		})
	}
}

func TestError(t *testing.T) {
	tests := []struct {
		s    string
		err  error
		want string
	}{
		{"no e rror", nil, ""},
		{"one error", fmt.Errorf("Test-Error"), "Number of errors: 1\n1: Test-Error"},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			errList := &errorList{}
			errList.add(tt.err)
			require.Equalf(t, tt.want, errList.Error(), "Error() returns \"%s\", want \"%s\"", errList.Error(), tt.want)
		})
	}
}
