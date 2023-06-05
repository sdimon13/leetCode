# 180. Consecutive Numbers

Level: `Medium`

https://leetcode.com/problems/consecutive-numbers/

---

# Description

Table: `Logs`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | id          | int     |
    | num         | varchar |
    +-------------+---------+
    id is the primary key for this table.
    id is an autoincrement column.

Write an SQL query to find all numbers that appear at least three times consecutively.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Logs table:
    +----+-----+
    | id | num |
    +----+-----+
    | 1  | 1   |
    | 2  | 1   |
    | 3  | 1   |
    | 4  | 2   |
    | 5  | 1   |
    | 6  | 2   |
    | 7  | 2   |
    +----+-----+
    Output:
    +-----------------+
    | ConsecutiveNums |
    +-----------------+
    | 1               |
    +-----------------+
    Explanation: 1 is the only number that appears consecutively for at least three times.

---

# Solution

## Intuition

The problem is asking us to find all numbers that appear consecutively at least three times. Since the id field is
auto-incremented, we can assume that the logs are already sorted in the order they were inserted. Therefore, we can use
the SQL window function **LAG** to look at previous rows and compare their 'num' values with the current row.

## Approach

The solution involves writing a subquery which adds two new columns to the Logs table: 'prev_num' and 'prev_prev_num'.
These represent the 'num' value of the previous row and the row before the previous one respectively. This is done by
using the **LAG** window function.

After obtaining this table with the two new columns, we then look for rows where the current 'num' value is the same
as 'prev_num' and 'prev_prev_num'. If both conditions are true, this means that the current number appears consecutively
three times. We use **SELECT DISTINCT** to ensure that we don't return duplicate numbers.

## Complexity

- Time complexity:
  The query involves two passes over the data: one for calculating the lagged values, and another for the comparison.
  Since each pass operates on every row once, the time complexity is O(n), where n is the number of rows in the Logs
  table.

- Space complexity:
  The space complexity is also O(n), because the intermediate result of the subquery with the lagged values includes
  every row from the original table, potentially with some extra columns.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)