# 098. 验证二叉搜索树

## 描述

给定一个二叉树，判断其是否是一个有效的二叉搜索树。

假设一个二叉搜索树具有如下特征：

节点的左子树只包含小于当前节点的数。
节点的右子树只包含大于当前节点的数。
所有左子树和右子树自身必须也是二叉搜索树。
示例 1:

输入:
    2
   / \
  1   3
输出: true
示例 2:

输入:
    5
   / \
  1   4
     / \
    3   6
输出: false
解释: 输入为: [5,1,4,null,null,3,6]。
     根节点的值为 5 ，但是其右子节点值为 4 。

## 思路

1) 题目要求判断给定的二叉树是不是 BST，也就是需要判断这棵二叉树是否满足 BST 的特征。需要注意，判断时不仅要确认某个节点满足其左节点值比它大，

右节点值比它大，还要确认其左子树中所有节点值都比它小，右子树中所有节点值都比它大。

可以分别界定某个节点左右子树中节点值的范围，递归的来判断。

2) 还有一种思路，因为中序遍历 BST 的结果是一个升序序列，所以可以将 root.Val 于它的前一个节点值比较，如果每个节点都符合此特点，则为 BST。
