package day03

import "testing"

func TestDay01ExamplePart1(t *testing.T) {
	exampleString := "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

	if ans := solvePart1([]string{exampleString}); ans != 161 {
		t.Errorf("expected 161, got %d", ans)
	}
}

func TestDay01ExamplePart2(t *testing.T) {
	exampleString := "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

	if ans := solvePart2([]string{exampleString}); ans != 48 {
		t.Errorf("expected 48, got %d", ans)
	}
}
