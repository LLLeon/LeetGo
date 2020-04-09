/*************************************************************************
+Author   : ChenhuiJia
+Data     : 2020-04-09
+************************************************************************/

package main

// 题目给出的是固定天数的每天温度的数组, 要求计算每天会在几天后会升温.
// 问题可以转化为, 找到数组中比当前索引 index1 值 T(index1) 大的值 T(index2) 的索引 index2.
func DailyTemperatures(T []int) []int {
	n := len(T)
	if n == 0 {
		return []int{}
	}

	stk := []int{}
	res := make([]int, n)

	for i := 0; i < n; i++ {
		// 不断循环计算当前值 T[i] 是否比栈顶索引位置的值大并记录结果, 将得出结果的索引出栈
		for len(stk) != 0 && T[i] > T[stk[len(stk)-1]] {
			res[stk[len(stk)-1]] = i - stk[len(stk)-1]
			stk = stk[:len(stk)-1]
		}

		// 将当前索引 i 入栈, 以后再计算.
		stk = append(stk, i)
	}

	for len(stk) != 0 {
		res[stk[len(stk)-1]] = 0
		stk = stk[:len(stk)-1]
	}

	return res
}

/*
prepare:
  T = [73, 74, 75, 71, 69, 72, 76, 73]
  n = 8
  skt = []

process:
1). i=0, stk=[0]
2). i=1, T1=74, T0=73, res[0]=1-0=1 -> res=[1], stk=[], stk=[1]
3). i=2, T2=75, T1=74, res[1]=2-1=1 -> res=[1, 1], stk=[], stk=[2]
4). i=3, T3=71, T2=75, stk=[2, 3]
5). i=4, T4=69, T3=71, stk=[2, 3, 4]
6). i=5, T5=72, T4=69, res[4]=5-4=1 -> res=[1, 1, *, *, 1], stk=[2, 3]
  6.1). T5=72, T3=71, res[3]=5-3=2 -> res=[1, 1, *, 2, 1], stk=[2]
  6.2). T5=72, T2=75, stk=[2, 5]
7). i=6, T6=76, T5=72, res[5]=6-5=1 -> res=[1, 1, *, 2, 1, 1], stk=[2]
  7.1). T6=76, T2=75, res[2]=6-2=4 -> res=[1, 1, 4, 2, 1, 1], stk=[], stk=[6]
8). i=7, T7=73, T6=76, stk=[6, 7]
9). res[7]=0 -> res=[1, 1, 4, 2, 1, 1, *, 0], stk=[6]
  9.1). res[6]=0 -> res=[1, 1, 4, 2, 1, 1, 0, 0], stk=[]
10). return res = [1, 1, 4, 2, 1, 1, 0, 0]
*/
