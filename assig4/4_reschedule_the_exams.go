package assig4

func Paint3(edges [][2]int, vertIndexToForbiddenColor []Color) (vertIndexToColor []Color) {
	verticesCount := len(vertIndexToForbiddenColor)
	clauses := make([][]int, 0, 3*verticesCount+3*len(edges))

	// a vertex must have one of non forbidden colors
	for vertIndex, forbiddenColor := range vertIndexToForbiddenColor {
		vertNo := vertIndex + 1

		forbiddenColorVar := getVar(vertNo, forbiddenColor, verticesCount)
		clauses = append(clauses, []int{-forbiddenColorVar})

		otherColor1, otherColor2 := getOtherColors(forbiddenColor)
		otherColor1Var := getVar(vertNo, otherColor1, verticesCount)
		otherColor2Var := getVar(vertNo, otherColor2, verticesCount)

		clauses = append(clauses, []int{otherColor1Var, otherColor2Var})
		clauses = append(clauses, []int{-otherColor1Var, -otherColor2Var})
	}

	// adjacent vertices must not have same colors
	for _, edge := range edges {
		vertNo1 := edge[0]
		vertNo2 := edge[1]

		for _, color := range [...]Color{R, G, B} {
			// TODO: the color can be already forbidden to either vert1 or vert2 or even both
			vert1Var := getVar(vertNo1, color, verticesCount)
			vert2Var := getVar(vertNo2, color, verticesCount)
			clauses = append(clauses, []int{-vert1Var, -vert2Var})
		}
	}

	vertIndexToBool := Solve2CNF(verticesCount*3, clauses)
	if len(vertIndexToBool) == 0 {
		return
	}

	vertIndexToColor = make([]Color, verticesCount)
	assignColorsFromVars(vertIndexToColor, vertIndexToBool, R)
	assignColorsFromVars(vertIndexToColor, vertIndexToBool, G)
	assignColorsFromVars(vertIndexToColor, vertIndexToBool, B)

	return
}

type Color byte

const (
	R = Color(iota)
	G
	B
)

func getOtherColors(color Color) (Color, Color) {
	switch color {
	case R:
		return G, B
	case G:
		return R, B
	case B:
		return R, G
	default:
		panic("Unknown Color")
	}
}

func getVar(vertNo int, color Color, verticesCount int) int {
	return int(color)*verticesCount + vertNo
}

func assignColorsFromVars(vertIndexToColor []Color, vertIndexToBool []bool, colorRange Color) {
	verticesCount := len(vertIndexToColor)
	colorRangeInt := int(colorRange)
	for i := colorRangeInt * verticesCount; i < (colorRangeInt+1)*verticesCount; i++ {
		if vertIndexToBool[i] {
			vertIndexToColor[i-colorRangeInt*verticesCount] = colorRange
		}
	}
}
