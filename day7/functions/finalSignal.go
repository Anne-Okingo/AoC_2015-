package functions

import (
	"strconv"
	"strings"
)

// Global map to store the signals for each wire
var signals map[string]uint16

// A map to store the parsed instructions
var instructions map[string]string

// Function to FinalSignal the signal for a wire or a number
func FinalSignal(wire string) uint16 {
	// If the wire is a digit, return it directly as a number
	if value, err := strconv.Atoi(wire); err == nil {
		return uint16(value)
	}

	// If the wire's value is already computed, return it
	if val, ok := signals[wire]; ok {
		return val
	}

	// Get the instruction for the wire
	instruction := instructions[wire]
	parts := strings.Fields(instruction)

	var signal uint16

	// Parse based on the type of instruction
	switch len(parts) {
	case 1: // e.g., "123 -> x" or "x -> y"
		signal = FinalSignal(parts[0])

	case 2: // e.g., "NOT x -> y"
		// NOT operation
		signal = ^FinalSignal(parts[1])

	case 3: // e.g., "x AND y -> z", "x LSHIFT 2 -> y"
		left := FinalSignal(parts[0])
		right := FinalSignal(parts[2])

		switch parts[1] {
		case "AND":
			signal = left & right
		case "OR":
			signal = left | right
		case "LSHIFT":
			signal = left << right
		case "RSHIFT":
			signal = left >> right
		}
	}

	// Memoize the result for future use
	signals[wire] = signal

	return signal
}


