# Write your MySQL query statement below

SELECT MAX(num) AS num
FROM
  (SELECT num
   FROM
     `my_numbers`
   GROUP BY num
   HAVING COUNT(num) = 1) AS t;
