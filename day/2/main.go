package main

import (
	"fmt"
	"slices"

	utl "github.com/kjabin/aoc2024/utils"
)

func main() {
	lines := utl.ReadFile("input")
	fmt.Println(calcSafeReports(lines))
}

func calcSafeReports(lines []string) int {
	totalSafeReports := 0
	for _, line := range lines {
		report := utl.ToIntArr(line, " ")
		if isReportSafe(report) || isReportDampened(report) {
			totalSafeReports += 1
		}
	}
	return totalSafeReports
}

func isReportDampened(report []int) bool {
	for i := 0; i < len(report); i++ {
		arr := slices.Concat(report[:i], report[i+1:])
		if isReportSafe(arr) {
			return true
		}
	}
	return false
}

func isReportSafe(report []int) bool {
	sign := report[1] - report[0]
	for i := 1; i < len(report); i++ {
		// diff limit exceeded
		if utl.Abs(report[i]-report[i-1]) > 3 {
			return false
		}
		// report is not monotonically increasing/decreasing
		if sign*(report[i]-report[i-1]) <= 0 {
			return false
		}
	}
	return true
}
