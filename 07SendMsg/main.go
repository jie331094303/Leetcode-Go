package main

import "fmt"

func main() {
	ary := [][]int{
		{0, 2},
		{0, 1},
		{1, 2},
		{2, 1},
		{2, 0},
		{1, 0},
	}
	a := numWays(3, ary, 1)
	fmt.Println(a)

}

func numWays(n int, relation [][]int, k int) int {
	target := n - 1
	//total := 0
	return ForAlway(target, relation, 0, k)
}

func ForAlway(target int, relation [][]int, startK int, tarGetK int) int {
	result := 0
	if target == 0 && startK == tarGetK {
		return 1
	} else if startK > tarGetK {
		return 0
	}
	for _, v := range relation {
		if v[1] == target {
			result += ForAlway(v[0], relation, startK+1, tarGetK)
		}
	}
	return result
}
