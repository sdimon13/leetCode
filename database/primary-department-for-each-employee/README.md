# 1789. Primary Department for Each Employee

Level: `Easy`

https://leetcode.com/problems/primary-department-for-each-employee/

---

# Description

Table: `Employee`

    +---------------+---------+
    | Column Name   |  Type   |
    +---------------+---------+
    | employee_id   | int     |
    | department_id | int     |
    | primary_flag  | varchar |
    +---------------+---------+
    (employee_id, department_id) is the primary key for this table.
    employee_id is the id of the employee.
    department_id is the id of the department to which the employee belongs.
    primary_flag is an ENUM of type ('Y', 'N'). If the flag is 'Y', the department is the primary department for the employee. If the flag is 'N', the department is not the primary.

Employees can belong to multiple departments. When the employee joins other departments, they need to decide which
department is their primary department. Note that when an employee belongs to only one department, their primary column
is `'N'`.

Write an SQL query to report all the employees with their primary department. For employees who belong to one
department, report their only department.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Employee table:
    +-------------+---------------+--------------+
    | employee_id | department_id | primary_flag |
    +-------------+---------------+--------------+
    | 1           | 1             | N            |
    | 2           | 1             | Y            |
    | 2           | 2             | N            |
    | 3           | 3             | N            |
    | 4           | 2             | N            |
    | 4           | 3             | Y            |
    | 4           | 4             | N            |
    +-------------+---------------+--------------+
    Output:
    +-------------+---------------+
    | employee_id | department_id |
    +-------------+---------------+
    | 1           | 1             |
    | 2           | 1             |
    | 3           | 3             |
    | 4           | 3             |
    +-------------+---------------+
    Explanation:
    - The Primary department for employee 1 is 1.
    - The Primary department for employee 2 is 1.
    - The Primary department for employee 3 is 3.
    - The Primary department for employee 4 is 3.

---

# Solution

## Intuition

From the problem statement, we need to return a list of all employees along with their primary department. For employees
who belong to only one department, we must return that as their primary department, even if the primary_flag for it is '
N'.

## Approach

We can solve this problem using SQL UNION operation. The UNION operator combines the result set of two or more SELECT
statements. In our case, we need two SELECT statements:

1. In the first SELECT statement, we select all employees with their primary department where the primary_flag is 'Y'.
2. In the second SELECT statement, we select all employees who belong to only one department. To get this, we choose
   employees that are not present in the result of the first SELECT statement.

The final result will combine these two SELECT statements, giving us a list of all employees along with their primary
department.

## Complexity

- Time complexity:
  The time complexity for this problem will mainly depend on the execution time of the SQL statements and can be
  considered as O(n), where n is the number of rows in the Employee table.

- Space complexity:
  The space complexity is O(n), where n is the number of rows in the Employee table. The SQL operation does not use
  additional space that scales with input size.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)