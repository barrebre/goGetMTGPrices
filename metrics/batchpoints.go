package metrics

import (
	"fmt"

	influx "github.com/influxdata/influxdb/client/v2"
)

const (
	myDB         = "CardPrices"
	maxQueueSize = 6
)

type batchPointsStruct struct {
	Client influx.Client
	Points []*influx.Point
}

func createBatchPointsStruct(client influx.Client) (*batchPointsStruct, error) {
	pointsArray := []*influx.Point{}
	bps := batchPointsStruct{
		Client: client,
		Points: pointsArray,
	}

	return &bps, nil
}

func (bp *batchPointsStruct) addPoint(pt influx.Point) error {
	bp.Points = append(bp.Points, &pt)

	if len(bp.Points) >= maxQueueSize {
		err := bp.sendPoints()
		if err != nil {
			return fmt.Errorf("couldn't send points - %v", err)
		}
		bp.removePoints()
	}

	return nil
}

func (bp *batchPointsStruct) removePoints() {
	bp.Points = bp.Points[:0]
}

func (bp batchPointsStruct) sendPoints() error {
	influxBP, err := createInfluxBatchPoints(bp)
	if err != nil {
		return fmt.Errorf("couldn't create influxdb data points - %v", err)
	}

	if err := bp.Client.Write(influxBP); err != nil {
		return fmt.Errorf("couldn't send points using the client - %v", err)
	}

	return nil
}
