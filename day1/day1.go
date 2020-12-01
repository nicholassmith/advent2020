package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input.txt")
	if err != nil {
		fmt.Println("error: ", err)
		return
	}
	strs := strings.Split(string(data), "\n")
	valArray := make([]int, len(strs))
	for i := range valArray {
		valArray[i], _ = strconv.Atoi(strs[i])
	}

	for _, n := range valArray {
		searchVal := 2020 - n
		if findMatch(valArray, searchVal) {
			fmt.Println("return value part 1:", searchVal*n)
			break
		}
	}

	subsetSum(valArray, 2020, make([]int, 0))
}

func findMatch(a []int, x int) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func subsetSum(a []int, target int, partial []int) {
	sumPartial := sum(partial)

	if sumPartial == target {
		if len(partial) == 3 {
			fmt.Println("return value part 2:", product(partial))
			return
		}

	}

	if sumPartial > target {
		return
	}

	for i, n := range a {
		subsetSum(a[i+1:], target, append(partial, n))
	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func product(array []int) int {
	result := 1
	for _, v := range array {
		result *= v
	}

	return result
}
