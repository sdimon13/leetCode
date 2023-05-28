# 1978. Employees Whose Manager Left the Company

Level: `Easy`

https://leetcode.com/problems/combine-two-tables/

---

# Description

Table: `Employees`

    +-------------+----------+
    | Column Name | Type     |
    +-------------+----------+
    | employee_id | int      |
    | name        | varchar  |
    | manager_id  | int      |
    | salary      | int      |
    +-------------+----------+
    employee_id is the primary key for this table.
    This table contains information about the employees, their salary, and the ID of their manager. Some employees do not have a manager (manager_id is null).

Write an SQL query to report the IDs of the employees whose salary is strictly less than `$30000` and whose manager left
the company. When a manager leaves the company, their information is deleted from the `Employees` table, but the reports
still have their `manager_id` set to the manager that left.

Return the result table ordered by `employee_id`.

The query result format is in the following example.

## Example 1:

    Input:  
    Employees table:
    +-------------+-----------+------------+--------+
    | employee_id | name      | manager_id | salary |
    +-------------+-----------+------------+--------+
    | 3           | Mila      | 9          | 60301  |
    | 12          | Antonella | null       | 31000  |
    | 13          | Emery     | null       | 67084  |
    | 1           | Kalel     | 11         | 21241  |
    | 9           | Mikaela   | null       | 50937  |
    | 11          | Joziah    | 6          | 28485  |
    +-------------+-----------+------------+--------+
    Output:
    +-------------+
    | employee_id |
    +-------------+
    | 11          |
    +-------------+
    
    Explanation:
    The employees with a salary less than $30000 are 1 (Kalel) and 11 (Joziah).
    Kalel's manager is employee 11, who is still in the company (Joziah).
    Joziah's manager is employee 6, who left the company because there is no row for employee 6 as it was deleted.

---

# Solution

## Intuition

This problem can be solved using a simple SQL query. We need to identify employees with a salary less than 30000 whose
manager is no longer in the company.

## Approach

We can solve this problem by joining the Employees table with itself on the manager_id and employee_id fields. Then, we
filter for rows where the employee's salary is less than 30000, the employee's manager_id is not NULL (indicating that
the employee has a manager), and the manager's employee_id is NULL (indicating that the manager has left the company).
Finally, we select the employee_id field from the result.

## Complexity

- Time complexity:
  The time complexity is O(n<sup>2</sup>), where n is the number of rows in the Employees table. This is because the
  JOIN operation compares each row with every other row in the table.

- Space complexity:
  The space complexity is O(n), where n is the number of rows in the Employees table. This is the space required to
  store the intermediate and final results of the query.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)