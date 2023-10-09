package api

import (
	"github.com/getsentry/sentry-go"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/api/storage"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/core"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/core/acquisition"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/core/equation"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/storer"
	"log"
	"strings"
)

type Worker struct {
	influxRead  *acquisition.Influx
	postgres    *storage.PostgresStore
	equations   []equation.Equation
	influxStore *storer.InfluxStorer
}

func NewWorker(influx *acquisition.Influx, postgres *storage.PostgresStore, equations []equation.Equation, store *storer.InfluxStorer) *Worker {
	return &Worker{
		influxRead:  influx,
		postgres:    postgres,
		equations:   equations,
		influxStore: store,
	}
}

func (w *Worker) GetBuckets() []string {
	return w.influxRead.GetBuckets()
}

func (w *Worker) GetActions() []data.Action {
	return w.postgres.GetAllActions()
}

func (w *Worker) GetFunctions() []string {
	res := make([]string, 0)
	for _, eq := range w.equations {
		res = append(res, eq.Name())
	}
	return res
}

func (w *Worker) AddAction(action data.Action) {
	w.postgres.AddAction(action)
}

func (w *Worker) UpdateAction(action data.Action) {
	w.postgres.UpdateAction(action)
}

func (w *Worker) DeleteAction(id string) {
	action := w.postgres.FindAction(id)
	w.postgres.DeleteAction(action)
}

func (w *Worker) TriggerAction(id string) error {
	action := w.postgres.FindAction(id)
	err := w.processAction(action)
	return err
}

func (w *Worker) TriggerBucket(name string) error {
	actions := w.postgres.FindActionsByBucket(name)
	for _, action := range actions {
		err := w.processAction(action)
		if err != nil {
			return err
		}
	}
	return nil
}

func (w *Worker) processAction(action data.Action) error {
	log.Printf("Processing action %s\n", action.ID)
	captors, variables, err := core.ParseVariables(action.ListVariables)
	if err != nil {
		return err
	}
	resultMap, t := w.influxRead.GetLastValue(action.BucketName, captors)
	results := equation.ComputeAll(variables, resultMap, core.GetEquation(w.equations, action.EquationName))
	measures := core.BuildMeasure(variables, results, t, action.EquationName)
	err = w.influxStore.Store(strings.Split(action.BucketName, ".")[1], "computed", measures)
	if err != nil {
		sentry.CaptureException(err)
		return err
	}
	return nil
}
