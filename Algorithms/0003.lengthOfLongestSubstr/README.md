# 0002. Add Two Numbers

## 描述
Given a string, find the length of the longest substring without repeating characters.

**Example:**

Given "abcabcbb", the answer is "abc", which the length is 3.

Given "bbbbb", the answer is "b", with the length of 1.

Given "pwwkew", the answer is "wke", with the length of 3.Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

## 解答

先考虑暴力法，枚举出所有子字符串，找到最大的无重复字符的子字符串长度即可。
需要一个函数 IsAllUnique() 来判断子字符串中是否有重复字符，可利用 map 的 key 不可重复特性来记录遍历的字符。
如果此函数返回 true，则更新无重复字符的子字符串的最大长度。
