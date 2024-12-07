package day06

import (
	"bufio"
	"fmt"
	"os"
)

type position struct {
	row int
	col int
}

type grid [][]int

type direction int

const (
	dir_up    = 0
	dir_right = 1
	dir_down  = 2
	dir_left  = 3
)

type result int

const (
	res_ok            = 0
	res_blocked       = 1
	res_out_of_bounds = 2
)

func step(grid grid, pos position, dir direction) (result, position) {
	new_pos := position{pos.row, pos.col}

	switch dir {
	case dir_up:
		new_pos.row--
	case dir_down:
		new_pos.row++
	case dir_left:
		new_pos.col--
	case dir_right:
		new_pos.col++
	}

	if new_pos.row < 0 || new_pos.row >= len(grid) || new_pos.col < 0 || new_pos.col >= len(grid[0]) {
		return res_out_of_bounds, new_pos
	}

	if grid[new_pos.row][new_pos.col] == 1 {
		return res_blocked, pos
	}

	return res_ok, new_pos
}

func drawGrid(grid grid, visited map[position]int) {
	for row := range grid {
		var s string
		for col := range grid[row] {
			pos := position{row, col}
			if visited[pos] > 0 {
				s += "X"
			} else if grid[row][col] == 1 {
				s += "#"
			} else {
				s += "."
			}
		}
		fmt.Println(s)
	}
	fmt.Println()
}

func solvePart1(grid grid, start position) int {
	pos := start
	var res result
	var dir direction = dir_up
	visited := make(map[position]int)
	visited[start]++
	// drawGrid(grid, visited)

	run := true

	for run {
		fmt.Printf("pos: %v, dir: %v\n", pos, dir)
		res, pos = step(grid, pos, dir)
		fmt.Printf("res: %v, pos: %v\n", res, pos)

		switch res {
		case res_ok:
			visited[pos]++
		case res_blocked:
			dir = (dir + 1) % 4
		case res_out_of_bounds:
			run = false
		}
		// drawGrid(grid, visited)
	}

	drawGrid(grid, visited)
	return len(visited)
}

func solvePart2(grid grid, start position) int {
	ans := 0

	return ans
}

func parseGrid(lines []string) (grid, position) {
	ans := grid{}
	pos := position{-1, -1}

	for row, line := range lines {
		newRow := []int{}
		for col, ch := range line {
			if ch == '.' {
				newRow = append(newRow, 0)
			} else if ch == '#' {
				newRow = append(newRow, 1)
			} else if ch == '^' {
				newRow = append(newRow, 0)
				pos = position{row, col}
			}
		}
		ans = append(ans, newRow)
	}

	return ans, pos
}

func Run() (int, int) {
	// Parse input
	file, _ := os.Open("./days/day06/input.txt")
	defer file.Close()

	r := bufio.NewReader(file)

	lines := []string{}

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			line := string(line)
			lines = append(lines, line)
		}
		if err != nil {
			break
		}
	}

	grid, start := parseGrid(lines)

	// Part 1
	part1 := solvePart1(grid, start)

	// Part 2
	part2 := solvePart2(grid, start)

	return part1, part2
}
