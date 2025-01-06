-- https://leetcode.com/problems/managers-with-at-least-5-direct-reports/description

select e.name from employee e
join employee ee
on e.id = ee.managerId
group by e.id
having count(e.id) >= 5
