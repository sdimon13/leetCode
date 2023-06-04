# 550. Game Play Analysis IV

Level: `Medium`

https://leetcode.com/problems/game-play-analysis-iv/

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

Write an SQL query to report the **fraction** of players that logged in again on the day after the day they first logged
in, **rounded to 2 decimal places**. In other words, you need to count the number of players that logged in for at least
two consecutive days starting from their first login date, then divide that number by the total number of players.

The query result format is in the following example.

## Example 1:

    Input:
    Activity table:
    +-----------+-----------+------------+--------------+
    | player_id | device_id | event_date | games_played |
    +-----------+-----------+------------+--------------+
    | 1         | 2         | 2016-03-01 | 5            |
    | 1         | 2         | 2016-03-02 | 6            |
    | 2         | 3         | 2017-06-25 | 1            |
    | 3         | 1         | 2016-03-02 | 0            |
    | 3         | 4         | 2018-07-03 | 5            |
    +-----------+-----------+------------+--------------+
    Output:
    +-----------+
    | fraction  |
    +-----------+
    | 0.33      |
    +-----------+
    Explanation:
    Only the player with id 1 logged back in after the first day he had logged in so the answer is 1/3 = 0.33

---

# Solution

## Intuition

This problem requires us to determine the proportion of players who have logged in for at least two consecutive days
starting from their first login date. We need to make use of the window function and perform some date manipulations in
SQL.

## Approach

1. First, we create a subquery that lists all players along with the date of their first login using the MIN function in
   conjunction with the window function OVER. This allows us to treat each player separately.
2. After having the date of first login for each player, we filter out only those rows where the difference between the
   event_date and the first login date is exactly 1, which means the player logged in the next day.
3. We then count the distinct number of players that fulfill this criteria.
4. Finally, we divide this count by the total number of distinct players in the Activity table to get the required
   fraction.

## Complexity

- Time complexity:
  The time complexity is O(n) because we traverse through the entire Activity table.

- Space complexity:
  The space complexity is also O(n) because the window function creates a new table with the same number of rows as the
  Activity table.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)