create table if not exists Transactions (id int, country varchar(4), state enum('approved', 'declined'), amount int, trans_date date)
;
create table if not exists Chargebacks (trans_id int, trans_date date)
;
Truncate table Transactions;
insert into Transactions (id, country, state, amount, trans_date) values ('101', 'US', 'approved', '1000', '2019-05-18');
insert into Transactions (id, country, state, amount, trans_date) values ('102', 'US', 'declined', '2000', '2019-05-19');
insert into Transactions (id, country, state, amount, trans_date) values ('103', 'US', 'approved', '3000', '2019-06-10');
insert into Transactions (id, country, state, amount, trans_date) values ('104', 'US', 'declined', '4000', '2019-06-13');
insert into Transactions (id, country, state, amount, trans_date) values ('105', 'US', 'approved', '5000', '2019-06-15');
Truncate table Chargebacks;
insert into Chargebacks (trans_id, trans_date) values ('102', '2019-05-29');
insert into Chargebacks (trans_id, trans_date) values ('101', '2019-06-30');
insert into Chargebacks (trans_id, trans_date) values ('105', '2019-09-18');
