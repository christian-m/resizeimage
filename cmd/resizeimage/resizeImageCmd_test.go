package main

import (
	. "bitbucket.org/christian-m/resizeimage/internal/resize"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

func TestParseSize(t *testing.T) {
	tests := []struct {
		s       string
		want    PicSize
		wantErr bool
	}{
		{"500x500", PicSize{500, 500, 0}, false},
		{"500-500", PicSize{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			got, err := parseSize(tt.s)
			if tt.wantErr {
				require.Falsef(t, err == nil, "parseSize(%s) error = %v, wantErr %v", tt.s, err, tt.wantErr)
			} else {
				require.Truef(t, err == nil, "parseSize(%s) error = %v, wantErr %v", tt.s, err, tt.wantErr)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("parseSize(%s) error = %v, wantErr %v", tt.s, err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseSize(%s)=%v want %v", tt.s, got, tt.want)
			}
		})
	}
}
