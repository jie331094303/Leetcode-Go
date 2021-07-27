package main

import (
	"fmt"	"math"
)


func main() {
	a := [][]int{
		{-10},
		{-7},
		{-5},
	}
	fmt.Println(searchMatrix(a, -10))
}

func searchMatrix(matrix [][]int, target int) bool {
	isOK := false
	isSortOK, valMap := checkSortAndGetValMap(matrix)
	if isSortOK && checkTargetNumIsContains(matrix, valMap, target) {
		isOK = true
	}
	return isOK
}

// Check Number is sort and get each arry scope
// return valMap is [2][]int
func checkSortAndGetValMap(matrix [][]int) (isOK bool, valMap [][]int) {
	isOK = true
	singleArrayLen := len(matrix[0]) - 1 
	lastVal := math.MinInt64
	valMap = [][]int{} //[2][]array
	for _, array := range matrix {
		curValScope := []int{}
		for i, nextValue := range array {
			if nextValue >= lastVal {
				lastVal = nextValue
				if i == 0 || i == singleArrayLen {
					curValScope = append(curValScope, lastVal)
				}
			} else {
				return false, valMap
			}
		}
		valMap = append(valMap, curValScope)
	}
	return isOK, valMap
}

func checkTargetNumIsContains(matrix [][]int, valMap [][]int, target int) bool {
	isOK := false
	matchIndex := -1
	for i, array := range valMap {
		startNum := array[0]
		var endNum int
		if len(array) > 1 {
			endNum = array[1]
		} else {
			endNum = array[0]
		}
		if target >= startNum && target <= endNum {
			matchIndex = i
			break
		}
	}
	if matchIndex > -1 {
		matchArray := matrix[matchIndex]
		for _, v := range matchArray {
			if v == target {
				isOK = true
			}
		}
	}
	return isOK
}
