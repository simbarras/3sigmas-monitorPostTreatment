package equation

import (
	ownData "github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"math"
)

type Gauche struct{}

func subPart(a, b, c, d, e, f, g float64) float64 {
	return (a - b) * c / math.Sqrt((d-e)*(d-e)+(f-g)*(f-g))
}

func (Gauche) Compute(variables []ownData.CaptorValue, values map[string]float64) float64 {
	const1 := values[variables[0].String()]
	const2 := values[variables[1].String()]
	var1 := values[variables[2].String()]
	var2 := values[variables[3].String()]
	var3 := values[variables[4].String()]
	var4 := values[variables[5].String()]
	var5 := values[variables[6].String()]
	var6 := values[variables[7].String()]
	var7 := values[variables[8].String()]
	var8 := values[variables[9].String()]
	var9 := values[variables[10].String()]
	var10 := values[variables[11].String()]
	var11 := values[variables[12].String()]
	var12 := values[variables[13].String()]

	av1 := average(var1, var4)
	av2 := average(var2, var5)
	av7 := average(var7, var10)
	av8 := average(var8, var11)

	partLeft := subPart(var3, var6, const2, var1, var4, var2, var5)
	partRight := subPart(var9, var12, const2, var7, var10, var8, var11)
	sum1 := av1 - av7
	sum2 := av2 - av8
	lp1234 := math.Sqrt(sum1*sum1 + sum2*sum2)

	return const1 + (partLeft-partRight)*1000/lp1234
}

func (Gauche) Name() string {
	return "Gauche"
}
