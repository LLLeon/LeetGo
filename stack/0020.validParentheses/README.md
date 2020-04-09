Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

```$xslt
Input: "()"

Output: true
```


Example 2:

```$xslt
Input: "()[]{}"
Output: true
```
Example 3:

```$xslt
Input: "(]"
Output: false
```

Example 4:
```$xslt
Input: "([)]"
Output: false
```

Example 5:

```$xslt
Input: "{[]}"
Output: true
```

解析:

1. 利用栈先入后出的特性, 遍历字符串, 遇到左括号入栈, 遇到右括号出栈, 遍历结束时如果栈中没有元素则合法;

2. 用 map 构建左右括号的映射关系, 这样查询左右括号是否匹配只需 O(1) 的时间复杂度.