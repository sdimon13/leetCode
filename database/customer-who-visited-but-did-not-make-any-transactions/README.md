# 1581. Customer Who Visited but Did Not Make Any Transactions

Level: `Easy`

https://leetcode.com/problems/customer-who-visited-but-did-not-make-any-transactions/

---

# Description

Table: `Visits`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | visit_id    | int     |
    | customer_id | int     |
    +-------------+---------+
    visit_id is the primary key for this table.
    This table contains information about the customers who visited the mall.

Table: `Transactions`

    +----------------+---------+
    | Column Name    | Type    |
    +----------------+---------+
    | transaction_id | int     |
    | visit_id       | int     |
    | amount         | int     |
    +----------------+---------+
    transaction_id is the primary key for this table.
    This table contains information about the transactions made during the visit_id.

Write a SQL query to find the IDs of the users who visited without making any transactions and the number of times they
made these types of visits.

Return the result table sorted in **any order**.

The query result format is in the following example.

## Example 1:

    Input: 
    Visits
    +----------+-------------+
    | visit_id | customer_id |
    +----------+-------------+
    | 1        | 23          |
    | 2        | 9           |
    | 4        | 30          |
    | 5        | 54          |
    | 6        | 96          |
    | 7        | 54          |
    | 8        | 54          |
    +----------+-------------+
    Transactions
    +----------------+----------+--------+
    | transaction_id | visit_id | amount |
    +----------------+----------+--------+
    | 2              | 5        | 310    |
    | 3              | 5        | 300    |
    | 9              | 5        | 200    |
    | 12             | 1        | 910    |
    | 13             | 2        | 970    |
    +----------------+----------+--------+
    Output:
    +-------------+----------------+
    | customer_id | count_no_trans |
    +-------------+----------------+
    | 54          | 2              |
    | 30          | 1              |
    | 96          | 1              |
    +-------------+----------------+
    Explanation:
    Customer with id = 23 visited the mall once and made one transaction during the visit with id = 12.
    Customer with id = 9 visited the mall once and made one transaction during the visit with id = 13.
    Customer with id = 30 visited the mall once and did not make any transactions.
    Customer with id = 54 visited the mall three times. During 2 visits they did not make any transactions, and during one visit they made 3 transactions.
    Customer with id = 96 visited the mall once and did not make any transactions.
    As we can see, users with IDs 30 and 96 visited the mall one time without making any transactions. Also, user 54 visited the mall twice and did not make any transactions.

---

# Solution

## Intuition

The problem is about finding the IDs of customers who visited but didn't make any transactions. For this, we have two
tables, **Visits** and **Transactions**. The idea is to join these two tables based on **visit_id** and then find the
rows where **transaction_id** is NULL. This indicates that the visit didn't result in a transaction. Counting these
occurrences for each customer should give us the required result.

## Approach

- Use a **LEFT JOIN** on the tables **Visits** and **Transactions** on **visit_id**. This gives us all visits, including
  those that didn't result in a transaction (which are represented as NULL in the joined table).
- Use a **WHERE** clause to filter out the visits that didn't have a corresponding transaction (i.e., **transaction_id**
  is NULL).
- Use **GROUP BY** to group these visits by **customer_id**.
- Use **COUNT** to get the number of such visits for each customer.
- Select **customer_id** and the count as **count_no_trans**.

## Complexity

- Time complexity:
  The time complexity is O(N) where N is the number of rows in the Visits table. This is because we are effectively
  scanning through all the rows in the Visits table in our SQL query.

- Space complexity:
  The space complexity is also O(N) because in the worst case scenario, every visit might not have a transaction, so we
  might need space to store all these visits in the output.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)