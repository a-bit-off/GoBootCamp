package main

import (
	"sort"
)

func main() {

}
func minCoins(val int, coins []int) []int {
	res := make([]int, 0)
	i := len(coins) - 1
	for i >= 0 {
		for val >= coins[i] {
			val -= coins[i]
			res = append(res, coins[i])
		}
		i -= 1
	}
	return res
}

func minCoins2(val int, coins []int) []int {
	res := make([]int, 0)
	sort.SliceStable(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})

	for _, c := range coins {
		for c <= val {
			res = append(res, c)
			val -= c
		}
		if val == 0 {
			break
		}
	}

	if val != 0 {
		return []int{}
	}
	return res
}
