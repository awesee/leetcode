# Write your MySQL query statement below

# Method 1
UPDATE salary
SET sex = IF(sex = 'm', 'f', 'm');

# Method 2
UPDATE salary
SET sex = CASE sex
          WHEN 'm'
            THEN 'f'
          ELSE 'm' END;
