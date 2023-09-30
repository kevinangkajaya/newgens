package src

import (
	"testing"
)

func TestReadLines(t *testing.T) {
	var tests = []struct {
		input string
		want  bool
	}{
		{"../files/202MEP_52A_57A_58A.txt", true},
		{"../files/202MEP_52D_57D_58A.txt", true},
		{"../files", false},
		{"../models", false},
		{"../model", false},
	}
	// The execution loop
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result, err := ReadLines(tt.input)
			if err != nil {
				if result != nil && tt.want != false {
					t.Errorf("Error %s", err)
				}
			}

			lengthValid := len(result) > 0
			if lengthValid != tt.want {
				t.Errorf("Length valid: %v, but want: %v", lengthValid, tt.want)
			}
		})
	}
}
