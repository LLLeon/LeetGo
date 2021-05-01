# 015. 三数之和

## 描述

给你一个包含 n 个整数的数组 nums，判断 nums 中是否存在三个元素 a，b，c ，使得 a + b + c = 0 ？请你找出所有和为 0 且不重复的三元组。

注意：答案中不可以包含重复的三元组。

## 示例

示例 1：

```
输入：nums = [-1,0,1,2,-1,-4]
输出：[[-1,-1,2],[-1,0,1]]
```

示例 2：

```
输入：nums = []
输出：[]
```

示例 3：

```
输入：nums = [0]
输出：[]
```


## 思路

根据题目要求，难点在如何去除重复的结果，比如 `[1, 0, -1]` 和 `[-1, 0, 1]` 就是重复的。

可以考虑先将数组进行去重和并排序，这样在寻找结果的过程中就不会对数字重复计算。先利用一个额外的 map 来存储每个数字出现的次数，

随后对 map 的 key 进行排序组成一个新数组。遍历新数组，找到与当前数字之和为 0 的两个数字。