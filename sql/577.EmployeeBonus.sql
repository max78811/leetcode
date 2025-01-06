-- https://leetcode.com/problems/employee-bonus/description

select employee.name, bonus.bonus from employee
left join bonus
on employee.empId = bonus.empId
where bonus < 1000 or bonus is null
