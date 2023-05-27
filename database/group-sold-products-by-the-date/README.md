# 1484. Group Sold Products By The Date

Level: `Easy`

https://leetcode.com/problems/fix-names-in-a-table/

---

# Description

Table `Activities`:

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | sell_date   | date    |
    | product     | varchar |
    +-------------+---------+
    There is no primary key for this table, it may contain duplicates.
    Each row of this table contains the product name and the date it was sold in a market.

Write an SQL query to find for each date the number of different products sold and their names.

The sold products names for each date should be sorted lexicographically.

Return the result table ordered by `sell_date`.

The query result format is in the following example.

## Example 1:

    Input:
    Activities table:
    +------------+------------+
    | sell_date  | product     |
    +------------+------------+
    | 2020-05-30 | Headphone  |
    | 2020-06-01 | Pencil     |
    | 2020-06-02 | Mask       |
    | 2020-05-30 | Basketball |
    | 2020-06-01 | Bible      |
    | 2020-06-02 | Mask       |
    | 2020-05-30 | T-Shirt    |
    +------------+------------+
    Output:
    +------------+----------+------------------------------+
    | sell_date  | num_sold | products                     |
    +------------+----------+------------------------------+
    | 2020-05-30 | 3        | Basketball,Headphone,T-shirt |
    | 2020-06-01 | 2        | Bible,Pencil                 |
    | 2020-06-02 | 1        | Mask                         |
    +------------+----------+------------------------------+
    Explanation:
    For 2020-05-30, Sold items were (Headphone, Basketball, T-shirt), we sort them lexicographically and separate them by a comma.
    For 2020-06-01, Sold items were (Pencil, Bible), we sort them lexicographically and separate them by a comma.
    For 2020-06-02, the Sold item is (Mask), we just return it.

---

# Solution

## Intuition

Given a table of sales data containing a date and a product, the problem asks for each date, the number of different
products sold and their names, sorted lexicographically. SQL is well-suited for these types of operations as it provides
**GROUP BY** clause for grouping results by a particular column (in this case, **sell_date**), and aggregation functions
like **COUNT** for counting distinct items and **GROUP_CONCAT** for concatenating strings.

## Approach

- The query groups the data by **sell_date** using the **GROUP BY** clause.
- Then, for each date, it counts the number of distinct products sold using **COUNT(DISTINCT(product))**. This gives the
  total number of unique products sold each day.
- To get the product names, it concatenates the distinct product names lexicographically using **GROUP_CONCAT(DISTINCT
  product ORDER BY product separator ',')**. This provides a comma-separated list of unique product names sold each day,
  sorted lexicographically.
- Finally, the query orders the result by **sell_date** to ensure the output is sorted by date.

## Complexity

- Time complexity:
  The query involves a group by operation, which requires a full scan of the table. Assuming there are n records in the
  table, the time complexity is O(n).

- Space complexity:
  The result set will contain as many rows as there are unique sell dates. In the worst case scenario, each row in the
  table has a unique sell date, so the space complexity is also O(n).

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)