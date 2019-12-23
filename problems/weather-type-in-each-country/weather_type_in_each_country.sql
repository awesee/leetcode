# Write your MySQL query statement below

SELECT country_name,
       CASE
           WHEN avg(weather_state) <= 15 THEN "Cold"
           WHEN avg(weather_state) >= 25 THEN "Hot"
           ELSE "Warm" END AS weather_type
FROM Countries
         INNER JOIN Weather ON Countries.country_id = Weather.country_id
WHERE LEFT(DAY, 7) = '2019-11'
GROUP BY country_name
