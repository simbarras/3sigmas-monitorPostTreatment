package core

import (
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/core/equation"
	ownData "github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/data"
	"strings"
	"time"
)

func BuildMeasure(listVariables [][]ownData.CaptorValue, result []float64, dateTime time.Time, function string) []data.Measure {
	measures := make([]data.Measure, 0)
	for i, variables := range listVariables {
		variablesString := make([]string, len(variables))
		for j, v := range variables {
			variablesString[j] = v.String()

		}
		measures = append(measures, ownData.ComputedMeasure{
			DateTime:  dateTime,
			Value:     result[i],
			Captor:    function,
			Variables: variablesString,
		})
	}
	return measures
}

func isIn(is string, in []ownData.CaptorValue) bool {
	for _, s := range in {
		if is == s.String() {
			return true
		}
	}
	return false
}

func ParseVariables(brute string) ([]ownData.CaptorValue, [][]ownData.CaptorValue) {
	brute = strings.ReplaceAll(brute, "\n", "")
	brute = strings.ReplaceAll(brute, "\r", "")
	brute = strings.ReplaceAll(brute, " ", "")
	captors := make([]ownData.CaptorValue, 0)
	variables := make([][]ownData.CaptorValue, 0)
	for i, vs := range strings.Split(brute, ";") {
		split := strings.Split(vs, ",")
		if split[0] == "" {
			break
		}
		variables = append(variables, make([]ownData.CaptorValue, len(split)))
		for j, v := range split {
			err := variables[i][j].FromString(v)
			if err != nil {
				return nil, nil
			}
			if !isIn(v, captors) {
				captors = append(captors, variables[i][j])
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
