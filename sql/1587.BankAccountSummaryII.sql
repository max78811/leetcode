-- https://leetcode.com/problems/bank-account-summary-ii/description/

select users.name, sum(transactions.amount) as balance from users
join transactions
on users.account = transactions.account
group by users.name
having sum(transactions.amount) > 10000
