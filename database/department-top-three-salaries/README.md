# 185. Department Top Three Salaries

Level: `Hard`

https://leetcode.com/problems/department-top-three-salaries/

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
    id is the primary key column for this table.
    Each row of this table indicates the ID of a department and its name.


A company's executives are interested in seeing who earns the most money in each of the company's departments. A **high earner** in a department is an employee who has a salary in the **top three unique** salaries for that department.

Write an SQL query to find the employees who are **high earners** in each of the departments.

Return the result table **in any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Employee table:
    +----+-------+--------+--------------+
    | id | name  | salary | departmentId |
    +----+-------+--------+--------------+
    | 1  | Joe   | 85000  | 1            |
    | 2  | Henry | 80000  | 2            |
    | 3  | Sam   | 60000  | 2            |
    | 4  | Max   | 90000  | 1            |
    | 5  | Janet | 69000  | 1            |
    | 6  | Randy | 85000  | 1            |
    | 7  | Will  | 70000  | 1            |
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
    | IT         | Max      | 90000  |
    | IT         | Joe      | 85000  |
    | IT         | Randy    | 85000  |
    | IT         | Will     | 70000  |
    | Sales      | Henry    | 80000  |
    | Sales      | Sam      | 60000  |
    +------------+----------+--------+
    Explanation:
    In the IT department:
    - Max earns the highest unique salary
    - Both Randy and Joe earn the second-highest unique salary
    - Will earns the third-highest unique salary
    
    In the Sales department:
    - Henry earns the highest salary
    - Sam earns the second-highest salary
    - There is no third-highest salary as there are only two employees

---

# Solution

## Intuition
The problem is to find employees with the top three highest salaries in each department. This means we must rank the salaries within each department and select only those employees with ranks 1, 2, or 3. MySQL does not have a built-in rank function, but we can emulate this with a subquery.

## Approach
Our approach here uses a subquery in the WHERE clause to emulate a ranking of salaries within each department. For each employee, the subquery counts the number of distinct salaries in the same department that are higher than the employee's salary. If this count is less than 3, that means the employee's salary is one of the top three in their department, so we include this employee in the output.

We join the Employee and Department tables on the DepartmentId field to get the name of each department.

## Complexity
- Time complexity:
  The time complexity is O(n^2), where n is the number of employees. The subquery is run for each employee, and in each subquery, we may scan up to n rows from the Employee table.

- Space complexity:
  The space complexity is O(m), where m is the number of departments. The final result set can have up to 3 rows per department, so it uses space proportional to the number of departments.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)