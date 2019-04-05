0680.Valid Palindrome II

##  描述
Given a non-empty string s, you may delete at most one character. Judge whether you can make it a palindrome.

Example 1:
```
Input: "aba"
Output: True
```

Example 2:
```
Input: "abca"
Output: True
```
Explanation: You could delete the character 'c'.

Note:
The string will only contain lowercase characters a-z. The maximum length of the string is 50000.

## 解答

使用头尾双指针，从左右两端开始验证是否是回文串，验证的过程中，若两个字符不等，再左右各加一或减一，验证一遍。
此时子串范围为 (i+1, j) 或 (i, j-1) 的俩子串只要有任意一个是回文串，则结果就是回文串，否则就不是。

思路：根据题目要求，可以删除一个字符，那么遇到不相等的字符时，就可以往后或往前跳过一个字符，即前后各尝试删除了一个字符，只要有一个是回文串即满足题目要求。
时间复杂度：O(n)
