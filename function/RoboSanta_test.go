package functions

import(
"testing"
)

func TestRoboSanta(t*testing.T){
	path := "^^<<v<<v><>><<"

	want := 7

	got:=RoboSanta(path)
	if want != got{
		t.Errorf("Test RoboSanta Failed Want : %d, Got : %d",want, got)
	}
}