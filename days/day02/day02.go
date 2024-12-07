package day02

import (
	"bufio"
	"bytes"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	"github.com/jtgs/aoc24-go/util"
)

type direction int

const (
	ascending  direction = 0
	descending direction = 1
	invalid    direction = 3
)

func tryPair(a int, b int, dir direction) bool {
	switch dir {
	case ascending:
		if a < b {
			slog.Debug(fmt.Sprintf("  %v less than %v - invalid\n", a, b))
			return false
		}
	case descending:
		if a > b {
			slog.Debug(fmt.Sprintf("  %v greater than %v - invalid\n", a, b))
			return false
		}
	}

	diff := util.AbsInt(a - b)
	if diff < 1 || diff > 3 {
		slog.Debug(fmt.Sprintf("  diff %v not in range - invalid\n", diff))
		return false
	}

	return true
}

func getDir(a int, b int) direction {
	// Work out direction by comparing first two numbers
	x := a - b
	switch {
	case x < 0:
		slog.Debug("  ascending")
		return ascending
	case x > 0:
		slog.Debug("  descending")
		return descending
	case x == 0:
		// invalid as not different by at least 1
		slog.Debug("  no change - invalid")
	}

	return invalid
}

func checkReport(report []int) bool {
	slog.Debug(fmt.Sprintf("%v", report))

	dir := getDir(report[0], report[1])
	if dir == invalid {
		return false
	}

	good := true
	for x := 1; x < len(report); x++ {
		good = tryPair(report[x], report[x-1], dir)
		if !good {
			break
		}
	}

	return good
}

func solvePart1(reports [][]int) int {
	ans := 0

	for _, report := range reports {
		if checkReport(report) {
			// This report is good
			slog.Debug("  good report")
			ans++
		}
	}

	return ans
}

func solvePart2(reports [][]int) int {
	ans := 0

	for _, report := range reports {
		if checkReport(report) {
			// This report is good
			slog.Debug("  good report")
			ans++
		} else {
			for ix := range report {
				report_copy := make([]int, len(report))
				_ = copy(report_copy, report)
				slog.Debug(fmt.Sprintf("%v", report_copy))
				test := append(report_copy[:ix], report_copy[ix+1:]...)
				if checkReport(test) {
					slog.Debug("  good report")
					ans++
					break
				}
			}
		}
	}

	return ans
}

func Run() (int, int) {
	// Parse input
	file, _ := os.Open("./days/day02/input.txt")
	defer file.Close()

	r := bufio.NewReader(file)

	reports := [][]int{}

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			report := []int{}
			parts := bytes.Split(line, []byte(" "))

			for _, part := range parts {
				num, _ := strconv.Atoi(string(part))
				report = append(report, num)
			}
			// slog.Debug(report)
			reports = append(reports, report)
		}
		if err != nil {
			break
		}
	}

	// Part 1
	part1 := solvePart1(reports)

	// Part 2
	part2 := solvePart2(reports)

	return part1, part2
}
