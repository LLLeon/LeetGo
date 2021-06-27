# 0108. 将有序数组转换为二叉搜索树

## 描述

给你一个整数数组 nums ，其中元素已经按 升序 排列，请你将其转换为一棵 高度平衡 二叉搜索树。

高度平衡 二叉树是一棵满足「每个节点的左右两个子树的高度差的绝对值不超过 1 」的二叉树。

示例 1：

输入：nums = [-10,-3,0,5,9]
输出：[0,-3,9,-10,null,5]
解释：[0,-10,5,null,-3,null,9] 也将被视为正确答案：

示例 2：

输入：nums = [1,3]
输出：[3,1]
解释：[1,3] 和 [3,1] 都是高度平衡二叉搜索树。


## 思路

BST 按中序遍历的结果是升序序列，可以考虑递归构建 BST，找到 root 节点，递归构建其左右子节点即可。这里选择数组中间位置的元素作为 root 节点。