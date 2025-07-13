-- user
-- id | firstname | lastname | birth
-- 1 | Ivan | Petrov | 1996-05-01
-- 2 | Anna | Petrova | 1999-06-01
-- 3 | Anna | Petrova | 1990-10-02
--
-- purchase
-- sku| price | user_id | date
-- 1  | 5500 | 1 | 2021-02-15
-- 1  | 5700 | 1 | 2021-01-15
-- 2  | 4000 | 1 | 2021-02-14
-- 3  | 8000 | 2 | 2021-03-01
-- 4  | 400 | 2 | 2021-03-02
--
-- ban_list
-- user_id | date_from
-- 1 | 2021-03-08
--
-- Нужно вывести:
-- Вывести уникальные комбинации пользователя и id товара для всех покупок, совершенных пользователями до того, как их забанили.
-- Отсортировать сначала по имени пользователя, потом по SKU

select distinct users.id
from users
         join purchase
              on users.id = purchase.user_id
         join ban_list
              on users.id = ban_list.user_id
where purchase.date < ban_list.date_from
order by (users.firstname, purchase.sku);

-- Найти пользователей, которые совершили покупок на сумму больше 5000р.
-- Вывести их имена в формате id пользователя | имя | фамилия | сумма покупок

select users.id, users.firstname, users.lastname, sum(purchase.price)
from users join purchase
on users.id = purchase.user_id
group by (users.id, users.firstname, users.lastname)
having sum(purchase.price) > 5000;
