package core

import "testing"

func testCaptorVariables(expectedCaptors []string, captors []string, expectedVariables [][]string, variables [][]string, t *testing.T) {
	if len(captors) != len(expectedCaptors) {
		t.Errorf("Expected %v captors, got %v", len(expectedCaptors), len(captors))
		return
	}
	for i := range captors {
		if captors[i] != expectedCaptors[i] {
			t.Errorf("Expected %v captor, got %v at index %d", expectedCaptors[i], captors[i], i)
			return
		}
	}
	if len(variables) != len(expectedVariables) {
		t.Errorf("Expected %v variables, got %v", len(expectedVariables), len(variables))
		return
	}
	for i := range variables {
		if len(variables[i]) != len(expectedVariables[i]) {
			t.Errorf("Expected %v variables, got %v", len(expectedVariables[i]), len(variables[i]))
			return
		}
		for j := range variables[i] {
			if variables[i][j] != expectedVariables[i][j] {
				t.Errorf("Expected %v variable, got %v", expectedVariables[i][j], variables[i][j])
				return
			}
		}
	}
}

func TestParseVariables(t *testing.T) {
	brute := "A;B;C;"
	expectedCaptors := []string{"A", "B", "C"}
	expectedVariables := [][]string{
		{"A"},
		{"B"},
		{"C"},
	}
	captors, variables := ParseVariables(brute)
	testCaptorVariables(expectedCaptors, captors, expectedVariables, variables, t)

	brute = "A,B,C;D,E,F;"
	expectedCaptors = []string{"A", "B", "C", "D", "E", "F"}
	expectedVariables = [][]string{
		{"A", "B", "C"},
		{"D", "E", "F"},
	}
	captors, variables = ParseVariables(brute)
	testCaptorVariables(expectedCaptors, captors, expectedVariables, variables, t)

	brute = "A,B,C;B,C,D;"
	expectedCaptors = []string{"A", "B", "C", "D"}
	expectedVariables = [][]string{
		{"A", "B", "C"},
		{"B", "C", "D"},
	}
	captors, variables = ParseVariables(brute)
	testCaptorVariables(expectedCaptors, captors, expectedVariables, variables, t)

	brute = "A,B,C;\nB,C,D;\n"
	expectedCaptors = []string{"A", "B", "C", "D"}
	expectedVariables = [][]string{
		{"A", "B", "C"},
		{"B", "C", "D"},
	}
	captors, variables = ParseVariables(brute)
	testCaptorVariables(expectedCaptors, captors, expectedVariables, variables, t)

	brute = "A,B,C; B, C,D ;"
	expectedCaptors = []string{"A", "B", "C", "D"}
	expectedVariables = [][]string{
		{"A", "B", "C"},
		{"B", "C", "D"},
	}
	captors, variables = ParseVariables(brute)
	testCaptorVariables(expectedCaptors, captors, expectedVariables, variables, t)
}

func FuzzParseVariables(f *testing.F) {
	f.Add("A;B;C;")
	f.Add("A,B,C;D,E,F;")
	f.Add("A,B,C;B,C,D;")
	f.Add("A,B,C;\nB,C,D;\n")
	f.Add("A,B,C; B, C,D ;")
	f.Fuzz(func(t *testing.T, brute string) {
		ParseVariables(brute)
	})
}
