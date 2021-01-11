Create table If Not Exists UserVisits(user_id int, visit_date date);
Truncate table UserVisits;
insert into UserVisits (user_id, visit_date) values ('1', '2020-11-28');
insert into UserVisits (user_id, visit_date) values ('1', '2020-10-20');
insert into UserVisits (user_id, visit_date) values ('1', '2020-12-3');
insert into UserVisits (user_id, visit_date) values ('2', '2020-10-5');
insert into UserVisits (user_id, visit_date) values ('2', '2020-12-9');
insert into UserVisits (user_id, visit_date) values ('3', '2020-11-11');
