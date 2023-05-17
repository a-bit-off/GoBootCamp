package parser

import (
	"bufio"
	"errors"
	"io"
	"math"
	"strconv"
)

const (
	MAX = 100000
	MIN = -100000
)

func Parser(reader io.Reader) ([]float64, error) {
	data := make([]float64, 0)
	scanner := bufio.NewScanner(reader)

	for scanner.Scan() {
		text, err := scanner.Text(), scanner.Err()
		if err != nil {
			return nil, err
		} else if err == io.EOF {
			if len(data) == 0 {
				return nil, errors.New("String is empty")
			}
			break
		}
		var num float64
		if num, err = strconv.ParseFloat(text, 64); err != nil {
			return nil, err
		}
		if num > MAX || num < MIN {
			return nil, errors.New("Out of range")
		}
		data = append(data, 0.01*math.Round(num*100))
	}
	return data, nil
}
