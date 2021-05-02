create schema if not exists province;
drop table if exists province.coordinates cascade;
create table province.coordinates
(province_name varchar(50) not NULL,
x_coordinate numeric(10,5),
y_coordinate numeric(10,5));



