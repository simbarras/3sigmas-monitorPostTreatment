package equation

import ownData "github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"

type Equation interface {
	Compute(variables []ownData.CaptorValue, values map[string]float64) float64
	Name() string
}

func ComputeAll(listVariables [][]ownData.CaptorValue, values map[string]float64, eq Equation) []float64 {
	result := make([]float64, len(listVariables))
	for i, variables := range listVariables {
		result[i] = eq.Compute(variables, values)
	}
	return result
}
