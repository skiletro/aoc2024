package main

import (
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) <= 1 {
		fmt.Println("Usage: " + os.Args[0] + " <fileName>")
		os.Exit(1)
	}

	buffer, error := os.ReadFile(os.Args[1])
	if error != nil {
		log.Fatal(error)
	}

	corruptedMemory := string(buffer)

	partOneAnswer := partOne(corruptedMemory)
	partTwoAnswer := partTwo(corruptedMemory)

	fmt.Printf("Part One: %d\nPart Two: %d\n", partOneAnswer, partTwoAnswer)
}

func partOne(corruptedMemory string) int {
	// Regex: mul(x, y) where x and y are any number between 0 and inf.
	regex, _ := regexp.Compile(`mul\([0-9]+\,[0-9]+\)`)

	recoveredInstructions := regex.FindAllString(corruptedMemory, -1)

	multiplicationResults := 0
	for _, instruction := range recoveredInstructions {
		var number1, number2 int
		fmt.Sscanf(instruction, "mul(%d,%d)", &number1, &number2)
		multiplicationResults += (number1 * number2)
	}

	return multiplicationResults
}

func partTwo(corruptedMemory string) int {
	// Regex: mul(x, y) where x and y are any number between 0 and inf,
	//        do() and don't() instructions.
	regex, _ := regexp.Compile(`(mul\([0-9]+\,[0-9]+\))|(do\(\))|(don't\(\))`)

	recoveredInstructions := regex.FindAllString(corruptedMemory, -1)

	multiplicationResults := 0
	instructionsAreEnabled := true
	for _, instruction := range recoveredInstructions {
		if instruction == "do()" {
			instructionsAreEnabled = true
			continue
		}

		if instruction == "don't()" {
			instructionsAreEnabled = false
			continue
		}

		// The only instruction that gets to this point should be mul()
		if instructionsAreEnabled {
			var number1, number2 int
			fmt.Sscanf(instruction, "mul(%d,%d)", &number1, &number2)
			multiplicationResults += (number1 * number2)
		}
	}

	return multiplicationResults
}
