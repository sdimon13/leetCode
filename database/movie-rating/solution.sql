(SELECT u.name AS results
 FROM Users AS u
          LEFT JOIN MovieRating AS mr ON u.user_id = mr.user_id
 GROUP BY u.name
 ORDER BY COUNT(mr.movie_id) DESC, u.name
 LIMIT 1)
UNION ALL
(SELECT m.title AS results
 FROM MovieRating AS mr
          LEFT JOIN Movies AS m ON mr.movie_id = m.movie_id
 WHERE YEAR(mr.created_at) = 2020
   AND MONTH(mr.created_at) = 2
 GROUP BY m.title
 ORDER BY AVG(mr.rating) DESC, m.title
 LIMIT 1)