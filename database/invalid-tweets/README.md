# 1683. Invalid Tweets

Level: `Easy`

https://leetcode.com/problems/invalid-tweets/

---

# Description

Table: `Tweets`

    +----------------+---------+
    | Column Name    | Type    |
    +----------------+---------+
    | tweet_id       | int     |
    | content        | varchar |
    +----------------+---------+
    tweet_id is the primary key for this table.
    This table contains all the tweets in a social media app.

Write an SQL query to find the IDs of the invalid tweets. The tweet is invalid if the number of characters used in the
content of the tweet is **strictly greater** than `15`.

Return the result table in **any order**.

The query result format is in the following example.

## Example 1:

    Input:
    Tweets table:
    +----------+----------------------------------+
    | tweet_id | content                          |
    +----------+----------------------------------+
    | 1        | Vote for Biden                   |
    | 2        | Let us make America great again! |
    +----------+----------------------------------+
    Output:
    +----------+
    | tweet_id |
    +----------+
    | 2        |
    +----------+
    Explanation:
    Tweet 1 has length = 14. It is a valid tweet.
    Tweet 2 has length = 32. It is an invalid tweet.

---

# Solution

## Intuition

To solve this problem, we need to identify the tweets with content that has more than 15 characters. We can achieve this
by using the **length()** function in SQL to determine the length of the content and then filter out the tweets that
have a length greater than 15.

## Approach

The approach is straightforward. We can write a SQL query that selects the **tweet_id** from the **Tweets** table where
the length of the **content** is greater than 15.

## Complexity

- Time complexity:
  The time complexity of this approach is O(n), where n is the number of tweets in the table. The query needs to check
  the length of each tweet's content.

- Space complexity:
  The space complexity is O(1) because we are only returning the IDs of the invalid tweets and not storing any
  additional data.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)