package main

import (
	"testing"
)

var tests = []struct {
	name     string
	coins    []int
	value    int
	expected []int
}{
	{name: "test1", coins: []int{1, 2, 10}, value: 21, expected: []int{10, 10, 1}},
	{name: "test2", coins: []int{1, 2, 5, 10}, value: 18, expected: []int{10, 5, 2, 1}},
	{name: "test3", coins: []int{1, 3, 5}, value: 7, expected: []int{5, 1, 1}},
	{name: "test4", coins: []int{1, 2, 5}, value: 11, expected: []int{5, 5, 1}},
	{name: "test5", coins: []int{2, 5, 10, 20}, value: 17, expected: []int{10, 5, 2}},
	{name: "test6", coins: []int{1, 2, 5, 10}, value: 3, expected: []int{2, 1}},
	{name: "test7", coins: []int{1, 2, 5, 10}, value: 0, expected: []int{}},
	{name: "test8", coins: []int{2, 4, 8}, value: 7, expected: []int{}},
	{name: "test9", coins: []int{1, 1, 2, 2, 5, 5, 10, 10}, value: 16, expected: []int{10, 5, 1}},
	{name: "test10", coins: []int{1, 2, 5}, value: 9, expected: []int{5, 2, 2}},
}

func BenchmarkMinCoins(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				minCoins(test.value, test.expected)
			}
		})
	}
}

func BenchmarkMinCoins2(b *testing.B) {
	for _, test := range tests {
		b.Run(test.name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				minCoins2(test.value, test.expected)
			}
		})
	}
}
