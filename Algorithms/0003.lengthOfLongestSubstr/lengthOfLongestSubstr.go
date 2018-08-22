package main

import (
	"fmt"
	"time"
)

/*
  暴力法
*/

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

/*
  滑动窗口
*/

func SlidingWindow(s string) int {
	var i, j, resp int

	n := len(s)
	unique := make(map[uint8]int)

	for i < n && j < n {
		if _, ok := unique[s[j]]; !ok {
			unique[s[j]] = j
			j++
			length := j - i

			if length > resp {
				resp = length
			}
		} else {
			delete(unique, s[i])
			i++
		}
	}

	return resp
}

func main() {
	s := "abcabcdabc"

	nowStart1 := time.Now()
	fmt.Println(LengthOfLongestSubstr(s))
	nowEnd1 := time.Now()
	fmt.Println(nowEnd1.Sub(nowStart1))

	nowStart2 := time.Now()
	fmt.Println(SlidingWindow(s))
	nowEnd2 := time.Now()
	fmt.Println(nowEnd2.Sub(nowStart2))
}
