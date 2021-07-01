package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// source: https://www.quora.com/How-would-you-explain-an-algorithm-that-generates-permutations-using-lexicographic-ordering

func main() {
	// s, err := getString()
	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// fmt.Println(s)
	// uniqueSlice := unique(s)
	// fmt.Println(uniqueSlice)
	// fmt.Println(uniqueSlice[0])

	printLexiogrphicOrder([]int{0, 1, 2})
}

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
	arry = swapValue(arry, largestI, largestJ)

	// STEP 4: reverse from largestI + 1 to the end of array
	startArry := make([]int, largestI+1)
	copy(startArry, arry)

	endArry := make([]int, 0)
	for i := largestI + 1; i < len(arry); i++ {
		endArry = append(endArry, arry[i])
	}

	endArry = reverseArray(endArry)
	arry = append(startArry, endArry...)

	// Print and recursively looping
	fmt.Println(arry)
	printLexiogrphicOrder(arry)
}

func swapValue(arry []int, i int, j int) []int {
	temp := arry[i]
	arry[i] = arry[j]
	arry[j] = temp

	return arry
}

func reverseArray(arr []int) []int {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func unique(s string) []string {
	keys := make(map[rune]bool)
	list := make([]string, 0)

	for _, entry := range s {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, strconv.QuoteRune(entry))
		}
	}
	return list
}

func getString() (string, error) {
	file, err := os.Open("./string2.txt")
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	var s string
	for scanner.Scan() {
		s = scanner.Text()
	}

	return s, scanner.Err()
}
