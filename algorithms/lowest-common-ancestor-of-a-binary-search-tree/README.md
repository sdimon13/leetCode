# 235. Lowest Common Ancestor of a Binary Search Tree
Level: `Medium`

https://leetcode.com/problems/lowest-common-ancestor-of-a-binary-search-tree

---

# Description

Given a binary search tree (BST), find the lowest common ancestor (LCA) node of two given nodes in the BST.

According to the definition of LCA on Wikipedia: “The lowest common ancestor is defined between two nodes `p` and `q` as the lowest node in `T` that has both `p` and `q` as descendants (where we allow **a node to be a descendant of itself**).”

## Example 1:

![Example 1](https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png)

    Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 8
    Output: 6
    Explanation: The LCA of nodes 2 and 8 is 6.

## Example 2:

![Example 2](https://assets.leetcode.com/uploads/2018/12/14/binarysearchtree_improved.png)

    Input: root = [6,2,8,0,4,7,9,null,null,3,5], p = 2, q = 4
    Output: 2
    Explanation: The LCA of nodes 2 and 4 is 2, since a node can be a descendant of itself according to the LCA definition.

## Example 3:

    Input: root = [2,1], p = 2, q = 1
    Output: 2

## Constraints:

- The number of nodes in the tree is in the range `[2, 10`<sup>`5`</sup>`]`.
- `-10`<sup>`9`</sup>` <= Node.val <= 10`<sup>`9`</sup>
- All `Node.val` are **unique**.
- `p != q`
- `p` and `q` will exist in the BST.

---

# Solution

## Intuition
Given the nature of a Binary Search Tree (BST), we can leverage its properties to find the lowest common ancestor (LCA). The LCA of two nodes p and q in a BST is the node where one node is in the left subtree and the other is in the right subtree, or when one of the nodes is itself the LCA.

## Approach
We start from the root node, then traverse the BST based on the following rules:

- If the values of both p and q are less than the value of the current node, we move to the left child of the current node.
- If the values of both p and q are greater than the value of the current node, we move to the right child of the current node.
- Otherwise, the current node is the LCA of p and q.

This approach works because in a BST, all nodes in the left subtree are less than the root, and all nodes in the right subtree are greater than the root.

## Complexity
- Time complexity:
  O(H), where H is the height of the tree. In the worst case, we might have to visit all nodes from the root to the deepest leaf node.

- Space complexity:
  O(H), to account for the stack space during the recursive function calls. In the worst case, a skewed tree, H equals N, where N is the number of nodes, therefore the space complexity could be O(N).

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.go)
- [Test Code](./solution_test.go)