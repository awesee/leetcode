# Write your MySQL query statement below

SELECT `name` AS 'Name'
FROM
  Candidate
  JOIN (
         SELECT Candidateid
         FROM
           Vote
         GROUP BY
           Candidateid
         ORDER BY
           COUNT(*) DESC
         LIMIT 1
       ) AS winner
WHERE
  Candidate.id = winner.Candidateid;
