// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// 	"strings"

// 	"day7/functions"
// )

// func main() {
// 	if len(os.Args) != 2 {
// 		fmt.Println("Error Usage should be : go run . lift.txt")
// 		os.Exit(1)
// 	}

// 	fileName := os.Args[1]
// 	suffix := ".txt"

// 	if !strings.HasSuffix(fileName, suffix) {
// 		fmt.Println("Only .txt files accepted")
// 		return
// 	}

// 	fileContent, err := os.Open(fileName)
// 	if err != nil {
// 		fmt.Println("errr", err)
// 		return
// 	}
// 	defer fileContent.Close()

// 	scanner := bufio.NewScanner(fileContent)

// 	for scanner.Scan() {
// 		line := scanner.Text()

// 		signal := functions.Split(line)
// 		functions.FinalSignal(signal)

// 	}
// }

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"day7/functions"
)

// Global map to store the signals for each wire
var signals map[string]uint16

// A map to store the parsed instructions
var instructions map[string]string

func main() {
	// Initialize maps
	signals = make(map[string]uint16)
	instructions = make(map[string]string)

	if len(os.Args) != 2 {
		fmt.Println("Error Usage should be : go run . lift.txt")
		os.Exit(1)
	}

	fileName := os.Args[1]
	suffix := ".txt"
	if !strings.HasSuffix(fileName, suffix) {
		fmt.Println("Error only .txt file accepted")
	}

	// Open the input file
	fileContent, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file", err)
		return
	}
	defer fileContent.Close()

	// Read the file line by line
	scanner := bufio.NewScanner(fileContent)

	for scanner.Scan(){
		line := scanner.Text()
		parts := functions.Separator(line,"->")
		instructions[parts[1]] =parts[0]

	}

	// Check for scanning errors
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Evaluate the signal for wire 'a' (change this to any wire you want to FinalSignal)
	fmt.Println("Signal at wire a:", functions.FinalSignal("a"))
}
