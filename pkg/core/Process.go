package core

import (
	"errors"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/core/equation"
	ownData "github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/data"
	"strings"
	"time"
)

func BuildMeasure(listVariables []ownData.EquationCaptor, result []float64, dateTime time.Time, function string) []data.Measure {
	measures := make([]data.Measure, 0)
	for i, variables := range listVariables {
		measures = append(measures, ownData.ComputedMeasure{
			DateTime:  dateTime,
			Value:     result[i],
			Captor:    function,
			Variables: variables.Name,
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

func ParseVariables(brute string) ([]ownData.CaptorValue, []ownData.EquationCaptor, error) {
	brute = strings.ReplaceAll(brute, "\n", "")
	brute = strings.ReplaceAll(brute, "\r", "")
	brute = strings.ReplaceAll(brute, " ", "")
	brute = strings.ReplaceAll(brute, "\t", "")
	captors := make([]ownData.CaptorValue, 0)
	variables := make([]ownData.EquationCaptor, 0)
	for i, line := range strings.Split(brute, ";") {
		if line == "" {
			continue
		}
		split := strings.Split(line, ":")
		if len(split) != 2 {
			return nil, nil, errors.New("name not provided with ':'")
		}
		name := split[0]
		vs := split[1]
		split = strings.Split(vs, ",")
		if split[0] == "" {
			break
		}
		variables = append(variables, ownData.EquationCaptor{
			Name:    name,
			Captors: make([]ownData.CaptorValue, len(split)),
		})
		for j, v := range split {
			err := variables[i].GetCaptor(j).FromString(v)
			if err != nil {
				return nil, nil, errors.New("invalid captor value")
			}
			if !isIn(v, captors) {
				captors = append(captors, *variables[i].GetCaptor(j))
			}
		}
	}
	return captors, variables, nil
}

func GetEquation(equations []equation.Equation, name string) equation.Equation {
	for _, eq := range equations {
		if eq.Name() == name {
			return eq
		}
	}
	return nil
}
