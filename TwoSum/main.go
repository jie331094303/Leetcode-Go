package main

import (
	"fmt"
)

func main() {
	testArry := []int{1, 3, 7}
	resultArry := twoSum(testArry, 8)
	fmt.Println(resultArry)
}

func twoSum(nums []int, target int) []int {
	var isFound bool = false
	rtArry := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		if isFound {
			break
		}
		firstVa := nums[i]
		for j := i + 1; j < len(nums); j++ {
			secondVa := nums[j]
			sum := firstVa + secondVa
			if sum == target {
				isFound = true
				rtArry[0] = i
				rtArry[1] = j
				break
			}
		}
	}
	return rtArry
}
