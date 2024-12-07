package day01

import (
	"bufio"
	"bytes"
	"os"
	"sort"
	"strconv"

	"github.com/jtgs/aoc24-go/util"
)

func solvePart1(listA []int, listB []int) int {
	ans := 0

	sort.Ints(listA)
	sort.Ints(listB)

	for ix := range listA {
		ans += util.AbsInt(listA[ix] - listB[ix])
	}

	return ans
}

func solvePart2(listA []int, listB []int) int {
	ans := 0
	counts := make(map[int]int)

	// First make a map of the counts of each number in listB
	for _, val := range listB {
		counts[val] += 1
	}

	// Then run through each number in listA, and calculate the
	// similarity score by looking up in the map
	for _, val := range listA {
		ans += counts[val] * val
	}

	return ans
}

func Run() (int, int) {
	// Parse input
	file, _ := os.Open("./days/day01/input.txt")
	defer file.Close()

	r := bufio.NewReader(file)

	listA := []int{}
	listB := []int{}

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			parts := bytes.Split(line, []byte("   "))
			numA, _ := strconv.Atoi(string(parts[0]))
			numB, _ := strconv.Atoi(string(parts[1]))
			listA = append(listA, numA)
			listB = append(listB, numB)
		}
		if err != nil {
			break
		}
	}

	// Part 1
	part1 := solvePart1(listA, listB)

	// Part 2
	part2 := solvePart2(listA, listB)

	return part1, part2
}
