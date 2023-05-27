# 1164. Product Price at a Given Date

Level: `Medium`

https://leetcode.com/problems/product-price-at-a-given-date/

---

# Description

Table: `Products`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | product_id    | int     |
    | new_price     | int     |
    | change_date   | date    |
    +---------------+---------+
    (product_id, change_date) is the primary key of this table.
    Each row of this table indicates that the price of some product was changed to a new price at some date.

Write an SQL query to find the prices of all products on 2019-08-16. Assume the price of all products before any change
is 10.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Products table:
    +------------+-----------+-------------+
    | product_id | new_price | change_date |
    +------------+-----------+-------------+
    | 1          | 20        | 2019-08-14  |
    | 2          | 50        | 2019-08-14  |
    | 1          | 30        | 2019-08-15  |
    | 1          | 35        | 2019-08-16  |
    | 2          | 65        | 2019-08-17  |
    | 3          | 20        | 2019-08-18  |
    +------------+-----------+-------------+
    Output:
    +------------+-------+
    | product_id | price |
    +------------+-------+
    | 2          | 50    |
    | 1          | 35    |
    | 3          | 10    |
    +------------+-------+

---

# Solution

## Intuition

Given the **Products** table, the task requires us to find the prices of all products on a particular date, which is "
2019-08-16". The prices of all products are initially assumed to be 10 before any change. We can achieve this by running
a SQL query that checks for the most recent change that doesn't exceed the given date for each product, and if there's
no such change, we use the default price of 10.

## Approach

We can solve this problem using a subquery. For each product, we can select the price that was most recently changed on
or before the target date. If there is no such change, we use the default price.

1. Use the **SELECT DISTINCT** statement to get each **product_id** once.
2. Use a **CASE** statement to set the price to 10 if the minimum change date is greater than the target date ("
   2019-08-16").
3. Otherwise, use a subquery to get the **new_price** from the **Products** table where the **change_date** is less than
   or equal to the target date. This subquery should select only the most recent **change_date** for each product, which
   can be achieved by ordering by change_date in descending order and then using **LIMIT 1** to get only the first (most
   recent) row.
4. Use the **GROUP BY** clause to ensure each product appears once in the output.

## Complexity

- Time complexity:
  O(n^2). The time complexity is quadratic because for each product, we perform a separate query to find the most recent price change before the target date.

- Space complexity:
  O(n). The space complexity is linear with respect to the number of different products, because we need to store the output for each product.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)