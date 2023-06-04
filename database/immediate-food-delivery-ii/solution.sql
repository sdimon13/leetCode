SELECT ROUND(100 * SUM(CASE WHEN D.order_date = D.customer_pref_delivery_date THEN 1 ELSE 0 END) / COUNT(*),
             2) AS immediate_percentage
FROM Delivery D
WHERE (D.customer_id, D.order_date) IN (SELECT customer_id, MIN(order_date) FROM Delivery GROUP BY customer_id);