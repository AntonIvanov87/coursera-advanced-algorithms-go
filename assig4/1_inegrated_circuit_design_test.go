package assig4

import (
	"testing"
	"math/rand"
)

func Test2CNF1(t *testing.T) {
	clauses := [][]int{
		{1, -3},
		{-1, 2},
		{-2, -3},
	}

	result := Solve2CNF(3, clauses)

	// TODO: other results are also possible
	expectedResult := []bool{true, true, false}
	if len(result) != len(expectedResult) {
		t.Fatalf("Expected: %v, got: %v", expectedResult, result)
	}
	for i := range result {
		if result[i] != expectedResult[i] {
			t.Fatalf("Expected: %v, got: %v", expectedResult, result)
		}
	}
}

func Test2CNF2(t *testing.T) {
	clauses := [][]int{
		{1, 1},
		{-1, -1},
	}

	result := Solve2CNF(1, clauses)

	if len(result) > 0 {
		t.Errorf("Expected: unsatisfiable, got: %v", result)
	}
}

func Test2CNF3(t *testing.T) {
	clauses := [][]int{
		{2, -3},
		{-1, -3},
		{-2, -1},
	}

	result := Solve2CNF(3, clauses)

	expectedResult := []bool{true, false, false}
	if len(result) != len(expectedResult) {
		t.Fatalf("Expected: %v, got: %v", expectedResult, result)
	}
	for i := range result {
		if result[i] != expectedResult[i] {
			t.Fatalf("Expected: %v, got: %v", expectedResult, result)
		}
	}
}

func Test2CNFRand(t *testing.T) {
	for i := 0; i < 1000; i++ {
		varsCount, clauses := genClauses()

		result := Solve2CNF(varsCount, clauses)

		if len(result) == 0 {
			bfResult := solve2CNFBruteForce(varsCount, clauses)
			if len(bfResult) != 0 {
				t.Errorf("Expected some solution, got none\n"+
					"2CNF: %v\n"+
					"Solution: %v",
					clauses, bfResult)
			}
		} else {
			for _, clause := range clauses {
				clauseIsTrue := checkClause(clause, result)
				if !clauseIsTrue {
					t.Errorf("Wrong solution:\n"+
						"2CNF: %v\n"+
						"Solution: %v",
						clauses, result)
				}
				checkClause(clause, result)
			}
		}

	}
}

func genClauses() (varsCount int, clauses [][]int) {
	varsCount = rand.Intn(30) + 1
	clauses = make([][]int, varsCount)
	for c := range clauses {
		variable := genVar(varsCount)
		clauses[c] = append(clauses[c], variable)
		if rand.Intn(10) > 0 {
			variable = genVar(varsCount)
			clauses[c] = append(clauses[c], variable)
		}
	}
	return
}

func genVar(varsCount int) int {
	varNo := rand.Intn(varsCount) + 1
	if rand.Intn(2) == 1 {
		return -varNo
	}
	return varNo
}

func solve2CNFBruteForce(varsCount int, clauses [][]int) []bool {
	solution := solveBranch(varsCount, clauses, true)
	if len(solution) > 0 {
		return solution
	}
	return solveBranch(varsCount, clauses, false)
}

func solveBranch(varsCount int, clauses [][]int, lastVarVal bool) (solution []bool) {
	rewrittenClauses, ok := rewriteClauses(varsCount, clauses, lastVarVal)
	if !ok {
		return
	}

	if varsCount == 1 {
		return append(solution, lastVarVal)
	}

	smallerSolution := solve2CNFBruteForce(varsCount-1, rewrittenClauses)
	if len(smallerSolution) == 0 {
		return
	}

	return append(smallerSolution, lastVarVal)
}

func rewriteClauses(varsCount int, clauses [][]int, lastVarVal bool) (rewrittenClauses [][]int, ok bool) {
	lastVar := varsCount
	for _, clause := range clauses {
		rewrittenClause, rewriteOk := rewriteClause(clause, lastVar, lastVarVal)
		if !rewriteOk {
			ok = false
			return
		}
		if len(rewrittenClause) > 0 {
			rewrittenClauses = append(rewrittenClauses, rewrittenClause)
		}
	}
	ok = true
	return
}

func rewriteClause(clause []int, varNo int, varVal bool) (rewrittenClause []int, ok bool) {
	var0No := abs(clause[0])
	if var0No == varNo {
		if clause[0] > 0 && varVal || clause[0] < 0 && !varVal {
			ok = true
			return
		}
		if len(clause) == 1 {
			ok = false
			return
		}
	}
	if len(clause) == 1 {
		rewrittenClause = clause
		ok = true
		return
	}
	var1No := abs(clause[1])
	if var1No == varNo {
		if clause[1] > 0 && varVal || clause[1] < 0 && !varVal {
			ok = true
			return
		}
	}
	if var0No == varNo {
		if var1No == varNo {
			ok = false
			return
		}
		rewrittenClause = append(rewrittenClause, clause[1])
		ok = true
		return
	}
	if var1No == varNo {
		rewrittenClause = append(rewrittenClause, clause[0])
		ok = true
		return
	}
	rewrittenClause = clause
	ok = true
	return
}

func checkClause(clause []int, varIndexToVal []bool) bool {
	for _, variable := range clause {
		if variable > 0 {
			if varIndexToVal[variable-1] {
				return true
			}
		} else {
			if !varIndexToVal[abs(variable)-1] {
				return true
			}
		}
	}
	return false
}
