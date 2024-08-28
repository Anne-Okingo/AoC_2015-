package functions

func Atoi(str string) int {
	multiplier := 1
	result := 0
	for i, ch := range str {
		if i == 0 && ch == '-' {
			multiplier = -1
			str = str[1:]
		}
		if i == 0 && ch == '+' {
			multiplier = 1
			str = str[1:]
		}
		if ch < '0' || ch > '9' {
			continue
		} else {
			result = result*10 + int(ch-'0')
		}

	}
	return multiplier * result
}
