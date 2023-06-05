# 1070. Product Sales Analysis III

Level: `Medium`

https://leetcode.com/problems/product-sales-analysis-iii/

---

# Description

Table: `Sales`

    +-------------+-------+
    | Column Name | Type  |
    +-------------+-------+
    | sale_id     | int   |
    | product_id  | int   |
    | year        | int   |
    | quantity    | int   |
    | price       | int   |
    +-------------+-------+
    (sale_id, year) is the primary key of this table.
    product_id is a foreign key to Product table.
    Each row of this table shows a sale on the product product_id in a certain year.
    Note that the price is per unit.

Table: `Product`

    +--------------+---------+
    | Column Name  | Type    |
    +--------------+---------+
    | product_id   | int     |
    | product_name | varchar |
    +--------------+---------+
    product_id is the primary key of this table.
    Each row of this table indicates the product name of each product.

Write an SQL query that selects the **product id, year, quantity**, and **price** for the **first year** of every
product sold.

Return the resulting table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Sales table:
    +---------+------------+------+----------+-------+
    | sale_id | product_id | year | quantity | price |
    +---------+------------+------+----------+-------+
    | 1       | 100        | 2008 | 10       | 5000  |
    | 2       | 100        | 2009 | 12       | 5000  |
    | 7       | 200        | 2011 | 15       | 9000  |
    +---------+------------+------+----------+-------+
    Product table:
    +------------+--------------+
    | product_id | product_name |
    +------------+--------------+
    | 100        | Nokia        |
    | 200        | Apple        |
    | 300        | Samsung      |
    +------------+--------------+
    Output:
    +------------+------------+----------+-------+
    | product_id | first_year | quantity | price |
    +------------+------------+----------+-------+
    | 100        | 2008       | 10       | 5000  |
    | 200        | 2011       | 15       | 9000  |
    +------------+------------+----------+-------+

---

# Solution

## Intuition

The task requires us to find the first year of sale for each product. This is essentially asking us to retrieve the
product's minimum 'year' value in the Sales table. This is a SQL-based problem and can be solved using built-in SQL
functions and concepts such as JOIN, GROUP BY and MIN.

## Approach

We will be using a subquery in this problem. In our subquery, we will find the minimum 'year' for each 'product_id' in
the Sales table by grouping by 'product_id'. This subquery will return the earliest year each product was sold.

Then in the outer query, we join this subquery with the original Sales table. The join condition will be based on both '
product_id' and 'year'. The reason we join on both columns is because we want to filter the Sales table to only the rows
where the 'product_id' and 'year' match with the 'product_id' and 'first_year' in our subquery.

Finally, we select the 'product_id', 'year' (renamed as 'first_year'), 'quantity', and 'price' from this filtered Sales
table.

## Complexity

- Time complexity:
  The time complexity is O(n), where n is the number of rows in the Sales table. This is because we need to scan through
  the entire Sales table to find the earliest year for each product, and then again to join the Sales table with our
  subquery.

- Space complexity:
  The space complexity is O(n), where n is the number of unique 'product_id's. This space is used to store the results
  from our subquery, which consists of one row for each unique 'product_id'.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)