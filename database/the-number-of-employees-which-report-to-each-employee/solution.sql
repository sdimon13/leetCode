SELECT e2.employee_id, e2.name, COUNT(e1.reports_to) AS reports_count, ROUND(AVG(e1.age)) AS average_age
FROM Employees AS e1
         JOIN Employees AS e2 ON e1.reports_to = e2.employee_id
GROUP BY e1.reports_to
ORDER BY employee_id