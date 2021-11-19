Create table If Not Exists Transactions (account_id int, day date, type ENUM('Deposit', 'Withdraw'), amount int);
Truncate table Transactions;
insert into Transactions (account_id, day, type, amount) values ('1', '2021-11-07', 'Deposit', '2000');
insert into Transactions (account_id, day, type, amount) values ('1', '2021-11-09', 'Withdraw', '1000');
insert into Transactions (account_id, day, type, amount) values ('1', '2021-11-11', 'Deposit', '3000');
insert into Transactions (account_id, day, type, amount) values ('2', '2021-12-07', 'Deposit', '7000');
insert into Transactions (account_id, day, type, amount) values ('2', '2021-12-12', 'Withdraw', '7000');
