package word

import "strings"

func FindLongestWord(s string, d []string) string {
	maxLength := 0
	maxIndex := -1

	for i := 0; i < len(d); i++ {
		word := d[i]
		// 找到 d 中包含在 s 中的字符串，记录最长的字符串
		if IsContain(s, word) {
			if len(word) > maxLength {
				maxLength = len(word)
				maxIndex = i
			}

			// 相同长度的字符串按照字典序排序
			if len(word) == maxLength && strings.Compare(word, d[maxIndex]) < 0 {
				maxIndex = i
			}
		}
	}

	// 说明 d 中所以字符串都没有包含在 s 中
	if maxIndex == -1 {
		return ""
	}

	return d[maxIndex]
}

// IsContain returns whether src contains sub.
func IsContain(src, sub string) bool {
	i, j := 0, 0
	m, n := len(src), len(sub)

	for i < m && j < n {
		// 因为是检查 src 中是否包含 sub，所以 src 可以跳过字符
		for (i < m) && (src[i] != sub[j]) {
			i++
		}

		// src 已经遍历完
		if i == m {
			break
		}
		if src[i] == sub[j] {
			i++
			j++
		}
	}

	// sub 还没有遍历完，说明 sub 比 src 要长
	if j != n {
		return false
	}

	return true
}
