package main

import ("testing" "mean")

func TestMean(t *testing.T) {
	// Arrange
	
	var data = []float64{1, 2, 3, 4, 5, 6, 7, 8, 9}
	expected := 5

	// Act
	num,mean.Mean(data)
}
