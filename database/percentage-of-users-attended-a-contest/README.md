# 1633. Percentage of Users Attended a Contest

Level: `Easy`

https://leetcode.com/problems/percentage-of-users-attended-a-contest/

---

# Description

Table: `Users`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | user_id     | int     |
    | user_name   | varchar |
    +-------------+---------+
    user_id is the primary key for this table.
    Each row of this table contains the name and the id of a user.

Table: `Register`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | contest_id  | int     |
    | user_id     | int     |
    +-------------+---------+
    (contest_id, user_id) is the primary key for this table.
    Each row of this table contains the id of a user and the contest they registered into.

Write an SQL query to find the percentage of the users registered in each contest rounded to **two decimals**.

Return the result table ordered by `percentage` in **descending order**. In case of a tie, order it by `contest_id` in *
*ascending order**.

The query result format is in the following example.

## Example 1:

    Input:
    Users table:
    +---------+-----------+
    | user_id | user_name |
    +---------+-----------+
    | 6       | Alice     |
    | 2       | Bob       |
    | 7       | Alex      |
    +---------+-----------+
    Register table:
    +------------+---------+
    | contest_id | user_id |
    +------------+---------+
    | 215        | 6       |
    | 209        | 2       |
    | 208        | 2       |
    | 210        | 6       |
    | 208        | 6       |
    | 209        | 7       |
    | 209        | 6       |
    | 215        | 7       |
    | 208        | 7       |
    | 210        | 2       |
    | 207        | 2       |
    | 210        | 7       |
    +------------+---------+
    Output:
    +------------+------------+
    | contest_id | percentage |
    +------------+------------+
    | 208        | 100.0      |
    | 209        | 100.0      |
    | 210        | 100.0      |
    | 215        | 66.67      |
    | 207        | 33.33      |
    +------------+------------+
    Explanation:
    All the users registered in contests 208, 209, and 210. The percentage is 100% and we sort them in the answer table by contest_id in ascending order.
    Alice and Alex registered in contest 215 and the percentage is ((2/3) * 100) = 66.67%
    Bob registered in contest 207 and the percentage is ((1/3) * 100) = 33.33%

---

# Solution

## Intuition

The problem is about calculating the percentage of users who registered for each contest. We will first calculate the
total distinct users in each contest and then divide it by the total number of users in the Users table to get the
percentage. The MySQL function ROUND is used to limit the result to two decimal places.

## Approach

1. Use the COUNT DISTINCT function to find the unique number of users in each contest.
2. Use a subquery to find the total number of distinct users in the Users table.
3. Calculate the percentage of users for each contest by dividing the number of distinct users in each contest by the
   total number of users. Multiply this by 100 and round to two decimal places.
4. Group by the contest_id and order the result by percentage in descending order and by contest_id in ascending order.

## Complexity

- Time complexity:
  The time complexity is O(n) where n is the number of rows in the Register table. This is because we have to scan each
  row in the Register table once.

- Space complexity:
  The space complexity is O(m), where m is the number of unique contest_id values in the Register table. This is because
  we store each unique contest_id and its corresponding count in the result set.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)