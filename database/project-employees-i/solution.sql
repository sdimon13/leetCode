SELECT p.project_id, ROUND(SUM(e.experience_years) / COUNT(e.experience_years), 2) AS average_years
FROM Project AS p
         JOIN Employee AS e ON p.employee_id = e.employee_id
GROUP BY p.project_id