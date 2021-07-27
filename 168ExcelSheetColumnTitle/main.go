package main

import "fmt"

func main() {
	str := convertToTitle(702)
	fmt.Println(str)
}

func convertToTitle(columnNumber int) string {
	allLetter := map[int]string{
		0:  "Z",
		1:  "A",
		2:  "B",
		3:  "C",
		4:  "D",
		5:  "E",
		6:  "F",
		7:  "G",
		8:  "H",
		9:  "I",
		10: "J",
		11: "K",
		12: "L",
		13: "M",
		14: "N",
		15: "O",
		16: "P",
		17: "Q",
		18: "R",
		19: "S",
		20: "T",
		21: "U",
		22: "V",
		23: "W",
		24: "X",
		25: "Y",
	}
	tagrgetStr := ""
	num := columnNumber
	for num > 0 { //说明前面还有值
		resultInt := num / 26
		yu := num % 26
		tagrgetStr = allLetter[yu] + tagrgetStr //从后往前计算

		if yu == 0 { //说明后一位是Z，一个Z是26，两个ZZ是702=26*26 + 26(26*27),702/26 = 27,所以要减一
			resultInt -= 1
		}
		num = resultInt
	}
	return tagrgetStr
}
