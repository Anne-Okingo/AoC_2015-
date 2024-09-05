package functions

import (
	"testing"
)

func TestMD5(t *testing.T) {
	input := "ckczppom"

	want := 117946

	got := MD5(input)
	if want != got {
		t.Errorf("Test MD5 Failed want : %d got : %d", want, got)
	}
}

func TestMD5b(t *testing.T) {
	input := "ckczppom"

	want := 3938038

	got := MD5b(input)
	if want != got {
		t.Errorf("Test MD5 Failed want : %d got : %d", want, got)
	}
}
