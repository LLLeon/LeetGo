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

	for left := 0; left < n; left++ {
		for right := left + 1; right <= n; right++ {
			if IsAllUnique(s, left, right) {
				length := right - left
				resp = Max(length, resp)
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

func Max(x, y int) int {
	if x >= y {
		return x
	}

	return y
}

/*
  滑动窗口算法
*/

func SlidingWindow(s string) int {
	n := len(s)
	left, right, resp := 0, 0, 0
	unique := make(map[uint8]int)

	// 刚开始时 left == right
	for left < n && right < n {
		// 检查索引 right 的值是否在 map 中，如果没有则存储，并向右滑动索引 right；
		// 记录当前最大子字符串长度。
		if _, ok := unique[s[right]]; !ok {
			unique[s[right]] = right
			right++

			length := right - left
			resp = Max(length, resp)
		} else {
			// 如果索引 right 的值已经存在于 map 中，则删除索引 left 的值并向右滑动 left。
			delete(unique, s[left])
			left++
		}
	}

	return resp
}

/*
  优化的滑动窗口算法
*/

func OptimizedSlidingWindow(s string) int {
	n := len(s)
	left, right, resp := 0, 0, 0
	unique := make(map[uint8]int)

	for ; right < n; right++ {
		if index, ok := unique[s[right]]; ok {
			// 遇到在 map 中已经存在的字符时，直接移动 left 到该字符的下一位的索引来跳过当前窗口，
			// 避免了逐个移动 left 索引。
			left = Max(index, left)
		}

		length := right - left + 1
		resp = Max(length, resp)

		// 这里存储的 value 是 right 的下一位索引，
		unique[s[right]] = right + 1
	}

	return resp
}

func main() {
	s := "ababcabceabc"

	nowStart1 := time.Now()
	fmt.Println(LengthOfLongestSubstr(s))
	nowEnd1 := time.Now()
	fmt.Println(nowEnd1.Sub(nowStart1))

	nowStart2 := time.Now()
	fmt.Println(SlidingWindow(s))
	nowEnd2 := time.Now()
	fmt.Println(nowEnd2.Sub(nowStart2))

	nowStart3 := time.Now()
	fmt.Println(OptimizedSlidingWindow(s))
	nowEnd3 := time.Now()
	fmt.Println(nowEnd3.Sub(nowStart3))
}
