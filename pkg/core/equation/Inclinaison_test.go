package equation

import (
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"math"
	"testing"
)

func TestInclinaison_Compute(t *testing.T) {
	variables := []data.CaptorValue{
		{
			Captor: "c1",
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
		"c1.":  60,
		"v1.":  2538517.00449813,
		"v2.":  1182022.11430351,
		"v3.":  2538517.052,
		"v4.":  1182022.091,
		"v5.":  500.01,
		"v6.":  2538518.37302081,
		"v7.":  1182022.73832296,
		"v8.":  2538518.357,
		"v9.":  1182022.685,
		"v10.": 510.03,
	}
	expectedResult := -6.88981215571969
	inclinaison := Inclinaison{}
	result := inclinaison.Compute(variables, vals)
	if math.Round(result*10) != math.Round(expectedResult*10) {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}
