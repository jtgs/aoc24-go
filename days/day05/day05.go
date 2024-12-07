package day05

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

type rule struct {
	before int
	after  int
}

type update []int

func solvePart1(ruleset []rule, updates []update) int {
	ans := 0

	for _, update := range updates {
		fmt.Println(update)
		good := true
		for _, rule := range ruleset {
			fmt.Printf("  %v\n", rule)
			if beforeIx := slices.Index(update, rule.before); beforeIx != -1 {
				afterIx := slices.Index(update, rule.after)
				fmt.Printf("    %v : %v\n", beforeIx, afterIx)
				if afterIx != -1 && afterIx <= beforeIx {
					good = false
					break
				}
			}
		}

		if good {
			ans += update[len(update)/2]
		}
	}

	return ans
}

func solvePart2(ruleset []rule, updates []update) int {
	ans := 0

	return ans
}

func Run() (int, int) {
	// Parse input
	file, _ := os.Open("./days/day05/input.txt")
	defer file.Close()

	r := bufio.NewReader(file)

	ruleset := []rule{}
	updates := []update{}

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			line := string(line)
			if strings.Contains(line, "|") {
				// rule
				parts := strings.Split(line, "|")
				before, _ := strconv.Atoi(parts[0])
				after, _ := strconv.Atoi(parts[1])
				rule := rule{before, after}
				ruleset = append(ruleset, rule)
			} else {
				// update
				parts := strings.Split(line, ",")
				var newUpdate update
				for _, part := range parts {
					num, _ := strconv.Atoi(part)
					newUpdate = append(newUpdate, num)
				}
				updates = append(updates, newUpdate)
			}
		}
		if err != nil {
			break
		}
	}

	fmt.Println(ruleset)
	fmt.Println(updates)

	// Part 1
	part1 := solvePart1(ruleset, updates)

	// Part 2
	part2 := solvePart2(ruleset, updates)

	return part1, part2
}
