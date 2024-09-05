package functions

import (
	"testing"
)

// Test for Contains function (true case)
func TestContainsTrue(t *testing.T) {
	input := "rthkunfaakmwmush"
	want := true
	
	got := Contains(input)
	if got != want {
		t.Errorf("TestContainsTrue failled : got : %v want : %v", want, got)
	}
}

// Test for Contains function (false case)
func TestContainsFalse(t *testing.T) {
	input := "abcdef" // Contains "ab", which is forbidden
	want := false

	if got := Contains(input); got != want {
		t.Errorf("TestContainsFalse failled : got : %v want : %v", want, got)
	}
}

// Test for IsVowel function (true case)
func TestIsVowelTrue(t *testing.T) {
	input := "rthkunfaakmwmush"
	want := true

	if got := IsVowel(input); got != want {
		t.Errorf("TestIsVowelTrue failled : got : %v want : %v", want, got)
	}
}

// Test for IsVowel function (false case)
func TestIsVowelFalse(t *testing.T) {
	input := "bcdfghjklmnpqrstvwxyz" // Contains no vowels
	want := false

	if got := IsVowel(input); got != want {
		t.Errorf("TestIsVowelFalse failled : got : %v want : %v", want, got)
	}
}

// Test for IsDouble function (true case)
func TestIsDoubleTrue(t *testing.T) {
	input := "rthkunfaakmwmush"
	want := true

	if got := IsDouble(input); got != want {
		t.Errorf("TestIsDoubleTrue failled : got : %v want : %v", want, got)
	}
}

// Test for IsDouble function (false case)
func TestIsDoubleFalse(t *testing.T) {
	input := "abcdef" // No repeating characters
	want := false

	if got := IsDouble(input); got != want {
		t.Errorf("TestIsDoubleFalse failled : got : %v want : %v", want, got)
	}
}

// Test for IsNiceString function (true case)
func TestIsNiceStringTrue(t *testing.T) {
	input := "rthkunfaakmwmush"
	want := true

	if got := IsNiceString(input); got != want {
		t.Errorf("TestIsNiceStringTrue failled : got : %v want : %v", want, got)
	}
}

// Test for IsNiceString function (false case)
func TestIsNiceStringFalse(t *testing.T) {
	input := "abcdefgh" // Contains "ab" and no repeating characters
	want := false

	if got := IsNiceString(input); got != want {
		t.Errorf("TestIsNiceStringFalse failled : got : %v want : %v", want, got)
	}
}

// Test for CountNiceString function (true case)
func TestCountNiceStringTrue(t *testing.T) {
	input := []string{"rthkunfaakmwmush"}
	want := 1

	if got := CountNiceString(input); got != want {
		t.Errorf("TestCountNiceStringTrue failled : got : %v want : %v", want, got)
	}
}

// Test for CountNiceString function (false case)
func TestCountNiceStringFalse(t *testing.T) {
	input := []string{"abcdefgh", "pqxyz"} // Both are not "nice" strings
	want := 0

	if got := CountNiceString(input); got != want {
		t.Errorf("TestCountNiceStringFalse failled : got : %v want : %v", want, got)
	}
}

