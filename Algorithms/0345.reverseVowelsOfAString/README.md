0345.Reverse Vowels of a String

##  描述

Write a function that takes a string as input and reverse only the vowels of a string.

Example 1:
```
Input: "hello"
Output: "holle"
```

Example 2:
```
Input: "leetcode"
Output: "leotcede"
```

## 解答

使用头尾双指针，一个指针从头向尾遍历，一个指针从尾到头遍历，遇到元音字符则交换。
哈希表查找的时间复杂度是 O(1)，
所以时间复杂度是 O(n)。