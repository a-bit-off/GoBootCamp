package main

import (
	"github.com/stretchr/testify/assert"
	"mean"
	"median"
	"mode"
	"sd"
	"testing"
)

func TestMean(t *testing.T) {
	// Arrange
	testTable := []struct {
		data     []float64
		expected float64
	}{
		{
			data:     []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 5.0,
		},
		{
			data:     []float64{3, 2, 5, 6, 3, 4, 5, 9, 6, 5},
			expected: 4.8,
		},
		{
			data:     []float64{3, 2, 1},
			expected: 2.0,
		},
		{
			data:     []float64{1, 1},
			expected: 1,
		},
		{
			data:     []float64{},
			expected: 0,
		},
	}

	// Act
	for _, testCase := range testTable {
		if result, err := mean.Mean(testCase.data); err == nil {
			// Assert
			assert.Equal(t, testCase.expected, result)
		}

	}
}

func TestMedian(t *testing.T) {
	// Arrange
	testTable := []struct {
		data     []float64
		expected float64
	}{
		{
			data:     []float64{1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: 5.0,
		},
		{
			data:     []float64{3, 2, 1, 1, 1, 1, 9, 6, 5},
			expected: 2.0,
		},
		{
			data:     []float64{3, 2, 1, 1, 1, 1, 1},
			expected: 1.0,
		},
		{
			data:     []float64{1, 1},
			expected: 1.0,
		},
		{
			data:     []float64{},
			expected: 0,
		},
	}

	// Act
	for _, testCase := range testTable {
		if result, err := median.Median(testCase.data); err == nil {
			// Assert
			assert.Equal(t, testCase.expected, result)
		}

	}
}

func TestMode(t *testing.T) {
	// Arrange
	testTable := []struct {
		data     []float64
		expected float64
	}{
		{
			data:     []float64{1, 2, 3, 4, 5, 6, 7, 7, 8, 8, 9},
			expected: 7.0,
		},
		{
			data:     []float64{3, 2, 1, 1, 1, 1, 9, 6, 5},
			expected: 1.0,
		},
		{
			data:     []float64{3, 2, 1, 1, 1, 1, 1},
			expected: 1.0,
		},
		{
			data:     []float64{1, 1},
			expected: 1.0,
		},
		{
			data:     []float64{},
			expected: 0,
		},
	}

	// Act
	for _, testCase := range testTable {
		if result, err := mode.Mode(testCase.data); err == nil {
			// Assert
			assert.Equal(t, testCase.expected, result)
		}

	}
}

func TestSD(t *testing.T) {
	// Arrange
	testTable := []struct {
		data     []float64
		expected float64
	}{
		{
			data:     []float64{1, 2, 3, 4, 5, 6, 7, 7, 8, 8, 9},
			expected: 2.5356955783602,
		},
		{
			data:     []float64{3, 2, 1, 1, 1, 1, 9, 6, 5},
			expected: 2.6988795114425,
		},
		{
			data:     []float64{3, 2, 1, 1, 1, 1, 1},
			expected: 0.72843135908468,
		},
		{
			data:     []float64{1, 1},
			expected: 0,
		},
		{
			data:     []float64{},
			expected: 0,
		},
	}

	// Act
	for _, testCase := range testTable {
		if result, err := sd.SD(testCase.data); err == nil {
			// Assert
			assert.InDelta(t, testCase.expected, result, 0.01)
		}

	}
}
