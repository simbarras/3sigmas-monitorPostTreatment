package main

import (
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/api"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/api/storage"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/core/acquisition"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/core/equation"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/storer"
	"log"
)

const Version = "0.0.2"
const ApiPrefix = "/api/v0"

func main() {

	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           "https://abad48b1066aff88bd3cdc1cef48a17b@o4505048574001152.ingest.sentry.io/4505987080388608",
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
		Release:          Version,
	}); err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	// Then create your app
	app := gin.Default()
	// Once it's done, you can attach the handler as one of your middleware
	app.Use(sentrygin.New(sentrygin.Options{}))
	// Set cors policy
	app.Use(api.CustomMiddleware())

	log.Printf("App started in release %s\n", Version)
	environment := data.ReadEnv()

	store := storage.NewPostgres(environment)
	influxStore := storer.NewInfluxStorer(environment.ExternEnv)
	equations := []equation.Equation{equation.FlecheV{}}
	worker := api.NewWorker(acquisition.NewInflux(environment.ExternEnv), store, equations, influxStore)

	// Set up routes
	api.SetRoutes(app, ApiPrefix, worker)

	// And run it
	err := app.Run("0.0.0.0:3001")
	if err != nil {
		return
	}

}
