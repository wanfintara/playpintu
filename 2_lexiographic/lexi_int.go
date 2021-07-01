package main

import "fmt"

// source: https://www.quora.com/How-would-you-explain-an-algorithm-that-generates-permutations-using-lexicographic-ordering

func printLexiogrphicOrder(arry []int) {
	// STEP 1
	largestI := -1
	for i := 0; i < len(arry)-1; i++ {
		if arry[i] < arry[i+1] {
			largestI = i
		}
	}

	if largestI == -1 {
		fmt.Println("FINISHED")
		return
	}

	// STEP 2
	largestJ := -1
	for j := 0; j < len(arry); j++ {
		if arry[largestI] < arry[j] {
			largestJ = j
		}
	}

	// STEP 3
	arry = swapValueInt(arry, largestI, largestJ)

	// STEP 4: reverse from largestI + 1 to the end of array
	startArry := make([]int, largestI+1)
	copy(startArry, arry)

	endArry := make([]int, 0)
	for i := largestI + 1; i < len(arry); i++ {
		endArry = append(endArry, arry[i])
	}

	endArry = reverseArrayInt(endArry)
	arry = append(startArry, endArry...)

	// Print and recursively looping
	fmt.Println(arry)
	printLexiogrphicOrder(arry)
}

func swapValueInt(arry []int, i int, j int) []int {
	temp := arry[i]
	arry[i] = arry[j]
	arry[j] = temp

	return arry
}

func reverseArrayInt(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}
