Create table If Not Exists SurveyLog (id int, action varchar(255), question_id int, answer_id int, q_num int, timestamp int);
Truncate table SurveyLog;
insert into SurveyLog (id, action, question_id, answer_id, q_num, timestamp) values ('5', 'show', '285', 'None', '1', '123');
insert into SurveyLog (id, action, question_id, answer_id, q_num, timestamp) values ('5', 'answer', '285', '124124', '1', '124');
insert into SurveyLog (id, action, question_id, answer_id, q_num, timestamp) values ('5', 'show', '369', 'None', '2', '125');
insert into SurveyLog (id, action, question_id, answer_id, q_num, timestamp) values ('5', 'skip', '369', 'None', '2', '126');
