package main

import (
	"fmt"
	"log"
	"os"
	"strings"
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

	var wordSearchGrid [][]rune
	for _, row := range strings.Split(string(buffer), "\n") {
		wordSearchRow := []rune(row)
		wordSearchGrid = append(wordSearchGrid, wordSearchRow)
	}

	partOneAnswer := partOne("XMAS", wordSearchGrid)
	partTwoAnswer := partTwo(wordSearchGrid)

	fmt.Printf("Part One: %d\nPart Two: %d\n", partOneAnswer, partTwoAnswer)
}

// Directions for searching: right, down, diagonal down-right, and their reverses
var directionsMatrix = [][2]int{
	{0, 1},   // right
	{1, 0},   // down
	{1, 1},   // diagonal down-right
	{0, -1},  // left
	{-1, 0},  // up
	{-1, -1}, // diagonal up-left
	{-1, 1},  // diagonal up-right
	{1, -1},  // diagonal down-left
}

func partOne(targetWord string, wordSearchGrid [][]rune) int {
	amountOfOccurances := 0
	rows := len(wordSearchGrid)
	columns := len(wordSearchGrid[0])

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			for _, direction := range directionsMatrix {
				if searchWord(targetWord, wordSearchGrid, row, column, direction[0], direction[1]) {
					amountOfOccurances++
				}
			}
		}
	}
	return amountOfOccurances
}

func searchWord(targetWord string, grid [][]rune, startX, startY, dirX, dirY int) bool {
	for i := 0; i < len(targetWord); i++ {
		x := startX + (i * dirX)
		y := startY + (i * dirY)

		// Check if the position is out of bounds
		if x < 0 || x >= len(grid) || y < 0 || y >= len(grid[0]) {
			return false
		}

		// Check if the character matches
		if grid[x][y] != rune(targetWord[i]) {
			return false
		}
	}
	return true
}

// Hardcoded the X-MAS values, because there is no need to overcomplicate.
func partTwo(wordSearchGrid [][]rune) int {
	amountOfOccurances := 0
	rows := len(wordSearchGrid)
	columns := len(wordSearchGrid[0])

	for row := 0; row < rows; row++ {
		for column := 0; column < columns; column++ {
			if searchMASinShapeOfX(wordSearchGrid, row, column) {
				amountOfOccurances++
			}
		}
	}
	return amountOfOccurances
}

func searchMASinShapeOfX(grid [][]rune, x, y int) bool {
	// Avoid the corners of the grid as checking would make us go out of bounds.
	if x < 1 || x >= len(grid)-1 || y < 1 || y >= len(grid[0])-1 {
		return false
	}

	// If the middle element isn't an A, then it can't possibly be an X-MAS
	if grid[x][y] != 'A' {
		return false
	}

	// Calculate each possible valid position
	validPositions := []bool{
		grid[x-1][y-1] == 'M' && grid[x+1][y-1] == 'S' && grid[x-1][y+1] == 'M' && grid[x+1][y+1] == 'S',
		grid[x-1][y-1] == 'S' && grid[x+1][y-1] == 'S' && grid[x-1][y+1] == 'M' && grid[x+1][y+1] == 'M',
		grid[x-1][y-1] == 'M' && grid[x+1][y-1] == 'M' && grid[x-1][y+1] == 'S' && grid[x+1][y+1] == 'S',
		grid[x-1][y-1] == 'S' && grid[x+1][y-1] == 'M' && grid[x-1][y+1] == 'S' && grid[x+1][y+1] == 'M',
	}

	// Check if any one of these positions match
	for _, position := range validPositions {
		if position {
			return true // If we get a match, we can return true
		}
	}
	return false // If none match, then we return false.
}
