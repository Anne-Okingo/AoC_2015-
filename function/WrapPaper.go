package functions

func WrapPaper(l, w, h int) int {
	A1 := l * w
	A2 := l * h
	A3 := w * h

	Small := A1
	if A2 < Small {
		Small = A2
	}
	if A3 < Small {
		Small = A3
	}

	SurfaceArea := ((2 * (l * w)) + (2 * (w * h)) + (2 * (h * l)))
	return SurfaceArea + Small
}
