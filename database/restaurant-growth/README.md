# 1321. Restaurant Growth

Level: `Medium`

https://leetcode.com/problems/restaurant-growth/

---

# Description

Table: `Customer`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | customer_id   | int     |
    | name          | varchar |
    | visited_on    | date    |
    | amount        | int     |
    +---------------+---------+
    (customer_id, visited_on) is the primary key for this table.
    This table contains data about customer transactions in a restaurant.
    visited_on is the date on which the customer with ID (customer_id) has visited the restaurant.
    amount is the total paid by a customer.

You are the restaurant owner and you want to analyze a possible expansion (there will be at least one customer every
day).

Write an SQL query to compute the moving average of how much the customer paid in a seven days window (i.e., current
day + 6 days before). `average_amount` should be **rounded to two decimal places**.

Return result table ordered by `visited_on` **in ascending order**.

The query result format is in the following example.

## Example 1:

    Input:
    Customer table:
    +-------------+--------------+--------------+-------------+
    | customer_id | name         | visited_on   | amount      |
    +-------------+--------------+--------------+-------------+
    | 1           | Jhon         | 2019-01-01   | 100         |
    | 2           | Daniel       | 2019-01-02   | 110         |
    | 3           | Jade         | 2019-01-03   | 120         |
    | 4           | Khaled       | 2019-01-04   | 130         |
    | 5           | Winston      | 2019-01-05   | 110         |
    | 6           | Elvis        | 2019-01-06   | 140         |
    | 7           | Anna         | 2019-01-07   | 150         |
    | 8           | Maria        | 2019-01-08   | 80          |
    | 9           | Jaze         | 2019-01-09   | 110         |
    | 1           | Jhon         | 2019-01-10   | 130         |
    | 3           | Jade         | 2019-01-10   | 150         |
    +-------------+--------------+--------------+-------------+
    Output:
    +--------------+--------------+----------------+
    | visited_on   | amount       | average_amount |
    +--------------+--------------+----------------+
    | 2019-01-07   | 860          | 122.86         |
    | 2019-01-08   | 840          | 120            |
    | 2019-01-09   | 840          | 120            |
    | 2019-01-10   | 1000         | 142.86         |
    +--------------+--------------+----------------+
    Explanation:
    1st moving average from 2019-01-01 to 2019-01-07 has an average_amount of (100 + 110 + 120 + 130 + 110 + 140 + 150)/7 = 122.86
    2nd moving average from 2019-01-02 to 2019-01-08 has an average_amount of (110 + 120 + 130 + 110 + 140 + 150 + 80)/7 = 120
    3rd moving average from 2019-01-03 to 2019-01-09 has an average_amount of (120 + 130 + 110 + 140 + 150 + 80 + 110)/7 = 120
    4th moving average from 2019-01-04 to 2019-01-10 has an average_amount of (130 + 110 + 140 + 150 + 80 + 110 + 130 + 150)/7 = 142.86

---

# Solution

## Intuition

The problem requires calculating a moving average over a window of seven days for the total amount paid by customers.
SQL can calculate this by using a self-join and built-in functions to compute the average.

## Approach

The general idea here is to join the **Customer** table to itself based on a condition that aligns the dates within a
7-day window.

1. First, a self-join is created using **JOIN**, connecting **Customer** with itself (**Customer AS c1 JOIN Customer AS
   c2**). The condition **c2.visited_on BETWEEN DATE_SUB(c1.visited_on, INTERVAL 6 DAY) AND c1.visited_on** makes sure
   to only join rows where the date **c2.visited_on** is within a 7-day window of **c1.visited_on**.

2. The **WHERE** clause is used to exclude rows where the 7-day window would not be full, i.e., **c1.visited_on** must
   be at least 6 days after the first visit in the records. This is accomplished by **c1.visited_on >= DATE_ADD((SELECT
   MIN(visited_on) FROM CUSTOMER), INTERVAL 6 DAY)**.

3. Then, the total amount and the average amount for each date are calculated using the **SUM** and **ROUND** functions,
   respectively. The average is rounded to 2 decimal places as required.

4. The **GROUP BY** clause groups the result by **c1.visited_on**, **c1.customer_id** because for each unique day and
   customer id pair, we need a separate row in the output. The grouping allows us to calculate the sum and average of
   the amount column for each of these groups.

5. Finally, the **ORDER BY** clause sorts the output by the **visited_on** column in ascending order.

## Complexity

- Time complexity:
  The query's time complexity is approximately O(n^2), where n is the number of rows in the Customer table. This is
  because for each row in the table, we potentially need to examine every other row to find those that are within the
  7-day window.

- Space complexity:
  The space complexity of the query is O(n), where n is the number of rows in the Customer table. This is because the
  result set could potentially include every row from the Customer table.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)