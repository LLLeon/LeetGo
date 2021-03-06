0125.Valid Palindrome

##  描述
给定一个字符串，验证它是否是回文串，只考虑字母和数字字符，可以忽略字母的大小写。

说明：本题中，我们将空字符串定义为有效的回文串。

示例 1:

输入: "A man, a plan, a canal: Panama"
输出: true
示例 2:

输入: "race a car"
输出: false

## 解答
### 普通解法
两次遍历，第一次先把原始字符串中非字母非数字的元素去除，同时把元素转化为大写或小写。
第二次循环则利用头尾指针依次比较元素，限制条件为头指针小于等于尾指针，若元素相同则头指针后移，
尾指针前移，若碰到不相同的元素，返回 false。
时间复杂度，第一次遍历 O(n)，第二次 O(n/2)，共需 O(1.5n)。

### 优化解法
只遍历一次，遇到不是字母或数字的元素直接移动头尾指针，随后比较元素是否相等即可。
时间复杂度 O(n/2)。
