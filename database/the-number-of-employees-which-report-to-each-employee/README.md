# 1731. The Number of Employees Which Report to Each Employee

Level: `Easy`

https://leetcode.com/problems/the-number-of-employees-which-report-to-each-employee/

---

# Description

Table: `Employees`

    +-------------+----------+
    | Column Name | Type     |
    +-------------+----------+
    | employee_id | int      |
    | name        | varchar  |
    | reports_to  | int      |
    | age         | int      |
    +-------------+----------+
    employee_id is the primary key for this table.
    This table contains information about the employees and the id of the manager they report to. Some employees do not report to anyone (reports_to is null).

For this problem, we will consider a **manager** an employee who has at least 1 other employee reporting to them.

Write an SQL query to report the ids and the names of all **managers**, the number of employees who report **directly**
to them,
and the average age of the reports rounded to the nearest integer.

Return the result table ordered by **employee_id**.

The query result format is in the following example.

## Example 1:

    Input:
    Employees table:
    +-------------+---------+------------+-----+
    | employee_id | name    | reports_to | age |
    +-------------+---------+------------+-----+
    | 9           | Hercy   | null       | 43  |
    | 6           | Alice   | 9          | 41  |
    | 4           | Bob     | 9          | 36  |
    | 2           | Winston | null       | 37  |
    +-------------+---------+------------+-----+
    Output:
    +-------------+-------+---------------+-------------+
    | employee_id | name  | reports_count | average_age |
    +-------------+-------+---------------+-------------+
    | 9           | Hercy | 2             | 39          |
    +-------------+-------+---------------+-------------+
    Explanation: Hercy has 2 people report directly to him, Alice and Bob. Their average age is (41+36)/2 = 38.5, which is 39 after rounding it to the nearest integer.

---

# Solution

## Intuition

Given the problem, our task is to write an SQL query that gives the ids and names of all managers, the number of
employees who report directly to them, and the average age of those employees (rounded to the nearest integer).

We recognize that each manager's id is stored in the 'reports_to' column of the Employees table for all employees who
report to them. To get the data we want, we need to join the table with itself, once for the manager (employee_id in the
first table = reports_to in the second table) and again for the employees who report to them. This can be achieved by
using a self-join on the 'employees' table.

## Approach

1. The first step is to join the Employees table to itself. We call the first instance of the table 'e1' and the second
   instance 'e2'. The join condition is that the 'employee_id' of 'e2' matches the 'reports_to' column of 'e1'. This
   effectively pairs every employee with the manager they report to.

2. The next step is to group the results by 'reports_to', effectively grouping the results by manager.

3. With these groups, we can now count the number of employees in each group (representing the number of employees that
   report to each manager) and calculate the average age of the employees in each group. The COUNT function is used for
   the former and AVG function for the latter. We also use ROUND function to round the average age to the nearest
   integer.

4. The final step is to order the result by 'employee_id'.

## Complexity

- Time complexity:
  The time complexity is O(nlogn) because the query involves a join operation, which can take up to n log n time.
  However, actual time complexity can vary depending on the specifics of the SQL database engine, as it may use
  optimization techniques to speed up operations.

- Space complexity:
  The space complexity is O(n) where n is the number of rows in the Employees table, since in worst case the self join
  may need to hold all rows in memory.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)