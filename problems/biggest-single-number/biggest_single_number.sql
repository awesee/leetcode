# Write your MySQL query statement below

# Method 1
SELECT num
FROM
  `number`
GROUP BY num
HAVING COUNT(num) = 1;

# Method 2
SELECT MAX(num) AS num
FROM
  (SELECT num
   FROM
     number
   GROUP BY num
   HAVING COUNT(num) = 1) AS t;
