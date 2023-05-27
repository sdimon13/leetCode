# 1965. Employees With Missing Information

Level: `Easy`

https://leetcode.com/problems/user-activity-for-the-past-30-days-i/

---

# Description

Table: `Employees`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | employee_id | int     |
    | name        | varchar |
    +-------------+---------+
    employee_id is the primary key for this table.
    Each row of this table indicates the name of the employee whose ID is employee_id.

Table: `Salaries`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | employee_id | int     |
    | salary      | int     |
    +-------------+---------+
    employee_id is the primary key for this table.
    Each row of this table indicates the salary of the employee whose ID is employee_id.

Write an SQL query to report the IDs of all the employees with **missing information**. The information of an employee
is missing if:

The employee's **name** is missing, or
The employee's **salary** is missing.
Return the result table ordered by `employee_id` **in ascending order**.

The query result format is in the following example.

## Example 1:

    Input:
    Employees table:
    +-------------+----------+
    | employee_id | name     |
    +-------------+----------+
    | 2           | Crew     |
    | 4           | Haven    |
    | 5           | Kristian |
    +-------------+----------+
    Salaries table:
    +-------------+--------+
    | employee_id | salary |
    +-------------+--------+
    | 5           | 76071  |
    | 1           | 22517  |
    | 4           | 63539  |
    +-------------+--------+
    Output:
    +-------------+
    | employee_id |
    +-------------+
    | 1           |
    | 2           |
    +-------------+
    Explanation:
    Employees 1, 2, 4, and 5 are working at this company.
    The name of employee 1 is missing.
    The salary of employee 2 is missing.

---

# Solution

## Intuition

We are given two tables: **Employees** and **Salaries**. We need to find all the employees who either don't have a
name (i.e., they are present in the **Salaries** table but not in the **Employees** table) or don't have a salary (i.e.,
they are present in the **Employees** table but not in the **Salaries** table).

## Approach

To solve this problem, we can use the SQL operator **NOT IN** and the **UNION** clause.

1. We first find the **employee_id** from the **Employees** table that is not present in the **Salaries** table. This
   will give us the employees who don't have a salary.

2. Then we find the **employee_id** from the **Salaries** table that is not present in the **Employees** table. This
   will give us the employees who don't have a name.

3. Finally, we combine the results from steps 1 and 2 using the **UNION** clause and order the results by **employee_id
   **.

## Complexity

- Time complexity:
  The time complexity is O(n + m), where n is the number of rows in the **Employees** table and m is the number of rows
  in the **Salaries** table. We are scanning both tables completely.

- Space complexity:
  The space complexity is O(n + m), where n is the number of rows in the **Employees** table and m is the number of rows
  in the **Salaries** table. The space is used to store the intermediate and final results.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)