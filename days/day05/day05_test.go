package day05

import "testing"

var exampleRuleset []rule = []rule{
	{before: 47, after: 53},
	{before: 97, after: 13},
	{before: 97, after: 61},
	{before: 97, after: 47},
	{before: 75, after: 29},
	{before: 61, after: 13},
	{before: 75, after: 53},
	{before: 29, after: 13},
	{before: 97, after: 29},
	{before: 53, after: 29},
	{before: 61, after: 53},
	{before: 97, after: 53},
	{before: 61, after: 29},
	{before: 47, after: 13},
	{before: 75, after: 47},
	{before: 97, after: 75},
	{before: 47, after: 61},
	{before: 75, after: 61},
	{before: 47, after: 29},
	{before: 75, after: 13},
	{before: 53, after: 13},
}

var exampleUpdates []update = []update{
	{75, 47, 61, 53, 29},
	{97, 61, 53, 29, 13},
	{75, 29, 13},
	{75, 97, 47, 61, 53},
	{61, 13, 29},
	{97, 13, 75, 29, 47},
}

func TestDay05ExamplePart1(t *testing.T) {
	if ans := solvePart1(exampleRuleset, exampleUpdates); ans != 143 {
		t.Errorf("expected 143, got %d", ans)
	}
}

func TestDay05ExamplePart2(t *testing.T) {
	if ans := solvePart2(exampleRuleset, exampleUpdates); ans != 9 {
		t.Errorf("expected 9, got %d", ans)
	}
}
