package equation

type Inclinaison struct {
}

func (Inclinaison) Compute(variables []string, values map[string]float64) float64 {
	return values[variables[0]] + values[variables[1]]
}

func (Inclinaison) Name() string {
	return "Inclinaison"
}
