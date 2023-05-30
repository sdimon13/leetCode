# 1075. Project Employees I

Level: `Easy`

https://leetcode.com/problems/project-employees-i/

---

# Description

Table: `Project`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | project_id  | int     |
    | employee_id | int     |
    +-------------+---------+
    (project_id, employee_id) is the primary key of this table.
    employee_id is a foreign key to Employee table.
    Each row of this table indicates that the employee with employee_id is working on the project with project_id.

Table: `Employee`

    +------------------+---------+
    | Column Name      | Type    |
    +------------------+---------+
    | employee_id      | int     |
    | name             | varchar |
    | experience_years | int     |
    +------------------+---------+
    employee_id is the primary key of this table. It's guaranteed that experience_years is not NULL.
    Each row of this table contains information about one employee.

Write an SQL query that reports the **average** experience years of all the employees for each project, **rounded to 2
digits**.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Project table:
    +-------------+-------------+
    | project_id  | employee_id |
    +-------------+-------------+
    | 1           | 1           |
    | 1           | 2           |
    | 1           | 3           |
    | 2           | 1           |
    | 2           | 4           |
    +-------------+-------------+
    Employee table:
    +-------------+--------+------------------+
    | employee_id | name   | experience_years |
    +-------------+--------+------------------+
    | 1           | Khaled | 3                |
    | 2           | Ali    | 2                |
    | 3           | John   | 1                |
    | 4           | Doe    | 2                |
    +-------------+--------+------------------+
    Output:
    +-------------+---------------+
    | project_id  | average_years |
    +-------------+---------------+
    | 1           | 2.00          |
    | 2           | 2.50          |
    +-------------+---------------+
    Explanation: The average experience years for the first project is (3 + 2 + 1) / 3 = 2.00 and for the second project is (3 + 2) / 2 = 2.50

---

# Solution

## Intuition

The problem is asking to calculate the average experience years of employees per project. The information about which
employee is working on which project is stored in the "Project" table, and the years of experience of each employee are
stored in the "Employee" table. Therefore, we need to combine these two tables. SQL's JOIN clause can be used to combine
rows from two or more tables, based on a related column between them, which in this case is the "employee_id".

## Approach

Firstly, we will join the "Project" and "Employee" tables on the "employee_id" field which is common in both tables.
This gives us a combined table where we have the "project_id", "employee_id" and the corresponding "experience_years".

Next, we group by the "project_id". This groups together all rows that have the same "project_id".

Then, for each group, we calculate the average "experience_years". In SQL, the average can be calculated using the AVG
function. However, the problem specifically asks us to round the result to 2 decimal places. Therefore, we use the ROUND
function to round the result of the average calculation.

## Complexity

- Time complexity:
  The time complexity is O(n), where n is the total number of rows in the "Project" table. This is because we are
  iterating over all rows once.

- Space complexity:
  The space complexity is O(1), since we are not using any additional space that scales with the size of the input.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)