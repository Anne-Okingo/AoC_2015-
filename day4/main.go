package main

import (
	f "AoC_2015-/function"
	"fmt"
)

func main() {
	input := "ckczppom"

	result1 := f.MD5(input)
	result2 := f.MD5b(input)
	fmt.Println(result1)
	fmt.Println(result2)
}
