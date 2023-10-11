package acquisition

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	ownData "github.com/simbarras/3sigmas-monitorPostTreatment/pkg/data"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/data"
	"log"
	"time"
)

type Influx struct {
	client       influxdb2.Client
	organization *domain.Organization
}

func NewInflux(env data.Env) *Influx {
	client := influxdb2.NewClient(env.InfluxUrl, env.InfluxToken)
	org, err := client.OrganizationsAPI().FindOrganizationByName(context.Background(), env.InfluxOrg)
	if err != nil {
		sentry.CaptureException(err)
		return nil
	}
	return &Influx{
		client:       client,
		organization: org,
	}
}

func (i *Influx) GetBuckets() []string {
	buckets, err := i.client.BucketsAPI().GetBuckets(context.Background())
	if err != nil {
		sentry.CaptureException(err)
		log.Fatal(err)
	}
	result := make([]string, len(*buckets))
	for i, bucket := range *buckets {
		result[i] = bucket.Name
	}
	return result
}

func (i *Influx) GetLastValue(bucketName string, captor ownData.CaptorValue) (float64, time.Time) {
	bucket, err := i.client.BucketsAPI().FindBucketByName(context.Background(), bucketName)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatal(err)
	}
	// Create a query API client
	queryAPI := i.client.QueryAPI(i.organization.Name)

	// Define the InfluxQL query to retrieve the last value from the measurement
	query := fmt.Sprintf("from(bucket:\"%s\") |> range(start: 0) |> filter(fn: (r) => r[\"_measurement\"] == \"%s\") |> filter(fn: (r) => r[\"_field\"] == \"%s\") |> last()", bucket.Name, captor.Captor, captor.Field)

	// Execute the query
	result, err := queryAPI.Query(context.Background(), query)
	if err != nil {
		sentry.CaptureException(err)
		log.Fatal(err)
	}

	// Process the query result
	var value float64
	var measureTime time.Time
	for result.Next() {
		value = result.Record().Value().(float64)
		measureTime = result.Record().Time()
	}

	// Check for errors from iterating over the result
	if result.Err() != nil {
		sentry.CaptureException(err)
		log.Fatal(result.Err())
	}

	return value, measureTime
}
