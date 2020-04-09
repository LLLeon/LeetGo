/*************************************************************************
+Author   : ChenhuiJia
+Data     : 2020-04-08
+************************************************************************/

package main

// 原则: 左括号入栈, 右括号出栈, 如果遍历结束时栈中没有元素, 说明合法.
func IsValid(s string) bool {
	m := map[string]string{
		")": "(",
		"}": "{",
		"]": "[",
	}

	stk := []string{}

	for _, v := range s {
		vs := string(v)
		x, vIsRight := m[vs]

		// 说明 vs 不是右括号, 可能是左括号或空字符串, 入栈
		if !vIsRight {
			stk = append(stk, vs)
		} else if len(stk) == 0 {
			// 此时需要执行 pop stack, 但里面却没有元素, 说明括号数量不匹配
			return false
		} else {
			// 出栈的元素与 x 比较是否相等
			if x != stk[len(stk)-1] {
				return false
			}
			stk = stk[:len(stk)-1]
		}
	}

	return len(stk) == 0
}
