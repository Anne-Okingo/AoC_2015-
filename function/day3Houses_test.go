package functions

import(
"testing"
)

func TestHouses(t *testing.T){
	input := "^^<<v<<v>"
	want := 10

	got := Houses(input)
	if want != got {
		t.Errorf("Test Houses failed : want :%d, got : %d", want, got)
	}
}