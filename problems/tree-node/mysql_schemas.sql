Create table If Not Exists tree (id int, p_id int);
Truncate table tree;
insert into tree (id, p_id) values ('1', 'None');
insert into tree (id, p_id) values ('2', '1');
insert into tree (id, p_id) values ('3', '1');
insert into tree (id, p_id) values ('4', '2');
insert into tree (id, p_id) values ('5', '2');
