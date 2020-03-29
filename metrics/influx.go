package metrics

import (
	"fmt"

	influx "github.com/influxdata/influxdb/client/v2"
)

func createInfluxClient() (influx.Client, error) {
	c, err := influx.NewHTTPClient(influx.HTTPConfig{
		Addr: "http://influxdb-service:8086",
	})
	if err != nil {
		return nil, fmt.Errorf("couldn't create influx client - %v", err)
	}
	defer c.Close()

	return c, nil
}

func createInfluxBatchPoints(bp *batchPointsStruct) (influx.BatchPoints, error) {
	influxBP, err := createBatchPointConfig(bp.Client)
	if err != nil {
		return nil, fmt.Errorf("couldn't create batch point config - %v", err)
	}

	for _, point := range bp.Points {
		influxBP.AddPoint(point)
	}

	return influxBP, nil
}

// Creates a new Batch Point Config
func createBatchPointConfig(client influx.Client) (influx.BatchPoints, error) {
	bp, err := influx.NewBatchPoints(influx.BatchPointsConfig{
		Database:  myDB,
		Precision: "s",
	})
	if err != nil {
		return nil, fmt.Errorf("couldn't create Influx batch point - %v", err)
	}

	return bp, nil
}
