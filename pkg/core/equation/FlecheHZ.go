package equation

import (
	ownData "github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"math"
)

type FlecheHZ struct{}

func (FlecheHZ) Compute(variables []ownData.CaptorValue, values map[string]float64) float64 {
	var1 := values[variables[0].String()]
	var2 := values[variables[1].String()]
	var3 := values[variables[2].String()]
	var4 := values[variables[3].String()]
	var5 := values[variables[4].String()]
	var6 := values[variables[5].String()]
	var7 := values[variables[6].String()]
	var8 := values[variables[7].String()]
	var9 := values[variables[8].String()]
	var10 := values[variables[9].String()]
	var11 := values[variables[10].String()]
	var12 := values[variables[11].String()]
	var13 := values[variables[12].String()]
	var14 := values[variables[13].String()]
	var15 := values[variables[14].String()]
	var16 := values[variables[15].String()]
	var17 := values[variables[16].String()]
	var18 := values[variables[17].String()]
	var19 := values[variables[18].String()]
	var20 := values[variables[19].String()]
	var21 := values[variables[20].String()]
	var22 := values[variables[21].String()]
	var23 := values[variables[22].String()]
	var24 := values[variables[23].String()]

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

	phi13 := math.Atan((av19 - av3) / (av20 - av4))
	phi12 := math.Atan((av11 - av3) / (av12 - av4))
	dist12 := math.Sqrt((av10-av2)*(av10-av2) + (av9-av1)*(av9-av1))

	phi13Ref := math.Atan((av17 - av1) / (av18 - av2))
	phi12Ref := math.Atan((av9 - av1) / (av10 - av2))

	return (((phi13 - phi12) * dist12) - ((phi13Ref - phi12Ref) * dist12)) * 1000.0
}

func (FlecheHZ) Name() string {
	return "FlecheHz"
}
