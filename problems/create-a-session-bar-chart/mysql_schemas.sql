Create table If Not Exists Sessions (session_id int, duration int);
Truncate table Sessions;
insert into Sessions (session_id, duration) values ('1', '30');
insert into Sessions (session_id, duration) values ('2', '199');
insert into Sessions (session_id, duration) values ('3', '299');
insert into Sessions (session_id, duration) values ('4', '580');
insert into Sessions (session_id, duration) values ('5', '1000');
