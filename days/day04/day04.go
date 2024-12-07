package day04

import (
	"bufio"
	"fmt"
	"log/slog"
	"os"
)

func solvePart1(grid [][]rune) int {
	ans := 0

	maxWidth := len(grid[0])
	maxHeight := len(grid)

	for row := range grid {
		for col := range grid {
			// L -> R
			if col+3 < maxWidth {
				if grid[row][col] == 'X' && grid[row][col+1] == 'M' && grid[row][col+2] == 'A' && grid[row][col+3] == 'S' {
					slog.Debug(fmt.Sprintf("[%v][%v]: '%v'\n", row, col, string(grid[row][col])))
					slog.Debug("  L -> R")
					ans++
				}
			}
			// R -> L
			if col-3 >= 0 {
				if grid[row][col-3] == 'S' && grid[row][col-2] == 'A' && grid[row][col-1] == 'M' && grid[row][col] == 'X' {
					slog.Debug(fmt.Sprintf("[%v][%v]: '%v'\n", row, col, string(grid[row][col])))
					slog.Debug("  R -> L")
					ans++
				}
			}

			// T -> B
			if row+3 < maxHeight {
				if grid[row][col] == 'X' && grid[row+1][col] == 'M' && grid[row+2][col] == 'A' && grid[row+3][col] == 'S' {
					slog.Debug(fmt.Sprintf("[%v][%v]: '%v'\n", row, col, string(grid[row][col])))
					slog.Debug("  T -> B")
					ans++
				}
			}
			// B -> T
			if row-3 >= 0 {
				if grid[row-3][col] == 'S' && grid[row-2][col] == 'A' && grid[row-1][col] == 'M' && grid[row][col] == 'X' {
					slog.Debug(fmt.Sprintf("[%v][%v]: '%v'\n", row, col, string(grid[row][col])))
					slog.Debug("  B -> T")
					ans++
				}
			}

			// TL -> BR
			if col+3 < maxWidth && row+3 < maxHeight {
				if grid[row][col] == 'X' && grid[row+1][col+1] == 'M' && grid[row+2][col+2] == 'A' && grid[row+3][col+3] == 'S' {
					slog.Debug(fmt.Sprintf("[%v][%v]: '%v'\n", row, col, string(grid[row][col])))
					slog.Debug("  TL -> BR")
					ans++
				}
			}
			// BR -> TL
			if col+3 < maxWidth && row+3 < maxHeight {
				if grid[row+3][col+3] == 'X' && grid[row+2][col+2] == 'M' && grid[row+1][col+1] == 'A' && grid[row][col] == 'S' {
					slog.Debug(fmt.Sprintf("[%v][%v]: '%v'\n", row, col, string(grid[row][col])))
					slog.Debug("  BR -> TL")
					ans++
				}
			}

			// TR -> BL
			if col-3 >= 0 && row+3 < maxHeight {
				if grid[row][col] == 'X' && grid[row+1][col-1] == 'M' && grid[row+2][col-2] == 'A' && grid[row+3][col-3] == 'S' {
					slog.Debug(fmt.Sprintf("[%v][%v]: '%v'\n", row, col, string(grid[row][col])))
					slog.Debug("  TR -> BL")
					ans++
				}
			}
			// BL -> TR
			if col+3 < maxWidth && row-3 >= 0 {
				if grid[row][col] == 'X' && grid[row-1][col+1] == 'M' && grid[row-2][col+2] == 'A' && grid[row-3][col+3] == 'S' {
					slog.Debug(fmt.Sprintf("[%v][%v]: '%v'\n", row, col, string(grid[row][col])))
					slog.Debug("  BL -> TR")
					ans++
				}
			}
		}
	}

	return ans
}

func solvePart2(grid [][]rune) int {
	ans := 0
	maxWidth := len(grid[0])
	maxHeight := len(grid)

	for row := range grid {
		if row-1 < 0 || row+1 >= maxHeight {
			continue
		}

		for col := range grid {
			if col-1 < 0 || col+1 >= maxWidth {
				continue
			}

			if grid[row][col] == 'A' {
				if (grid[row+1][col+1] == 'S' && grid[row-1][col-1] == 'M' && grid[row-1][col+1] == 'M' && grid[row+1][col-1] == 'S') ||
					(grid[row+1][col+1] == 'S' && grid[row-1][col-1] == 'M' && grid[row-1][col+1] == 'S' && grid[row+1][col-1] == 'M') ||
					(grid[row+1][col+1] == 'M' && grid[row-1][col-1] == 'S' && grid[row-1][col+1] == 'S' && grid[row+1][col-1] == 'M') ||
					(grid[row+1][col+1] == 'M' && grid[row-1][col-1] == 'S' && grid[row-1][col+1] == 'M' && grid[row+1][col-1] == 'S') {
					ans++
				}
			}
		}
	}

	return ans
}

func Run() (int, int) {
	// Parse input
	file, _ := os.Open("./days/day04/input.txt")
	defer file.Close()

	r := bufio.NewReader(file)

	grid := [][]rune{}

	for {
		line, _, err := r.ReadLine()
		if len(line) > 0 {
			row := []rune{}
			for _, ch := range line {
				row = append(row, rune(ch))
			}

			grid = append(grid, row)
		}
		if err != nil {
			break
		}
	}

	slog.Debug(fmt.Sprintf("%v x %v\n", len(grid), len(grid[0])))

	// Part 1
	part1 := solvePart1(grid)

	// Part 2
	part2 := solvePart2(grid)

	return part1, part2
}
