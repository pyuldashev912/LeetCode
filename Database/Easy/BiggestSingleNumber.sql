/* Write your T-SQL query statement below */
SELECT MAX(num) as 'num'
FROM (
    SELECT num
    FROM mynumbers
    GROUP BY num
    HAVING COUNT(num) = 1
) sub_query