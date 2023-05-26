# 184. Department Highest Salary

Level: `Medium`

https://leetcode.com/problems/department-highest-salary/

---

# Description

Table: `Employee`

    +--------------+---------+
    | Column Name  | Type    |
    +--------------+---------+
    | id           | int     |
    | name         | varchar |
    | salary       | int     |
    | departmentId | int     |
    +--------------+---------+
    id is the primary key column for this table.
    departmentId is a foreign key of the ID from the Department table.
    Each row of this table indicates the ID, name, and salary of an employee. It also contains the ID of their department.


Table: `Department`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | id          | int     |
    | name        | varchar |
    +-------------+---------+
    id is the primary key column for this table. It is guaranteed that department name is not NULL.
    Each row of this table indicates the ID of a department and its name.


Write an SQL query to find employees who have the highest salary in each of the departments.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Employee table:
    +----+-------+--------+--------------+
    | id | name  | salary | departmentId |
    +----+-------+--------+--------------+
    | 1  | Joe   | 70000  | 1            |
    | 2  | Jim   | 90000  | 1            |
    | 3  | Henry | 80000  | 2            |
    | 4  | Sam   | 60000  | 2            |
    | 5  | Max   | 90000  | 1            |
    +----+-------+--------+--------------+
    Department table:
    +----+-------+
    | id | name  |
    +----+-------+
    | 1  | IT    |
    | 2  | Sales |
    +----+-------+
    Output:
    +------------+----------+--------+
    | Department | Employee | Salary |
    +------------+----------+--------+
    | IT         | Jim      | 90000  |
    | Sales      | Henry    | 80000  |
    | IT         | Max      | 90000  |
    +------------+----------+--------+
    Explanation: Max and Jim both have the highest salary in the IT department and Henry has the highest salary in the Sales department.

---

# Solution

## Intuition
The problem is to find the highest salary per department and return the details of that employee and department. This requires aggregating salaries within each department and identifying the highest one. MySQL provides aggregation functions like MAX() to achieve this. To identify the specific employee details, we can use a WHERE clause and a subquery.

## Approach
Our approach here uses a subquery with a GROUP BY clause to find the highest salary in each department. The subquery is run on the Employee table and groups the data by the DepartmentId, finding the maximum salary in each group. The output of this subquery is a list of DepartmentIds and their respective maximum salaries.

We then use this subquery in the main query's WHERE clause, selecting only those rows from the Employee table which match the DepartmentId and Salary pairs returned by the subquery. We join this with the Department table on the DepartmentId field to get the name of each department.

## Complexity
- Time complexity:
  The time complexity is O(n), where n is the number of employees. The GROUP BY clause used in the subquery scans all employees once, and the main query also scans all employees once.

- Space complexity:
  The space complexity is O(m), where m is the number of departments. The subquery generates a result set with one row per department, so it uses space proportional to the number of departments.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)