package db

import (
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/wonyus/update-status-service/contexts"
)

type Influx struct {
	Org    string
	Bucket string
}

type InfluxApiClient struct {
	influxdb2.Client
}

func NewInflux(ctx contexts.Resource) *InfluxApiClient {
	client := influxdb2.NewClient(ctx.InfluxUrl, ctx.InfluxToken)
	return &InfluxApiClient{client}
}
