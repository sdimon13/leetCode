SELECT ROUND(SUM(i2.tiv_2016), 2) AS tiv_2016
FROM Insurance i2
WHERE i2.tiv_2015 IN (SELECT i1.tiv_2015
                      FROM Insurance i1
                      GROUP BY i1.tiv_2015
                      HAVING COUNT(*) > 1)
  AND NOT EXISTS (SELECT *
                  FROM Insurance i3
                  WHERE i3.lat = i2.lat
                    AND i3.lon = i2.lon
                    AND i3.pid <> i2.pid)