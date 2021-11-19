Create table If Not Exists NewYork (student_id int, score int);
Create table If Not Exists California (student_id int, score int);
Truncate table NewYork;
insert into NewYork (student_id, score) values ('1', '90');
insert into NewYork (student_id, score) values ('2', '87');
Truncate table California;
insert into California (student_id, score) values ('2', '89');
insert into California (student_id, score) values ('3', '88');
