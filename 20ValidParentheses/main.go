package main

import "fmt"

func main() {
	isOK := isValid("){")
	fmt.Println(isOK)
}

func isValid(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	ok := true
	arry := []string{}
	matchDic := map[string]string{
		"}": "{",
		"]": "[",
		")": "(",
	}
	for i := 0; i < len(s); i++ {
		curStr := string(s[i])
		if curStr == "(" || curStr == "{" || curStr == "[" {
			arry = append(arry, curStr)
		} else {
			expectStr := matchDic[curStr]
			if len(arry) == 0 { //"){"
				return false
			}
			popStr := arry[len(arry)-1:][0]
			if expectStr != popStr {
				return false
			} else {
				arry = arry[:len(arry)-1] //移除当前最后一个元素
			}
		}
	}
	if len(arry) > 0 { //匹配完应该都为空
		ok = false
	}
	return ok
}
