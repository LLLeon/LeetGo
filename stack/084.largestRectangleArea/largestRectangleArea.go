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

// 利用单调栈来求解
func LargestRegAreaII(heights []int) int {
	n := len(heights)
	if n == 0 {
		return 0
	}

	// 栈用来存放矩形的索引, 从栈底到栈顶元素的矩形高度是递增的
	maxArea, stack, height := 0, []int{}, 0

	for i := 0; i <= n; i++ {
		if i == n {
			height = 0
		} else {
			height = heights[i]
		}

		if len(stack) == 0 || height >= heights[stack[len(stack)-1]] {
			// 如果栈中没有数据或当前矩形高度大于栈顶矩形高度, 则直接将当前矩形的索引入栈, 继续遍历
			stack = append(stack, i)
		} else {
			// 如果栈中有数据且当前矩形高度小于栈顶矩形高度, 此时栈顶元素矩形的最大可达面积可以确定了(因为在栈中栈顶元素下面的元素高度肯定比它低),
			// 取出栈顶元素, 单独计算这个栈顶元素矩形的可达宽度
			tmpIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			length := 0

			if len(stack) == 0 {
				// 此时栈中没有元素了, 而且当前遍历到的元素比栈顶元素矩形低, 所以栈顶元素的可达宽度就是数组中当前元素左边的所有元素个数
				length = i
			} else {
				// 与上面情况不同的是, 因为弹出栈顶元素后的新栈顶元素比 tmpIndex 元素的高度低, 所以需要减去数组中新栈顶元素左边的元素个数
				length = i - 1 - stack[len(stack)-1]
			}

			maxArea = max(maxArea, heights[tmpIndex]*length)
			// 由于当前遍历到的元素没有计算, 所以下次遍历还需要再次计算它
			i--
		}
	}

	return maxArea
}

func max(a, b int) int {
	if a > b {
		return a
	}

	return b
}
