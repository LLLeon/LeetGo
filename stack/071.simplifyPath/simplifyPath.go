package main

import "strings"

// . 和 .. 分别表示当前目录和上一级目录, 根据题目要求, 也就是需要返回目录跳转后的 path.
// 思路: 可以先将 path 按 / 切分为目录数组, 然后逐个判断各个目录是否需要跳转. 这里可以利用栈的特性来存储遍历过的目录:
//     1. 如果是 . 不需要跳转;
//     2. 如果是 .. 需要跳转到上一级目录, 将栈顶目录出栈;
//     3. 将目录入栈;
//     4. 这样操作过后, 将栈中剩余的目录拼接即可.

// 时间复杂度 O(N), 空间复杂度 O(N).
func SimplifyPath(path string) string {
	if path == "" {
		return ""
	}

	stack := []string{}
	res := "/"

	pList := strings.Split(path, "/")

	for i := 0; i < len(pList); i++ {
		if pList[i] == "." || pList[i] == "" {
			continue
		}

		if pList[i] == ".." {
			if len(stack) != 0 {
				stack = stack[:len(stack)-1]
			}
			continue
		}

		stack = append(stack, pList[i])
	}

	for i := 0; i < len(stack); i++ {
		res = res + stack[i]
		if i != len(stack)-1 {
			res = res + "/"
		}
	}

	return res
}
