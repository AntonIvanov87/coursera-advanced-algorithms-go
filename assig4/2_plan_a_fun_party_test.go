package assig4

import (
	"testing"
	"math/rand"
)

func TestFindMaxIndependentSet1(t *testing.T) {
	funFactors := []int{1000}
	var edges [][2]int

	maxFunFactor, invited := FindMaxIndependentSet(funFactors, edges)

	if maxFunFactor != 1000 {
		t.Errorf("Expected fun factor %v, got %v", 1000, maxFunFactor)
	}

	expectedInvited := []int{1}
	if !slicesEqual(invited, expectedInvited) {
		t.Errorf("Expected invited %v, got %v", expectedInvited, invited)
	}
}

func TestFindMaxIndependentSet2(t *testing.T) {
	funFactors := []int{1, 2}
	edges := [][2]int{{1, 2}}

	maxFunFactor, invited := FindMaxIndependentSet(funFactors, edges)

	if maxFunFactor != 2 {
		t.Errorf("Expected fun factor %v, got %v", 2, maxFunFactor)
	}

	expectedInvited := []int{2}
	if !slicesEqual(invited, expectedInvited) {
		t.Errorf("Expected invited %v, got %v", expectedInvited, invited)
	}
}

func TestFindMaxIndependentSet3(t *testing.T) {
	funFactors := []int{1, 5, 3, 7, 5}
	edges := [][2]int{
		{5, 4},
		{2, 3},
		{4, 2},
		{1, 2},
	}

	maxFunFactor, invited := FindMaxIndependentSet(funFactors, edges)

	if maxFunFactor != 11 {
		t.Errorf("Expected fun factor %v, got %v", 11, maxFunFactor)
	}

	expectedInvited := []int{1, 3, 4}
	if !slicesEqual(invited, expectedInvited) {
		t.Errorf("Expected invited %v, got %v", expectedInvited, invited)
	}
}

func TestFindMaxIndependentSet4(t *testing.T) {
	funFactors := []int{1, 3, 1}
	edges := [][2]int{
		{1, 2},
		{2, 3},
	}

	maxFunFactor, invited := FindMaxIndependentSet(funFactors, edges)

	if maxFunFactor != 3 {
		t.Errorf("Expected fun factor %v, got %v", 3, maxFunFactor)
	}

	expectedInvited := []int{2}
	if !slicesEqual(invited, expectedInvited) {
		t.Errorf("Expected invited %v, got %v", expectedInvited, invited)
	}
}

func TestFindMaxIndependentSetRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		verticesCount := rand.Intn(1000) + 1
		funFactors := genFunFactors(verticesCount)
		edges := genEdges(verticesCount)

		maxFunFactor, invited := FindMaxIndependentSet(funFactors, edges)

		invitedFunFactor := 0
		for _, person := range invited {
			invitedFunFactor += funFactors[person-1]
		}
		if maxFunFactor != invitedFunFactor {
			t.Errorf("Max fun factor %v differs from fun factor of invited people %v", maxFunFactor, invitedFunFactor)
		}

		adjacentPeople := findAdjacentPeople(invited, edges, verticesCount)
		for _, edge := range adjacentPeople {
			t.Errorf("Found adjacent invited people %v and %v", edge[0], edge[1])
		}

		if t.Failed() {
			t.Logf("Fun factors: %v", funFactors)
			t.Logf("Edges: %v", edges)
		}

	}
}
func slicesEqual(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	for i, val := range slice1 {
		if val != slice2[i] {
			return false
		}
	}

	return true
}

func genFunFactors(verticesCount int) []int {
	funFactors := make([]int, verticesCount)
	for i := range funFactors {
		funFactors[i] = rand.Intn(1000) + 1
	}
	return funFactors
}

func genEdges(verticesCount int) [][2]int {
	vertsNotInGraph := make([]int, verticesCount)
	for i := range vertsNotInGraph {
		vertsNotInGraph[i] = i
	}

	notInGraphLen := len(vertsNotInGraph)
	notInGraphI := rand.Intn(notInGraphLen)
	vertToGraph := vertsNotInGraph[notInGraphI]
	vertsNotInGraph[notInGraphI] = vertsNotInGraph[notInGraphLen-1]
	notInGraphLen--

	vertsInGraph := make([]int, 0, verticesCount)
	vertsInGraph = append(vertsInGraph, vertToGraph)

	graph := make([][]int, verticesCount)

	for ; notInGraphLen > 0; notInGraphLen-- {
		notInGraphI = rand.Intn(notInGraphLen)
		vertToGraph = vertsNotInGraph[notInGraphI]
		vertsNotInGraph[notInGraphI] = vertsNotInGraph[notInGraphLen-1]

		inGraphI := rand.Intn(len(vertsInGraph))
		vertFromGraph := vertsInGraph[inGraphI]

		graph[vertFromGraph] = append(graph[vertFromGraph], vertToGraph)
		vertsInGraph = append(vertsInGraph, vertToGraph)
	}

	edges := make([][2]int, 0, verticesCount-1)
	for fromVert, toVerts := range graph {
		for _, toVert := range toVerts {
			edges = append(edges, [2]int{fromVert + 1, toVert + 1})
		}
	}
	return edges
}

func findAdjacentPeople(invited []int, edges [][2]int, verticesCount int) (adjacentPeople [][2]int) {
	graph := make([]map[int]bool, verticesCount)
	for i := range graph {
		graph[i] = make(map[int]bool)
	}
	for _, edge := range edges {
		vert1 := edge[0] - 1
		vert2 := edge[1] - 1
		graph[vert1][vert2] = true
		graph[vert2][vert1] = true
	}

	for i := 0; i < len(invited)-1; i++ {
		for j := i + 1; j < len(invited); j++ {
			vert1 := invited[i] - 1
			vert2 := invited[j] - 1
			if graph[vert1][vert2] || graph[vert2][vert1] {
				adjacentPeople = append(adjacentPeople, [2]int{vert1 + 1, vert2 + 1})
			}
		}
	}
	return
}
