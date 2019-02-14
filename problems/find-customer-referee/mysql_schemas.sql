CREATE TABLE IF NOT EXISTS customer (id INT,name VARCHAR(25),referee_id INT);;
Truncate table customer;
insert into customer (id, name, referee_id) values ('1', 'Will', 'None');
insert into customer (id, name, referee_id) values ('2', 'Jane', 'None');
insert into customer (id, name, referee_id) values ('3', 'Alex', '2');
insert into customer (id, name, referee_id) values ('4', 'Bill', 'None');
insert into customer (id, name, referee_id) values ('5', 'Zack', '1');
insert into customer (id, name, referee_id) values ('6', 'Mark', '2');
