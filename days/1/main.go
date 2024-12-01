package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
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

	var leftList []int
	var rightList []int

	lines := strings.Split(string(buffer), "\n")
	for _, line := range lines {
		if line == "" {
			continue
		}

		var leftValue, rightValue int
		fmt.Sscanf(line, "%d   %d", &leftValue, &rightValue)

		leftList = append(leftList, leftValue)
		rightList = append(rightList, rightValue)
	}

	partOneAnswer := partOne(leftList, rightList)
	partTwoAnswer := partTwo(leftList, rightList)

	fmt.Println(partOneAnswer)
	fmt.Println(partTwoAnswer)
}

func partOne(leftList, rightList []int) int {
	// Create duplicates of the lists to avoid modifying the originals
	sortedLeftList := make([]int, len(leftList))
	sortedRightList := make([]int, len(rightList))

	copy(sortedLeftList, leftList)
	copy(sortedRightList, rightList)

	// Sort the lists from smallest number to biggest
	sort.Slice(sortedLeftList, func(i, j int) bool {
		return sortedLeftList[i] < sortedLeftList[j]
	})

	sort.Slice(sortedRightList, func(i, j int) bool {
		return sortedRightList[i] < sortedRightList[j]
	})

	// Calculate how far apart the two numbers are
	totalDistance := 0
	for index := range sortedLeftList {
		totalDistance += int(math.Abs(float64(sortedLeftList[index]) - float64(sortedRightList[index])))
	}

	return totalDistance
}

func count(slice []int, numberToCheck int) int {
	count := 0
	for _, number := range slice {
		if number == numberToCheck {
			count++
		}
	}
	return count
}

func partTwo(leftList, rightList []int) int {
	totalSimilarityScore := 0

	for _, number := range leftList {
		similarityScore := number * count(rightList, number)
		totalSimilarityScore += similarityScore
	}

	return totalSimilarityScore
}
