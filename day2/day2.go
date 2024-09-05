package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	f "AoC_2015-/function"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Error Usage should be : go run . <name of file>.txt")
		os.Exit(1)
	}
	filepath := os.Args[1]
	suffix := ".txt"
	if !strings.HasSuffix(filepath, suffix) {
		fmt.Println("Error : Filepath should be <name of file>.txt")
		return
	}

	file, err := os.Open(filepath)
	if err != nil {
		fmt.Println("Error", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	TotalWrapPaper := 0
	TotalRibon := 0

	for scanner.Scan() {
		str := scanner.Text()
		stri := strings.TrimSpace(str)
		if stri == "" {
			continue
		}
		slice := []string{}
		s := ""
		for _, ch := range stri {
			if ch != 'x' {
				s += string(ch)
			} else if s != "" {
				slice = append(slice, s)
				s = ""
			}
		}
		if s != "" {
			slice = append(slice, s)
		}
		if len(slice) != 3 {
			fmt.Println("Error: Invalid dimensions format. Expected format is LxWxH.")
			return
		}

		length := f.Atoi(slice[0])
		width := f.Atoi(slice[1])
		height := f.Atoi(slice[2])

		WrapPaper := (f.WrapPaper(length, width, height))
		Ribbon := f.Ribbon(length, width, height)
		TotalWrapPaper += WrapPaper
		TotalRibon += Ribbon

	}
	fmt.Printf("TotalWrapPaper : %v", TotalWrapPaper)
	fmt.Println()
	fmt.Printf("TotalRibbon: %v", TotalRibon)
	fmt.Println()
}
