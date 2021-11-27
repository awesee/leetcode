Create table If Not Exists Follow (followee varchar(255), follower varchar(255));
Truncate table Follow;
insert into Follow (followee, follower) values ('Alice', 'Bob');
insert into Follow (followee, follower) values ('Bob', 'Cena');
insert into Follow (followee, follower) values ('Bob', 'Donald');
insert into Follow (followee, follower) values ('Donald', 'Edward');
