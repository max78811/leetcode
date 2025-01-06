-- https://leetcode.com/problems/duplicate-emails/description/

select email as Email from person
group by email
having count(email) > 1