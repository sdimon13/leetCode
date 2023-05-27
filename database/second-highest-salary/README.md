# 176. Second Highest Salary

Level: `Medium`

https://leetcode.com/problems/department-top-three-salaries/

---

# Description

Table: `Employee`

    +-------------+------+
    | Column Name | Type |
    +-------------+------+
    | id          | int  |
    | salary      | int  |
    +-------------+------+
    id is the primary key column for this table.
    Each row of this table contains information about the salary of an employee.


Write an SQL query to report the second highest salary from the `Employee` table. If there is no second highest salary, the query should report `null`.

The query result format is in the following example.

## Example 1:

    Input:
    Employee table:
    +----+--------+
    | id | salary |
    +----+--------+
    | 1  | 100    |
    | 2  | 200    |
    | 3  | 300    |
    +----+--------+
    Output:
    +---------------------+
    | SecondHighestSalary |
    +---------------------+
    | 200                 |
    +---------------------+

## Example 2:

    Input:
    Employee table:
    +----+--------+
    | id | salary |
    +----+--------+
    | 1  | 100    |
    +----+--------+
    Output:
    +---------------------+
    | SecondHighestSalary |
    +---------------------+
    | null                |
    +---------------------+

---

# Solution

## Intuition
We need to find the second highest salary. SQL provides order by and limit clauses that we can use to solve this problem. The **ORDER BY** clause can sort the salaries in descending order and the **LIMIT** clause can be used to limit the results to the top 2. However, we need the second highest salary, so we will use **OFFSET** to skip the first (highest) salary.

## Approach
We will use a nested SQL query. In the inner query, we will first make salaries distinct using the **DISTINCT** keyword, then order them in descending order, then use **LIMIT 1 OFFSET 1** to get the second highest salary. The outer query then uses **IFNULL** to handle the case when there is no second highest salary (i.e., there is only one distinct salary in the table).

## Complexity
- Time complexity:
  The time complexity is O(n log n) because the **ORDER BY** clause generally uses a sorting algorithm, and sorting is generally an O(n log n) operation. Here, n is the number of employees.

- Space complexity:
  The space complexity is O(n) because the **ORDER BY** clause needs to store the data in memory to sort it. Here, n is the number of employees.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)