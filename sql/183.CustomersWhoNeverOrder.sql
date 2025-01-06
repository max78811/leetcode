-- https://leetcode.com/problems/customers-who-never-order/description/

select customers.name as Customers from orders
right join customers
on customers.id = orders.customerId
where orders.id is null
