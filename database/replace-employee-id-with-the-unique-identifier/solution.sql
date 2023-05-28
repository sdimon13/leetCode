SELECT IFNULL(uni.unique_id, null) AS unique_id, e.name
FROM Employees AS e
         LEFT JOIN EmployeeUNI AS uni ON e.id = uni.id