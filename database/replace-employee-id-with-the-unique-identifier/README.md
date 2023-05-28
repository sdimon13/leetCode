# 1378. Replace Employee ID With The Unique Identifier

Level: `Easy`

https://leetcode.com/problems/replace-employee-id-with-the-unique-identifier/

---

# Description

Table: `Employees`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | id            | int     |
    | name          | varchar |
    +---------------+---------+
    id is the primary key for this table.
    Each row of this table contains the id and the name of an employee in a company.

Table: `EmployeeUNI`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | id            | int     |
    | unique_id     | int     |
    +---------------+---------+
    (id, unique_id) is the primary key for this table.
    Each row of this table contains the id and the corresponding unique id of an employee in the company.

Write an SQL query to show the **unique ID** of each user, If a user does not have a unique ID replace just show `null`.

Return the result table in **any** order.

The query result format is in the following example.

## Example 1:

    Input:
    Employees table:
    +----+----------+
    | id | name     |
    +----+----------+
    | 1  | Alice    |
    | 7  | Bob      |
    | 11 | Meir     |
    | 90 | Winston  |
    | 3  | Jonathan |
    +----+----------+
    EmployeeUNI table:
    +----+-----------+
    | id | unique_id |
    +----+-----------+
    | 3  | 1         |
    | 11 | 2         |
    | 90 | 3         |
    +----+-----------+
    Output:
    +-----------+----------+
    | unique_id | name     |
    +-----------+----------+
    | null      | Alice    |
    | null      | Bob      |
    | 2         | Meir     |
    | 3         | Winston  |
    | 1         | Jonathan |
    +-----------+----------+
    Explanation:
    Alice and Bob do not have a unique ID, We will show null instead.
    The unique ID of Meir is 2.
    The unique ID of Winston is 3.
    The unique ID of Jonathan is 1.

---

# Solution

## Intuition

The problem asks to return the **unique_id** and **name** for each employee from two tables: **Employees** and *
*EmployeeUNI**. If an employee doesn't have a **unique_id**, we should show **NULL** instead. This is a typical SQL JOIN
operation problem where we have to join two tables based on a common attribute, in this case **id**.

## Approach

The best approach to this problem is to use a LEFT JOIN operation on the **id** column, which is common in both tables.
The **Employees** table is on the left side and **EmployeeUNI** on the right. By using a LEFT JOIN, we can ensure that
we get all the records from the **Employees** table and corresponding records from **EmployeeUN**I where the **id**
matches. If there is no match, the result is **NULL** from the right side (**EmployeeUNI**). Hence, if an employee
doesn't have a **unique_id**, **NULL** will be shown.

## Complexity

- Time complexity:
  The time complexity of a SQL query heavily depends on various factors like the DBMS, indexes, the number of records,
  etc. However, in general, a LEFT JOIN operation could be considered as O(N), where N is the number of records in the
  Employees table.

- Space complexity:
  The space complexity is O(M), where M is the number of records in the result of the JOIN operation.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)