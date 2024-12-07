package day07

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Equation struct {
	target  int
	numbers []int
}

func recurse(target int, nums []int, concat bool) int {
	if len(nums) == 1 {
		if nums[0] == target {
			return target
		} else {
			return 0
		}
	}

	lhs := nums[0]
	rhs := nums[1]
	// fmt.Printf("  %v\n", nums)

	new_nums := make([]int, len(nums)-2)
	_ = copy(new_nums, nums[2:])
	new_nums = append([]int{lhs + rhs}, new_nums...)
	// fmt.Printf("    %v + %v = %v\n", lhs, rhs, lhs+rhs)
	if recurse(target, new_nums, concat) != 0 {
		return target
	}

	new_nums_2 := make([]int, len(nums)-2)
	_ = copy(new_nums_2, nums[2:])
	new_nums_2 = append([]int{lhs * rhs}, new_nums_2...)
	// fmt.Printf("    %v * %v = %v\n", lhs, rhs, lhs*rhs)
	if recurse(target, new_nums_2, concat) != 0 {
		return target
	}

	if concat {
		new_nums_3 := make([]int, len(nums)-2)
		_ = copy(new_nums_3, nums[2:])
		concat_num, _ := strconv.Atoi(fmt.Sprintf("%d%d", lhs, rhs))
		new_nums_3 = append([]int{concat_num}, new_nums_3...)
		// fmt.Printf("    %v || %v = %v\n", lhs, rhs, concat_num)
		if recurse(target, new_nums_3, concat) != 0 {
			return target
		}
	}

	return 0
}

func solvePart1(equations []Equation) int {
	ans := 0

	for _, eq := range equations {
		// fmt.Println(eq)
		ans += recurse(eq.target, eq.numbers, false)
		// fmt.Printf("ans is now %v\n", ans)
	}

	return ans
}

func solvePart2(equations []Equation) int {
	ans := 0

	for _, eq := range equations {
		// fmt.Println(eq)
		ans += recurse(eq.target, eq.numbers, true)
		// fmt.Printf("ans is now %v\n", ans)
	}

	return ans
}

var equationRegex = regexp.MustCompile(`([0-9]+): ([0-9 ]+)`)

func Run() (int, int) {
	// Parse input
	file, _ := os.Open("./days/day07/input.txt")
	defer file.Close()

	r := bufio.NewReader(file)

	equations := []Equation{}

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			line := string(line)
			match := equationRegex.FindStringSubmatch(line)

			if match != nil {
				target, _ := strconv.Atoi(match[1])
				numbers := []int{}
				for _, num := range strings.Split(match[2], " ") {
					num, _ := strconv.Atoi(num)
					numbers = append(numbers, num)
				}

				newEq := Equation{target, numbers}
				equations = append(equations, newEq)
			}

		}
		if err != nil {
			break
		}
	}

	// fmt.Println(equations)

	// Part 1
	part1 := solvePart1(equations)

	// Part 2
	part2 := solvePart2(equations)

	return part1, part2
}
