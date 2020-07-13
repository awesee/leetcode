Create table If Not Exists Customers (customer_id int, name varchar(30), country varchar(30));
Create table If Not Exists Product (product_id int, description varchar(30), price int)
;
Create table If Not Exists Orders (order_id int, customer_id int, product_id int, order_date date, quantity int)
;
Truncate table Customers;
insert into Customers (customer_id, name, country) values ('1', 'Winston', 'USA');
insert into Customers (customer_id, name, country) values ('2', 'Jonathan', 'Peru');
insert into Customers (customer_id, name, country) values ('3', 'Moustafa', 'Egypt');
Truncate table Product;
insert into Product (product_id, description, price) values ('10', 'LC Phone', '300');
insert into Product (product_id, description, price) values ('20', 'LC T-Shirt', '10');
insert into Product (product_id, description, price) values ('30', 'LC Book', '45');
insert into Product (product_id, description, price) values ('40', 'LC Keychain', '2');
Truncate table Orders;
insert into Orders (order_id, customer_id, product_id, order_date, quantity) values ('1', '1', '10', '2020-06-10', '1');
insert into Orders (order_id, customer_id, product_id, order_date, quantity) values ('2', '1', '20', '2020-07-01', '1');
insert into Orders (order_id, customer_id, product_id, order_date, quantity) values ('3', '1', '30', '2020-07-08', '2');
insert into Orders (order_id, customer_id, product_id, order_date, quantity) values ('4', '2', '10', '2020-06-15', '2');
insert into Orders (order_id, customer_id, product_id, order_date, quantity) values ('5', '2', '40', '2020-07-01', '10');
insert into Orders (order_id, customer_id, product_id, order_date, quantity) values ('6', '3', '20', '2020-06-24', '2');
insert into Orders (order_id, customer_id, product_id, order_date, quantity) values ('7', '3', '30', '2020-06-25', '2');
insert into Orders (order_id, customer_id, product_id, order_date, quantity) values ('9', '3', '30', '2020-05-08', '3');
