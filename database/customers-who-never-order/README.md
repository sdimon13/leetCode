# 183. Customers Who Never Order

Level: `Easy`

https://leetcode.com/problems/customers-who-never-order/

---

# Description

Table: `Customers`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | id          | int     |
    | name        | varchar |
    +-------------+---------+
    id is the primary key column for this table.
    Each row of this table indicates the ID and name of a customer.


Table: `Orders`

    +-------------+------+
    | Column Name | Type |
    +-------------+------+
    | id          | int  |
    | customerId  | int  |
    +-------------+------+
    id is the primary key column for this table.
    customerId is a foreign key of the ID from the Customers table.
    Each row of this table indicates the ID of an order and the ID of the customer who ordered it.


Write an SQL query to report all customers who never order anything.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Customers table:
    +----+-------+
    | id | name  |
    +----+-------+
    | 1  | Joe   |
    | 2  | Henry |
    | 3  | Sam   |
    | 4  | Max   |
    +----+-------+
    Orders table:
    +----+------------+
    | id | customerId |
    +----+------------+
    | 1  | 3          |
    | 2  | 1          |
    +----+------------+
    Output:
    +-----------+
    | Customers |
    +-----------+
    | Henry     |
    | Max       |
    +-----------+

---

# Solution

## Intuition
When given two tables of **Customers** and **Orders**, and the goal is to select the customers who have never placed an order, we need to find the **Customer** records that do not have a matching record in the **Orders** table. This is a typical case for utilizing SQL joins.

## Approach
We can use a **LEFT JOIN** to solve this problem.

The **LEFT JOIN** keyword returns all records from the left table (**Customers**), and the matched records from the right table (**Orders**). The result is **NULL** from the right side if there is no match. This means if a customer never orders, the result from the **Orders** side will be **NULL** after the **LEFT JOIN**.

So, we can select those customers whose id from the **Orders** table is **NULL** after performing **LEFT JOIN**.

## Complexity
- Time complexity:
  The time complexity is O(n), where n is the number of customers. The LEFT JOIN operation needs to go through each customer.

- Space complexity:
  The space complexity is O(1), assuming that the database is responsible for handling the storage of queries, and we only count the space used by our final output.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)