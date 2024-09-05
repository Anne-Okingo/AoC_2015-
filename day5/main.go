package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	fs "day5/functions"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage should be : go run  . <name of file.txt> ")
		return
	}
	filepath := os.Args[1]
	suffix := ".txt"

	if !strings.HasSuffix(filepath, suffix) {
		fmt.Println("Only .txt file accepted")
		return
	}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error", err)
		os.Exit(1)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	inputSlice := []string{}
	for scanner.Scan() {
		line := scanner.Text()
		inputSlice = append(inputSlice, line)
	}
	if len (inputSlice)== 0{
		fmt.Println("Error : empty file")
		return
	}
	result := fs.CountNiceString(inputSlice)

	fmt.Println(result)
}
