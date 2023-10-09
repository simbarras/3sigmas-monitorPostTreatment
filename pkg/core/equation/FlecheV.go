package equation

import "math"

type FlecheV struct{}

func average(a, b float64) float64 {
	return a/2.0 + b/2.0
}

func coefficient(av1, av9, av2, av10, av17, av18 float64) float64 {
	up := math.Sqrt((av1-av9)*(av1-av9) + (av2-av10)*(av2-av10))
	down := math.Sqrt((av1-av17)*(av1-av17) + (av2-av18)*(av2-av18))
	return up / down
}

func part(a, b, c, coef float64) float64 {
	return a - b + (c-b)*coef
}

func (FlecheV) Compute(variables []string, values map[string]float64) float64 {
	var1 := values[variables[0]]
	var2 := values[variables[1]]
	var3 := values[variables[2]]
	var4 := values[variables[3]]
	var5 := values[variables[4]]
	var6 := values[variables[5]]
	var7 := values[variables[6]]
	var8 := values[variables[7]]
	var9 := values[variables[8]]
	var10 := values[variables[9]]
	var11 := values[variables[10]]
	var12 := values[variables[11]]
	var13 := values[variables[12]]
	var14 := values[variables[13]]
	var15 := values[variables[14]]
	var16 := values[variables[15]]
	var17 := values[variables[16]]
	var18 := values[variables[17]]
	var19 := values[variables[18]]
	var20 := values[variables[19]]
	var21 := values[variables[20]]
	var22 := values[variables[21]]
	var23 := values[variables[22]]
	var24 := values[variables[23]]

	av1 := average(var1, var5)
	av2 := average(var2, var6)
	av3 := average(var3, var7)
	av4 := average(var4, var8)
	av9 := average(var9, var13)
	av10 := average(var10, var14)
	av11 := average(var11, var15)
	av12 := average(var12, var16)
	av17 := average(var17, var21)
	av18 := average(var18, var22)
	av19 := average(var19, var23)
	av20 := average(var20, var24)

	coef := coefficient(av1, av9, av2, av10, av17, av18)
	pLeft := part(av12, av4, av20, coef)
	pRight := part(av11, av3, av19, coef)
	return pLeft - pRight
}

func (FlecheV) Name() string {
	return "FlecheV"
}
