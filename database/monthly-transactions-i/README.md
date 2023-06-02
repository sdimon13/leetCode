# 1193. Monthly Transactions I

Level: `Medium`

https://leetcode.com/problems/monthly-transactions-i/

---

# Description

Table: `Transactions`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | id            | int     |
    | country       | varchar |
    | state         | enum    |
    | amount        | int     |
    | trans_date    | date    |
    +---------------+---------+
    id is the primary key of this table.
    The table has information about incoming transactions.
    The state column is an enum of type ["approved", "declined"].

Write an SQL query to find for each month and country, the number of transactions and their total amount, the number of
approved transactions and their total amount.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Transactions table:
    +------+---------+----------+--------+------------+
    | id   | country | state    | amount | trans_date |
    +------+---------+----------+--------+------------+
    | 121  | US      | approved | 1000   | 2018-12-18 |
    | 122  | US      | declined | 2000   | 2018-12-19 |
    | 123  | US      | approved | 2000   | 2019-01-01 |
    | 124  | DE      | approved | 2000   | 2019-01-07 |
    +------+---------+----------+--------+------------+
    Output:
    +----------+---------+-------------+----------------+--------------------+-----------------------+
    | month    | country | trans_count | approved_count | trans_total_amount | approved_total_amount |
    +----------+---------+-------------+----------------+--------------------+-----------------------+
    | 2018-12  | US      | 2           | 1              | 3000               | 1000                  |
    | 2019-01  | US      | 1           | 1              | 2000               | 2000                  |
    | 2019-01  | DE      | 1           | 1              | 2000               | 2000                  |
    +----------+---------+-------------+----------------+--------------------+-----------------------+

---

# Solution

## Intuition

The task is to find for each month and country, the number of transactions and their total amount, the number of
approved transactions and their total amount. This is a SQL aggregation problem where we have to group the transactions
by country and month and then calculate various statistics for each group.

## Approach

To solve this problem, we can leverage SQL's powerful aggregation features, which allow us to perform calculations on
groups of rows in a table. Specifically, we will need to use the GROUP BY clause to divide the transactions into groups
based on their country and month.

To get the month from the trans_date, we can use the DATE_FORMAT function in MySQL. It formats a date as specified by a
format string. The format string supports a variety of placeholders that let you control how the date is formatted.

For each group, we then use the COUNT and SUM functions to calculate the total number of transactions, the total amount
of transactions, the total number of approved transactions, and the total amount of approved transactions. For counting
and summing up the approved transactions we will use conditional aggregation with the help of CASE WHEN statement. It
allows us to perform different calculations for different rows in a group.

## Complexity

- Time complexity:
  The time complexity is O(n) where n is the number of rows in the Transactions table, as the query needs to go through
  all rows once.

- Space complexity:
  The space complexity is O(m) where m is the number of unique country-month combinations in the Transactions table.
  This is because the result of the GROUP BY clause is a new table where each row corresponds to a unique country-month
  combination.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)