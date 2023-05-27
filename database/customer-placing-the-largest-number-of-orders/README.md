# 586. Customer Placing the Largest Number of Orders

Level: `Easy`

https://leetcode.com/problems/customer-placing-the-largest-number-of-orders/

---

# Description

Table: `Orders`

    +-----------------+----------+
    | Column Name     | Type     |
    +-----------------+----------+
    | order_number    | int      |
    | customer_number | int      |
    +-----------------+----------+
    order_number is the primary key for this table.
    This table contains information about the order ID and the customer ID.

Write an SQL query to find the customer_number for the customer who has placed **the largest number of orders**.

The test cases are generated so that **exactly one customer** will have placed more orders than any other customer.

The query result format is in the following example.

## Example 1:

    Input:
    Orders table:
    +--------------+-----------------+
    | order_number | customer_number |
    +--------------+-----------------+
    | 1            | 1               |
    | 2            | 2               |
    | 3            | 3               |
    | 4            | 3               |
    +--------------+-----------------+
    Output:
    +-----------------+
    | customer_number |
    +-----------------+
    | 3               |
    +-----------------+
    Explanation:
    The customer with number 3 has two orders, which is greater than either customer 1 or 2 because each of them only has one order.
    So the result is customer_number 3.

---

# Solution

## Intuition

This problem is asking us to find the customer who has placed the most orders. We can approach this problem by grouping
the data by **customer_numbe**r and counting the number of orders each customer placed.

## Approach

1. Group data: First, we group the data by **customer_number** using the SQL **GROUP BY** statement. This gives us the
   number of orders placed by each customer.

2. Order data: After grouping the data, we need to order it to find the customer who placed the most orders. The **ORDER
   BY** statement in SQL can be used for this purpose. We will order the data by the count of **customer_number** in
   descending order so that the customer with the most orders will be on top.

3. Limit results: We only want to find the customer who placed the most orders, so we will limit our output to 1 using
   the **LIMIT** statement in SQL.

## Complexity

- Time complexity:
  The time complexity of this query is O(n log n), where n is the number of rows in the Orders table. This is due to the
  sorting operation involved in the ORDER BY statement.

- Space complexity:
  The space complexity is O(n), where n is the number of rows in the Orders table. This is because the GROUP BY
  statement needs to create a separate group for each unique customer_number.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)