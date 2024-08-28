package functions

import (
	"testing"
)

func TestFloor(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{"()()", 0},
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

func TestPosition(t *testing.T) {
	tests := []struct {
		input        string
		wantposition int
	}{
		{"())", 3},
		{"(()))", 5},
		{"())", 3},
		{")(", 1},
	}
	for _, test := range tests {
		gotposition := Position(test.input)
		if test.wantposition != gotposition {
			t.Errorf("TestPostion Failed : wantposition %v ; gotposition %v", test.wantposition, gotposition)
		}
	}
}
