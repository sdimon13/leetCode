# 1517. Find Users With Valid E-Mails

Level: `Easy`

https://leetcode.com/problems/find-users-with-valid-e-mails/

---

# Description

Table: `Users`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | user_id       | int     |
    | name          | varchar |
    | mail          | varchar |
    +---------------+---------+
    user_id is the primary key for this table.
    This table contains information of the users signed up in a website. Some e-mails are invalid.

Write an SQL query to find the users who have **valid emails**.

A valid e-mail has a prefix name and a domain where:

**The prefix name** is a string that may contain letters (upper or lower case), digits, underscore `'_'`, period `'.'`,
and/or dash `'-'`. The prefix name **must** start with a letter.
**The domain** is `'@leetcode.com'`.
Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Users table:
    +---------+-----------+-------------------------+
    | user_id | name      | mail                    |
    +---------+-----------+-------------------------+
    | 1       | Winston   | winston@leetcode.com    |
    | 2       | Jonathan  | jonathanisgreat         |
    | 3       | Annabelle | bella-@leetcode.com     |
    | 4       | Sally     | sally.come@leetcode.com |
    | 5       | Marwan    | quarz#2020@leetcode.com |
    | 6       | David     | david69@gmail.com       |
    | 7       | Shapiro   | .shapo@leetcode.com     |
    +---------+-----------+-------------------------+
    Output:
    +---------+-----------+-------------------------+
    | user_id | name      | mail                    |
    +---------+-----------+-------------------------+
    | 1       | Winston   | winston@leetcode.com    |
    | 3       | Annabelle | bella-@leetcode.com     |
    | 4       | Sally     | sally.come@leetcode.com |
    +---------+-----------+-------------------------+
    Explanation:
    The mail of user 2 does not have a domain.
    The mail of user 5 has the # sign which is not allowed.
    The mail of user 6 does not have the leetcode domain.
    The mail of user 7 starts with a period.

---

# Solution

## Intuition

The problem involves searching for valid emails in a Users database. To solve this problem, we need to use SQL's pattern
matching feature - regular expressions (REGEXP). We need to construct a regular expression pattern that matches the
conditions given for a valid email.

## Approach

The conditions for a valid email are as follows:

1. The prefix name can contain letters (upper or lower case), digits, underscore '_', period '.', and/or dash '-'.
2. The prefix name must start with a letter.
3. The domain should be '@leetcode.com'.

Given these conditions, we can construct a regular expression as follows:

```
'^[a-zA-Z][a-zA-Z0-9._-]*@leetcode\\.com$'
```

Let's break down this regular expression:

- `^[a-zA-Z]`: This matches the start of the string. It should start with either a lowercase or uppercase letter.

- `[a-zA-Z0-9._-]*`: This matches any number (including zero) of lowercase letters, uppercase letters, digits,
  underscores, periods, and dashes. This constitutes the rest of the prefix name.

- `@leetcode\\.com$`: This matches the exact domain '@leetcode.com' at the end of the string. The double backslashes are
  needed to escape the dot character because a single dot has a special meaning in regular expressions.

The SQL query will select all rows from the Users table where the 'mail' column matches this regular expression.

## Complexity

- Time complexity:
  The time complexity is O(n) because the REGEXP operator needs to scan through all the characters in the 'mail' column
  for each row in the Users table. Here, n is the total number of characters in the 'mail' column across all rows.

- Space complexity:
  The space complexity is O(1) because the REGEXP operation does not require any additional space that grows with the
  size of the input. The result of the operation is a boolean value for each row.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)