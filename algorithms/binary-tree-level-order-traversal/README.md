# 102. Binary Tree Level Order Traversal
Level: `Medium`

https://leetcode.com/problems/binary-tree-level-order-traversal/

---

# Description

Given the `root` of a binary tree, return the level order traversal of its nodes' values. (i.e., from left to right, level by level).

## Example 1:

![Example 1](https://assets.leetcode.com/uploads/2021/02/19/tree1.jpg)

    Input: root = [3,9,20,null,null,15,7]
    Output: [[3],[9,20],[15,7]]

## Example 2:

    Input: root = [1]
    Output: [[1]]

## Example 3:

    Input: root = []
    Output: []

## Constraints:

 - The number of nodes in the tree is in the range `[0, 2000]`.

 - `-1000 <= Node.val <= 1000`

---

# Solution

## Intuition
The initial thought was to use Breadth-First Search (BFS) to solve this problem. BFS is perfectly suited for traversing the levels of a tree, starting from the root.

## Approach
To implement BFS, a queue data structure was used. Initially, the root node is added to the queue. Then, in a loop, nodes of one level are dequeued and their values are stored in a temporary slice. The child nodes of each dequeued node are enqueued for later processing. The process continues until all nodes in the tree have been processed.

## Complexity
- Time complexity: O(n), where n is the number of nodes in the tree. Each node is processed exactly once.


- Space complexity: O(n), where n is the number of nodes in the tree. In the worst-case scenario, when the tree is a complete binary tree, all nodes at the most populated level might be in the queue at the same time, which makes up approximately n/2 nodes.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)