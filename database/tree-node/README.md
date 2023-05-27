# 608. Tree Node

Level: `Medium`

https://leetcode.com/problems/tree-node/

---

# Description

Table: `Tree`

    +-------------+------+
    | Column Name | Type |
    +-------------+------+
    | id          | int  |
    | p_id        | int  |
    +-------------+------+
    id is the primary key column for this table.
    Each row of this table contains information about the id of a node and the id of its parent node in a tree.
    The given structure is always a valid tree.

Each node in the tree can be one of three types:

- **"Leaf"**: if the node is a leaf node.
- **"Root"**: if the node is the root of the tree.
- **"Inner"**: If the node is neither a leaf node nor a root node.
  Write an SQL query to report the type of each node in the tree.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

![Example 1](https://assets.leetcode.com/uploads/2021/10/22/tree1.jpg)

    Input:
    Tree table:
    +----+------+
    | id | p_id |
    +----+------+
    | 1  | null |
    | 2  | 1    |
    | 3  | 1    |
    | 4  | 2    |
    | 5  | 2    |
    +----+------+
    Output:
    +----+-------+
    | id | type  |
    +----+-------+
    | 1  | Root  |
    | 2  | Inner |
    | 3  | Leaf  |
    | 4  | Leaf  |
    | 5  | Leaf  |
    +----+-------+
    Explanation:
    Node 1 is the root node because its parent node is null and it has child nodes 2 and 3.
    Node 2 is an inner node because it has parent node 1 and child node 4 and 5.
    Nodes 3, 4, and 5 are leaf nodes because they have parent nodes and they do not have child nodes.

## Example 2:

![Example 2](https://assets.leetcode.com/uploads/2021/10/22/tree2.jpg)

    Input: 
    Tree table:
    +----+------+
    | id | p_id |
    +----+------+
    | 1  | null |
    +----+------+
    Output:
    +----+-------+
    | id | type  |
    +----+-------+
    | 1  | Root  |
    +----+-------+
    Explanation: If there is only one node on the tree, you only need to output its root attributes.

---

# Solution

## Intuition

We have a tree structured database where each row represents a node of the tree. The column **id** is the node's own id
and **p_id** is the id of its parent. We are to classify each node as **Root** if it's the root node, **Leaf** if it's a
leaf node (i.e., it has a parent but no children), or **Inner** if it's neither root nor leaf (i.e., it has both a
parent and children).

## Approach

Given that **p_id** indicates the parent of a node, the root will be the only node that has a **NULL p_id**. We can thus
directly identify it using a **CASE** clause to check for a **NULL p_id**.

For the other two types, we have to look at the **p_id** column as well. Nodes that appear in the p_id column are nodes
that have children, and are thus either **Root** or **Inner** nodes. We have already identified **Root**, so now, if a
node appears in the **p_id** column and is not **Root**, it must be an **Inner** node.

By elimination, any remaining nodes must be **Leaf** nodes. These are nodes that do not appear in the **p_id** column (
i.e., they have no children), but do have a **non-NULL p_id** (i.e., they have a parent).

## Complexity

- Time complexity:
  The time complexity of this query is O(n^2), as we're performing a subquery for each node in the tree to check if its **id** appears in the **p_id** column.

- Space complexity:
  The space complexity of this query is O(1), as we're not using any additional storage.


---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)