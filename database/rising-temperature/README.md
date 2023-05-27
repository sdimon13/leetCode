# 197. Rising Temperature

Level: `Easy`

https://leetcode.com/problems/rising-temperature/

---

# Description

Table: `Weather`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | id            | int     |
    | recordDate    | date    |
    | temperature   | int     |
    +---------------+---------+
    id is the primary key for this table.
    This table contains information about the temperature on a certain day.

Write an SQL query to find all dates' `Id` with higher temperatures compared to its previous dates (yesterday).

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Weather table:
    +----+------------+-------------+
    | id | recordDate | temperature |
    +----+------------+-------------+
    | 1  | 2015-01-01 | 10          |
    | 2  | 2015-01-02 | 25          |
    | 3  | 2015-01-03 | 20          |
    | 4  | 2015-01-04 | 30          |
    +----+------------+-------------+
    Output:
    +----+
    | id |
    +----+
    | 2  |
    | 4  |
    +----+
    Explanation:
    In 2015-01-02, the temperature was higher than the previous day (10 -> 25).
    In 2015-01-04, the temperature was higher than the previous day (20 -> 30).

---

# Solution

## Intuition

To solve this problem, we need to find the cases where the temperature on a specific day is higher than the temperature
on the previous day. This requires us to compare the temperature of each record with the temperature of the record of
the previous day.

## Approach

The problem can be solved by self-joining the Weather table. For every record in the Weather table, we look for another
record with a **recordDate** exactly one day before and the temperature lower than the current record.

Here, we make use of the **DATE_ADD** function to get the date of the previous day and then compare the temperatures. If
the current day's temperature is higher than the previous day's, the id of that record is selected.

## Complexity

- Time complexity:
  The time complexity for this problem would be dependent on the underlying SQL engine's ability to join the table to
  itself. In the worst case, this could potentially be O(n<sup>2</sup>), where n is the number of records in the table.

- Space complexity:
  The space complexity is typically not considered for SQL queries. But for completeness, this SQL query doesn't require
  additional space that grows with input size, so it can be considered to have O(1) space complexity.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)