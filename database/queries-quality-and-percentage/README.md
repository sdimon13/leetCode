# 1211. Queries Quality and Percentage

Level: `Easy`

https://leetcode.com/problems/queries-quality-and-percentage/

---

# Description

Table: `Queries`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | query_name  | varchar |
    | result      | varchar |
    | position    | int     |
    | rating      | int     |
    +-------------+---------+
    There is no primary key for this table, it may have duplicate rows.
    This table contains information collected from some queries on a database.
    The position column has a value from 1 to 500.
    The rating column has a value from 1 to 5. Query with rating less than 3 is a poor query.

We define query `quality` as:

The average of the ratio between query rating and its position.

We also define` poor query percentage` as:

The percentage of all queries with rating less than 3.

Write an SQL query to find each `query_name`, the quality and `poor_query_percentage`.

Both `quality` and `poor_query_percentage` should be **rounded to 2 decimal places**.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Queries table:
    +------------+-------------------+----------+--------+
    | query_name | result            | position | rating |
    +------------+-------------------+----------+--------+
    | Dog        | Golden Retriever  | 1        | 5      |
    | Dog        | German Shepherd   | 2        | 5      |
    | Dog        | Mule              | 200      | 1      |
    | Cat        | Shirazi           | 5        | 2      |
    | Cat        | Siamese           | 3        | 3      |
    | Cat        | Sphynx            | 7        | 4      |
    +------------+-------------------+----------+--------+
    Output:
    +------------+---------+-----------------------+
    | query_name | quality | poor_query_percentage |
    +------------+---------+-----------------------+
    | Dog        | 2.50    | 33.33                 |
    | Cat        | 0.66    | 33.33                 |
    +------------+---------+-----------------------+
    Explanation:
    Dog queries quality is ((5 / 1) + (5 / 2) + (1 / 200)) / 3 = 2.50
    Dog queries poor_ query_percentage is (1 / 3) * 100 = 33.33
    
    Cat queries quality equals ((2 / 5) + (3 / 3) + (4 / 7)) / 3 = 0.66
    Cat queries poor_ query_percentage is (1 / 3) * 100 = 33.33

---

# Solution

## Intuition

Given the problem, we can see that we need to calculate two values: the quality and the poor query percentage for each
query_name. The quality is the average ratio of the rating and position for each query, while the poor query percentage
is the percentage of queries that have a rating less than 3.

## Approach

Since we need to group results by the **query_name**, we will use the GROUP BY clause. For the **quality**, we will
calculate the average of the ratio between **rating** and **position**. For the **poor_query_percentage**, we calculate
the percentage of queries with a rating less than 3.

To achieve this, we will use conditional aggregation, where we sum up all the cases where the rating is less than 3,
divide by the total count of queries for each query_name and then multiply by 100 to convert to percentage. The ROUND
function is used to round these values to two decimal places as requested.

## Complexity

- Time complexity:
  The time complexity is O(n) because we need to iterate over all rows in the table, where n is the number of rows in
  the table.

- Space complexity:
  The space complexity is O(m) where m is the number of unique query_name entries in the table, as that will determine
  the number of groups, and thus, the amount of intermediate data stored for aggregation.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)