package core

import (
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/core/equation"
	ownData "github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/data"
	"strings"
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

func isIn(is string, in []string) bool {
	for _, s := range in {
		if is == s {
			return true
		}
	}
	return false
}

func ParseVariables(brute string) ([]string, [][]string) {
	brute = strings.ReplaceAll(brute, "\n", "")
	brute = strings.ReplaceAll(brute, "\r", "")
	brute = strings.ReplaceAll(brute, " ", "")
	captors := make([]string, 0)
	variables := make([][]string, 0)
	for i, vs := range strings.Split(brute, ";") {
		split := strings.Split(vs, ",")
		if split[0] == "" {
			break
		}
		variables = append(variables, make([]string, len(split)))
		for j, v := range split {
			variables[i][j] = v
			if !isIn(v, captors) {
				captors = append(captors, v)
			}
		}
	}
	return captors, variables
}

func GetEquation(equations []equation.Equation, name string) equation.Equation {
	for _, eq := range equations {
		if eq.Name() == name {
			return eq
		}
	}
	return nil
}
