package metrics

import (
	"fmt"
	"log"
	"time"

	influx "github.com/influxdata/influxdb/client/v2"
)

const (
	// Influx DB to write metrics to
	myDB = "CardPrices"

	// Max size of batch metrics before they send
	maxQueueSize = 100

	// How frequently to flush the queue
	queueFlushTimer = 60 * time.Second
)

type batchPointsStruct struct {
	Client influx.Client
	Points []*influx.Point
}

func createBatchPointsStruct(client influx.Client) (*batchPointsStruct, error) {
	pointsArray := []*influx.Point{}
	bp := batchPointsStruct{
		Client: client,
		Points: pointsArray,
	}

	bp.addFlushTimer()

	return &bp, nil
}

func (bp *batchPointsStruct) addFlushTimer() {
	go func() {
		for {
			time.Sleep(queueFlushTimer)

			log.Println("INFO - Sending points due to timer.")
			bp.sendPoints()
		}
	}()
}

func (bp *batchPointsStruct) addPoint(pt influx.Point) error {
	bp.Points = append(bp.Points, &pt)

	if len(bp.Points) >= maxQueueSize {
		log.Println("INFO - Sending points due to maxQueueSize.")
		err := bp.sendPoints()
		if err != nil {
			return fmt.Errorf("couldn't send points - %v", err)
		}
	}

	return nil
}

func (bp *batchPointsStruct) createInfluxBatchPoints() (influx.BatchPoints, error) {
	influxBP, err := createBatchPointConfig(bp.Client)
	if err != nil {
		return nil, fmt.Errorf("couldn't create batch point config - %v", err)
	}

	for _, point := range bp.Points {
		influxBP.AddPoint(point)
	}

	return influxBP, nil
}

func (bp *batchPointsStruct) removePoints() {
	bp.Points = bp.Points[:0]
}

func (bp *batchPointsStruct) sendPoints() error {
	pointCount := len(bp.Points)

	if pointCount > 0 {
		influxBP, err := bp.createInfluxBatchPoints()
		if err != nil {
			return fmt.Errorf("couldn't create influxdb data points - %v", err)
		}

		if err := bp.Client.Write(influxBP); err != nil {
			return fmt.Errorf("couldn't send points using the client - %v", err)
		}

		log.Printf("INFO - Sent %v batched metrics.\n", pointCount)

		bp.removePoints()
	}

	return nil
}
