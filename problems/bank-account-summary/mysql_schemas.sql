Create table If Not Exists Users (user_id int, user_name varchar(20), credit int)
;
Create table If Not Exists Transaction (trans_id int, paid_by int, paid_to int, amount int, transacted_on date);
Truncate table Users;
insert into Users (user_id, user_name, credit) values ('1', 'Moustafa', '100');
insert into Users (user_id, user_name, credit) values ('2', 'Jonathan', '200');
insert into Users (user_id, user_name, credit) values ('3', 'Winston', '10000');
insert into Users (user_id, user_name, credit) values ('4', 'Luis', '800');
Truncate table Transaction;
insert into Transaction (trans_id, paid_by, paid_to, amount, transacted_on) values ('1', '1', '3', '400', '2020-08-01');
insert into Transaction (trans_id, paid_by, paid_to, amount, transacted_on) values ('2', '3', '2', '500', '2020-08-02');
insert into Transaction (trans_id, paid_by, paid_to, amount, transacted_on) values ('3', '2', '1', '200', '2020-08-03');
