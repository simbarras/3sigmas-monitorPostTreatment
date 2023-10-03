package equation

type Addition struct{}

func (Addition) Compute(a int, b int) int {
	return a + b
}
