# 1667. Fix Names in a Table

Level: `Easy`

https://leetcode.com/problems/fix-names-in-a-table/

---

# Description

Table: Users

    +----------------+---------+
    | Column Name    | Type    |
    +----------------+---------+
    | user_id        | int     |
    | name           | varchar |
    +----------------+---------+
    user_id is the primary key for this table.
    This table contains the ID and the name of the user. The name consists of only lowercase and uppercase characters.

Write an SQL query to fix the names so that only the first character is uppercase and the rest are lowercase.

Return the result table ordered by `user_id`.

The query result format is in the following example.

## Example 1:

    Input:
    Users table:
    +---------+-------+
    | user_id | name  |
    +---------+-------+
    | 1       | aLice |
    | 2       | bOB   |
    +---------+-------+
    Output:
    +---------+-------+
    | user_id | name  |
    +---------+-------+
    | 1       | Alice |
    | 2       | Bob   |
    +---------+-------+

---

# Solution

## Intuition

The problem is straightforward as it's mainly about string manipulation. The requirement is to format user names in a
specific way - only the first character should be in uppercase and the rest should be in lowercase.

## Approach

We can leverage built-in SQL functions to solve this problem. Specifically, **UPPER()**, **LOWER()**, **SUBSTR()**, and
**CONCAT()** functions will be handy.

The **SUBSTR()** function is used to extract a part of the string. We will extract the first character and the rest of
the string separately.

We will use **UPPER()** on the first character to make sure it's in uppercase, and **LOWER()** on the rest of the string
to convert it to lowercase.

The **CONCAT()** function will help us combine the first character (which is now in uppercase) and the rest of the
string (which is now in lowercase) together.

Finally, we sort the result by **user_id**.

## Complexity

- Time complexity:
  The time complexity of this SQL query is O(n), where n is the total number of users in the table. This is because the
  query needs to process each user's name to make sure it meets the specified conditions.

- Space complexity:
  The space complexity is O(1). We are not using any extra space that scales with the size of the input. The memory used
  to store the result set is not typically counted towards space complexity.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)