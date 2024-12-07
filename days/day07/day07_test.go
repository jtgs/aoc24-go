package day07

import "testing"

var exampleEquations []Equation = []Equation{
	{190, []int{10, 19}},
	{3267, []int{81, 40, 27}},
	{83, []int{17, 5}},
	{156, []int{15, 6}},
	{7290, []int{6, 8, 6, 15}},
	{161011, []int{16, 10, 13}},
	{192, []int{17, 8, 14}},
	{21037, []int{9, 7, 18, 13}},
	{292, []int{11, 6, 16, 20}},
}

func TestDay05ExamplePart1(t *testing.T) {
	if ans := solvePart1(exampleEquations); ans != 3749 {
		t.Errorf("expected 3749, got %d", ans)
	}
}

func TestDay05ExamplePart2(t *testing.T) {
	if ans := solvePart2(exampleEquations); ans != 11387 {
		t.Errorf("expected 11387, got %d", ans)
	}
}
