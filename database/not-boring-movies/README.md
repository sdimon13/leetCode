# 620. Not Boring Movies

Level: `Easy`

https://leetcode.com/problems/not-boring-movies/

---

# Description

Table: `Cinema`

    +----------------+----------+
    | Column Name    | Type     |
    +----------------+----------+
    | id             | int      |
    | movie          | varchar  |
    | description    | varchar  |
    | rating         | float    |
    +----------------+----------+
    id is the primary key for this table.
    Each row contains information about the name of a movie, its genre, and its rating.
    rating is a 2 decimal places float in the range [0, 10]

Write an SQL query to report the movies with an odd-numbered ID and a description that is not `"boring"`.

Return the result table ordered by `rating` in **descending order**.

The query result format is in the following example.

## Example 1:

    Input:
    Cinema table:
    +----+------------+-------------+--------+
    | id | movie      | description | rating |
    +----+------------+-------------+--------+
    | 1  | War        | great 3D    | 8.9    |
    | 2  | Science    | fiction     | 8.5    |
    | 3  | irish      | boring      | 6.2    |
    | 4  | Ice song   | Fantacy     | 8.6    |
    | 5  | House card | Interesting | 9.1    |
    +----+------------+-------------+--------+
    Output:
    +----+------------+-------------+--------+
    | id | movie      | description | rating |
    +----+------------+-------------+--------+
    | 5  | House card | Interesting | 9.1    |
    | 1  | War        | great 3D    | 8.9    |
    +----+------------+-------------+--------+
    Explanation:
    We have three movies with odd-numbered IDs: 1, 3, and 5. The movie with ID = 3 is boring so we do not include it in the answer.

---

# Solution

## Intuition

This problem can be solved by writing a SQL query to filter the rows in the given Cinema table. The criteria to filter
is whether the movie's id is odd-numbered and its description is not "boring".

## Approach

To solve this problem, we can follow these steps:

1. Use the WHERE clause in the SQL query to filter out rows from the Cinema table based on the criteria. We can use the
   modulo operation to check whether the id is odd, and we can compare the description to the string "boring" to see
   whether it is not equal.

2. Order the remaining rows by the rating column in descending order using the ORDER BY clause. By specifying DESC, we
   ensure that the highest ratings come first.

## Complexity

- Time complexity:
  The time complexity is O(n log n), where n is the number of rows in the table. This is because sorting the table based
  on the rating is an O(n log n) operation.

- Space complexity:
  The space complexity is O(n), where n is the number of rows in the table. This is because the query creates a new
  table to store the output.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)