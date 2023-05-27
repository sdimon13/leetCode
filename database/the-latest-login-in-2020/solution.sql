SELECT user_id, MAX(time_stamp) AS last_stamp
FROM Logins
Where YEAR(time_stamp) = 2020
GROUP BY user_id