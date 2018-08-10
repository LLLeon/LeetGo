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

先用的暴力方法，两个 for 循环，相近的两个元素依次计算和再与 target 比较。