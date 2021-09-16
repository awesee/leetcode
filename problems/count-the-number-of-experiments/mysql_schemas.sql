Create table If Not Exists Experiments (experiment_id int, platform ENUM('Android', 'IOS', 'Web'), experiment_name ENUM('Reading', 'Sports', 'Programming'));
Truncate table Experiments;
insert into Experiments (experiment_id, platform, experiment_name) values ('4', 'IOS', 'Programming');
insert into Experiments (experiment_id, platform, experiment_name) values ('13', 'IOS', 'Sports');
insert into Experiments (experiment_id, platform, experiment_name) values ('14', 'Android', 'Reading');
insert into Experiments (experiment_id, platform, experiment_name) values ('8', 'Web', 'Reading');
insert into Experiments (experiment_id, platform, experiment_name) values ('12', 'Web', 'Reading');
insert into Experiments (experiment_id, platform, experiment_name) values ('18', 'Web', 'Programming');
