# 1141. User Activity for the Past 30 Days I

Level: `Easy`

https://leetcode.com/problems/user-activity-for-the-past-30-days-i/

---

# Description

Table: `Activity`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | user_id       | int     |
    | session_id    | int     |
    | activity_date | date    |
    | activity_type | enum    |
    +---------------+---------+
    There is no primary key for this table, it may have duplicate rows.
    The activity_type column is an ENUM of type ('open_session', 'end_session', 'scroll_down', 'send_message').
    The table shows the user activities for a social media website.
    Note that each session belongs to exactly one user.

Write an SQL query to find the daily active user count for a period of 30 days ending 2019-07-27 inclusively. A user was
active on someday if they made at least one activity on that day.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Activity table:
    +---------+------------+---------------+---------------+
    | user_id | session_id | activity_date | activity_type |
    +---------+------------+---------------+---------------+
    | 1       | 1          | 2019-07-20    | open_session  |
    | 1       | 1          | 2019-07-20    | scroll_down   |
    | 1       | 1          | 2019-07-20    | end_session   |
    | 2       | 4          | 2019-07-20    | open_session  |
    | 2       | 4          | 2019-07-21    | send_message  |
    | 2       | 4          | 2019-07-21    | end_session   |
    | 3       | 2          | 2019-07-21    | open_session  |
    | 3       | 2          | 2019-07-21    | send_message  |
    | 3       | 2          | 2019-07-21    | end_session   |
    | 4       | 3          | 2019-06-25    | open_session  |
    | 4       | 3          | 2019-06-25    | end_session   |
    +---------+------------+---------------+---------------+
    Output:
    +------------+--------------+
    | day        | active_users |
    +------------+--------------+
    | 2019-07-20 | 2            |
    | 2019-07-21 | 2            |
    +------------+--------------+
    Explanation: Note that we do not care about days with zero active users.

---

# Solution

## Intuition

The task is to find the daily active user count for a period of 30 days. A user is considered active if they performed
at least one activity in a given day. Therefore, we need to count the number of unique users for each day within the
specified period.

## Approach

We can solve this problem by using MySQL's built-in COUNT DISTINCT functionality to count unique users for each day.
Here's the step by step approach:

1. We need to filter out records that are outside the desired 30 days period. We can do this by using the **DATEDIFF**
   function, which calculates the number of days difference between two dates. We want to include only those records
   where the difference between the given end date '2019-07-27' and **activity_date** is less than 30 days.

2. After that, we need to group records by **activity_date**. This will allow us to count active users for each day
   separately.

3. In each group, we use **COUNT DISTINCT** to count the number of unique **user_ids**. This gives us the number of
   active users for each day.

## Complexity

- Time complexity:
  The time complexity is O(n), where n is the number of rows in the table. This is because we need to scan every row in
  the table to perform the grouping and counting.

- Space complexity:
  The space complexity is O(m), where m is the number of unique activity_dates in the table. This is because each unique
  date forms a group, and we need to maintain these groups in memory during the grouping and counting operations.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)