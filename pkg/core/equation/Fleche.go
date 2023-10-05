package equation

type Fleche struct{}

func (Fleche) Compute(variables []string, values map[string]float64) float64 {
	return values[variables[0]] - values[variables[1]]
}

func (Fleche) Name() string {
	return "Fleche"
}
