package equation

type Equation interface {
	Compute(a float64, b float64) float64
}
