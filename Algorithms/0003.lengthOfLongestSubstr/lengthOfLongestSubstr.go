package main

import (
	"fmt"
)

// 暴力法
func LengthOfLongestSubstr(s string) int {
	resp := 0
	n := len(s)

	for i := 0; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			if IsAllUnique(s, i, j) {
				length := j - i

				if length > resp {
					resp = length
				}
			}
		}
	}

	return resp
}

func IsAllUnique(s string, start, end int) bool {
	unique := make(map[uint8]int)

	for i := start; i < end; i++ {
		// 检查 map 中是否有该元素，没有的话再放入当前元素
		if _, ok := unique[s[i]]; ok {
			return false
		}

		unique[s[i]] = i
	}

	return true
}

func main() {
	s := "hello"

	fmt.Println(LengthOfLongestSubstr(s))
}
