# 0003. Longest Substring Without Repeating Characters

## 描述
Given a string, find the length of the longest substring without repeating characters.

**Example:**

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3.Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

## 解答
### 暴力法
枚举出所有子字符串，找到最大的无重复字符的子字符串长度即可。
需要一个函数 IsAllUnique() 来判断子字符串中是否有重复字符，可利用 map 的 key 不可重复特性来记录遍历的字符。
如果此函数返回 true，则更新无重复字符的子字符串的最大长度。

### 滑动窗口
可用 map 作为滑动窗口来存不重复的字符，即窗口 [i, j) 中不包含重复字符，查找的时间复杂度为 O(1)。
这样只要向右滑动索引 j，直到检查到字符 `s[j]` 存在于窗口中，再向右滑动索引 i，直到遍历完所有的 i，
这样就找到了以索引 i 为起点的最长子字符串。

### 优化的滑动窗口
按照滑动窗口的思路继续优化，如果遇到重复字符时，直接跳过该窗口，即移动 left 索引到 map 中已存储的重复字符的下一位，
即可避免逐个增加 left 的索引。
