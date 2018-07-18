
package main

import (
	"fmt"
	"math"
)


func main() {
	one := "abcabcbb"
	two := "bbbbb"
	third := "pwwkew"
	four := "abcdabaqio"
	fmt.Println("hello, world"+one+two+third+four)

	lengthOfLongestSubstring(one)
}

/**
* 解决思路: 利用窗口移动的思路;
* 1. 假设有两个端点，在这两个端点中不存在重复的元素;
* 
*/
func lengthOfLongestSubstring(s string) int {
	var max float64 = 0
	
	startSides := 0

	record := make(map[byte]int)

	for endSides:= 0; endSides < len(s); endSides++ {
		charStartSides, exist := record[s[endSides]]

		if exist {
			startSides = int(math.Max(float64(startSides),float64(charStartSides)))
		}

		max = math.Max(float64(endSides - startSides+1),max)
		record[s[endSides]]  = int(endSides+1)
	
	}
	fmt.Println(max)
	return int(max)
}
