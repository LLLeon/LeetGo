# 0002. Add Two Numbers

## 描述

You are given two **non-empty** linked lists representing two non-negative integers. The digits are stored in **reverse order** and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

**Example:**

```bash
Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
```

## 解答

链表第一个节点表示个位，是倒序的，可以迭代两个链表来逐个计算每个节点的和。

可以新建一个节点来记录 l1 和 l2 各位计算的和，最后返回该节点。

需要注意，进位的值在迭代到下一位时要加上，以及迭代到最后时，不要忘记加上最终进位的值。

见代码注释。