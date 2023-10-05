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
