/*************************************************************************
+Author   : ChenhuiJia
+Data     : 2020-04-09
+************************************************************************/

package main

import "testing"

func TestDailyTemperatures(t *testing.T) {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}

	// [1 1 4 2 1 1 0 0]
	res := DailyTemperatures(T)
	t.Log("res: ", res)
}
