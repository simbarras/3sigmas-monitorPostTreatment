package equation

type Addition struct{}

func (Addition) Compute(variables []string, values map[string]float64) float64 {
	return values[variables[0]] + values[variables[1]]
}

func (Addition) Name() string {
	return "Addition"
}
