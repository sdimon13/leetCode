SELECT p.product_name, sum(o.unit) AS unit
FROM Products AS p
         LEFT JOIN Orders AS o ON p.product_id = o.product_id
WHERE order_date BETWEEN '2020-02-01' AND '2020-02-29'
GROUP BY o.product_id
HAVING unit >= 100