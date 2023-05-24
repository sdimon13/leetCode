# 182. Duplicate Emails

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


Write an SQL query to report all the duplicate emails. Note that it's guaranteed that the email field is not NULL.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Person table:
    +----+---------+
    | id | email   |
    +----+---------+
    | 1  | a@b.com |
    | 2  | c@d.com |
    | 3  | a@b.com |
    +----+---------+
    Output:
    +---------+
    | Email   |
    +---------+
    | a@b.com |
    +---------+
    Explanation: a@b.com is repeated two times.

---

# Solution

## Intuition
The problem is asking us to find duplicate emails from a given table named "Person". The table has two columns, "Id" and "Email". We don't care about the "Id", we just need to find which emails appear more than once in the table. A direct approach to this problem would be to use SQL's **GROUP BY** clause to group the data by "Email", and then to filter out only those groups where the count of "Email" is greater than 1, which indicates the email is duplicated.

## Approach
To solve this problem, we can use a simple SQL query. The query makes use of the **GROUP BY** clause to create groups of emails, and then uses the **HAVING** clause to filter out only those groups where the count of emails is more than 1.

We're not considering the "Id" column in this problem, just the "Email" column. Therefore, the space and time complexity would largely depend on the database management system's implementation of the **GROUP BY** and **HAVING** clauses.

Here's a step-by-step breakdown of how the SQL query works:

1. **SELECT Email FROM Person** - this selects the Email column from the Person table.
2. **GROUP BY Email** - this groups the results by the Email column. If there are two rows with the same email, they will be in the same group.
3. **HAVING COUNT(*) > 1** - this filters the groups and only keeps those where the count of rows in the group (i.e., the number of duplicate emails) is more than 1.

## Complexity
- Time complexity:
  The time complexity is generally O(n log n) due to the grouping operation. But it is important to mention that the actual time complexity can vary depending on the DBMS's internal implementation, as well as the number of records in the database and the distribution of the data.

- Space complexity:
  The space complexity is O(n), where n is the number of different emails. This is because in the worst-case scenario, the DBMS might need to store all distinct emails when performing the **GROUP BY** operation. Again, it is important to note that the actual space complexity can vary based on the DBMS's internal implementation.

---

# Code
Here are the solution and test files:
- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)