# 196. Delete Duplicate Emails

Level: `Easy`

https://leetcode.com/problems/delete-duplicate-emails/

---

# Description


    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | id          | int     |
    | email       | varchar |
    +-------------+---------+
    
    id is the primary key column for this table.
    Each row of this table contains an email. The emails will not contain uppercase letters.


Write an SQL query to **delete** all the duplicate emails, keeping only one unique email with the smallest `id`. Note that you are supposed to write a `DELETE` statement and not a `SELECT` one.

After running your script, the answer shown is the `Person` table. The driver will first compile and run your piece of code and then show the `Person` table. The final order of the `Person` table **does not matter**.

The query result format is in the following example.

## Example 1:

    Input:
    Person table:
    +----+------------------+
    | id | email            |
    +----+------------------+
    | 1  | john@example.com |
    | 2  | bob@example.com  |
    | 3  | john@example.com |
    +----+------------------+
    Output:
    +----+------------------+
    | id | email            |
    +----+------------------+
    | 1  | john@example.com |
    | 2  | bob@example.com  |
    +----+------------------+
    Explanation: john@example.com is repeated two times. We keep the row with the smallest Id = 1.

---

# Solution

## Intuition
The problem seems to require a solution that handles duplicate data in a database table. In this case, we need to remove duplicate rows based on the "Email" field from the "Person" table.

## Approach
To solve this problem, we can use a MySQL DELETE statement. The key idea is to use a self-join on the "Person" table. A self-join is a regular join, but the table is joined with itself.

In the WHERE clause, we check for rows that have the same "Email" but a higher "Id". By doing so, we ensure that we keep only the person with the lowest "Id" for each email.

Finally, we use the DELETE statement to remove the unwanted rows from the table.

## Complexity
- Time complexity:
  The time complexity is O(n^2), where n is the number of rows in the "Person" table. This is because for every row in the table, we are checking all other rows for the same email.

- Space complexity:
  The space complexity is O(1) as we are not using any additional data structures in our query.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)