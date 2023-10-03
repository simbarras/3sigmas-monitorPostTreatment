package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/acquisition"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/env"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/equation"
	"log"
)

const Version = "0.0.1"

func main() {

	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           "https://90bf7224527f47d66c14459f656c7a6c@o4505048574001152.ingest.sentry.io/4505985728512000",
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
		Release:          Version,
	}); err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}

	//// Create an instance of sentryfasthttp
	//sentryHandler := sentryfasthttp.New(sentryfasthttp.Options{})
	//
	//// After creating the instance, you can attach the handler as one of your middleware
	//fastHTTPHandler := sentryHandler.Handle(func(ctx *fasthttp.RequestCtx) {
	//	panic("y tho")
	//})
	//
	//fmt.Println("Listening and serving HTTP on :3000")
	//
	//// And run it
	//if err := fasthttp.ListenAndServe(":3000", fastHTTPHandler); err != nil {
	//	panic(err)
	//}

	log.Printf("App started in release %s\n", Version)
	environment := env.Read()

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

}
