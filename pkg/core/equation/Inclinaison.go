package equation

import (
	ownData "github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"math"
)

type Inclinaison struct{}

func (Inclinaison) Compute(variables []ownData.CaptorValue, values map[string]float64) float64 {
	const1 := values[variables[0].String()]
	var1 := values[variables[1].String()]
	var2 := values[variables[2].String()]
	var3 := values[variables[3].String()]
	var4 := values[variables[4].String()]
	var5 := values[variables[5].String()]
	var6 := values[variables[6].String()]
	var7 := values[variables[7].String()]
	var8 := values[variables[8].String()]
	var9 := values[variables[9].String()]
	var10 := values[variables[10].String()]

	deb := var3 - var1
	dnb := var4 - var2
	dtransb := deb*math.Sin(const1*math.Pi/200) + dnb*math.Cos(const1*math.Pi/200)

	deh := var8 - var6
	dnh := var9 - var7
	dtransh := deh*math.Sin(const1*math.Pi/200) + dnh*math.Cos(const1*math.Pi/200)

	dhbh := var10 - var5

	return (dtransh - dtransb) * 1000 / dhbh
}

func (Inclinaison) Name() string {
	return "Inclinaison"
}
