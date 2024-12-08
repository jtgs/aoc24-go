package day08

import (
	"testing"
)

var exampleGrid = []string{
	"............",
	"........0...",
	".....0......",
	".......0....",
	"....0.......",
	"......A.....",
	"............",
	"............",
	"........A...",
	".........A..",
	"............",
	"............",
}

func TestDay08InputParsing(t *testing.T) {
	ant, ht, wt := parseGrid(exampleGrid)
	expected := position{row: 5, col: 6}
	if ht != 12 {
		t.Errorf("wrong height: expected 12, got %v", ht)
	}
	if wt != 12 {
		t.Errorf("wrong width: expected 12, got %v", wt)
	}
	if ant['A'][0] != expected {
		t.Errorf("antenna in wrong position: expected (5, 6), got %v", ant['A'][0])
	}
}

func TestDay08ExamplePart1(t *testing.T) {
	ant, ht, wt := parseGrid(exampleGrid)
	if ans := solvePart1(ant, ht, wt); ans != 14 {
		t.Errorf("expected 14, got %d", ans)
	}
}

func TestDay08ExamplePart2(t *testing.T) {
	ant, ht, wt := parseGrid(exampleGrid)
	if ans := solvePart2(ant, ht, wt); ans != 34 {
		t.Errorf("expected 34, got %d", ans)
	}
}
