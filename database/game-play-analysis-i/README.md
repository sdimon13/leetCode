# 511. Game Play Analysis I

Level: `Easy`

https://leetcode.com/problems/game-play-analysis-i/

---

# Description

Table: `Activity`

    +--------------+---------+
    | Column Name  | Type    |
    +--------------+---------+
    | player_id    | int     |
    | device_id    | int     |
    | event_date   | date    |
    | games_played | int     |
    +--------------+---------+
    (player_id, event_date) is the primary key of this table.
    This table shows the activity of players of some games.
    Each row is a record of a player who logged in and played a number of games (possibly 0) before logging out on someday using some device.

Write an SQL query to report the **first login date** for each player.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Activity table:
    +-----------+-----------+------------+--------------+
    | player_id | device_id | event_date | games_played |
    +-----------+-----------+------------+--------------+
    | 1         | 2         | 2016-03-01 | 5            |
    | 1         | 2         | 2016-05-02 | 6            |
    | 2         | 3         | 2017-06-25 | 1            |
    | 3         | 1         | 2016-03-02 | 0            |
    | 3         | 4         | 2018-07-03 | 5            |
    +-----------+-----------+------------+--------------+
    Output:
    +-----------+-------------+
    | player_id | first_login |
    +-----------+-------------+
    | 1         | 2016-03-01  |
    | 2         | 2017-06-25  |
    | 3         | 2016-03-02  |
    +-----------+-------------+

---

# Solution

## Intuition

This problem asks for the first login date for each player. The "first" login date implies the earliest date. Therefore,
we need to find the minimum date for each player. SQL has built-in functions to aggregate data, which are ideal for this
scenario. Here we would use the **MIN()** function and **GROUP BY** clause.

## Approach

To get the earliest login date for each player, we will use the **GROUP BY** clause on the **player_id** column. This
groups all records with the same **player_id** together. Then, we will use the **MIN()** function on the **event_date**
column. This function will select the earliest date from each group.

## Complexity

- Time complexity:
  O(n), where n is the number of records in the Activity table. This is because we are grouping the whole table by
  player_id.

- Space complexity:
  O(n), where n is the number of unique player_ids. This is because we will have a unique row for each player_id in our
  result.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)