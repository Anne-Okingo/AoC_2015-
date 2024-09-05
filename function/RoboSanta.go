package functions

import (
	"fmt"
)

func RoboSanta(s string) int {
	mappedH := make(map[string]bool)

	sx, sy := 0, 0
	rx, ry := 0, 0
	// mappedH[fmt.Sprintf("%d, %d", sx, sy)] = true
	// mappedH[fmt.Sprintf("%d, %d", rx, ry)] = true
	mappedH["0,0"] = true

	for i := 0; i < len(s); i++ {
		if i%2 == 0 {
			switch s[i] {
			case '^':
				sy++
			case 'v':
				sy--
			case '>':
				sx++
			case '<':
				sx--
			}
			mappedH[fmt.Sprintf("%d,%d", sx, sy)] = true
		} else {
			switch s[i] {
			case '^':
				ry++
			case 'v':
				ry--
			case '>':
				rx++
			case '<':
				rx--
			}
			mappedH[fmt.Sprintf("%d,%d", rx, ry)] = true
		}
	}
	return len(mappedH)
}
