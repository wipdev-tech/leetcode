SELECT
    CASE
        WHEN (SELECT COUNT(DISTINCT(salary)) FROM Employee) < 2 THEN null
        ELSE
            (SELECT DISTINCT(salary)
            FROM Employee
            ORDER BY salary DESC
            OFFSET 1
            LIMIT 1)
    END
    AS SecondHighestSalary; 
