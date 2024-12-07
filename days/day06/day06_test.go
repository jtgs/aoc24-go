package day06

import (
	"testing"
)

var exampleGrid = []string{
	"....#.....",
	".........#",
	"..........",
	"..#.......",
	".......#..",
	"..........",
	".#..^.....",
	"........#.",
	"#.........",
	"......#...",
}

func TestDay06InputParsing(t *testing.T) {
	_, start := parseGrid(exampleGrid)
	expected := position{row: 6, col: 4}
	if start != expected {
		t.Errorf("guard in wrong position; expected (6, 4), got %v", start)
	}
}

func TestDay06ExamplePart1(t *testing.T) {
	grid, start := parseGrid(exampleGrid)
	if ans := solvePart1(grid, start); ans != 41 {
		t.Errorf("expected 41, got %d", ans)
	}
}

func TestDay06ExamplePart2(t *testing.T) {
	grid, start := parseGrid(exampleGrid)
	if ans := solvePart2(grid, start); ans != 6 {
		t.Errorf("expected 6, got %d", ans)
	}
}
