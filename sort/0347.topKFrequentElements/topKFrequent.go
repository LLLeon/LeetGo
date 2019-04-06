package frequent

func TopKFrequent(nums []int, k int) []int {
	var (
		n            = len(nums)
		frequencyNum = make(map[int]int)
		buckets      = make([][]int, n+1)
		topK         = make([]int, 0)
	)

	// 用哈希表统计每个数字出现的频率
	for _, num := range nums {
		if frequency, ok := frequencyNum[num]; !ok {
			frequencyNum[num] = 1
		} else {
			frequencyNum[num] = frequency + 1
		}
	}

	// 把哈希表中统计好的数字放到桶中其相应频率的下标位置上
	for num := range frequencyNum {
		frequency := frequencyNum[num]

		if buckets[frequency] == nil {
			buckets[frequency] = []int{}
		}

		buckets[frequency] = append(buckets[frequency], num)
	}

	// 从尾到头遍历桶，找到出现频率最多的 k 个数字
	for i := n; i >= 0 && len(topK) < k; i-- {
		if buckets[i] == nil {
			continue
		}

		topK = append(topK, buckets[i]...)
	}

	return topK
}
