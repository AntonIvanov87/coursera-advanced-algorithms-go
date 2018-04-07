package assig4

import (
	"testing"
	"math/rand"
)

func Test1Paint3(t *testing.T) {
	edges := [][2]int{
		{1, 3},
		{1, 4},
		{3, 4},
		{2, 4},
		{2, 3},
	}
	vertIndexToForbiddenColor := []Color{R, R, R, G}

	vertIndexToColor := Paint3(edges, vertIndexToForbiddenColor)

	expectedColors := []Color{G, G, B, R}
	if !coloringsEqual(vertIndexToColor, expectedColors) {
		t.Errorf("Expected colors %v, got %v", expectedColors, vertIndexToColor)
	}
}

func Test2Paint3(t *testing.T) {
	edges := [][2]int{
		{1, 3},
		{1, 4},
		{3, 4},
		{2, 4},
		{2, 3},
	}
	vertIndexToForbiddenColor := []Color{R, G, R, R}

	vertIndexToColor := Paint3(edges, vertIndexToForbiddenColor)

	if len(vertIndexToColor) != 0 {
		t.Errorf("Expected impossible coloring, got %v", vertIndexToColor)
	}
}

func Test3Paint3(t *testing.T) {
	var edges [][2]int
	vertIndexToForbiddenColor := []Color{R}

	vertIndexToColor := Paint3(edges, vertIndexToForbiddenColor)

	expectedColoring1 := []Color{G}
	expectedColoring2 := []Color{B}
	if !coloringsEqual(vertIndexToColor, expectedColoring1) && !coloringsEqual(vertIndexToColor, expectedColoring2) {
		t.Errorf("Expected coloring %v or %v, got %v", expectedColoring1, expectedColoring2, vertIndexToColor)
	}
}

func Test4Pain3(t *testing.T) {
	edges := [][2]int{
		{1, 2},
	}
	vertIndexToForbiddenColor := []Color{R, G}

	vertIndexToColor := Paint3(edges, vertIndexToForbiddenColor)

	expectedColorings := [][]Color{
		{G, R},
		{G, B},
		{B, R},
	}
	if !oneOfExpectedColorings(vertIndexToColor, expectedColorings) {
		t.Errorf("Expected colorings %v, got %v", expectedColorings, vertIndexToColor)
	}
}

func Test4Pain4(t *testing.T) {
	edges := [][2]int{
		{1, 2},
	}
	vertIndexToForbiddenColor := []Color{R, R}

	vertIndexToColor := Paint3(edges, vertIndexToForbiddenColor)

	expectedColorings := [][]Color{
		{G, B},
		{B, R},
	}
	if !oneOfExpectedColorings(vertIndexToColor, expectedColorings) {
		t.Errorf("Expected colorings %v, got %v", expectedColorings, vertIndexToColor)
	}
}

func Test5Paint3(t *testing.T) {
	edges := [][2]int{
		{1, 2},
		{1, 3},
	}
	vertIndexToForbiddenColor := []Color{R, R, G}

	vertIndexToColor := Paint3(edges, vertIndexToForbiddenColor)

	expectedColorings := [][]Color{
		{G, B, R},
		{G, B, B},
		{B, G, R},
	}
	if !oneOfExpectedColorings(vertIndexToColor, expectedColorings) {
		t.Errorf("Expected colorings %v, got %v", expectedColorings, vertIndexToColor)
	}
}

func Test6Pain3(t *testing.T) {
	edges := [][2]int{
		{1, 2},
	}
	vertIndexToForbiddenColor := []Color{G, G}

	vertIndexToColor := Paint3(edges, vertIndexToForbiddenColor)

	expectedColorings := [][]Color{
		{R, B},
		{B, R},
	}
	if !oneOfExpectedColorings(vertIndexToColor, expectedColorings) {
		t.Errorf("Expected colorings %v, got %v", expectedColorings, vertIndexToColor)
	}
}

func TestRandPaint3(t *testing.T) {
	for i := 0; i < 1000; i++ {
		vertIndexToForbiddenColor := genForbiddenColors()
		edges := genPaint3Edges(len(vertIndexToForbiddenColor))

		vertIndexToColor := Paint3(edges, vertIndexToForbiddenColor)

		if len(vertIndexToColor) == 0 {
			bfSolution := paint3BruteForce(edges, vertIndexToForbiddenColor)
			if len(bfSolution) != 0 {
				t.Errorf("Expected solution %v, but found none.", bfSolution)
			}
		} else {
			checkSolution(edges, vertIndexToForbiddenColor, vertIndexToColor, t)
		}

		if t.Failed() {
			t.Logf("Edges: %v\n"+
				"Forbidden colors %v\n"+
				"Solution: %v",
				edges, vertIndexToForbiddenColor, vertIndexToColor)
			return
		}
	}
}

func coloringsEqual(coloring1 []Color, coloring2 []Color) bool {
	if len(coloring1) != len(coloring2) {
		return false
	}

	for i, color1 := range coloring1 {
		if coloring2[i] != color1 {
			return false
		}
	}

	return true
}

func oneOfExpectedColorings(actualColoring []Color, expectedColorings [][]Color) bool {
	for _, expectedColoring := range expectedColorings {
		if coloringsEqual(actualColoring, expectedColoring) {
			return true
		}
	}
	return false
}

func genForbiddenColors() []Color {
	verticesCount := rand.Intn(50) + 1
	forbiddenColors := make([]Color, verticesCount)
	for i := range forbiddenColors {
		forbiddenColors[i] = Color(rand.Intn(3))
	}
	return forbiddenColors
}

func genPaint3Edges(verticesCount int) (edges [][2]int) {
	if verticesCount == 1 {
		return
	}

	maxPossibleEdgesCount := verticesCount * (verticesCount - 1) / 2
	possibleEdges := make([][2]int, 0, maxPossibleEdgesCount)
	for from := 1; from < verticesCount; from++ {
		for to := from + 1; to <= verticesCount; to++ {
			possibleEdges = append(possibleEdges, [2]int{from, to})
		}
	}

	edgesCount := rand.Intn(verticesCount*20) + 1
	if edgesCount > maxPossibleEdgesCount {
		edgesCount = maxPossibleEdgesCount
	}

	edges = make([][2]int, 0, edgesCount)
	remainingPossibleEdgesCount := len(possibleEdges)
	for {
		i := rand.Intn(remainingPossibleEdgesCount)
		edges = append(edges, possibleEdges[i])
		possibleEdges[i] = possibleEdges[remainingPossibleEdgesCount-1]
		remainingPossibleEdgesCount--
		if len(edges) == edgesCount {
			break
		}
	}
	return
}

func checkSolution(edges [][2]int, vertIToForbiddenColor []Color, vertIToColor []Color, t *testing.T) {
	if len(vertIToColor) != len(vertIToForbiddenColor) {
		t.Errorf("Expected %v verticies in solution, got %v", len(vertIToForbiddenColor), len(vertIToColor))
		return
	}

	for i, color := range vertIToColor {
		if color == vertIToForbiddenColor[i] {
			t.Errorf("Vertex %v has forbidden color %v in solution", i+1, color)
		}
	}

	for _, edge := range edges {
		if vertIToColor[edge[0]-1] == vertIToColor[edge[1]-1] {
			t.Errorf("Adjacent vertices %v and %v share the same color %v", edge[0], edge[1], vertIToColor[edge[0]-1])
		}
	}
}

func paint3BruteForce(edges [][2]int, vertIToForbiddenColor []Color) []Color {

	graph := make([][]int, len(vertIToForbiddenColor))
	for _, edge := range edges {
		vertI1 := edge[0] - 1
		vertI2 := edge[1] - 1
		graph[vertI1] = append(graph[vertI1], vertI2)
		graph[vertI2] = append(graph[vertI2], vertI1)
	}

	vertIToColor := make([]Color, 0, len(vertIToForbiddenColor))

	return paint3GraphBruteForce(graph, vertIToForbiddenColor, vertIToColor)

}

func paint3GraphBruteForce(graph [][]int, vertIToForbiddenColor []Color, vertIToColor []Color) []Color {
	if len(vertIToColor) == len(vertIToForbiddenColor) {
		return vertIToColor
	}

	vertI := len(vertIToColor)
	forbiddenColors := make(map[Color]bool, 1)
	forbiddenColors[vertIToForbiddenColor[vertI]] = true
	for _, adjacentVertI := range graph[vertI] {
		if adjacentVertI < len(vertIToColor) {
			forbiddenColors[vertIToColor[adjacentVertI]] = true
		}
	}

	if len(forbiddenColors) == 3 {
		return nil
	}

	for colorI := 0; colorI < 3; colorI++ {
		color := Color(colorI)
		if !forbiddenColors[color] {
			newVertIToColor := make([]Color, len(vertIToColor)+1)
			copy(newVertIToColor, vertIToColor)
			newVertIToColor[len(newVertIToColor)-1] = color
			solution := paint3GraphBruteForce(graph, vertIToForbiddenColor, newVertIToColor)
			if len(solution) != 0 {
				return solution
			}
		}
	}

	return nil
}
