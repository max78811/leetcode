create table dbo.clientStore
(
    id   int primary key,
    name varchar(50) not null
);

create table dbo.orderStore
(
    id       int primary key,
    clientId int   not null,
    amount   money not null
);

insert into dbo.clientStore (id, name)
values (1, 'Петя'),
       (2, 'Вася'),
       (3, 'Саша');

insert into dbo.orderStore (id, clientId, amount)
values (1, 1, 100),
       (2, 2, 200);


-- Напишите запрос, которые возвращает ид клиентов и сумму их заказов, у которых сумма заказов больше 150 рублей?

select orderStore.clientId, sum(amount)
from orderStore
group by orderStore.clientId
having sum(amount) > 150;

--     Какой результат вернут запросы
select c.id as clientId, c.name, o.id as orderId, o.amount
from dbo.clientStore as c
         left join dbo.orderStore as o on o.clientId = c.id and o.amount > 100


-- 1 Петя   null
-- 2 Вася 1 200
-- 3 Саша   null

select c.id as clientId, c.name, o.id as orderId, o.amount
from dbo.clientStore as c
         left join dbo.orderStore as o on o.clientId = c.id
where o.amount > 100

--     2 Вася 1 200