# 1045. Customers Who Bought All Products

Level: `Medium`

https://leetcode.com/problems/customers-who-bought-all-products/

---

# Description

Table: `Customer`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | customer_id | int     |
    | product_key | int     |
    +-------------+---------+
    There is no primary key for this table. It may contain duplicates. customer_id is not NULL.
    product_key is a foreign key to Product table.

Table: `Product`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | product_key | int     |
    +-------------+---------+
    product_key is the primary key column for this table.

Write an SQL query to report the customer ids from the `Customer` table that bought all the products in the `Product`
table.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Customer table:
    +-------------+-------------+
    | customer_id | product_key |
    +-------------+-------------+
    | 1           | 5           |
    | 2           | 6           |
    | 3           | 5           |
    | 3           | 6           |
    | 1           | 6           |
    +-------------+-------------+
    Product table:
    +-------------+
    | product_key |
    +-------------+
    | 5           |
    | 6           |
    +-------------+
    Output:
    +-------------+
    | customer_id |
    +-------------+
    | 1           |
    | 3           |
    +-------------+
    Explanation:
    The customers who bought all the products (5 and 6) are customers with IDs 1 and 3.

---

# Solution

## Intuition

Given the requirements of the problem, we need to identify customers who have bought all the products. This suggests
that we will have to aggregate data on a per-customer basis and check if the number of unique products bought by each
customer matches the total number of products.

## Approach

To achieve this, we will use a combination of GROUP BY, HAVING, and subquery in SQL.

1. First, we will group the Customer table by **customer_id**.

2. For each group, we use the **HAVING** clause to filter out the groups based on a condition. In this case, we want to
   ensure that the count of distinct **product_key** is equal to the total number of products.

3. To get the total number of products, we run a subquery on the Product table to count all the products.

## Complexity

- Time complexity:
  This approach essentially goes through all entries in the Customer table once, so it has a time complexity of O(n),
  where n is the number of entries in the Customer table.

- Space complexity:
  SQL queries run on the database server and generally do not have a space complexity in the traditional sense. However,
  this query should be relatively light on memory usage as it is not creating any temporary tables or storing
  substantial intermediate results.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)