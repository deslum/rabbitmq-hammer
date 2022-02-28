package statistic

import (
	"fmt"
	"github.com/montanaflynn/stats"

	"time"
)

type Data struct {
	Min         float64
	Max         float64
	Mean        float64
	Median      float64
	Percentiles []string
	RPS         float64
	Time        time.Duration
}

func CalcData(data []float64) *Data {

	var percentiles = []float64{50.0, 66.0, 75.0, 80.0, 90.0, 95.0, 98.0, 99.0, 100.0}

	min, _ := stats.Min(data)
	max, _ := stats.Max(data)
	mean, _ := stats.Mean(data)
	median, _ := stats.Median(data)

	calcData := &Data{
		Min:    min,
		Max:    max,
		Median: median,
		Mean:   mean,
	}

	for _, percentile := range percentiles {
		result, _ := stats.Percentile(data, percentile)

		res := fmt.Sprintf("%3v%% %7.1f μs", percentile, result)
		calcData.Percentiles = append(calcData.Percentiles, res)

	}

	if mean > 0 {
		calcData.RPS = float64(time.Second.Microseconds()) / mean
	}

	allTime, _ := stats.Sum(data)

	calcData.Time = time.Duration(allTime) * time.Microsecond

	return calcData
}
