package day08

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jtgs/aoc24-go/util"
	combinations "github.com/mxschmitt/golang-combinations"
)

type position struct {
	row int
	col int
}

func inBounds(pos position, height int, width int) bool {
	if pos.row < 0 || pos.col < 0 || pos.row >= height || pos.col >= width {
		return false
	}
	return true
}

func solvePart1(antennas map[rune][]position, height int, width int) int {
	antinodes := make(map[position]int)

	for lt, ant := range antennas {
		fmt.Printf("Antenna %v: %v\n", lt, ant)
		combis := combinations.Combinations(ant, 2)
		for _, combi := range combis {
			fmt.Printf("  A: %v B: %v\n", combi[0], combi[1])

			var row_1, col_1 int
			// B.row will always be greater than A.row due to sorting
			row_1 = combi[0].row - util.AbsInt(combi[0].row-combi[1].row)

			// Col is trickier
			if combi[0].col > combi[1].col {
				col_1 = combi[0].col + util.AbsInt(combi[0].col-combi[1].col)
			} else {
				col_1 = combi[0].col - util.AbsInt(combi[0].col-combi[1].col)
			}

			pos_1 := position{row_1, col_1}
			fmt.Printf("    pos_1: %v\n", pos_1)

			if inBounds(pos_1, height, width) {
				antinodes[pos_1]++
			}

			var row_2, col_2 int
			row_2 = combi[1].row + util.AbsInt(combi[0].row-combi[1].row)
			if combi[0].col > combi[1].col {
				col_2 = combi[1].col - util.AbsInt(combi[0].col-combi[1].col)
			} else {
				col_2 = combi[1].col + util.AbsInt(combi[0].col-combi[1].col)
			}

			pos_2 := position{row_2, col_2}
			fmt.Printf("    pos_2: %v\n", pos_2)

			if inBounds(pos_2, height, width) {
				antinodes[pos_2]++
			}
		}
	}

	fmt.Println(antinodes)
	return len(antinodes)
}

func solvePart2(antennas map[rune][]position, height int, width int) int {
	antinodes := make(map[position]int)

	for lt, ant := range antennas {
		fmt.Printf("Antenna %c: %v\n", lt, ant)
		combis := combinations.Combinations(ant, 2)
		for _, combi := range combis {
			fmt.Printf("  A: %v B: %v\n", combi[0], combi[1])

			row_diff := util.AbsInt(combi[0].row - combi[1].row)
			col_diff := util.AbsInt(combi[0].col - combi[1].col)

			pos_1 := position{combi[0].row, combi[0].col}
			for {
				var row_1, col_1 int
				// B.row will always be greater than A.row due to sorting
				row_1 = pos_1.row - util.AbsInt(row_diff)

				// Col is trickier
				if pos_1.col > combi[1].col {
					col_1 = pos_1.col + util.AbsInt(col_diff)
				} else {
					col_1 = pos_1.col - util.AbsInt(col_diff)
				}

				pos_1 = position{row_1, col_1}
				fmt.Printf("    pos_1: %v\n", pos_1)

				if !inBounds(pos_1, height, width) {
					break
				}

				antinodes[pos_1]++
			}

			pos_2 := position{combi[1].row, combi[1].col}
			for {
				var row_2, col_2 int
				// B.row will always be greater than A.row due to sorting
				row_2 = pos_2.row + util.AbsInt(row_diff)

				// Col is trickier
				if pos_2.col < combi[0].col {
					col_2 = pos_2.col - util.AbsInt(col_diff)
				} else {
					col_2 = pos_2.col + util.AbsInt(col_diff)
				}

				pos_2 = position{row_2, col_2}
				fmt.Printf("    pos_2: %v\n", pos_2)

				if !inBounds(pos_2, height, width) {
					break
				}

				antinodes[pos_2]++
			}
		}

		// There's probably a clever way to add the antenna positions but I can't be bothered.
		if len(ant) > 1 {
			for _, pos := range ant {
				antinodes[pos]++
			}
		}
	}

	fmt.Println(antinodes)
	return len(antinodes)
}

func parseGrid(lines []string) (map[rune][]position, int, int) {
	antennas := make(map[rune][]position)

	for row, line := range lines {
		for col, ch := range line {
			if ch != '.' {
				pos := position{row, col}
				antennas[ch] = append(antennas[ch], pos)
			}
		}
	}

	return antennas, len(lines), len(lines[0])
}

func Run() (int, int) {
	// Parse input
	file, _ := os.Open("./days/day08/input.txt")
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

	antennas, height, width := parseGrid(lines)

	// Part 1
	part1 := solvePart1(antennas, height, width)

	// Part 2
	part2 := solvePart2(antennas, height, width)

	return part1, part2
}
