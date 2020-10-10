Create table If Not Exists Customer (customer_id int, customer_name varchar(20));
Create table If Not Exists Orders (order_id int, sale_date date, order_cost int, customer_id int, seller_id int)
;
Create table If Not Exists Seller (seller_id int, seller_name varchar(20))
;
Truncate table Customer;
insert into Customer (customer_id, customer_name) values ('101', 'Alice');
insert into Customer (customer_id, customer_name) values ('102', 'Bob');
insert into Customer (customer_id, customer_name) values ('103', 'Charlie');
Truncate table Orders;
insert into Orders (order_id, sale_date, order_cost, customer_id, seller_id) values ('1', '2020-03-01', '1500', '101', '1');
insert into Orders (order_id, sale_date, order_cost, customer_id, seller_id) values ('2', '2020-05-25', '2400', '102', '2');
insert into Orders (order_id, sale_date, order_cost, customer_id, seller_id) values ('3', '2019-05-25', '800', '101', '3');
insert into Orders (order_id, sale_date, order_cost, customer_id, seller_id) values ('4', '2020-09-13', '1000', '103', '2');
insert into Orders (order_id, sale_date, order_cost, customer_id, seller_id) values ('5', '2019-02-11', '700', '101', '2');
Truncate table Seller;
insert into Seller (seller_id, seller_name) values ('1', 'Daniel');
insert into Seller (seller_id, seller_name) values ('2', 'Elizabeth');
insert into Seller (seller_id, seller_name) values ('3', 'Frank');
