package acquisition

import (
	"context"
	"fmt"
	"github.com/getsentry/sentry-go"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"github.com/simbarras/3sigmas-monitorVisualization/pkg/data"
	"log"
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

func (i *Influx) GetLastValue(bucketName string, captors []string) map[string]float64 {
	resultMap := make(map[string]float64)
	for _, captor := range captors {
		bucket, err := i.client.BucketsAPI().FindBucketByName(context.Background(), bucketName)
		if err != nil {
			sentry.CaptureException(err)
			log.Fatal(err)
		}
		// Create a query API client
		queryAPI := i.client.QueryAPI(i.organization.Name)

		// Define the InfluxQL query to retrieve the last value from the measurement
		query := fmt.Sprintf("from(bucket:\"%s\") |> range(start: 0) |> filter(fn: (r) => r[\"_measurement\"] == \"%s\") |> filter(fn: (r) => r[\"_field\"] == \"%s\") |> last()", bucket.Name, captor, "northing")

		// Execute the query
		result, err := queryAPI.Query(context.Background(), query)
		if err != nil {
			sentry.CaptureException(err)
			log.Fatal(err)
		}

		// Process the query result
		for result.Next() {
			resultMap[captor] = result.Record().Value().(float64)
		}

		// Check for errors from iterating over the result
		if result.Err() != nil {
			sentry.CaptureException(err)
			log.Fatal(result.Err())
		}
	}
	return resultMap
}
