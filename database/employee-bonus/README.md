# 577. Employee Bonus

Level: `Easy`

https://leetcode.com/problems/employee-bonus/

---

# Description

Table: `Employee`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | empId       | int     |
    | name        | varchar |
    | supervisor  | int     |
    | salary      | int     |
    +-------------+---------+
    empId is the primary key column for this table.
    Each row of this table indicates the name and the ID of an employee in addition to their salary and the id of their manager.

Table: `Bonus`

    +-------------+------+
    | Column Name | Type |
    +-------------+------+
    | empId       | int  |
    | bonus       | int  |
    +-------------+------+
    empId is the primary key column for this table.
    empId is a foreign key to empId from the Employee table.
    Each row of this table contains the id of an employee and their respective bonus.

Write an SQL query to report the name and bonus amount of each employee with a bonus **less than** `1000`.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Employee table:
    +-------+--------+------------+--------+
    | empId | name   | supervisor | salary |
    +-------+--------+------------+--------+
    | 3     | Brad   | null       | 4000   |
    | 1     | John   | 3          | 1000   |
    | 2     | Dan    | 3          | 2000   |
    | 4     | Thomas | 3          | 4000   |
    +-------+--------+------------+--------+
    Bonus table:
    +-------+-------+
    | empId | bonus |
    +-------+-------+
    | 2     | 500   |
    | 4     | 2000  |
    +-------+-------+
    Output:
    +------+-------+
    | name | bonus |
    +------+-------+
    | Brad | null  |
    | John | null  |
    | Dan  | 500   |
    +------+-------+

---

# Solution

## Intuition

# Intuition

The goal of this problem is to extract the names and bonuses of employees whose bonus is less than 1000. If an employee
doesn't have a bonus record, then we consider their bonus as null. This problem can be solved using a **LEFT JOIN**
operation to combine the Employee and Bonus tables based on the **empId** field, which is the primary key in both
tables.

## Approach

1. Using the **LEFT JOIN** operation, combine the Employee and Bonus tables based on the **empId** field. The result of
   this operation will contain all records from the Employee table and matching records from the Bonus table. If there
   is no match, the result is null on the Bonus side.

2. After the join operation, filter out the records using the **WHERE** clause. The condition here is that the bonus
   must be less than 1000 or null.

## Complexity

- Time complexity:
  Since the operation involves a JOIN operation and filtering, the time complexity is O(n+m), where n is the number of
  records in the Employee table and m is the number of records in the Bonus table.

- Space complexity:
  The space complexity is also O(n+m), since we need to store the result of the join operation in memory.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)