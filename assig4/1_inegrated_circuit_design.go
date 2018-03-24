package assig4

func Solve2CNF(varsCount int, clauses [][]int) []bool {
	graph := constructGraph(clauses, varsCount)

	vertToStrongComp, strongCompToVerts := findStrongComponents(graph)

	assignedComps := make([]bool, len(strongCompToVerts))
	varIndexToVal := make([]bool, varsCount)
	for strongComp := len(strongCompToVerts)-1; strongComp>=0; strongComp-- {
		if assignedComps[strongComp] {
			continue
		}

		assignedVars := make(map[int]bool)
		for _, vert := range strongCompToVerts[strongComp] {
			varIndex, val := getVarIndexAndVal(vert, varsCount)
			if assignedVars[varIndex] {
				return []bool{}
			}
			varIndexToVal[varIndex] = val
			assignedVars[varIndex] = true
		}
		assignedComps[strongComp] = true

		vert := strongCompToVerts[strongComp][0]
		var oppositeVert int
		if vert < varsCount {
			oppositeVert = varsCount + vert
		} else {
			oppositeVert = vert - varsCount
		}
		oppositeComp := vertToStrongComp[oppositeVert]
		assignedComps[oppositeComp] = true
	}

	return varIndexToVal
}

func constructGraph(clauses [][]int, varsCount int) [][]int {
	graph := make([][]int, varsCount*2)
	for _, clause := range clauses {
		fromVertex := getVertex(-clause[0], varsCount)
		if len(clause) == 2 {
			toVertex := getVertex(clause[1], varsCount)
			graph[fromVertex] = append(graph[fromVertex], toVertex)

			fromVertex = getVertex(-clause[1], varsCount)
			toVertex = getVertex(clause[0], varsCount)
			graph[fromVertex] = append(graph[fromVertex], toVertex)
		} else {
			toVertex := getVertex(clause[0], varsCount)
			graph[fromVertex] = append(graph[fromVertex], toVertex)
		}
	}
	return graph
}

func abs(i int) int {
	if i >= 0 {
		return i
	}
	return -i
}

func getVertex(variable int, varsCount int) int {
	if variable > 0 {
		return variable - 1
	} else {
		return varsCount + abs(variable) - 1
	}
}

func getVarIndexAndVal(vertex int, varsCount int) (varIndex int, val bool) {
	if vertex < varsCount {
		return vertex, true
	} else {
		return vertex - varsCount, false
	}
}

func findStrongComponents(graph [][]int) (vertToStrongComponent []int, strongCompToVerts [][]int) {
	revTopOrder := findRevTopolOrder(graph)
	transpGraph := transposeGraph(graph)

	vertToStrongComponent = make([]int, len(graph))
	for i := range vertToStrongComponent {
		vertToStrongComponent[i] = -1
	}

	strongComp := 0
	for i:=len(revTopOrder)-1; i>=0; i-- {
		vert := revTopOrder[i]
		if vertToStrongComponent[vert] != -1 {
			continue
		}

		vertToStrongComponent[vert] = strongComp
		strongCompToVerts = append(strongCompToVerts, []int{vert})
		vertsToVisit := []int{vert}
		for len(vertsToVisit) > 0 {
			fromVert := vertsToVisit[len(vertsToVisit)-1]
			vertsToVisit = vertsToVisit[:len(vertsToVisit)-1]

			for _, toVert := range transpGraph[fromVert] {
				if vertToStrongComponent[toVert] == -1 {
					vertToStrongComponent[toVert] = strongComp
					strongCompToVerts[strongComp] = append(strongCompToVerts[strongComp], toVert)
					vertsToVisit = append(vertsToVisit, toVert)
				}
			}
		}

		strongComp++
	}
	return
}

func findRevTopolOrder(graph [][]int) []int {
	revTopolOrder := make([]int, 0, len(graph))
	visitedVerts := make(map[int]bool)
	for vert := range graph {
		if visitedVerts[vert] {
			continue
		}
		revTopolOrder = dfs(vert, graph, visitedVerts, revTopolOrder)
	}

	return revTopolOrder
}

func dfs(fromVert int, graph [][]int, visitedVerts map[int]bool, revTopolOrder []int) []int {
	visitedVerts[fromVert] = true
	for _, toVertex := range graph[fromVert] {
		if !visitedVerts[toVertex] {
			revTopolOrder = dfs(toVertex, graph, visitedVerts, revTopolOrder)
		}
	}
	return append(revTopolOrder, fromVert)
}

func transposeGraph(graph [][]int) [][]int {
	transpGraph := make([][]int, len(graph))
	for fromVert, toVerts := range graph {
		for _, toVert := range toVerts {
			transpGraph[toVert] = append(transpGraph[toVert], fromVert)
		}
	}
	return transpGraph
}
