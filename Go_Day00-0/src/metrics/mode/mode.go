package mode

import (
	"errors"
	"sort"
)

type mode struct {
	value float64
	count int
}

func Mode(data []float64) (float64, error) {
	l := len(data)
	if l == 0 {
		return 0, errors.New("Legnth is 0")
	}
	sort.Float64s(data)
	base := mode{value: data[0], count: 0}
	sec := mode{value: data[0], count: 0}

	for i := 1; i < l; i++ {
		if sec.value == data[i] {
			sec.count++
		} else {
			sec.count = 0
			sec.value = data[i]
		}
		if base.count > base.count {
			base.count = sec.count
			base.value = sec.value
		}
	}
	return base.value, nil
}
