package main

import (
	"sort"
)

// нет проверки краевых случаев
// в случае когда прорамма должна использовать больший номинал она этого не делает
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

// minCoins2 сортирует слайс coins по убыванию
// в случае пустого слайса программа вернет пустой слайс
// в случае когда сумма отрицательная, программа вернет пустой слайс
// в случае когда не удалось набрать необходимую сумму, программа вернет пустой слайс
func minCoins2(val int, coins []int) []int {
	res := make([]int, 0)

	if len(coins) <= 0 || val <= 0 {
		return res
	}

	sort.SliceStable(coins, func(i, j int) bool {
		return coins[i] > coins[j]
	})

	for _, c := range coins {
		for c <= val {
			res = append(res, c)
			val -= c
		}
		if val == 0 {
			return res
		}
	}

	if val != 0 {
		return []int{}
	}
	return res
}

// запуск документации
// godoc -http:6060
// open http://localhost:6060/
