package main

import "fmt"

//最长前缀
func main() {
	testArray := []string{"abc", "abd", "ab"}
	targetStr := longestCommonPrefix(testArray)
	fmt.Println(targetStr)
}

func longestCommonPrefix(strs []string) string {
	minStr := findMinLenStr(strs)
	longestComStr := findLogest(strs, minStr)
	return longestComStr
}

func findMinLenStr(strs []string) string {
	minLenStr := strs[0]
	for _, v := range strs {
		minLen := len(minLenStr)
		curLen := len(v)
		if curLen < minLen {
			minLenStr = v
		}
	}
	return minLenStr
}

func findLogest(strs []string, minStr string) string {
	LongestStr := ""
	arryLen := len(strs)
	isStop := false
	if arryLen == 1 { //如果为1时返回本身就是最长字符串
		return strs[0]
	} else {
		for i := 0; i < len(minStr); i++ {
			if !isStop {
				firstArryStr := ""       //已第一个数组为准
				for j, v := range strs { //第一个数组
					curCommonPrefix := v[0 : i+1]
					if j == 0 {
						firstArryStr = curCommonPrefix
					} else if j == arryLen-1 { //最后一个数组
						if firstArryStr == curCommonPrefix {
							LongestStr = firstArryStr
						} else {
							isStop = true
						}
					} else {
						if firstArryStr != curCommonPrefix {
							isStop = true
							break
						}
					}

				}
			} else {
				break
			}
		}
	}
	return LongestStr
}
