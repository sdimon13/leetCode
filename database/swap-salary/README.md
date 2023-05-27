# 627. Swap Salary

Level: `Easy`

https://leetcode.com/problems/swap-salary/

---

# Description

Table: Salary

    +-------------+----------+
    | Column Name | Type     |
    +-------------+----------+
    | id          | int      |
    | name        | varchar  |
    | sex         | ENUM     |
    | salary      | int      |
    +-------------+----------+
    id is the primary key for this table.
    The sex column is ENUM value of type ('m', 'f').
    The table contains information about an employee.


Write an SQL query to swap all `'f'` and `'m'` values (i.e., change all `'f'` values to `'m'` and vice versa) with a **single update statement** and no intermediate temporary tables.

Note that you must write a single update statement, **do not** write any select statement for this problem.

The query result format is in the following example.

## Example 1:

    Input:
    Salary table:
    +----+------+-----+--------+
    | id | name | sex | salary |
    +----+------+-----+--------+
    | 1  | A    | m   | 2500   |
    | 2  | B    | f   | 1500   |
    | 3  | C    | m   | 5500   |
    | 4  | D    | f   | 500    |
    +----+------+-----+--------+
    Output:
    +----+------+-----+--------+
    | id | name | sex | salary |
    +----+------+-----+--------+
    | 1  | A    | f   | 2500   |
    | 2  | B    | m   | 1500   |
    | 3  | C    | f   | 5500   |
    | 4  | D    | m   | 500    |
    +----+------+-----+--------+
    Explanation:
    (1, A) and (3, C) were changed from 'm' to 'f'.
    (2, B) and (4, D) were changed from 'f' to 'm'.

---

# Solution

## Intuition
This problem requires us to swap the values of an ENUM type column (**sex**) in a SQL table. The ENUM type column **sex** has two possible values: 'm' and 'f'. We are required to change all 'f' values to 'm' and vice versa.

## Approach
The SQL **UPDATE** command is used to modify the existing records in a table. In combination with the **IF** function, we can set the new value of the **sex** column based on its current value. If the current value is 'm', we change it to 'f'; otherwise, we change it to 'm'.

## Complexity
- Time complexity:
  O(n), where n is the number of rows in the Salary table. The database needs to scan each row to evaluate the IF function and perform the update.

- Space complexity:
  The space complexity is O(n), where n is the number of different emails. This is because in the worst-case scenario, the DBMS might need to store all distinct emails when performing the **GROUP BY** operation. Again, it is important to note that the actual space complexity can vary based on the DBMS's internal implementation.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)