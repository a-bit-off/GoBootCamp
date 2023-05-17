package sd

import (
	"errors"
	"math"
	"mean"
)

func SD(data []float64) (float64, error) {
	l := len(data)
	if l == 0 {
		return 0, errors.New("Legnth is 0")
	}

	var variance float64
	var meanValue float64
	meanValue, err := mean.Mean(data)
	if err != nil {
		return 0, err
	}

	for i := 0; i < l; i++ {
		variance += math.Pow(data[i]-meanValue, 2)
	}
	return math.Sqrt(variance / float64(l)), nil
}
