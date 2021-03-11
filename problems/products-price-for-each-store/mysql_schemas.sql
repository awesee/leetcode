Create table If Not Exists Products (product_id int, store ENUM('store1', 'store2', 'store3'), price int);
Truncate table Products;
insert into Products (product_id, store, price) values ('0', 'store1', '95');
insert into Products (product_id, store, price) values ('0', 'store3', '105');
insert into Products (product_id, store, price) values ('0', 'store2', '100');
insert into Products (product_id, store, price) values ('1', 'store1', '70');
insert into Products (product_id, store, price) values ('1', 'store3', '80');
