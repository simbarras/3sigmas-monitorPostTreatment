package core

import (
	ownData "github.com/simbarras/3sigmas-monitorPostTreatment/pkg/core/data"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/data"
	"time"
)

func BuildMeasure(listVariables [][]string, result []float64, dateTime time.Time, function string) []data.Measure {
	measures := make([]data.Measure, 0)
	for i, variables := range listVariables {
		measures = append(measures, ownData.ComputedMeasure{
			DateTime:  dateTime,
			Value:     result[i],
			Captor:    function,
			Variables: variables,
		})
	}
	return measures
}
