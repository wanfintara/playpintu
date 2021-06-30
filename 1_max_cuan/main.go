package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	prices, err := getPrices()
	if err != nil {
		fmt.Println(err)
		return
	}

	maxCuan := findDiff(prices)
	fmt.Println("Max Cuan:", maxCuan)
}

func findDiff(arry []int) int {

	// Initialize
	n := len(arry)
	diff := arry[1] - arry[0]
	nowSum := diff
	maxSum := nowSum

	for i := 1; i < n-1; i++ {
		diff = arry[i+1] - arry[i]

		if nowSum > 0 {
			nowSum += diff

		} else {
			nowSum = diff
		}

		if nowSum > maxSum {
			maxSum = nowSum
		}
	}

	if maxSum <= 0 {
		return 0
	}
	return maxSum
}

func getPrices() ([]int, error) {
	file, err := os.Open("./prices.txt")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var prices []int
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return prices, err
		}
		prices = append(prices, x)
	}
	return prices, scanner.Err()
}
