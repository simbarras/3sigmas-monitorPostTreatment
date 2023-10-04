package main

import (
	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-gonic/gin"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/api"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/api/storage"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/core"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/data"
	"log"
)

const Version = "0.0.1"
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

	log.Printf("App started in release %s\n", Version)
	environment := data.ReadEnv()

	store := storage.NewPostgres()

	// Set up routes
	api.SetRoutes(app, ApiPrefix, store)

	// And run it
	err := app.Run("localhost:3000")
	if err != nil {
		return
	}

	core.DummyMain(environment)

}
