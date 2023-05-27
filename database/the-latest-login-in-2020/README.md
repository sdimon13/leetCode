# 1890. The Latest Login in 2020

Level: `Easy`

https://leetcode.com/problems/the-latest-login-in-2020/

---

# Description

Table: `Logins`

    +----------------+----------+
    | Column Name    | Type     |
    +----------------+----------+
    | user_id        | int      |
    | time_stamp     | datetime |
    +----------------+----------+
    (user_id, time_stamp) is the primary key for this table.
    Each row contains information about the login time for the user with ID user_id.

Write an SQL query to report the **latest** login for all users in the year `2020`. Do **not** include the users who did
not login in `2020`.

Return the result table **in any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Logins table:
    +---------+---------------------+
    | user_id | time_stamp          |
    +---------+---------------------+
    | 6       | 2020-06-30 15:06:07 |
    | 6       | 2021-04-21 14:06:06 |
    | 6       | 2019-03-07 00:18:15 |
    | 8       | 2020-02-01 05:10:53 |
    | 8       | 2020-12-30 00:46:50 |
    | 2       | 2020-01-16 02:49:50 |
    | 2       | 2019-08-25 07:59:08 |
    | 14      | 2019-07-14 09:00:00 |
    | 14      | 2021-01-06 11:59:59 |
    +---------+---------------------+
    Output:
    +---------+---------------------+
    | user_id | last_stamp          |
    +---------+---------------------+
    | 6       | 2020-06-30 15:06:07 |
    | 8       | 2020-12-30 00:46:50 |
    | 2       | 2020-01-16 02:49:50 |
    +---------+---------------------+
    Explanation:
    User 6 logged into their account 3 times but only once in 2020, so we include this login in the result table.
    User 8 logged into their account 2 times in 2020, once in February and once in December. We include only the latest one (December) in the result table.
    User 2 logged into their account 2 times but only once in 2020, so we include this login in the result table.
    User 14 did not login in 2020, so we do not include them in the result table.

---

# Solution

## Intuition

The task is to find the latest login timestamp for each user within the year 2020. This is a straightforward database
querying problem that requires grouping by user and filtering by the year of the timestamp. MySQL's **YEAR()** and *
*MAX()** functions can be used for this task.

## Approach

The approach is to filter the data by year, then group it by user id and find the maximum timestamp within each group.
This is done in two steps:

1. First, we filter the data to include only logins from 2020 using the **WHERE YEAR(time_stamp) = 2020** condition.
   This is done using the **YEAR()** function which extracts the year part from a date.

2. Then, we group the data by user id and find the latest login timestamp within each group. This is done using the *
   *GROUP BY user_id** clause and the **MAX()** function which finds the maximum value within each group. We select *
   *user_id** and **MAX(time_stamp)** to return the user id and the latest login timestamp.

## Complexity

- Time complexity:
  The time complexity is O(n), where n is the number of rows in the Logins table. This is because we need to scan each
  row in the table once to apply the WHERE clause and the GROUP BY clause.

- Space complexity:
  The space complexity is O(n), where n is the number of distinct users in the Logins table. This is because we need to
  store the group for each user when applying the GROUP BY clause.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)