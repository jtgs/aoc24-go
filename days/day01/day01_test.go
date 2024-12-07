package day01

import "testing"

func TestDay01ExamplePart1(t *testing.T) {
	exampleListA := []int{3, 4, 2, 1, 3, 3}
	exampleListB := []int{4, 3, 5, 3, 9, 3}

	if ans := solvePart1(exampleListA[:], exampleListB[:]); ans != 11 {
		t.Errorf("expected 11, got %d", ans)
	}
}

func TestDay01ExamplePart2(t *testing.T) {
	exampleListA := []int{3, 4, 2, 1, 3, 3}
	exampleListB := []int{4, 3, 5, 3, 9, 3}

	if ans := solvePart2(exampleListA[:], exampleListB[:]); ans != 31 {
		t.Errorf("expected 31, got %d", ans)
	}
}
