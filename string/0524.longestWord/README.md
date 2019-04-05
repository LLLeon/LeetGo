0524.Longest Word in Dictionary through Deleting

## 描述

Given a string and a string dictionary, find the longest string in the dictionary that can be formed by deleting some characters of the given string. If there are more than one possible results, return the longest word with the smallest lexicographical order. If there is no possible result, return the empty string.

Example 1:
```
Input:
s = "abpcplea", d = ["ale","apple","monkey","plea"]

Output:
"apple"
```

Example 2:
```
Input:
s = "abpcplea", d = ["a","b","c"]

Output:
"a"
```

Note:
1. All the strings in the input will only contain lower-case letters.
2. The size of the dictionary won't exceed 1,000.
3. The length of all the strings in the input won't exceed 1,000.


## 解答

这个题目可以换个说法：找 d 列表中包含在 s 中的子串（不一定连续）的最长的一个，如果有多个一样长的，那就返回字典序最小的。

思路：
1. 挨个遍历 d 中的字符串，判断它是否在 s 中。
2. 选出这些在 s 中的字符串中最长的那个，返回。如最长的不止一个，则返回字典序最小的那个。




