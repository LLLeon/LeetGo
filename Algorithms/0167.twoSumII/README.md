0125.Two Sum II - Input array is sorted

##  描述
Given an array of integers that is already sorted in ascending order, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

Note:

Your returned answers (both index1 and index2) are not zero-based.
You may assume that each input would have exactly one solution and you may not use the same element twice.

Example:

Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.

## 解答

因为给定数组是从小到大排好序的，可以使用头尾双指针，头指针向后移动，尾指针向前移动，依次比较两数之和：
- 如果和大于目标数字，则尾指针向前移动。
- 反之则头指针向后移动，直到找到目标数字。但要注意头指针不能超过尾指针。

时间复杂度：O(n)