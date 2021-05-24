package main

import "fmt"

func main() {
	fmt.Println(longestPalindrome("cbc"))
}

func longestPalindrome(s string) string {
	res := ""
	maxLen := len(s)
	if len(s) < 2 {
		return s
	}
	//假设回文从奇数和偶数都有回文
	for i := 0; i < maxLen; i++ {
		GetSubString(s, i, i, &res)
		GetSubString(s, i, i+1, &res)
	}
	return res
}

func GetSubString(s string, left int, right int, res *string) {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	//获取到的回文长度大于现有的长度
	if right-left-1 > len(*res) {
		*res = s[left+1 : right] //获取现在有的长度
	}
}
