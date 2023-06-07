package combine

func GenerateCombinations(n int) [][]int {
	combinations := [][]int{}
	generate([]int{}, 1, n, &combinations)
	return combinations[1:]
}

func generate(current []int, start, n int, combinations *[][]int) {
	temp := make([]int, len(current))
	copy(temp, current)
	*combinations = append(*combinations, temp)

	for i := start; i <= n; i++ {
		generate(append(current, i), i+1, n, combinations)
	}
}
