package main

import (
	"fmt"
	"strconv"
)

func main() {
	array := letterCombinations("23")
	fmt.Println(array)
}

func letterCombinations(digits string) []string {
	targetArray := []string{}
	allLetter := map[int][]string{
		2: {"a", "b", "c"},
		3: {"d", "e", "f"},
		4: {"g", "h", "i"},
		5: {"j", "k", "l"},
		6: {"m", "n", "o"},
		7: {"p", "q", "r", "s"},
		8: {"t", "u", "v"},
		9: {"w", "x", "y", "z"},
	}
	if len(digits) > 0 {
		arry1 := GetArray(digits, allLetter, 0)
		arry2 := GetArray(digits, allLetter, 1)
		arry3 := GetArray(digits, allLetter, 2)
		arry4 := GetArray(digits, allLetter, 3)

		switch len(digits) {
		case 1:
			targetArray = for1(arry1)
		case 2:
			targetArray = for2(arry1, arry2)
		case 3:
			targetArray = for3(arry1, arry2, arry3)
		case 4:
			targetArray = for4(arry1, arry2, arry3, arry4)
		}
	}
	return targetArray
}

func GetArray(digits string, allLetter map[int][]string, index int) []string {
	arry := []string{}
	arryLen := len(digits)
	if arryLen > index { //数组长度要大于需要取的下标才进行，不然会报错

		intVal, _ := strconv.Atoi(string(digits[index]))
		arry = allLetter[intVal]

	}
	return arry
}

func for1(arry []string) []string {
	array := []string{}
	for _, v := range arry {
		array = append(array, v)
	}
	return array
}

func for2(arry1, arry2 []string) []string {
	array := []string{}
	for _, v1 := range arry1 {
		for _, v2 := range arry2 {
			array = append(array, v1+v2)
		}
	}
	return array
}

func for3(arry1, arry2, arry3 []string) []string {
	array := []string{}
	for _, v1 := range arry1 {
		for _, v2 := range arry2 {
			for _, v3 := range arry3 {
				array = append(array, v1+v2+v3)
			}
		}
	}
	return array
}

func for4(arry1, arry2, arry3, arry4 []string) []string {
	array := []string{}
	for _, v1 := range arry1 {
		for _, v2 := range arry2 {
			for _, v3 := range arry3 {
				for _, v4 := range arry4 {
					array = append(array, v1+v2+v3+v4)
				}
			}
		}
	}
	return array
}
