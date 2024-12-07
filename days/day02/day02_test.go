package day02

import "testing"

func TestDay02ExamplePart1(t *testing.T) {
	exampleInput := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	if ans := solvePart1(exampleInput[:]); ans != 2 {
		t.Errorf("expected 2, got %d", ans)
	}
}

func TestDay02ExamplePart2(t *testing.T) {
	exampleInput := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	if ans := solvePart2(exampleInput[:]); ans != 4 {
		t.Errorf("expected 4, got %d", ans)
	}
}
