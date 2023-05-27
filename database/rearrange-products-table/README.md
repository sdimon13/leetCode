# 1795. Rearrange Products Table

Level: `Easy`

https://leetcode.com/problems/user-activity-for-the-past-30-days-i/

---

# Description

Table: `Products`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | product_id  | int     |
    | store1      | int     |
    | store2      | int     |
    | store3      | int     |
    +-------------+---------+
    product_id is the primary key for this table.
    Each row in this table indicates the product's price in 3 different stores: store1, store2, and store3.
    If the product is not available in a store, the price will be null in that store's column.

Write an SQL query to rearrange the `Products` table so that each row has (`product_id, store, price`). If a product is
not available in a store, do **not** include a row with that `product_id` and `store` combination in the result table.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Products table:
    +------------+--------+--------+--------+
    | product_id | store1 | store2 | store3 |
    +------------+--------+--------+--------+
    | 0          | 95     | 100    | 105    |
    | 1          | 70     | null   | 80     |
    +------------+--------+--------+--------+
    Output:
    +------------+--------+-------+
    | product_id | store  | price |
    +------------+--------+-------+
    | 0          | store1 | 95    |
    | 0          | store2 | 100   |
    | 0          | store3 | 105   |
    | 1          | store1 | 70    |
    | 1          | store3 | 80    |
    +------------+--------+-------+
    Explanation:
    Product 0 is available in all three stores with prices 95, 100, and 105 respectively.
    Product 1 is available in store1 with price 70 and store3 with price 80. The product is not available in store2.

---

# Solution

## Intuition

Given a table with products and their prices in three different stores, we are asked to reorganize this information into
a different format. Instead of having the prices in separate columns per store, we need to present each
product-store-price combination in a separate row. This is a classic case where we need to use SQL's UNION operator to
combine rows from multiple SELECT statements into a single result set.

## Approach

To achieve this, we need to perform three different SELECT operations, each one focusing on a different store column in
the Products table:

1. Select the product_id, a constant string `'store1'` as the store, and the store1 column as price from the Products
   table, but only where the store1 column is **not NULL**.
2. Repeat the above step for `'store2'` and `'store3'`.
   Since we need all the resulting rows from these three SELECT statements, we will use the **UNION** operator. This
   operator combines the result sets of two or more SELECT statements and removes duplicate rows. However, in our case,
   there should be no duplicate rows, as the product_id and store combination should be **unique**.

The ORDER BY clause is used at the end to sort the result set. In this case, it will sort the result first by
product_id (ORDER BY 1) and then by store name (ORDER BY 2).

## Complexity

- Time complexity:
  The time complexity is O(n), where n is the number of rows in the Products table. Each of the three SELECT operations
  scans the entire table.

- Space complexity:
  The space complexity is also O(n), where n is the number of rows in the Products table. The UNION operation combines
  the rows from the three SELECT operations into a single result set.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)