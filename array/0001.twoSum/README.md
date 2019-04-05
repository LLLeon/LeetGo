# 0001. Two Sum

## 描述

Given an array of integers, return **indices** of the two numbers such that they add up to a specific target.

You may assume that each input would have **exactly** one solution, and you may not use the *same* element twice.

**Example**：

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

## 解答

1. 先用的暴力方法，两个 for 循环，相近的两个元素依次计算和再与 target 比较。
2. 由于 map 查找的时间复杂度是 O(1)，先把 nums 的 值-index 存入 map，再遍历比较。有一点需要注意，目标元素不能是 nums[i] 自身。
3. 进一步，可以在第一次遍历时先进行比较，如果匹配到目标值就返回，否则就放入 map。不过返回时要注意两个 index 的顺序。