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

func Position(str string) int {
	steps := 0
	position := 0
	positions := []int{}

	for i, ch := range str {
		position = i + 1

		if ch == '(' {
			steps++
		} else {
			steps--
		}
		if i < len(str) && steps == -1 {
			positions = append(positions, position)
		}
	}
	return positions[0]
}
