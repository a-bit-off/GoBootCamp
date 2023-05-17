package median

import (
	"errors"
	"sort"
)

func Median(data []float64) (float64, error) {
	l := len(data)
	if l == 0 {
		return 0, errors.New("Legnth is 0")
	}
	sort.Float64s(data)
	if l%2 == 0 {
		return (data[l/2] + data[l/2-1]) / 2, nil
	}
	return data[l/2], nil
}
