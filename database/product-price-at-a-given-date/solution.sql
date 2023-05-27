SELECT DISTINCT product_id,
                CASE
                    WHEN min(p1.change_date) > DATE("2019-08-16") THEN 10
                    ELSE (SELECT new_price
                          FROM Products AS p2
                          WHERE p1.product_id = p2.product_id
                            AND p2.change_date <= DATE("2019-08-16")
                          ORDER BY change_date DESC
                          LIMIT 1)
                    END AS price
FROM Products AS p1
GROUP BY p1.product_id