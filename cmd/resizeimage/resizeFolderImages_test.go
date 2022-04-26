package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUseFile(t *testing.T) {
	tests := []struct {
		filename string
		want     bool
	}{
		{"a/b/file.JPG", true},
		{"file.jpeg", true},
		{"kein_bild.ABC", false},
		{"keine_endung", false},
	}
	for _, tt := range tests {
		t.Run(tt.filename, func(t *testing.T) {
			require.Equalf(t, tt.want, useFile(tt.filename), "useFile(%s) = %v, want %v", tt.filename, useFile(tt.filename), tt.want)
		})
	}

}
