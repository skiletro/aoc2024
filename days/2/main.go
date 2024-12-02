package main

import (
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strconv"
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

	reports := strings.Split(string(buffer), "\n")

	partOneAnswer := partOne(reports)

	fmt.Println(partOneAnswer)
}

func partOne(reports []string) int {
	totalSafeReports := 0

	for _, report := range reports {
		// Convert the strings into numbers that can be worked with
		reportNumbers := strings.Split(report, " ")
		var convertedReportNumbers []int
		for _, num := range reportNumbers {
			convertedNumber, _ := strconv.Atoi(num)
			convertedReportNumbers = append(convertedReportNumbers, convertedNumber)
		}

		// Check for each requirement
		isAscending := sort.SliceIsSorted(convertedReportNumbers, func(i, j int) bool {
			return convertedReportNumbers[i] < convertedReportNumbers[j]
		})

		isDescending := sort.SliceIsSorted(convertedReportNumbers, func(i, j int) bool {
			return convertedReportNumbers[i] > convertedReportNumbers[j]
		})

		eachAscNumberInRange := sort.SliceIsSorted(convertedReportNumbers, func(i, j int) bool {
			difference := math.Abs(float64(convertedReportNumbers[i] - convertedReportNumbers[j]))
			return difference > 3 || difference <= 0
		})

		// It can't be both ascending and descending, so this should work.
		if (isAscending || isDescending) && eachAscNumberInRange {
			totalSafeReports++
		}
	}

	return totalSafeReports
}
