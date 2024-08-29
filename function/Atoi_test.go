package functions

import (
	"testing"
)

func TestAtoi(t *testing.T) {

	tests := []struct {
		input string
		want  int
	}{
		{"-254", -254},
		{"254", 254},
		{"+254", 254},
	}

	for _, test := range tests {
		got := Atoi(test.input)
		if test.want != got {
			t.Errorf("TestAtoi Failed : Want : %v, Got : %v", test.want, got)
		}
	}
}
