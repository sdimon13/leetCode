# 1174. Immediate Food Delivery II

Level: `Medium`

https://leetcode.com/problems/immediate-food-delivery-ii/

---

# Description

Table: `Delivery`

    +-----------------------------+---------+
    | Column Name                 | Type    |
    +-----------------------------+---------+
    | delivery_id                 | int     |
    | customer_id                 | int     |
    | order_date                  | date    |
    | customer_pref_delivery_date | date    |
    +-----------------------------+---------+
    delivery_id is the primary key of this table.
    The table holds information about food delivery to customers that make orders at some date and specify a preferred delivery date (on the same order date or after it).

If the customer's preferred delivery date is the same as the order date, then the order is called **immediate**;
otherwise, it is called **scheduled**.

The **first order** of a customer is the order with the earliest order date that the customer made. It is guaranteed
that a customer has precisely one first order.

Write an SQL query to find the percentage of immediate orders in the first orders of all customers, **rounded to 2
decimal places**.

The query result format is in the following example.

## Example 1:

    Input:
    Delivery table:
    +-------------+-------------+------------+-----------------------------+
    | delivery_id | customer_id | order_date | customer_pref_delivery_date |
    +-------------+-------------+------------+-----------------------------+
    | 1           | 1           | 2019-08-01 | 2019-08-02                  |
    | 2           | 2           | 2019-08-02 | 2019-08-02                  |
    | 3           | 1           | 2019-08-11 | 2019-08-12                  |
    | 4           | 3           | 2019-08-24 | 2019-08-24                  |
    | 5           | 3           | 2019-08-21 | 2019-08-22                  |
    | 6           | 2           | 2019-08-11 | 2019-08-13                  |
    | 7           | 4           | 2019-08-09 | 2019-08-09                  |
    +-------------+-------------+------------+-----------------------------+
    Output:
    +----------------------+
    | immediate_percentage |
    +----------------------+
    | 50.00                |
    +----------------------+
    Explanation:
    The customer id 1 has a first order with delivery id 1 and it is scheduled.
    The customer id 2 has a first order with delivery id 2 and it is immediate.
    The customer id 3 has a first order with delivery id 5 and it is scheduled.
    The customer id 4 has a first order with delivery id 7 and it is immediate.
    Hence, half the customers have immediate first orders.

---

# Solution

## Intuition

Given the nature of the problem, it appears we need to focus on two main points:

1. Identify the first order of each customer.
2. Check if the first order is immediate or scheduled.

## Approach

The first order for each customer is the one with the earliest **order_date**. We can identify these orders by grouping
by **customer_id** and selecting the minimum **order_date**.

We can then check if these first orders are immediate or not by comparing the **order_date** to the *
*customer_pref_delivery_date**. If they are the same, the order is immediate.

We then take the sum of these results, divide by the total number of customers, and multiply by 100 to get the
percentage of immediate first orders.

## Complexity

- Time complexity:
  The time complexity of this solution is O(n), where n is the number of rows in the Delivery table. This is because we
  need to scan each row at least once.

- Space complexity:
  The space complexity is also O(n), due to the potential need to store the first order date for each customer in the
  temporary result set.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)