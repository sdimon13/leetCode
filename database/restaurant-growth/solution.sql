SELECT DISTINCT c1.visited_on,
                SUM(c2.amount)               AS amount,
                ROUND(SUM(c2.amount) / 7, 2) AS average_amount
FROM Customer AS c1
         JOIN Customer AS c2 ON c2.visited_on BETWEEN DATE_SUB(c1.visited_on, INTERVAL 6 DAY) AND c1.visited_on
WHERE c1.visited_on >= DATE_ADD(
        (SELECT MIN(visited_on) FROM CUSTOMER), INTERVAL 6 DAY
    )
GROUP BY c1.visited_on, c1.customer_id
ORDER BY c1.visited_on