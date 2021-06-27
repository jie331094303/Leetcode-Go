package main

import (
	"fmt"
	"sort"
)

func main() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	val := threeSum(nums)
	fmt.Println(val)
}

func threeSum(nums []int) [][]int {
	targetArray := [][]int{}
	sort.Ints(nums)
	arryLen := len(nums)
	for i, v := range nums {
		if v > 0 { //排除后了，后面的正数不可能相加等于0
			return targetArray
		}

		if i > 0 && nums[i] == nums[i-1] { //去除重复的解：例如[-1, 0, 1, 2, -1, -4] 去除i = -1 时重复的解
			continue
		}

		l := i + 1
		r := arryLen - 1
		for l < r {
			sum := v + nums[l] + nums[r]
			if sum > 0 {
				r--
			} else if sum < 0 {
				l++
			} else {
				zeroArray := []int{
					v,
					nums[l],
					nums[r],
				}
				targetArray = append(targetArray, zeroArray)
				for l < r && nums[l] == nums[l+1] { //去除:l+1 后重复的解
					l++
				}

				for l < r && nums[r] == nums[r-1] { //去除 r + 1后重复的解
					r--
				}

				l++
				r--
			}
		}
	}
	return targetArray
}
