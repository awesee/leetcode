# Write your MySQL query statement below

SELECT
  department_salary.pay_month,
  department_id,
  CASE
  WHEN department_avg > company_avg
    THEN 'higher'
  WHEN department_avg < company_avg
    THEN 'lower'
  ELSE 'same'
  END AS comparison
FROM
  (SELECT
     department_id,
     avg(amount)                    AS department_avg,
     date_format(pay_date, '%Y-%m') AS pay_month
   FROM salary
     JOIN employee ON salary.employee_id = employee.employee_id
   GROUP BY department_id, pay_month
  ) AS department_salary
  JOIN (
         SELECT
           avg(amount)                    AS company_avg,
           date_format(pay_date, '%Y-%m') AS pay_month
         FROM salary
         GROUP BY date_format(pay_date, '%Y-%m')
       ) AS company_salary
    ON department_salary.pay_month = company_salary.pay_month;
