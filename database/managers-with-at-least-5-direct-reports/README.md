# 570. Managers with at Least 5 Direct Reports

Level: `Medium`

https://leetcode.com/problems/managers-with-at-least-5-direct-reports/

---

# Description

Table: `Employee`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | id          | int     |
    | name        | varchar |
    | department  | varchar |
    | managerId   | int     |
    +-------------+---------+
    id is the primary key column for this table.
    Each row of this table indicates the name of an employee, their department, and the id of their manager.
    If managerId is null, then the employee does not have a manager.
    No employee will be the manager of themself.

Write an SQL query to report the managers with at least **five direct reports**.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Employee table:
    +-----+-------+------------+-----------+
    | id  | name  | department | managerId |
    +-----+-------+------------+-----------+
    | 101 | John  | A          | None      |
    | 102 | Dan   | A          | 101       |
    | 103 | James | A          | 101       |
    | 104 | Amy   | A          | 101       |
    | 105 | Anne  | A          | 101       |
    | 106 | Ron   | B          | 101       |
    +-----+-------+------------+-----------+
    Output:
    +------+
    | name |
    +------+
    | John |
    +------+

---

# Solution

## Intuition

Given the structure of the Employee table, we are looking for managers who have at least five direct reports. The direct
report relation is reflected by the managerId field of the Employee table. Therefore, we need to count how many times
the id of each employee appears as a managerId in the table.

## Approach

We can use SQL's GROUP BY and HAVING clauses to solve this problem. First, we join the Employee table with itself,
pairing the manager's information (in table e1) with their respective employees' information (in table e2) using the
managerId field. Then we group the result of the join by the manager's id and count the number of occurrences of each
manager's id in the joined table. We use the HAVING clause to filter out the managers who have fewer than five direct
reports.

## Complexity

- Time complexity:
  The time complexity of this query mainly depends on the underlying DBMS, its implementation, and the data
  distribution. However, in general, the JOIN operation can be expensive, as it might require to scan the table once for
  each of its rows, leading to an O(n^2) complexity. However, modern DBMS have optimizations that can significantly
  reduce this cost.

- Space complexity:
  Similarly, the space complexity also depends on the DBMS and the data distribution. The additional space is needed to
  store the intermediate result of the JOIN operation and the grouping operation.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)