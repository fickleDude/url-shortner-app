CREATE TABLE airports_data (
    airport_code varchar NOT NULL,
    airport_name varchar NOT NULL,
    city varchar NOT NULL,
	UNIQUE(airport_code)
);

CREATE TABLE flights (
    flight_id integer NOT NULL,
    scheduled_departure timestamp with time zone NOT NULL,
    scheduled_arrival timestamp with time zone NOT NULL,
    departure_airport varchar NOT NULL,
    arrival_airport varchar NOT NULL,
    aircraft_code varchar NOT NULL,
	UNIQUE(departure_airport, arrival_airport)
);

ALTER TABLE ONLY flights
    ADD CONSTRAINT flights_arrival_airport_fkey FOREIGN KEY (arrival_airport) REFERENCES airports_data(airport_code);
ALTER TABLE ONLY flights
    ADD CONSTRAINT flights_departure_airport_fkey FOREIGN KEY (departure_airport) REFERENCES airports_data(airport_code);

insert into airports_data (airport_code, airport_name, city) values
('YKS', 'Yakutsk Airport',	'Yakutsk'),
('MJZ','Mirny Airport','Mirnyj'),	
('KHV','Khabarovsk-Novy Airport','Khabarovsk'),	
('PKC','Yelizovo Airport', 'Petropavlovsk');
insert into flights (flight_id, scheduled_departure, scheduled_arrival, departure_airport, arrival_airport, aircraft_code) values
('1','2017-08-08 20:55','2017-08-09 03:20','PKC','YKS','PG0243'),
('2','2017-08-08 19:45','2017-08-09 07:00','YKS','PKC','PG0243'),
('3','2017-08-08 20:05','2017-08-09 07:10','KHV','YKS','PG0243'),
('4','2017-08-08 07:55','2017-08-09 10:08','YKS','KHV','PG0243'),
('5','2017-08-08 02:05','2017-08-09 23:20','MJZ','YKS','PG0243'),
('6','2017-08-08 20:15','2017-08-09 13:20','YKS','MJZ','PG0243');


 CREATE TABLE aircrafts_data (
    aircraft_code varchar NOT NULL UNIQUE,
    model varchar NOT NULL,
    range integer NOT NULL
);
CREATE TABLE seats (
    aircraft_code varchar NOT NULL,
    seat_no varchar NOT NULL,
    fare_conditions varchar NOT NULL
);

ALTER TABLE ONLY seats
ADD CONSTRAINT seats_aircraft_code_fkey FOREIGN KEY (aircraft_code) REFERENCES aircrafts_data(aircraft_code) ON DELETE CASCADE;

insert into aircrafts_data (aircraft_code, model, range) values
('773','Boeing 777-300','11100'),
('763','Sukhoi Superjet-100','7900'),
('SU9','Airbus A320-200','3000'),
('320','Airbus A321-200','5700');