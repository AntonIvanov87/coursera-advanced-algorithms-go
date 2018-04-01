package assig4

import "testing"

func TestTSP1(t *testing.T) {
	edges := []Edge{
		{1, 2, 20},
		{1, 3, 42},
		{1, 4, 35},
		{2, 3, 30},
		{2, 4, 34},
		{3, 4, 12},
	}

	minLen, minPath := SolveTSP(edges, 4)
	if minLen != 97 {
		t.Errorf("Expected min length %v, got %v", 97, minLen)
	}

	expectedShortestPath := []int{1, 2, 3, 4}
	if !slicesEqual(minPath, expectedShortestPath) {
		t.Errorf("Expected min path %v, got %v", expectedShortestPath, minPath)
	}
}

func TestTSP2(t *testing.T) {
	edges := []Edge{
		{1, 2, 1},
		{2, 3, 4},
		{3, 4, 5},
		{4, 2, 1},
	}

	minLen, minPath := SolveTSP(edges, 4)
	if minLen != -1 {
		t.Errorf("Expected min length %v, got %v", -1, minLen)
	}

	if len(minPath) != 0 {
		t.Errorf("Expected empty path, got %v", minPath)
	}
}
