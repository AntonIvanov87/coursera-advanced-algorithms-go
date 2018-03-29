package assig4

func FindMaxIndependentSet(funFactors []int, edges [][2]int) (maxFunFactor int, invited []int) {
	graph := makeGraph(edges, len(funFactors))
	vertToMaxFunParams := make([]maxFunParams, len(funFactors)+1)
	// TODO: make recursion not so deep, finding the center of the graph
	fillMaxFunParams(vertToMaxFunParams, graph, 0, funFactors)
	maxFunParams := vertToMaxFunParams[0]

	invited = make([]int, len(maxFunParams.invited))
	for i := range invited {
		invited[i] = maxFunParams.invited[i] + 1
	}
	return maxFunParams.factor, invited
}

func makeGraph(edges [][2]int, verticesCount int) [][]int {
	graph := make([][]int, verticesCount)
	for _, edge := range edges {
		vert1 := edge[0] - 1
		vert2 := edge[1] - 1
		graph[vert1] = append(graph[vert1], vert2)
		graph[vert2] = append(graph[vert2], vert1)
	}
	return graph
}

type maxFunParams struct {
	factor  int
	invited []int
}

func fillMaxFunParams(vertToMaxFunParams []maxFunParams, graph [][]int, vert int, funFactors []int) {
	nextVertsMaxFunParams := maxFunParams{}
	nextNextVertsMaxFunParams := maxFunParams{factor: funFactors[vert], invited: []int{vert}}

	// -1 - dummy value that means "in progress"
	vertToMaxFunParams[vert].factor = -1

	for _, nextVert := range graph[vert] {
		if vertToMaxFunParams[nextVert].factor == -1 {
			continue
		}
		if vertToMaxFunParams[nextVert].factor == 0 {
			fillMaxFunParams(vertToMaxFunParams, graph, nextVert, funFactors)
		}

		nextVertsMaxFunParams.factor += vertToMaxFunParams[nextVert].factor
		nextVertsMaxFunParams.invited = append(nextVertsMaxFunParams.invited, vertToMaxFunParams[nextVert].invited...)

		for _, nextNextVert := range graph[nextVert] {
			if vertToMaxFunParams[nextNextVert].factor == -1 {
				continue
			}
			nextNextVertsMaxFunParams.factor += vertToMaxFunParams[nextNextVert].factor
			nextNextVertsMaxFunParams.invited = append(nextNextVertsMaxFunParams.invited, vertToMaxFunParams[nextNextVert].invited...)
		}
	}

	if nextVertsMaxFunParams.factor > nextNextVertsMaxFunParams.factor {
		vertToMaxFunParams[vert] = nextVertsMaxFunParams
	} else {
		vertToMaxFunParams[vert] = nextNextVertsMaxFunParams
	}
}
