-- https://leetcode.com/problems/big-countries/description

select name, population, area from World
where World.area >= 3000000 or World.population >= 25000000
