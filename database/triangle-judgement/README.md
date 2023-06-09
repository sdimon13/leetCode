# 610. Triangle Judgement

Level: `Easy`

https://leetcode.com/problems/triangle-judgement/

---

# Description

Table: `Triangle`

    +-------------+------+
    | Column Name | Type |
    +-------------+------+
    | x           | int  |
    | y           | int  |
    | z           | int  |
    +-------------+------+
    (x, y, z) is the primary key column for this table.
    Each row of this table contains the lengths of three line segments.

Write an SQL query to report for every three line segments whether they can form a triangle.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Triangle table:
    +----+----+----+
    | x  | y  | z  |
    +----+----+----+
    | 13 | 15 | 30 |
    | 10 | 20 | 15 |
    +----+----+----+
    Output:
    +----+----+----+----------+
    | x  | y  | z  | triangle |
    +----+----+----+----------+
    | 13 | 15 | 30 | No       |
    | 10 | 20 | 15 | Yes      |
    +----+----+----+----------+

---

# Solution

## Intuition

Given three line segments, they can form a triangle if and only if the sum of the lengths of any two segments is greater
than the length of the remaining one. This is known as the triangle inequality theorem. In our case, we have to check if
x + y > z, x + z > y, and y + z > x.

## Approach

We can leverage the CASE statement in SQL to perform the triangle inequality checks directly in the query. For each row
in the table, we will check if the lengths satisfy the triangle inequality theorem. If they do, we will return 'Yes',
otherwise 'No'. This information will be returned in a new column named 'triangle'.

## Complexity

- Time complexity:
  The time complexity is O(n), where n is the number of rows in the table. This is because we have to go through each
  row to perform the triangle inequality theorem checks.

- Space complexity:
  The space complexity is O(1), since we are not using any additional space that scales with the input size. The result
  of our query is directly written into the output and does not require additional space.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)