package functions

import (
	"fmt"
)

func Houses(s string) int {
	mappedH := make(map[string]bool)

	sx, sy := 0, 0

	mappedH[fmt.Sprintf("%d,%d", sx, sy)] = true
	for _, ch := range s {
		if ch == '^' {
			sy++
		} else if ch == 'v' {
			sy--
		} else if ch == '>' {
			sx++
		} else if ch == '<' {
			sx--
		}
		mappedH[fmt.Sprintf("%d,%d", sx, sy)] = true
	}
	return len(mappedH)
}
