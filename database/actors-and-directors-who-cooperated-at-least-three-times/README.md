# 1050. Actors and Directors Who Cooperated At Least Three Times

Level: `Easy`

https://leetcode.com/problems/actors-and-directors-who-cooperated-at-least-three-times/

---

# Description

Table: `ActorDirector`

    +-------------+---------+
    | Column Name | Type    |
    +-------------+---------+
    | actor_id    | int     |
    | director_id | int     |
    | timestamp   | int     |
    +-------------+---------+
    timestamp is the primary key column for this table.

Write a SQL query for a report that provides the pairs `(actor_id, director_id)` where the actor has cooperated with the
director at least three times.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    ActorDirector table:
    +-------------+-------------+-------------+
    | actor_id    | director_id | timestamp   |
    +-------------+-------------+-------------+
    | 1           | 1           | 0           |
    | 1           | 1           | 1           |
    | 1           | 1           | 2           |
    | 1           | 2           | 3           |
    | 1           | 2           | 4           |
    | 2           | 1           | 5           |
    | 2           | 1           | 6           |
    +-------------+-------------+-------------+
    Output:
    +-------------+-------------+
    | actor_id    | director_id |
    +-------------+-------------+
    | 1           | 1           |
    +-------------+-------------+
    Explanation: The only pair is (1, 1) where they cooperated exactly 3 times.

---

# Solution

## Intuition

The problem is asking to find pairs of actors and directors who have worked together at least three times. Given that
the database table 'ActorDirector' holds information about the actor_id, director_id and the timestamp of their
cooperation, we can use SQL commands to group and filter out the required information.

## Approach

We need to use 'GROUP BY' SQL command to group rows that have the same values in specified columns into aggregated data.
Here, we are interested in groups formed by 'actor_id' and 'director_id'. The aggregate function 'COUNT' is used to
count the number of times each pair of 'actor_id' and 'director_id' appears, indicating their number of collaborations.

Finally, we need to filter the groups to only include those where the actor and the director have worked together at
least three times. This is achieved using 'HAVING' clause which allows us to filter on aggregate functions.

## Complexity

- Time complexity:
  The time complexity is O(n), where n is the number of rows in the 'ActorDirector' table. This is because we are
  grouping all rows based on 'actor_id' and 'director_id' and then counting the number of times each pair appears.

- Space complexity:
  The space complexity is also O(n), because in the worst-case scenario (every actor and director have worked together
  exactly once), we would need a separate group for each row.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)