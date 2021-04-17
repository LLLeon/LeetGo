package main

// 时间复杂度 O(N2), 空间复杂度 O(1)
// 暴力法, 向当前位置左右扩散寻找比其高或等高的矩形个数
func LargestRegAreaI(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	maxArea := 0

	for i, num := range heights {
		leftCount, rightCount := 0, 0

		// 计算当前位置左边比其高或等高的矩形个数, 但中间不能跨过比当前矩形低的位置
		for j := i - 1; j >= 0; j-- {
			if heights[j] >= num {
				leftCount++
			} else {
				break
			}
		}

		// 同理计算当前位置右边
		for k := i + 1; k < n; k++ {
			if heights[k] >= num {
				rightCount++
			} else {
				break
			}
		}

		currArea := (leftCount + 1 + rightCount) * num
		if currArea > maxArea {
			maxArea = currArea
		}
	}

	return maxArea
}
