# 1407. Top Travellers

Level: `Easy`

https://leetcode.com/problems/top-travellers/

---

# Description

Table: `Users`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | id            | int     |
    | name          | varchar |
    +---------------+---------+
    id is the primary key for this table.
    name is the name of the user.

Table: `Rides`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | id            | int     |
    | user_id       | int     |
    | distance      | int     |
    +---------------+---------+
    id is the primary key for this table.
    user_id is the id of the user who traveled the distance "distance".

Write an SQL query to report the distance traveled by each user.

Return the result table ordered by `travelled_distance` in **descending order**, if two or more users traveled the same
distance, order them by their `name` in **ascending order**.

The query result format is in the following example.

## Example 1:

    Input:
    Users table:
    +------+-----------+
    | id   | name      |
    +------+-----------+
    | 1    | Alice     |
    | 2    | Bob       |
    | 3    | Alex      |
    | 4    | Donald    |
    | 7    | Lee       |
    | 13   | Jonathan  |
    | 19   | Elvis     |
    +------+-----------+
    Rides table:
    +------+----------+----------+
    | id   | user_id  | distance |
    +------+----------+----------+
    | 1    | 1        | 120      |
    | 2    | 2        | 317      |
    | 3    | 3        | 222      |
    | 4    | 7        | 100      |
    | 5    | 13       | 312      |
    | 6    | 19       | 50       |
    | 7    | 7        | 120      |
    | 8    | 19       | 400      |
    | 9    | 7        | 230      |
    +------+----------+----------+
    Output:
    +----------+--------------------+
    | name     | travelled_distance |
    +----------+--------------------+
    | Elvis    | 450                |
    | Lee      | 450                |
    | Bob      | 317                |
    | Jonathan | 312                |
    | Alex     | 222                |
    | Alice    | 120                |
    | Donald   | 0                  |
    +----------+--------------------+
    Explanation:
    Elvis and Lee traveled 450 miles, Elvis is the top traveler as his name is alphabetically smaller than Lee.
    Bob, Jonathan, Alex, and Alice have only one ride and we just order them by the total distances of the ride.
    Donald did not have any rides, the distance traveled by him is 0.

---

# Solution

## Intuition

The problem requires us to retrieve the total distance travelled by each user and sort them according to the distance in
descending order. If the distance is the same, we need to sort them by their names in ascending order. For users who
didn't travel, we should display their travelled distance as 0. To achieve this, we can use SQL's JOIN, GROUP BY, and
ORDER BY clauses.

## Approach

We can solve this problem by performing a LEFT JOIN operation on the **Users** and **Rides** tables on the **user_id**
field. The LEFT JOIN ensures that all users are included in the result, even if they don't have any rides. We then group
the results by **user_id** and calculate the sum of the **distance** for each group using the SQL aggregate function *
*SUM**.

When summing the **distance**, we use the SQL **IFNULL** function to handle cases where the user doesn't have any
rides (in which case, **distance** would be **NULL**). The **IFNULL** function replaces **NULL** values with the
specified value (0 in this case).

Finally, we sort the results by **travelled_distance** in descending order and by **name** in ascending order.

## Complexity

- Time complexity:
  The time complexity depends on the implementation of the SQL engine. Generally, performing a join operation is O(n +
  m), where n and m are the number of rows in the Users and Rides tables respectively. The subsequent grouping and
  ordering operations can also contribute to the time complexity, but this varies based on the specific SQL engine.

- Space complexity:
  The space complexity also depends on the SQL engine implementation. In general, the space required would be enough to
  store the result set of the join operation, so it's O(n + m) as well.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)