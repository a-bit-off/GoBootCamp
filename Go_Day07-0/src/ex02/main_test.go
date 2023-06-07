package main

import (
	"reflect"
	"testing"
)

var tests = []struct {
	coins    []int
	value    int
	expected []int
}{
	{coins: []int{1, 2, 10}, value: 21, expected: []int{10, 10, 1}},
	{coins: []int{1, 2, 5, 10}, value: 18, expected: []int{10, 5, 2, 1}},
	{coins: []int{1, 3, 5}, value: 7, expected: []int{5, 1, 1}},
	{coins: []int{1, 2, 5}, value: 11, expected: []int{5, 5, 1}},
	{coins: []int{2, 5, 10, 20}, value: 17, expected: []int{10, 5, 2}},
	{coins: []int{1, 2, 5, 10}, value: 3, expected: []int{2, 1}},
	{coins: []int{1, 2, 5, 10}, value: 0, expected: []int{}},
	{coins: []int{2, 4, 8}, value: 7, expected: []int{}},
	{coins: []int{1, 1, 2, 2, 5, 5, 10, 10}, value: 16, expected: []int{10, 5, 1}},
	{coins: []int{1, 2, 5}, value: 9, expected: []int{5, 2, 2}},
}

func Test_minCoins(t *testing.T) {
	for _, test := range tests {
		combination := minCoins(test.value, test.coins)
		if !reflect.DeepEqual(combination, test.expected) {
			t.Errorf("Coins: %v, Value: %d, Expected: %v, Got: %v", test.coins, test.value, test.expected, combination)
		}
	}
}

func Test_minCoins2(t *testing.T) {
	for _, test := range tests {
		combination := minCoins2(test.value, test.coins)
		if !reflect.DeepEqual(combination, test.expected) {
			t.Errorf("Coins: %v, Value: %d, Expected: %v, Got: %v", test.coins, test.value, test.expected, combination)
		}
	}
}
