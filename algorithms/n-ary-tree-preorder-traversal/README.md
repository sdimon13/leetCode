# 589. N-ary Tree Preorder Traversal
Level: `Easy`

https://leetcode.com/problems/n-ary-tree-preorder-traversal/

---

# Description

Given the `root` of an n-ary tree, return _the preorder traversal of its nodes' values_.

Nary-Tree input serialization is represented in their level order traversal. Each group of children is separated by the null value (See examples)

## Example 1:

![Example 1](https://assets.leetcode.com/uploads/2018/10/12/narytreeexample.png)

    Input: root = [1,null,3,2,4,null,5,6]
    Output: [1,3,5,6,2,4]

## Example 2:

![Example 2](https://assets.leetcode.com/uploads/2019/11/08/sample_4_964.png)

    Input: root = [1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
    Output: [1,2,3,6,7,11,14,4,8,12,5,9,13,10]

## Constraints:

 - The number of nodes in the tree is in the range `[0, 10`<sup>`4`</sup>].

 - `0 <= Node.val <= 10`<sup>`4`</sup>

 - The height of the n-ary tree is less than or equal to `1000`.

---

# Solution

## Intuition
The problem requires us to perform a pre-order traversal of an N-ary tree. In a pre-order traversal, we visit the root node first, then visit its children from left to right. The two main approaches to tree traversal are recursive and iterative.

## Approach
1. Recursive Approach: We start with the root node, add its value to the result, and then recursively process each of its children. This process continues until we've processed all nodes in the tree.

2. Iterative Approach: We use a stack to manually implement the process that happens automatically in the call stack during recursion. We start by pushing the root node onto the stack. Then, we enter a loop that continues as long as the stack is not empty. In each iteration of the loop, we pop a node from the stack, add its value to the result, and then push its children onto the stack in reverse order (rightmost child first). This ensures that when we pop the next node from the stack, it will be the leftmost child, which is what we want for pre-order traversal.

## Complexity
Recursive Approach

- Time complexity: O(n), where n is the total number of nodes in the tree. This is because we visit each node exactly once.
- Space complexity: O(n), in the worst case, if the tree is very unbalanced and looks more like a linked list, the call stack could potentially hold all n nodes in the call stack at the same time.

Iterative Approach

- Time complexity: O(n), where n is the total number of nodes in the tree. This is because we visit each node exactly once.
- Space complexity: O(n), in the worst case, if the tree is very unbalanced, the stack could potentially hold all n nodes at the same time.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)