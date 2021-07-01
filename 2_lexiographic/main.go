package main

import (
	"bufio"
	"fmt"
	"os"
)

// source: https://www.quora.com/How-would-you-explain-an-algorithm-that-generates-permutations-using-lexicographic-ordering

func main() {
	// a := []string{"a", "b", "c", "a", "c"}
	// fmt.Println(a)
	// s := "abcacb"
	s, err := getString()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
	uniqueSlice := unique(s)
	fmt.Println(uniqueSlice)

	// aa := []int{0, 1, 2}
	// fmt.Println(aa)
	// printLexiogrphicOrder(aa)

	lexioString(uniqueSlice)
	// s := "abc"
	// for _, v := range s {
	// 	fmt.Println(int(v))
	// }
}

func lexioString(arry []string) {
	// STEP 1
	largestI := -1
	for i := 0; i < len(arry)-1; i++ {
		if getAscii(arry[i]) < getAscii(arry[i+1]) {
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
		if getAscii(arry[largestI]) < getAscii(arry[j]) {
			largestJ = j
		}
	}

	// STEP 3
	arry = swapValue(arry, largestI, largestJ)

	// STEP 4: reverse from largestI + 1 to the end of array
	startArry := make([]string, largestI+1)
	copy(startArry, arry)

	endArry := make([]string, 0)
	for i := largestI + 1; i < len(arry); i++ {
		endArry = append(endArry, arry[i])
	}

	endArry = reverseArray(endArry)
	arry = append(startArry, endArry...)

	// Print and recursively looping
	fmt.Println(arry)
	// lexioString(arry)
}

func getAscii(s string) (a int) {
	for _, v := range s {
		a = int(v)
		break
	}
	return
}

func swapValue(arry []string, i int, j int) []string {
	temp := arry[i]
	arry[i] = arry[j]
	arry[j] = temp

	return arry
}

func reverseArray(arry []string) []string {
	for i, j := 0, len(arry)-1; i < j; i, j = i+1, j-1 {
		arry[i], arry[j] = arry[j], arry[i]
	}
	return arry
}

func unique(s string) []string {
	keys := make(map[int]bool)
	list := make([]string, 0)

	for _, r := range s {
		if _, value := keys[int(r)]; !value {
			keys[int(r)] = true
			list = append(list, string(r))
		}
	}
	return list
}

func getString() (string, error) {
	file, err := os.Open("./string.txt")
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
