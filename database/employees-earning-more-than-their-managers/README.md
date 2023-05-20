# 181. Employees Earning More Than Their Managers

Level: `Easy`

https://leetcode.com/problems/employees-earning-more-than-their-managers/

---

# Description

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | id          | int     |
    | name        | varchar |
    | salary      | int     |
    | managerId   | int     |
    +-------------+---------+
    id is the primary key column for this table.
    Each row of this table indicates the ID of an employee, their name, salary, and the ID of their manager.


Write an SQL query to find the employees who earn more than their managers.

Return the result table in **any order**.

The query result format is in the following example.



## Example 1:

    Input:
    Employee table:
    +----+-------+--------+-----------+
    | id | name  | salary | managerId |
    +----+-------+--------+-----------+
    | 1  | Joe   | 70000  | 3         |
    | 2  | Henry | 80000  | 4         |
    | 3  | Sam   | 60000  | Null      |
    | 4  | Max   | 90000  | Null      |
    +----+-------+--------+-----------+
    Output:
    +----------+
    | Employee |
    +----------+
    | Joe      |
    +----------+
    Explanation: Joe is the only employee who earns more than his manager.

---

# Solution

## Intuition
This problem requires knowledge of SQL joins and conditionals. In this scenario, we are asked to select the employees who earn more than their managers. To solve this, we would join the Employee table with itself on the condition that the **managerId** of an employee matches the **id** of another employee (the manager). Then we would compare the salaries of each pair of employee and manager.

## Approach
1. First, we join the Employee table with itself, using the **managerId** of one as the **id** of the other. We can consider one instance of the Employee table as **e** (for employees) and the other as **m** (for managers).

2. Then, in our WHERE clause, we specify that we want only the rows where the salary of the employee (**e.salary**) is greater than the salary of the manager (**m.salary**).

3. Finally, we select the name of the employee (**e.Name**) who earns more than their manager.

## Complexity
The time complexity and space complexity of this solution would depend on the implementation of the SQL database and the specifics of its query optimization. However, generally, the join operation is expensive and could be in the order of O(n<sup>2</sup>) where **n** is the number of rows in the Employee table. The space complexity can also be significant as the join operation may create a new table in memory. Yet, this too depends largely on the specifics of the SQL database.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)