package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// source: https://www.quora.com/How-would-you-explain-an-algorithm-that-generates-permutations-using-lexicographic-ordering

func main() {
	s, err := getString()
	if err != nil {
		fmt.Println(err)
		return
	}
	// fmt.Println(s)
	uniqueSlice := unique(s)
	// fmt.Println(uniqueSlice)

	lexioString(uniqueSlice)
}

func lexioString(arry []string) {
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

	largestJ := -1
	for j := 0; j < len(arry); j++ {
		if getAscii(arry[largestI]) < getAscii(arry[j]) {
			largestJ = j
		}
	}

	arry = swapValue(arry, largestI, largestJ)

	startArry := make([]string, largestI+1)
	copy(startArry, arry)

	endArry := make([]string, 0)
	for i := largestI + 1; i < len(arry); i++ {
		endArry = append(endArry, arry[i])
	}

	endArry = reverseArray(endArry)
	arry = append(startArry, endArry...)

	// Print and recursively looping
	s := strings.Join(arry[:], "")
	fmt.Println(s)
	lexioString(arry)
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
