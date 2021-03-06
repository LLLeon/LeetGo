# 0297. 二叉树的序列化与反序列化

## 描述

序列化是将一个数据结构或者对象转换为连续的比特位的操作，进而可以将转换后的数据存储在一个文件或者内存中，同时也可以通过网络传输到另一个计算机环境，采取相反方式重构得到原数据。

请设计一个算法来实现二叉树的序列化与反序列化。这里不限定你的序列 / 反序列化算法执行逻辑，你只需要保证一个二叉树可以被序列化为一个字符串并且将这个字符串反序列化为原始的树结构。

示例 1：

输入：root = [1,2,3,null,null,4,5]
输出：[1,2,3,null,null,4,5]

示例 2：

输入：root = []
输出：[]

## 思路

根据题目要求，序列化与反序列化二叉树，也就是在二叉树节点与字符串之间进行转换，只要序列化与反序列化方法的逻辑能自洽即可。

序列化：给定的输入是二叉树的根节点，需要遍历二叉树，然后将每个节点的值转换为字符串，可以选择前序、后序遍历，看个人偏好。

反序列化：根据序列化的结果，提取出每个节点的值，然后根据序列化的方式（前序、后序）构建二叉树。

注意：中序可以序列化，但由于无法确定 root 节点的索引，所以无法实现中序反序列化。
