package main

import (
	f "AoC_2015-/function"
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error Usage should be : go run . lift.txt")
		os.Exit(1)
	}

	filepath := os.Args[1]

	FilePath, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error Opening File", err)
		return
	}
	defer FilePath.Close()

	scanner := bufio.NewScanner(FilePath)

	for scanner.Scan() {
		santasteps := scanner.Text()
		fmt.Println(f.Floor(santasteps))
	}
}
