# 1873. Calculate Special Bonus

Level: `Easy`

https://leetcode.com/problems/calculate-special-bonus/

---

# Description

Table: `Employees`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | employee_id | int     |
    | name        | varchar |
    | salary      | int     |
    +-------------+---------+
    employee_id is the primary key for this table.
    Each row of this table indicates the employee ID, employee name, and salary.

Write an SQL query to calculate the bonus of each employee. The bonus of an employee is `100%` of their salary if the ID
of the employee is **an odd number** and **the employee name does not start with the character** `'M'`. The bonus of an
employee is `0` otherwise.

Return the result table ordered by `employee_id`.

The query result format is in the following example.

## Example 1:

    Input:
    Employees table:
    +-------------+---------+--------+
    | employee_id | name    | salary |
    +-------------+---------+--------+
    | 2           | Meir    | 3000   |
    | 3           | Michael | 3800   |
    | 7           | Addilyn | 7400   |
    | 8           | Juan    | 6100   |
    | 9           | Kannon  | 7700   |
    +-------------+---------+--------+
    Output:
    +-------------+-------+
    | employee_id | bonus |
    +-------------+-------+
    | 2           | 0     |
    | 3           | 0     |
    | 7           | 7400  |
    | 8           | 0     |
    | 9           | 7700  |
    +-------------+-------+
    Explanation:
    The employees with IDs 2 and 8 get 0 bonus because they have an even employee_id.
    The employee with ID 3 gets 0 bonus because their name starts with 'M'.
    The rest of the employees get a 100% bonus.

---

# Solution

## Intuition

The task requires us to calculate a bonus for employees based on their **employee_id** and **name**. If an employee's ID
is an odd number and their name does not start with 'M', they should receive a bonus equal to their salary. Otherwise,
their bonus should be 0.

## Approach

We can use a **SELECT** statement to get the **employee_id** and calculate the **bonus** based on the conditions given
in the problem. We'll use the SQL **IF** function to set the value of **bonus**. If the **employee_id** is odd (*
*employee_id%2!=0**) and the **name** does not start with 'M' (**name NOT LIKE 'M%'**), the **bonus** is equal to the *
*salary**; otherwise, it's 0. The result should be ordered by **employee_id** using the **ORDER BY** clause.

## Complexity

- Time complexity:
  O(n), where n is the number of rows in the Employees table. The database needs to scan each row to evaluate the IF
  function and calculate the bonus.

- Space complexity:
  O(1), the space complexity does not depend on the size of the input table, and only a constant amount of space is
  used.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)