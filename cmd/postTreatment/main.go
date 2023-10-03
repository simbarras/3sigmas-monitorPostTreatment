package main

import (
	"github.com/getsentry/sentry-go"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg"
	"github.com/simbarras/3sigmas-monitorPostTreatment/pkg/equation"
	"log"
)

func main() {

	// To initialize Sentry's handler, you need to initialize Sentry itself beforehand
	if err := sentry.Init(sentry.ClientOptions{
		Dsn:           "https://90bf7224527f47d66c14459f656c7a6c@o4505048574001152.ingest.sentry.io/4505985728512000",
		EnableTracing: true,
		// Set TracesSampleRate to 1.0 to capture 100%
		// of transactions for performance monitoring.
		// We recommend adjusting this value in production,
		TracesSampleRate: 1.0,
		Release:          pkg.Version,
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

	log.Printf("App started in release %s\n", pkg.Version)
	add := equation.Addition{}
	log.Printf("10 + 30 = %d\n", add.Compute(10, 30))
}
