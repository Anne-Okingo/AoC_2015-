package functions

func Ribbon(l, w, h int) int {
	volume := l * w * h

	dimentions := []int{l, w, h}
	for i := 0; i < len(dimentions); i++ {
		for j := 0; j < len(dimentions)-i-1; j++ {
			if dimentions[j] > dimentions[j+1] {
				dimentions[j], dimentions[j+1] = dimentions[j+1], dimentions[j]
			}
		}
	}
	perimeter := (2 * dimentions[0]) + (2 * dimentions[1])
	return volume + perimeter
}
