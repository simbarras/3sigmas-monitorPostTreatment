package equation

import (
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"math"
	"testing"
)

func TestFlecheHZ_Compute(t *testing.T) {
	variables := []data.CaptorValue{
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
		{
			Captor: "v13",
			Field:  "",
		},
		{
			Captor: "v14",
			Field:  "",
		},
		{
			Captor: "v15",
			Field:  "",
		},
		{
			Captor: "v16",
			Field:  "",
		},
		{
			Captor: "v17",
			Field:  "",
		},
		{
			Captor: "v18",
			Field:  "",
		},
		{
			Captor: "v19",
			Field:  "",
		},
		{
			Captor: "v20",
			Field:  "",
		},
		{
			Captor: "v21",
			Field:  "",
		},
		{
			Captor: "v22",
			Field:  "",
		},
		{
			Captor: "v23",
			Field:  "",
		},
		{
			Captor: "v24",
			Field:  "",
		},
	}
	vals := map[string]float64{
		"v1.":  2538517.00449813,
		"v2.":  1182022.11430351,
		"v3.":  2538517.052,
		"v4.":  1182022.091,
		"v5.":  2538518.37302081,
		"v6.":  1182022.73832296,
		"v7.":  2538518.357,
		"v8.":  1182022.685,
		"v9.":  2538515.00855942,
		"v10.": 1182026.47941259,
		"v11.": 2538515.059,
		"v12.": 1182026.457,
		"v13.": 2538516.37587758,
		"v14.": 1182027.10606514,
		"v15.": 2538516.366,
		"v16.": 1182027.048,
		"v17.": 2538513.01130294,
		"v18.": 1182030.84508448,
		"v19.": 2538513.061,
		"v20.": 1182030.833,
		"v21.": 2538514.37919406,
		"v22.": 1182031.47048644,
		"v23.": 2538514.37,
		"v24.": 1182031.42,
	}
	expectedResult := 0.179900492366579
	flecheHZ := FlecheHZ{}
	result := flecheHZ.Compute(variables, vals)
	if math.Round(result*10)/10 != math.Round(expectedResult*10)/10 {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func FuzzFlecheHZ_Compute(f *testing.F) {
	variables := []data.CaptorValue{
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
		{
			Captor: "v13",
			Field:  "",
		},
		{
			Captor: "v14",
			Field:  "",
		},
		{
			Captor: "v15",
			Field:  "",
		},
		{
			Captor: "v16",
			Field:  "",
		},
		{
			Captor: "v17",
			Field:  "",
		},
		{
			Captor: "v18",
			Field:  "",
		},
		{
			Captor: "v19",
			Field:  "",
		},
		{
			Captor: "v20",
			Field:  "",
		},
		{
			Captor: "v21",
			Field:  "",
		},
		{
			Captor: "v22",
			Field:  "",
		},
		{
			Captor: "v23",
			Field:  "",
		},
		{
			Captor: "v24",
			Field:  "",
		},
	}
	vals := map[string]float64{
		"v1.":  0.0,
		"v2.":  0.0,
		"v3.":  0.0,
		"v4.":  0.0,
		"v5.":  0.0,
		"v6.":  0.0,
		"v7.":  0.0,
		"v8.":  0.0,
		"v9.":  0.0,
		"v10.": 0.0,
		"v11.": 0.0,
		"v12.": 0.0,
		"v13.": 0.0,
		"v14.": 0.0,
		"v15.": 0.0,
		"v16.": 0.0,
		"v17.": 0.0,
		"v18.": 0.0,
		"v19.": 0.0,
		"v20.": 0.0,
		"v21.": 0.0,
		"v22.": 0.0,
		"v23.": 0.0,
		"v24.": 0.0,
	}
	f.Add(0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0)
	f.Add(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0, 21.0, 22.0, 23.0, 24.0)
	f.Add(math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64)
	f.Fuzz(func(t *testing.T, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24 float64) {
		vals["v1."] = v1
		vals["v2."] = v2
		vals["v3."] = v3
		vals["v4."] = v4
		vals["v5."] = v5
		vals["v6."] = v6
		vals["v7."] = v7
		vals["v8."] = v8
		vals["v9."] = v9
		vals["v10."] = v10
		vals["v11."] = v11
		vals["v12."] = v12
		vals["v13."] = v13
		vals["v14."] = v14
		vals["v15."] = v15
		vals["v16."] = v16
		vals["v17."] = v17
		vals["v18."] = v18
		vals["v19."] = v19
		vals["v20."] = v20
		vals["v21."] = v21
		vals["v22."] = v22
		vals["v23."] = v23
		vals["v24."] = v24

		flecheHz := FlecheHZ{}
		flecheHz.Compute(variables, vals)
	})

}
