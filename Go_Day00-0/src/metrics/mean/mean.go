package mean

import (
	"errors"
)

func Mean(data []float64) (float64, error) {
	l := len(data)
	if l == 0 {
		return 0, errors.New("Legnth is 0")
	}
	var sum float64
	for i := 0; i < l; i++ {
		sum += data[i]
	}
	return sum / float64(l), nil
}
