package functions

import (
	"testing"
)

func TestFloor(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{{"()()", 0},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"))(", -1},
		{")())())", -3},
	}
	for _, test := range tests {
		got := Floor(test.input)
		if test.want != got {
			t.Errorf("TestFloor Failed : Want %v, Got %v", test.want, got)
		}
	}
}

