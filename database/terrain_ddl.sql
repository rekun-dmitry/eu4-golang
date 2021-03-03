create schema if not exists land_warfare;
drop table if exists land_warfare.terrain_war cascade;
create table land_warfare.terrain_war
(terrain_name varchar(50) not NULL,
supply_limit varchar(50),
local_defensiveness varchar(50),
movement_cost varchar(50),
attacker_penalty varchar(50),
local_development_cost varchar(50));



