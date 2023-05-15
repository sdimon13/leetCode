# 98. Validate Binary Search Tree
Level: `Medium`

https://leetcode.com/problems/validate-binary-search-tree/

---

# Description

Given the `root` of a binary tree, determine if it is a valid binary search tree (BST).

A **valid BST** is defined as follows:

The left **subtree** of a node contains only nodes with keys **less than** the node's key.
The right subtree of a node contains only nodes with keys **greater than** the node's key.
Both the left and right subtrees must also be binary search trees.

## Example 1:

![Example 1](https://assets.leetcode.com/uploads/2020/12/01/tree1.jpg)

    Input: root = [2,1,3]
    Output: true

## Example 2:

![Example 2](https://assets.leetcode.com/uploads/2020/12/01/tree2.jpg)

    Input: root = [5,1,4,null,null,3,6]
    Output: false
    Explanation: The root node's value is 5 but its right child's value is 4.

## Constraints:

 - The number of nodes in the tree is in the range `[1, 10`<sup>`4`</sup>].
 - `-2`<sup>`31`</sup>` <= Node.val <= 2`<sup>`31`</sup>` - 1`

---

# Solution

## Intuition
The problem is asking us to verify if a given binary tree is a valid Binary Search Tree (BST). A BST has the property that for any given node, all nodes in its left subtree are less than the node, and all nodes in its right subtree are greater than the node. In addition, both the left and right subtrees must also be BSTs. This leads us to think about a recursive approach, where we check each node and its subtrees.

## Approach
The basic idea is to perform a depth-first search (DFS) on the tree, starting from the root. For each node, we check if its value lies within the valid range (which initially is from negative infinity to positive infinity). If it does, we recursively validate its left and right subtrees, but with updated ranges. For the left subtree, the maximum valid value becomes the current node's value. For the right subtree, the minimum valid value becomes the current node's value. If at any point a node's value does not lie within the valid range or any subtree is invalid, we return false, otherwise, we return true.

## Complexity
- Time complexity:
  O(n), where n is the number of nodes in the tree. We need to visit each node once to check if it's valid.

- Space complexity:
  O(h), where h is the height of the tree. This is the maximum depth of the recursion, which is equal to the height of the tree in the worst case (when the tree is skewed).

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)