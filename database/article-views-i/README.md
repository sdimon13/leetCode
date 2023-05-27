# 1148. Article Views I

Level: `Easy`

https://leetcode.com/problems/article-views-i/

---

# Description

Table: `Views`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | article_id    | int     |
    | author_id     | int     |
    | viewer_id     | int     |
    | view_date     | date    |
    +---------------+---------+
    There is no primary key for this table, it may have duplicate rows.
    Each row of this table indicates that some viewer viewed an article (written by some author) on some date.
    Note that equal author_id and viewer_id indicate the same person.

Write an SQL query to find all the authors that viewed at least one of their own articles.

Return the result table sorted by `id` in ascending order.

The query result format is in the following example.

## Example 1:

    Input:
    Views table:
    +------------+-----------+-----------+------------+
    | article_id | author_id | viewer_id | view_date  |
    +------------+-----------+-----------+------------+
    | 1          | 3         | 5         | 2019-08-01 |
    | 1          | 3         | 6         | 2019-08-02 |
    | 2          | 7         | 7         | 2019-08-01 |
    | 2          | 7         | 6         | 2019-08-02 |
    | 4          | 7         | 1         | 2019-07-22 |
    | 3          | 4         | 4         | 2019-07-21 |
    | 3          | 4         | 4         | 2019-07-21 |
    +------------+-----------+-----------+------------+
    Output:
    +------+
    | id   |
    +------+
    | 4    |
    | 7    |
    +------+

---

# Solution

## Intuition

The problem asks to find all the authors who viewed their own articles. In the Views table, the 'author_id' column
represents the authors and 'viewer_id' represents the viewers. When an author views his/her own article, 'author_id'
and 'viewer_id' for that entry are equal. So we need to find distinct rows where 'author_id' equals 'viewer_id'.

## Approach

To solve this problem, we need to perform a SELECT operation with a condition. The condition being that the 'author_id'
is equal to 'viewer_id'. As we need distinct authors, we will use DISTINCT keyword to prevent duplicate entries. The
final result should be sorted by 'author_id' in ascending order.

## Complexity

- Time complexity:
  The time complexity of this solution is O(n), where n is the number of rows in the Views table. This is because the
  SQL operation will have to scan all the rows in the table to find where 'author_id' equals 'viewer_id'.

- Space complexity:
  The space complexity is O(m), where m is the number of unique authors who viewed their own articles. This space will
  be used to store the result set.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)