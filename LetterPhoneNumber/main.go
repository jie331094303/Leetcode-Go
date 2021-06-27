package main

import (
	"fmt"
)

func main() {
	array := letterCombinations("23")
	fmt.Println(array)
}

func letterCombinations(digits string) []string {
	targetArray := []string{}
	allLetter := map[string][]string{
		"2": {"a", "b", "c"},
		"3": {"d", "e", "f"},
		"4": {"g", "h", "i"},
		"5": {"j", "k", "l"},
		"6": {"m", "n", "o"},
		"7": {"p", "q", "r", "s"},
		"8": {"t", "u", "v"},
		"9": {"w", "x", "y", "z"},
	}
	if len(digits) > 0 {
		GoBack(digits, allLetter, &targetArray, 0, "")
	}
	return targetArray
}

func GoBack(digits string, allLetter map[string][]string, targetArray *[]string, index int, result string) {
	if index > len(digits)-1 {
		*targetArray = append(*targetArray, result)
		return
	}
	for _, v := range allLetter[string(digits[index])] {
		GoBack(digits, allLetter, targetArray, index+1, result+v)
	}
}
