Create table If Not Exists Queue (person_id int, person_name varchar(30), weight int, turn int);
Truncate table Queue;
insert into Queue (person_id, person_name, weight, turn) values ('5', 'George Washington', '250', '1');
insert into Queue (person_id, person_name, weight, turn) values ('4', 'Thomas Jefferson', '175', '5');
insert into Queue (person_id, person_name, weight, turn) values ('3', 'John Adams', '350', '2');
insert into Queue (person_id, person_name, weight, turn) values ('6', 'Thomas Jefferson', '400', '3');
insert into Queue (person_id, person_name, weight, turn) values ('1', 'James Elephant', '500', '6');
insert into Queue (person_id, person_name, weight, turn) values ('2', 'Will Johnliams', '200', '4');
