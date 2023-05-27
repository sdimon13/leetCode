# 1729. Find Followers Count

Level: `Easy`

https://leetcode.com/problems/find-followers-count/

---

# Description

Table: `Followers`

    +-------------+------+
    | Column Name | Type |
    +-------------+------+
    | user_id     | int  |
    | follower_id | int  |
    +-------------+------+
    (user_id, follower_id) is the primary key for this table.
    This table contains the IDs of a user and a follower in a social media app where the follower follows the user.

Write an SQL query that will, for each user, return the number of followers.

Return the result table ordered by user_id in ascending order.

The query result format is in the following example.

## Example 1:

    Input:
    Followers table:
    +---------+-------------+
    | user_id | follower_id |
    +---------+-------------+
    | 0       | 1           |
    | 1       | 0           |
    | 2       | 0           |
    | 2       | 1           |
    +---------+-------------+
    Output:
    +---------+----------------+
    | user_id | followers_count|
    +---------+----------------+
    | 0       | 1              |
    | 1       | 1              |
    | 2       | 2              |
    +---------+----------------+
    Explanation:
    The followers of 0 are {1}
    The followers of 1 are {0}
    The followers of 2 are {0,1}

---

# Solution

## Intuition

Given the problem description, we need to count the followers for each user. In the given "Followers" table, the
column "user_id" represents the user, and "follower_id" represents the user's followers. A simple way to solve this
problem would be to group the data by the user_id and count the number of follower_id for each user_id.

## Approach

We'll leverage SQL's **GROUP BY** clause to group the data by user_id, then use the **COUNT** function to count the
number of follower_id for each user_id. We'll order the result by user_id in ascending order as required by the problem
statement.

## Complexity

- Time complexity:
  The time complexity is O(n log n), where n is the number of rows in the table. This is because the **GROUP BY** clause
  involves sorting the table.

- Space complexity:
  The space complexity is O(n), where n is the number of unique user_ids. This is because the result of the **GROUP BY**
  clause is a table with a row for each unique user_id.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)