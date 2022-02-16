Create table If Not Exists Flights (departure_airport int, arrival_airport int, flights_count int);
Truncate table Flights;
insert into Flights (departure_airport, arrival_airport, flights_count) values ('1', '2', '4');
insert into Flights (departure_airport, arrival_airport, flights_count) values ('2', '1', '5');
insert into Flights (departure_airport, arrival_airport, flights_count) values ('2', '4', '5');
