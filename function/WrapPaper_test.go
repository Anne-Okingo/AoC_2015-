package functions

import (
	"testing"
)

func TestWrapPaper(t *testing.T) {
	l, w, h := 30, 28, 5

	want := 2400

	got := WrapPaper(l, w, h)
	if want != got {
		t.Errorf("TestWrapPaper Failed : want %v; got %v", want, got)
	}
}

func TestRibbon(t *testing.T) {
	l, w, h := 30, 28, 5

	want := 4266

	got := Ribbon(l, w, h)
	if want != got {
		t.Errorf("TestRibbon Failed : want %v; got %v", want, got)
	}
}
