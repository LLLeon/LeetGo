/*************************************************************************
+Author   : chenhuijia@deepglint.com
+Data     : 2020-03-14
+************************************************************************/

package power

// 1. 暴力法, 2 的 0 次幂是 1, 将 1, 2, 4, ... 逐个与 n 比较是否相等,
// 直到所枚举的数超过 n, 如果找到与 n 相等的数, 则表示 n 是 2 的某次幂, 否则不是.
func IsPowerOfTwoI(n int) bool {
	if n <= 0 {
		return false
	}

	power := 1

	for power <= n {
		if power == n {
			return true
		}

		power <<= 1
		// 注意 overflow int 时要 break
	}

	return false
}

// 2. 2 的次幂的二进制形式只有最高位是 1, 其余位是 0,
// 问题就转化为“判断一个数的二进制，除了最高位为 1，是否还有别的 1 存在”,
// 假设 N 是 2 的次幂, 那么 N-1 的所有位都是 1, 则会发现这样的规律: N&(N-1)=0.
func IsPowerOfTwoII(n int) bool {
	if n <= 0 {
		return false
	}

	return n&(n-1) == 0
}

// 假设 n 是 2 的次幂, 求一下 log2N.
// 非递归方式
func Log2I(n int) int {
	if n <= 0 {
		return -1
	}

	// 如果 n 是 1, 则幂为 0
	power := 0

	// 在到达 2 最小的次幂前不断右移, 计算 n 是 2 的几次幂.
	for n > 1 {
		n >>= 1
		power++
	}

	return power
}

// 递归方式
func Log2II(n int) int {
	if n <= 0 {
		return -1
	}

	// 递归函数的结束条件, 即 2 的 0 次方是 1.
	if n == 1 {
		return 0
	}

	// 函数的参数不断右移, 直到触及结束条件, 再加上外面的 1,
	// 递归计算出结果.
	return 1 + Log2II(n>>1)
}
