-- https://leetcode.com/problems/combine-two-tables/description/

select firstName, lastName, city, state from person
left join address
on person.personId = address.personId
