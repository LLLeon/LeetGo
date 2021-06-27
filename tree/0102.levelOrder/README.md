# 102. 二叉树的层序遍历

## 描述

给你一个二叉树，请你返回其按 层序遍历 得到的节点值。 （即逐层地，从左到右访问所有节点）。

 

示例：
二叉树：[3,9,20,null,null,15,7],

    3
   / \
  9  20
    /  \
   15   7
返回其层序遍历结果：

[
  [3],
  [9,20],
  [15,7]
]

## 思路

利用队列存储每层的节点，只要队列不为空就遍历队列，存储每层节点值到数组，并将其非 nil 子节点加入队列。