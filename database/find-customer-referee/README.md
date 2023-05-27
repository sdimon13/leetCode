# 584. Find Customer Referee

Level: `Easy`

https://leetcode.com/problems/find-customer-referee/

---

# Description

Table: `Customer`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | id          | int     |
    | name        | varchar |
    | referee_id  | int     |
    +-------------+---------+
    id is the primary key column for this table.
    Each row of this table indicates the id of a customer, their name, and the id of the customer who referred them.


Write an SQL query to report the names of the customer that are **not referred by** the customer with `id = 2`.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Customer table:
    +----+------+------------+
    | id | name | referee_id |
    +----+------+------------+
    | 1  | Will | null       |
    | 2  | Jane | null       |
    | 3  | Alex | 2          |
    | 4  | Bill | null       |
    | 5  | Zack | 1          |
    | 6  | Mark | 2          |
    +----+------+------------+
    Output:
    +------+
    | name |
    +------+
    | Will |
    | Jane |
    | Bill |
    | Zack |
    +------+

---

# Solution

## Intuition
Given the structure of the Customer table, we want to retrieve all customers who were not referred by the customer with id = 2. In SQL, this involves a simple **SELECT** query with a condition that checks the **referee_id**.

## Approach
We need to select all customers (name) where their referee_id is not equal to 2 or is null. This can be achieved by using the SELECT command to select the name column from the Customer table with a WHERE clause that checks if referee_id is not equal to 2 (referee_id != 2) or referee_id is null (referee_id IS NULL).

## Complexity
- Time complexity:
  O(n), where n is the number of rows in the Customer table. The database needs to scan each row to evaluate the WHERE clause.

- Space complexity:
  O(1), the space complexity does not depend on the size of the input table, and only a constant amount of space is used.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)