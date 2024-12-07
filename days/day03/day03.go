package day03

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var mul = regexp.MustCompile(`mul\(([0-9]+),([0-9]+)\)`)

func findAndMultiply(input string) int {
	ans := 0

	matches := mul.FindAllStringSubmatch(input, -1)

	for _, match := range matches {
		// fmt.Printf("match: (%s, %s)\n", match[1], match[2])
		numA, _ := strconv.Atoi(string(match[1]))
		numB, _ := strconv.Atoi(string(match[2]))
		ans += (numA * numB)
	}

	return ans
}

func solvePart1(lines []string) int {
	return findAndMultiply(strings.Join(lines, ""))
}

func solvePart2(lines []string) int {
	ans := 0
	do_to_dont := regexp.MustCompile(`do\(\)(.*?)(don't\(\)|\z)`)

	full_str := strings.Join(lines, "")
	full_str = "do()" + full_str

	matches := do_to_dont.FindAllStringSubmatch(full_str, -1)

	for _, match := range matches {
		ans += findAndMultiply(match[1])
	}

	return ans
}

func Run() (int, int) {
	// Parse input
	file, _ := os.Open("./days/day03/input.txt")
	defer file.Close()

	r := bufio.NewReader(file)

	lines := []string{}

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			lines = append(lines, string(line))
		}
		if err != nil {
			break
		}
	}

	// Part 1
	part1 := solvePart1(lines)

	// Part 2
	part2 := solvePart2(lines)

	return part1, part2
}
