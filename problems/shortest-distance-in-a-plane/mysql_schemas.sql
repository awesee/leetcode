Create Table If Not Exists Point2D (x int not null, y int not null);
Truncate table Point2D;
insert into Point2D (x, y) values ('-1', '-1');
insert into Point2D (x, y) values ('0', '0');
insert into Point2D (x, y) values ('-1', '-2');
