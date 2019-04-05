0633.Sum of Square Numbers

##  描述
Given a non-negative integer c, your task is to decide whether there're two integers a and b such that a2 + b2 = c.

Example 1:

Input: 5
Output: True
Explanation: 1 * 1 + 2 * 2 = 5


Example 2:

Input: 3
Output: False

## 解答

也是利用双指针，从 i=0 到目标数字的平方根 j，依次对 `i*i + j*j` 与目标数字进行比较：
- 如果比目标数字大则向前移动 j。
- 如果比目标数字小则向后移动 i。