package assig4

func SolveTSP(edges []Edge, verticesCount int) (minLen int, minPath []int) {
	graph := make([][]toEdge, verticesCount)
	for _, edge := range edges {
		vert1 := edge.vert1 - 1
		vert2 := edge.vert2 - 1
		graph[vert1] = append(graph[vert1], toEdge{vert2, edge.length})
		graph[vert2] = append(graph[vert2], toEdge{vert1, edge.length})
	}

	minLen = -1
	for startVert := 0; startVert < verticesCount; startVert++ {
		remainingVerts := make(map[int]bool, verticesCount-1)
		for vert := 0; vert < verticesCount; vert++ {
			if vert != startVert {
				remainingVerts[vert] = true
			}
		}
		length, path := solveTSPFrom(graph, []int{startVert}, 0, remainingVerts, -1)
		if length != -1 && (length < minLen || minLen == -1)  {
			minLen = length
			minPath = path
		}
	}

	// bring back original numbers of vertices instead of indices
	for i := range minPath {
		minPath[i] += 1
	}

	return
}

type Edge struct {
	vert1  int
	vert2  int
	length int
}

type toEdge struct {
	toVert int
	length int
}

func solveTSPFrom(graph [][]toEdge, startPath []int, startPathLen int, remainingVerts map[int]bool, minLenOfOtherPath int) (minLen int, minPath []int) {
	lastVert := startPath[len(startPath)-1]
	minLen = -1

	if minLenOfOtherPath != -1 && startPathLen > minLenOfOtherPath {
		return
	}

	if len(remainingVerts) == 0 {
		for _, edge := range graph[lastVert] {
			if edge.toVert == startPath[0] {
				minLen = startPathLen + edge.length
				minPath = startPath
				return
			}
		}
		// there is no edge from last vert int path to first vert in path
		return
	}

	for _, edge := range graph[lastVert] {
		if !remainingVerts[edge.toVert] {
			continue
		}
		newRemainingVerts := make(map[int]bool, len(remainingVerts)-1)
		for vert := range remainingVerts {
			if vert != edge.toVert {
				newRemainingVerts[vert] = true
			}
		}

		newLen := startPathLen + edge.length
		if minLenOfOtherPath != -1 && newLen > minLenOfOtherPath {
			continue
		}

		length, path := solveTSPFrom(graph, append(startPath, edge.toVert), startPathLen + edge.length, newRemainingVerts, minLenOfOtherPath)

		if length != -1 && (length < minLen || minLen == -1) {
			minLen = length
			minPath = path
			if minLen < minLenOfOtherPath {
				minLenOfOtherPath = minLen
			}
		}
	}
	return
}
