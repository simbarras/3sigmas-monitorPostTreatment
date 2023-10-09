package equation

import (
	"math"
	"testing"
)

func TestFlecheV_Compute(t *testing.T) {
	variables := []string{"v1", "v2", "v3", "v4", "v5", "v6", "v7", "v8", "v9", "v10", "v11", "v12", "v13", "v14", "v15", "v16", "v17", "v18", "v19", "v20", "v21", "v22", "v23", "v24"}
	vals := map[string]float64{
		"v1":  2538517.00449813,
		"v2":  1182022.11430351,
		"v3":  434.827,
		"v4":  434.827741537086,
		"v5":  2538518.37302081,
		"v6":  1182022.73832296,
		"v7":  434.82,
		"v8":  434.826492746764,
		"v9":  2538515.00855942,
		"v10": 1182026.47941259,
		"v11": 434.81,
		"v12": 434.813964758716,
		"v13": 2538516.37587758,
		"v14": 1182027.10606514,
		"v15": 434.81,
		"v16": 434.81523624641,
		"v17": 2538513.01130294,
		"v18": 1182030.84508448,
		"v19": 434.8,
		"v20": 434.803792618789,
		"v21": 2538514.37919406,
		"v22": 1182031.47048644,
		"v23": 434.8,
		"v24": 434.803754732423,
	}
	expectedResult := 0.000905085028932717
	flecheV := FlecheV{}
	result := flecheV.Compute(variables, vals)
	if math.Round(result*10)/10 != math.Round(expectedResult*10)/10 {
		t.Errorf("Expected %f, got %f", expectedResult, result)
	}
}

func FuzzFlecheV_Compute(f *testing.F) {
	variables := []string{"v1", "v2", "v3", "v4", "v5", "v6", "v7", "v8", "v9", "v10", "v11", "v12", "v13", "v14", "v15", "v16", "v17", "v18", "v19", "v20", "v21", "v22", "v23", "v24"}
	vals := map[string]float64{
		"v1":  0.0,
		"v2":  0.0,
		"v3":  0.0,
		"v4":  0.0,
		"v5":  0.0,
		"v6":  0.0,
		"v7":  0.0,
		"v8":  0.0,
		"v9":  0.0,
		"v10": 0.0,
		"v11": 0.0,
		"v12": 0.0,
		"v13": 0.0,
		"v14": 0.0,
		"v15": 0.0,
		"v16": 0.0,
		"v17": 0.0,
		"v18": 0.0,
		"v19": 0.0,
		"v20": 0.0,
		"v21": 0.0,
		"v22": 0.0,
		"v23": 0.0,
		"v24": 0.0,
	}
	f.Add(0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0, 0.0)
	f.Add(1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0, 12.0, 13.0, 14.0, 15.0, 16.0, 17.0, 18.0, 19.0, 20.0, 21.0, 22.0, 23.0, 24.0)
	f.Add(math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64, math.MaxFloat64)
	f.Fuzz(func(t *testing.T, v1, v2, v3, v4, v5, v6, v7, v8, v9, v10, v11, v12, v13, v14, v15, v16, v17, v18, v19, v20, v21, v22, v23, v24 float64) {
		vals["v1"] = v1
		vals["v2"] = v2
		vals["v3"] = v3
		vals["v4"] = v4
		vals["v5"] = v5
		vals["v6"] = v6
		vals["v7"] = v7
		vals["v8"] = v8
		vals["v9"] = v9
		vals["v10"] = v10
		vals["v11"] = v11
		vals["v12"] = v12
		vals["v13"] = v13
		vals["v14"] = v14
		vals["v15"] = v15
		vals["v16"] = v16
		vals["v17"] = v17
		vals["v18"] = v18
		vals["v19"] = v19
		vals["v20"] = v20
		vals["v21"] = v21
		vals["v22"] = v22
		vals["v23"] = v23
		vals["v24"] = v24

		flecheV := FlecheV{}
		flecheV.Compute(variables, vals)
	})

}
