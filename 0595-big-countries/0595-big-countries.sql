# Write your MySQL query statement below
-- create index indx_area on World(area)
-- create index indx_pop on World(population)
select name ,population,area from World where area>=3000000 or population>=25000000