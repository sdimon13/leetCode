# 1251. Average Selling Price

Level: `Easy`

https://leetcode.com/problems/average-selling-price/

---

# Description

Table: `Prices`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | product_id    | int     |
    | start_date    | date    |
    | end_date      | date    |
    | price         | int     |
    +---------------+---------+
    (product_id, start_date, end_date) is the primary key for this table.
    Each row of this table indicates the price of the product_id in the period from start_date to end_date.
    For each product_id there will be no two overlapping periods. That means there will be no two intersecting periods for the same product_id.

Table: `UnitsSold`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | product_id    | int     |
    | purchase_date | date    |
    | units         | int     |
    +---------------+---------+
    There is no primary key for this table, it may contain duplicates.
    Each row of this table indicates the date, units, and product_id of each product sold.

Write an SQL query to find the average selling price for each product. `average_price` should be **rounded to 2 decimal
places**.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Prices table:
    +------------+------------+------------+--------+
    | product_id | start_date | end_date   | price  |
    +------------+------------+------------+--------+
    | 1          | 2019-02-17 | 2019-02-28 | 5      |
    | 1          | 2019-03-01 | 2019-03-22 | 20     |
    | 2          | 2019-02-01 | 2019-02-20 | 15     |
    | 2          | 2019-02-21 | 2019-03-31 | 30     |
    +------------+------------+------------+--------+
    UnitsSold table:
    +------------+---------------+-------+
    | product_id | purchase_date | units |
    +------------+---------------+-------+
    | 1          | 2019-02-25    | 100   |
    | 1          | 2019-03-01    | 15    |
    | 2          | 2019-02-10    | 200   |
    | 2          | 2019-03-22    | 30    |
    +------------+---------------+-------+
    Output:
    +------------+---------------+
    | product_id | average_price |
    +------------+---------------+
    | 1          | 6.96          |
    | 2          | 16.96         |
    +------------+---------------+
    Explanation:
    Average selling price = Total Price of Product / Number of products sold.
    Average selling price for product 1 = ((100 * 5) + (15 * 20)) / 115 = 6.96
    Average selling price for product 2 = ((200 * 15) + (30 * 30)) / 230 = 16.96

---

# Solution

## Intuition

This problem is about calculating the average selling price for each product. Given that we have the price and units
sold information in two different tables, we need to find a way to relate them and perform necessary calculations.

## Approach

To solve this problem, we can leverage SQL JOIN operations to combine the necessary data from the **Prices** and *
*UnitsSold** tables. We specifically require the **JOIN** operation because the information we need is split across two
tables.

The SQL query will be constructed as follows:

1. **JOIN** the two tables **Prices** and **UnitsSold** on the **product_id** column and ensure the purchase date is
   within the price effective period. This provides a combined view of the prices and units sold for each product.

2. Then, we **SUM** the product of **price** and **units** for each **product_id** (which gives us the total revenue for
   each product) and also **SUM** the **units** for each **product_id** (which gives us the total units sold for each
   product).

3. Next, we calculate the average selling price by dividing the total revenue by the total units sold for each product.

4. Finally, we round off the average selling price to 2 decimal places.

## Complexity

- Time complexity:
  The time complexity depends on the underlying SQL engine that is executing the JOIN, SUM and GROUP BY operations. In
  most cases, this is likely to be O(n), where n is the total number of rows in the joined table view.

- Space complexity:
  The space complexity is largely dependent on the size of the dataset. In the worst case, if every row from the first
  table matches every row from the second table (a Cartesian product), the space complexity could be O(n<sup>2</sup>),
  where n is the number of rows in each table.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)