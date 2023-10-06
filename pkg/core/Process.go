package core

import (
	"github.com/getsentry/sentry-go"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/core/acquisition"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/core/equation"
	ownData "github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/data"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/storer"
	"log"
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

func DummyMain(environment data.Env) {
	captors := []string{"KM_000_D", "KM_000_G", "KM_035_D"}
	dataReader := acquisition.NewInflux(environment)
	resultMap := dataReader.GetLastValue("production.3s_230913.trimble", captors)
	log.Printf("Result: %v\n", resultMap)

	variables := [][]string{
		{"KM_000_D", "KM_000_G"},
		{"KM_000_G", "KM_035_D"},
		{"KM_035_D", "KM_035_G"},
	}
	results := equation.ComputeAll(variables, resultMap, equation.Addition{})
	log.Printf("Results: %v\n", results)

	measures := BuildMeasure(variables, results, time.Now(), "Addition")
	log.Printf("Measures: %v\n", measures)

	influxStorer := storer.NewInfluxStorer(environment)
	err := influxStorer.Store("3s_230913", "trimble.computed", measures)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatalf("Error storing measures: %s\n", err.Error())
	}
}
