# 1327. List the Products Ordered in a Period

Level: `Easy`

https://leetcode.com/problems/list-the-products-ordered-in-a-period/

---

# Description

Table: `Products`

    +------------------+---------+
    | Column Name      | Type    |
    +------------------+---------+
    | product_id       | int     |
    | product_name     | varchar |
    | product_category | varchar |
    +------------------+---------+
    product_id is the primary key for this table.
    This table contains data about the company's products.

Table: `Orders`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | product_id    | int     |
    | order_date    | date    |
    | unit          | int     |
    +---------------+---------+
    There is no primary key for this table. It may have duplicate rows.
    product_id is a foreign key to the Products table.
    unit is the number of products ordered in order_date.

Write an SQL query to get the names of products that have at least `100` units ordered in **February 2020** and their
amount.

Return result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Products table:
    +-------------+-----------------------+------------------+
    | product_id  | product_name          | product_category |
    +-------------+-----------------------+------------------+
    | 1           | Leetcode Solutions    | Book             |
    | 2           | Jewels of Stringology | Book             |
    | 3           | HP                    | Laptop           |
    | 4           | Lenovo                | Laptop           |
    | 5           | Leetcode Kit          | T-shirt          |
    +-------------+-----------------------+------------------+
    Orders table:
    +--------------+--------------+----------+
    | product_id   | order_date   | unit     |
    +--------------+--------------+----------+
    | 1            | 2020-02-05   | 60       |
    | 1            | 2020-02-10   | 70       |
    | 2            | 2020-01-18   | 30       |
    | 2            | 2020-02-11   | 80       |
    | 3            | 2020-02-17   | 2        |
    | 3            | 2020-02-24   | 3        |
    | 4            | 2020-03-01   | 20       |
    | 4            | 2020-03-04   | 30       |
    | 4            | 2020-03-04   | 60       |
    | 5            | 2020-02-25   | 50       |
    | 5            | 2020-02-27   | 50       |
    | 5            | 2020-03-01   | 50       |
    +--------------+--------------+----------+
    Output:
    +--------------------+---------+
    | product_name       | unit    |
    +--------------------+---------+
    | Leetcode Solutions | 130     |
    | Leetcode Kit       | 100     |
    +--------------------+---------+
    Explanation:
    Products with product_id = 1 is ordered in February a total of (60 + 70) = 130.
    Products with product_id = 2 is ordered in February a total of 80.
    Products with product_id = 3 is ordered in February a total of (2 + 3) = 5.
    Products with product_id = 4 was not ordered in February 2020.
    Products with product_id = 5 is ordered in February a total of (50 + 50) = 100.

---

# Solution

## Intuition

The task asks for products that were ordered at least 100 times in February 2020. We are given two tables - **Products**
and **Orders** - that can be linked by the **product_id** field. Our strategy here is to join these tables, then filter
for orders from February 2020, sum the units ordered for each product, and finally filter out products with less than
100 total units ordered.

## Approach

To solve this task, we can apply the following steps:

1. Perform a **LEFT JOIN** on the **Products** and **Orders** tables using product_id as the joining field. The LEFT
   JOIN ensures that all products are included in the result, even if they have no corresponding entries in the **Orders
   ** table.
2. Filter the results to include only orders from February 2020, using the **WHERE** clause to filter the **order_date**
   field.
3. Group the results by **product_id** using the **GROUP BY** clause. This is necessary because the same product may
   have multiple entries in the **Orders** table, and we want to consider the total units ordered for each product.
4. Calculate the sum of **unit** for each product using the **SUM** function.
5. Finally, use the **HAVING** clause to include only products where the total units ordered are at least 100.

## Complexity

- Time complexity:
  The time complexity is O(n), where n is the total number of rows in the Orders table, because the database has to go
  through all rows to apply the filter, grouping, and aggregation operations.

- Space complexity:
  The space complexity is O(m), where m is the number of distinct products that were ordered in February 2020, because
  that's the maximum number of rows the result can have.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)