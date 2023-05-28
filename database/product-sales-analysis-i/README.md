# 1068. Product Sales Analysis I

Level: `Easy`

https://leetcode.com/problems/product-sales-analysis-i/

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

Write an SQL query that reports the `product_name`, `year`, and `price` for each `sale_id` in the `Sales` table.

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
    +--------------+-------+-------+
    | product_name | year  | price |
    +--------------+-------+-------+
    | Nokia        | 2008  | 5000  |
    | Nokia        | 2009  | 5000  |
    | Apple        | 2011  | 9000  |
    +--------------+-------+-------+
    Explanation:
    From sale_id = 1, we can conclude that Nokia was sold for 5000 in the year 2008.
    From sale_id = 2, we can conclude that Nokia was sold for 5000 in the year 2009.
    From sale_id = 7, we can conclude that Apple was sold for 9000 in the year 2011.

---

# Solution

## Intuition

This problem requires us to join two tables on a common field and then select specific fields from the joined table. The
common field between the **Sales** and **Product** tables is **product_id**. We are asked to report the **product_name
**, **year**, and **price** for each **sale_id** in the **Sales** table. The **product_name** is located in the *
*Product** table and the **year** and **price** are located in the **Sales** table.

## Approach

The approach is to perform a LEFT JOIN on the **Sales** table and the **Product** table, with the common field being *
*product_id**. This will give us a combined table from which we can select the required fields. We use a LEFT JOIN
because we want to return all rows from the **Sales** table, even if there is no corresponding entry in the **Product**
table.

## Complexity

- Time complexity:
  The time complexity of this query depends on the database management system and its optimization techniques. However,
  generally speaking, a JOIN operation has a time complexity of O(N*M), where N and M are the number of rows in the
  Sales and Product tables, respectively.

- Space complexity:
  The space complexity of this query is also dependent on the database management system. The space complexity is
  typically proportional to the amount of data returned by the query, so in this case, it would be the number of rows in
  the Sales table.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)