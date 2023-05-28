# 602. Friend Requests II: Who Has the Most Friends

Level: `Medium`

https://leetcode.com/problems/friend-requests-ii-who-has-the-most-friends/

---

# Description

Table: `RequestAccepted`

    +----------------+---------+
    | Column Name    | Type    |
    +----------------+---------+
    | requester_id   | int     |
    | accepter_id    | int     |
    | accept_date    | date    |
    +----------------+---------+
    (requester_id, accepter_id) is the primary key for this table.
    This table contains the ID of the user who sent the request, the ID of the user who received the request, and the date when the request was accepted.

Write an SQL query to find the people who have the most friends and the most friends number.

The test cases are generated so that only one person has the most friends.

The query result format is in the following example.

## Example 1:

    Input:
    RequestAccepted table:
    +--------------+-------------+-------------+
    | requester_id | accepter_id | accept_date |
    +--------------+-------------+-------------+
    | 1            | 2           | 2016/06/03  |
    | 1            | 3           | 2016/06/08  |
    | 2            | 3           | 2016/06/08  |
    | 3            | 4           | 2016/06/09  |
    +--------------+-------------+-------------+
    Output:
    +----+-----+
    | id | num |
    +----+-----+
    | 3  | 3   |
    +----+-----+
    Explanation:
    The person with id 3 is a friend of people 1, 2, and 4, so he has three friends in total, which is the most number than any others.

---

# Solution

## Intuition

Given the problem's conditions, we need to find the user who has the most friends. Friends are defined as any user who
has accepted a friend request or sent one. Therefore, our initial idea was to get all user IDs involved in friend
requests and count how often each ID appears. The user with the most appearances will have the most friends.

## Approach

Our approach consists of several steps. First, we create a subquery that unites all **requester_id** and **accepter_id**
from the **RequestAccepted** table. We use the **UNION ALL** operator to include all IDs, including duplicates.

Next, we group the subquery results by ID and count the number of friends for each ID. This is done using the **COUNT(*)
** function, which counts the number of rows for each ID.

Finally, we sort the results by the number of friends in descending order and select the user with the most friends. If
there are multiple users with the same number of friends, only one is chosen, as the problem's conditions guarantee that
only one user has the most friends.

## Complexity

- Time complexity:
  O(n log n)

The time complexity of the query primarily depends on the sorting operation. In the worst-case scenario, sorting has a
time complexity of O(n log n), where n is the total number of rows in the RequestAccepted table.

- Space complexity:
  O(n)

The space complexity of the query is O(n), where n is the total number of rows in the RequestAccepted table. This is
because the friends subquery may require space proportional to the size of the RequestAccepted table.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)