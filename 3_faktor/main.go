package main

import (
	"fmt"
	"math"
)

// 262144 -> 13208

func main() {
	r := findAll6FactorN(134217728)
	fmt.Println(len(r))
	// fmt.Println(r)
}

func findFactor(n int) []int {
	result := make([]int, 0)
	sq := int(math.Floor(math.Sqrt(float64(n))))

	for i := 1; i <= sq; i++ {
		if n%i == 0 {
			if n/i == i {
				result = append(result, i)
			} else {
				result = append(result, i, n/i)
			}
		}

		if len(result) > 6 {
			break
		}
	}
	return result
}

func findAll6FactorN(n int) map[int][]int {

	// map[int][]int, int is the number, []int is factor of number
	result := make(map[int][]int)
	for i := 1; i < n+1; i++ {
		factor := findFactor(i)
		if len(factor) == 6 {
			result[i] = factor
		}
	}
	return result
}
