-- Построить оптимальный индекс

SELECT * FROM employee
WHERE gender = 'm' AND salary > 300000 AND age = 20
ORDER BY created_at

create index on employee (age, salary, created_at)
