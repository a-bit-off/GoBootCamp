package parser

import (
	"bufio"
	"errors"
	"io"
	"strconv"
	"strings"
)

const (
	MAX = 100000.00
	MIN = -100000.00
)

func ParseOrder(reader io.Reader) ([]int, error) {
	in := bufio.NewScanner(reader)
	in.Scan()
	if err := in.Err(); err != nil {
		return nil, err
	}
	order := make([]int, 0)
	strSplit := strings.Split(in.Text(), " ")
	for _, value := range strSplit {
		if num, err := strconv.Atoi(value); err == nil {
			order = append(order, num)
		}
	}
	return order, nil
}

func ParserData(reader io.Reader) ([]float64, error) {
	data := make([]float64, 0)
	in := bufio.NewScanner(reader)

	for in.Scan() {
		text, err := in.Text(), in.Err()
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
		data = append(data, num)
	}
	return data, nil
}
