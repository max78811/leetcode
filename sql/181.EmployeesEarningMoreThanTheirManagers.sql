-- https://leetcode.com/problems/employees-earning-more-than-their-managers/description/

select m.name as Employee from Employee m
join Employee e
on e.id = m.managerId
where m.salary > e.salary
