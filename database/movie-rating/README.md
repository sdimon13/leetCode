# 1341. Movie Rating

Level: `Medium`

https://leetcode.com/problems/movie-rating/

---

# Description

Table: `Movies`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | movie_id      | int     |
    | title         | varchar |
    +---------------+---------+
    movie_id is the primary key for this table.
    title is the name of the movie.

Table: `Users`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | user_id       | int     |
    | name          | varchar |
    +---------------+---------+
    user_id is the primary key for this table.

Table: `MovieRating`

    +---------------+---------+
    | Column Name   | Type    |
    +---------------+---------+
    | movie_id      | int     |
    | user_id       | int     |
    | rating        | int     |
    | created_at    | date    |
    +---------------+---------+
    (movie_id, user_id) is the primary key for this table.
    This table contains the rating of a movie by a user in their review.
    created_at is the user's review date.

Write an SQL query to:

Find the name of the user who has rated the greatest number of movies. In case of a tie, return the lexicographically
smaller user name.
Find the movie name with the **highest average** rating in `February 2020`. In case of a tie, return the
lexicographically smaller movie name.
The query result format is in the following example.

## Example 1:

    Input:
    Movies table:
    +-------------+--------------+
    | movie_id    |  title       |
    +-------------+--------------+
    | 1           | Avengers     |
    | 2           | Frozen 2     |
    | 3           | Joker        |
    +-------------+--------------+
    Users table:
    +-------------+--------------+
    | user_id     |  name        |
    +-------------+--------------+
    | 1           | Daniel       |
    | 2           | Monica       |
    | 3           | Maria        |
    | 4           | James        |
    +-------------+--------------+
    MovieRating table:
    +-------------+--------------+--------------+-------------+
    | movie_id    | user_id      | rating       | created_at  |
    +-------------+--------------+--------------+-------------+
    | 1           | 1            | 3            | 2020-01-12  |
    | 1           | 2            | 4            | 2020-02-11  |
    | 1           | 3            | 2            | 2020-02-12  |
    | 1           | 4            | 1            | 2020-01-01  |
    | 2           | 1            | 5            | 2020-02-17  |
    | 2           | 2            | 2            | 2020-02-01  |
    | 2           | 3            | 2            | 2020-03-01  |
    | 3           | 1            | 3            | 2020-02-22  |
    | 3           | 2            | 4            | 2020-02-25  |
    +-------------+--------------+--------------+-------------+
    Output:
    +--------------+
    | results      |
    +--------------+
    | Daniel       |
    | Frozen 2     |
    +--------------+
    Explanation:
    Daniel and Monica have rated 3 movies ("Avengers", "Frozen 2" and "Joker") but Daniel is smaller lexicographically.
    Frozen 2 and Joker have a rating average of 3.5 in February but Frozen 2 is smaller lexicographically.

---

# Solution

## Intuition

The task asks for two separate pieces of information - the name of the user who has rated the most movies and the name
of the movie with the highest average rating in February 2020. Thus, the solution would involve writing two separate
queries, and then combining them using UNION ALL to form a single query.

## Approach

1. For finding the name of the user who has rated the most movies, we join the Users table with the MovieRating table on
   the user_id column. Then, we group by the user name, and order the results by the count of the movie_id (descending),
   and by the user name (to handle the case of a tie). We limit the results to 1 to get only the top user.

2. For finding the name of the movie with the highest average rating in February 2020, we join the MovieRating table
   with the Movies table on the movie_id column. We add a WHERE clause to filter the reviews that were created in
   February 2020. Then, we group by the movie title, and order the results by the average rating (descending), and by
   the movie title (to handle the case of a tie). We limit the results to 1 to get only the top movie.

3. Finally, we combine these two queries using UNION ALL to get the final result.

## Complexity

- Time complexity:
  The time complexity could be up to O(n*m), where n is the number of rows in the MovieRating table, and m is the
  maximum number of ratings done by a single user or the maximum number of ratings for a single movie. This is because
  we're grouping the MovieRating table by user_id and movie_id, and calculating counts and averages on these groups. If
  there's an index on user_id and movie_id in the MovieRating table, the complexity could be reduced.

- Space complexity:
  The space complexity is likely to be O(u + m), where u is the number of unique users and m is the number of unique
  movies. The SQL engine needs to create temporary groups to calculate count and average for each user and movie. In
  worst case, this could be as large as the number of unique users and unique movies.

---

# Code

Here are the solution and test files:

- [Solution Code](./solution.sql)
- [Test Code](./solution_test.go)