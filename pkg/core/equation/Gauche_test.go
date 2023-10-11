package equation

import (
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"math"
	"testing"
)

func TestGauche_Compute(t *testing.T) {
	variables := []data.CaptorValue{
		{
			Captor: "c1",
			Field:  "",
		},
		{
			Captor: "c2",
			Field:  "",
		},
		{
			Captor: "v1",
			Field:  "",
		},
		{
			Captor: "v2",
			Field:  "",
		},
		{
			Captor: "v3",
			Field:  "",
		},
		{
			Captor: "v4",
			Field:  "",
		},
		{
			Captor: "v5",
			Field:  "",
		},
		{
			Captor: "v6",
			Field:  "",
		},
		{
			Captor: "v7",
			Field:  "",
		},
		{
			Captor: "v8",
			Field:  "",
		},
		{
			Captor: "v9",
			Field:  "",
		},
		{
			Captor: "v10",
			Field:  "",
		},
		{
			Captor: "v11",
			Field:  "",
		},
		{
			Captor: "v12",
			Field:  "",
		},
	}
	vals := map[string]float64{
		"c1.":  0.01,
		"c2.":  1.5,
		"v1.":  2538515.00855942,
		"v2.":  1182026.47941259,
		"v3.":  434.813964758716,
		"v4.":  2538516.37587758,
		"v5.":  1182027.10606514,
		"v6.":  434.81523624641,
		"v7.":  2538513.01130294,
		"v8.":  1182030.84508448,
		"v9.":  434.803792618789,
		"v10.": 2538514.37919406,
		"v11.": 1182031.47048644,
		"v12.": 434.803754732423,
	}
	expectedResult := -0.262046551318759
	gauche := Gauche{}
	result := gauche.Compute(variables, vals)
	if math.Round(result*10) != math.Round(expectedResult*10) {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}
