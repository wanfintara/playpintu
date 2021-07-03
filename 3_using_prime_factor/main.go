package main

// source: https://www.youtube.com/watch?v=dfZ8DjHeNII&list=LL&index=7
// find total number in a factorization

import (
	"fmt"
	"math"
	"sync"
)

func main() {
	// r := findPrimeNumbers(134217728)
	// for k, v := range r {
	// 	fmt.Println(k, "has appeared", v, "times")
	// }

	// n := findTotalFactorNumber(r)
	// fmt.Println("Total Factor number", n)

	// fmt.Println(findTotalNumbersHaveSixFactor(134217728))
	fmt.Println(findTotalNumbersHaveSixFactorII(134217728))
}

// Approximate 60 mins to finish, sad...
// 134217728 has 3969187 of 6 factor numbers
func findTotalNumbersHaveSixFactor(n int) int {
	if n < 12 {
		return 0
	}

	var result int
	for i := 12; i <= n; i++ {
		primeNumbers := findPrimeNumbers(i)
		totalFactorNumbers := findTotalFactorNumber(primeNumbers)

		if totalFactorNumbers == 6 {
			result += 1
			fmt.Println(result, ". ", i, "has", totalFactorNumbers)
		}
	}
	return result
}

// Using goroutine
// 11:56 PM - 12:11 AM
// 34217728 -> 3969187, 15 mins, better
func findTotalNumbersHaveSixFactorII(n int) int {
	var mutex = &sync.Mutex{}
	var wg sync.WaitGroup
	var result int

	if n < 12 {
		return 0
	}

	for i := 12; i <= n; i++ {
		wg.Add(1)
		go func(ig int) {
			defer wg.Done()

			primeNumbers := findPrimeNumbers(ig)
			totalFactorNumbers := findTotalFactorNumber(primeNumbers)

			if totalFactorNumbers == 6 {
				mutex.Lock()
				result += 1
				fmt.Println(result, ". ", ig, "has", totalFactorNumbers)
				mutex.Unlock()
			}
		}(i)
	}
	fmt.Println("Finding.... please wait bro :)")
	wg.Wait()
	fmt.Println("Finished")

	return result
}

func isToTalFactorNumberIsSix(n int) bool {
	primeNumbers := findPrimeNumbers(n)
	totalFactorNumbers := findTotalFactorNumber(primeNumbers)

	if totalFactorNumbers != 6 {
		return false
	}
	return true
}

func findTotalFactorNumber(m map[int]int) int {
	n := 1
	for _, v := range m {
		v += 1
		n = n * v
	}
	return n
}

// https://www.geeksforgeeks.org/print-all-prime-factors-of-a-given-number/
func findPrimeNumbers(n int) map[int]int {
	// int I  -> prime number
	// int II -> prime number appeared how many times
	result := make(map[int]int)

	for n%2 == 0 {
		result[2] += 1
		n = n / 2
	}

	sqrtN := int(math.Sqrt(float64(n)))
	for i := 3; i <= sqrtN; i += 2 {
		for n%i == 0 {
			result[i] += 1
			n /= i
		}
	}

	if n > 2 {
		result[n] += 1
	}

	return result
}
