package day04

import "testing"

var exampleGrid [][]rune = [][]rune{
	{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
	{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
	{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
	{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
	{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
	{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
	{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
	{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
	{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
	{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
}

func TestDay04ExamplePart1(t *testing.T) {
	if ans := solvePart1(exampleGrid); ans != 18 {
		t.Errorf("expected 18, got %d", ans)
	}
}

func TestDay04ExamplePart2(t *testing.T) {
	if ans := solvePart2(exampleGrid); ans != 9 {
		t.Errorf("expected 9, got %d", ans)
	}
}
