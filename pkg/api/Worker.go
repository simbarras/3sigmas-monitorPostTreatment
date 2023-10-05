package api

import (
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/api/storage"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/core/acquisition"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/core/equation"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
)

type Worker struct {
	influx    *acquisition.Influx
	postgres  *storage.PostgresStore
	equations []equation.Equation
}

func NewWorker(influx *acquisition.Influx, postgres *storage.PostgresStore, equations []equation.Equation) *Worker {
	return &Worker{
		influx:    influx,
		postgres:  postgres,
		equations: equations,
	}
}

func (w *Worker) GetBuckets() []string {
	return w.influx.GetBuckets()
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
