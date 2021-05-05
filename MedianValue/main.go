package main

import (
	"fmt"
	"sort"
)

func main() {
	var result float64 = findMedianSortedArrays([]int{1, 3}, []int{2, 7})
	fmt.Println(result)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var result float64
	targets := []int{} //make([]int,10)
	targets = append(targets, nums1...)
	targets = append(targets, nums2...)
	sort.Ints(targets)
	if len(targets)%2 == 0 {
		secondIndex := len(targets) / 2
		firstIndex := secondIndex - 1
		firstValue := targets[firstIndex]
		lastValue := targets[secondIndex]
		result = float64((firstValue + lastValue)) / 2.0
	} else {
		index := len(targets) / 2
		result = float64(targets[index])
	}
	return result
}
