package src

import (
	"testing"
)

func TestValidPath(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"../files/202MEP_52A_57A_58A.txt", true},
		{"../files/202MEP_52D_57D_58A.txt", true},
		{"../files", true},
		{"../models", true},
		{"../model", false},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			valid, err := isValid(tt.input)
			if err != nil && tt.want != false {
				t.Errorf("Error %s", err)
			}

			if valid != tt.want {
				t.Errorf("File exists: %v, but file exists should be: %v", valid, tt.want)
			}
		})
	}
}
