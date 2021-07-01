package main

import "fmt"

func main() {
	fmt.Println(findFactor2(128))
}

func findFactors(n int) []int {
	result := make([]int, 0)

	i := 1
	for i*i <= n {
		if n%i == 0 {
			result = append(result, i)
		}

		if n/i != i {
			result = append(result, n/i)
		}
		i++
	}
	return result
}

func findFactor2(n int) []int {
	result := make([]int, 0)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			result = append(result, i)
		}
	}
	return result
}
