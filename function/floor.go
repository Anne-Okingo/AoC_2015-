package functions

func Floor(str string) int {
	steps := 0

	for _, ch := range str {

		if ch == '(' {
			steps++
		} else {
			steps--
		}
	}
	return steps
}
