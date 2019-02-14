Create table If Not Exists orders (order_number int, customer_number int, order_date date, required_date date, shipped_date date, status char(15), comment char(200), key(order_number));
Truncate table orders;
insert into orders (order_number, customer_number) values ('1', '1');
insert into orders (order_number, customer_number) values ('2', '2');
insert into orders (order_number, customer_number) values ('3', '3');
insert into orders (order_number, customer_number) values ('4', '3');
