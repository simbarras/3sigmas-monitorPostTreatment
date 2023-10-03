package equation

type Equation interface {
	Compute(variables []string, values map[string]float64) float64
}

func ComputeAll(listVariables [][]string, values map[string]float64, eq Equation) []float64 {
	result := make([]float64, len(listVariables))
	for i, variables := range listVariables {
		result[i] = eq.Compute(variables, values)
	}
	return result
}
