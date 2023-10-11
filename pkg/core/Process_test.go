package core

import (
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"log"
	"testing"
)

func testCaptorVariables(expectedCaptors, captors []data.CaptorValue, expectedVariables, variables []data.EquationCaptor, t *testing.T) {
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
		if variables[i].Name != expectedVariables[i].Name {
			t.Errorf("Expected %v variable, got %v", expectedVariables[i].Name, variables[i].Name)
			return
		}
		if len(variables[i].Captors) != len(expectedVariables[i].Captors) {
			t.Errorf("Expected %v variables, got %v", len(expectedVariables[i].Captors), len(variables[i].Captors))
			return
		}
		for j := range variables[i].Captors {
			if variables[i].GetCaptor(j).String() != expectedVariables[i].GetCaptor(j).String() {
				t.Errorf("Expected %v variable, got %v", expectedVariables[i].GetCaptor(j), variables[i].GetCaptor(j))
				return
			}
		}
	}
}

func TestParseVariables(t *testing.T) {
	brute := "1:A.a;2:B.b;3:C.c;"
	log.Printf("Testing %s ", brute)
	expectedCaptors := []data.CaptorValue{
		{
			Captor: "A",
			Field:  "a",
		},
		{
			Captor: "B",
			Field:  "b",
		},
		{
			Captor: "C",
			Field:  "c",
		},
	}
	expectedVariables := []data.EquationCaptor{
		{
			Name: "1",
			Captors: []data.CaptorValue{
				{
					Captor: "A",
					Field:  "a",
				},
			},
		},
		{
			Name: "2",
			Captors: []data.CaptorValue{
				{
					Captor: "B",
					Field:  "b",
				},
			},
		},
		{
			Name: "3",
			Captors: []data.CaptorValue{
				{
					Captor: "C",
					Field:  "c",
				},
			},
		},
	}
	captors, variables, err := ParseVariables(brute)
	if err != nil {
		t.Error(err)
		return
	}
	testCaptorVariables(expectedCaptors, captors, expectedVariables, variables, t)
	log.Printf("passed\n")

	brute = "1:A.a,B.b,C.c;2:D.d,E.e,F.f;"
	log.Printf("Testing %s ", brute)
	expectedCaptors = []data.CaptorValue{
		{
			Captor: "A",
			Field:  "a",
		},
		{
			Captor: "B",
			Field:  "b",
		},
		{
			Captor: "C",
			Field:  "c",
		},
		{
			Captor: "D",
			Field:  "d",
		},
		{
			Captor: "E",
			Field:  "e",
		},
		{
			Captor: "F",
			Field:  "f",
		},
	}
	expectedVariables = []data.EquationCaptor{
		{
			Name: "1",
			Captors: []data.CaptorValue{
				{
					Captor: "A",
					Field:  "a",
				},
				{
					Captor: "B",
					Field:  "b",
				},
				{
					Captor: "C",
					Field:  "c",
				},
			},
		},
		{
			Name: "2",
			Captors: []data.CaptorValue{
				{
					Captor: "D",
					Field:  "d",
				},
				{
					Captor: "E",
					Field:  "e",
				},
				{
					Captor: "F",
					Field:  "f",
				},
			},
		},
	}
	captors, variables, err = ParseVariables(brute)
	testCaptorVariables(expectedCaptors, captors, expectedVariables, variables, t)
	log.Printf("passed\n")

	brute = "1:A.a,B.b,C.c;2:B.b,C.c,D.d;"
	log.Printf("Testing %s ", brute)
	expectedCaptors = []data.CaptorValue{
		{
			Captor: "A",
			Field:  "a",
		},
		{
			Captor: "B",
			Field:  "b",
		},
		{
			Captor: "C",
			Field:  "c",
		},
		{
			Captor: "D",
			Field:  "d",
		},
	}
	expectedVariables = []data.EquationCaptor{
		{
			Name: "1",
			Captors: []data.CaptorValue{
				{
					Captor: "A",
					Field:  "a",
				},
				{
					Captor: "B",
					Field:  "b",
				},
				{
					Captor: "C",
					Field:  "c",
				},
			},
		},
		{
			Name: "2",
			Captors: []data.CaptorValue{
				{
					Captor: "B",
					Field:  "b",
				},
				{
					Captor: "C",
					Field:  "c",
				},
				{
					Captor: "D",
					Field:  "d",
				},
			},
		},
	}
	captors, variables, err = ParseVariables(brute)
	testCaptorVariables(expectedCaptors, captors, expectedVariables, variables, t)
	log.Printf("passed\n")

	brute = "1:A.a,B.b,C.c;\n2:B.b,C.c,D.d;\n"
	log.Printf("Testing %s ", brute)
	expectedCaptors = []data.CaptorValue{
		{
			Captor: "A",
			Field:  "a",
		},
		{
			Captor: "B",
			Field:  "b",
		},
		{
			Captor: "C",
			Field:  "c",
		},
		{
			Captor: "D",
			Field:  "d",
		},
	}
	expectedVariables = []data.EquationCaptor{
		{
			Name: "1",
			Captors: []data.CaptorValue{
				{
					Captor: "A",
					Field:  "a",
				},
				{
					Captor: "B",
					Field:  "b",
				},
				{
					Captor: "C",
					Field:  "c",
				},
			},
		},
		{
			Name: "2",
			Captors: []data.CaptorValue{
				{
					Captor: "B",
					Field:  "b",
				},
				{
					Captor: "C",
					Field:  "c",
				},
				{
					Captor: "D",
					Field:  "d",
				},
			},
		},
	}
	captors, variables, err = ParseVariables(brute)
	testCaptorVariables(expectedCaptors, captors, expectedVariables, variables, t)
	log.Printf("passed\n")

	brute = "1:A.a,B.b,C.c; 2:B.b, C.c,D.d ;"
	log.Printf("Testing %s ", brute)
	expectedCaptors = []data.CaptorValue{
		{
			Captor: "A",
			Field:  "a",
		},
		{
			Captor: "B",
			Field:  "b",
		},
		{
			Captor: "C",
			Field:  "c",
		},
		{
			Captor: "D",
			Field:  "d",
		},
	}
	expectedVariables = []data.EquationCaptor{
		{
			Name: "1",
			Captors: []data.CaptorValue{
				{
					Captor: "A",
					Field:  "a",
				},
				{
					Captor: "B",
					Field:  "b",
				},
				{
					Captor: "C",
					Field:  "c",
				},
			},
		},
		{
			Name: "2",
			Captors: []data.CaptorValue{
				{
					Captor: "B",
					Field:  "b",
				},
				{
					Captor: "C",
					Field:  "c",
				},
				{
					Captor: "D",
					Field:  "d",
				},
			},
		},
	}
	captors, variables, err = ParseVariables(brute)
	testCaptorVariables(expectedCaptors, captors, expectedVariables, variables, t)
	log.Printf("passed\n")
}

func TestParseVariables_constant(t *testing.T) {
	brute := "1:${42};"
	log.Printf("Testing %s ", brute)
	expectedCaptors := []data.CaptorValue{
		{
			Captor: "$",
			Field:  "42",
		},
	}
	expectedVariables := []data.EquationCaptor{
		{
			Name: "1",
			Captors: []data.CaptorValue{
				{
					Captor: "$",
					Field:  "42",
				},
			},
		},
	}
	captors, variables, err := ParseVariables(brute)
	if err != nil {
		t.Error(err)
		return
	}
	testCaptorVariables(expectedCaptors, captors, expectedVariables, variables, t)
}

func FuzzParseVariables(f *testing.F) {
	f.Add("A;B;C;")
	f.Add("A,B,C;D,E,F;")
	f.Add("A,B,C;B,C,D;")
	f.Add("A,B,C;\nB,C,D;\n")
	f.Add("A,B,C; B, C,D ;")
	f.Fuzz(func(t *testing.T, brute string) {
		_, _, _ = ParseVariables(brute)
	})
}
