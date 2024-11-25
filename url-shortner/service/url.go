package services

import (
	"time"
)

type Flights struct {
	Flight_id           int       `json:"flight_id"`
	Scheduled_departure time.Time `json:"scheduled_departure"`
	Scheduled_arrival   time.Time `json:"scheduled_arrival"`
	Departure_airport   string    `json:"departure_airport"`
	Arrival_airport     string    `json:"arrival_airport"`
	Aircraft_code       string    `json:"aircraft_code"`
}

type Airports struct {
	Airport_code string `json:"airport_code"`
	Airport_name string `json:"airport_name"`
	City         string `json:"city"`
}

type Aircrafts struct {
	Airport_code string `json:"airport_code"`
	Model        string `json:"model"`
	Range        int    `json:"range"`
}

func (c *Airports) GetAllAirports() ([]*Airports, error) {

	query := `select airport_code,airport_name,city from airports_data`
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	var airports []*Airports
	for rows.Next() {
		var airport Airports
		err := rows.Scan(
			&airport.Airport_code,
			&airport.Airport_name,
			&airport.City,
		)

		if err != nil {
			return nil, err
		}
		airports = append(airports, &airport)
	}

	return airports, nil
}

func (c *Airports) GetAirportsByAirportCode(code string) (*Airports, error) {
	query := `select airport_code,airport_name,city from airports_data where airport_code = $1`
	var airport Airports

	row := db.QueryRow(query, code)
	err := row.Scan(
		&airport.Airport_code,
		&airport.Airport_name,
		&airport.City,
	)

	if err != nil {
		return nil, err
	}
	return &airport, nil
}

func (c *Airports) CreateAirport(airport Airports) (*Airports, error) {

	query := `insert into airports_data (airport_code, airport_name, city) values ($1, $2, $3) returning *`

	_, err := db.Exec(
		query,
		&airport.Airport_code,
		&airport.Airport_name,
		&airport.City,
	)

	if err != nil {
		return nil, err
	}

	return &airport, nil
}

func (c *Airports) UpdateAirports(code string, body Airports) (*Airports, error) {
	query := `update airports_data set airport_name = $1,city = $2 where airport_code = $3 returning *`

	_, err := db.Exec(
		query,
		body.Airport_name,
		body.City,
		code,
	)
	if err != nil {
		return nil, err
	}
	return &body, nil
}

func (c *Airports) DeleteAirports(code string) error {
	query := `DELETE FROM airports_data WHERE airport_code = $1`
	_, err := db.Exec(query, code)
	if err != nil {
		return err
	}
	return nil
}

func (c *Flights) GetDepartureFlights(airport_code string) ([]*Flights, error) {

	query := `SELECT f.flight_id, f.scheduled_departure, f.scheduled_arrival,f.departure_airport, 
ad.airport_name, f.aircraft_code FROM flights f join airports_data ad on f.arrival_airport = ad.airport_code
where  f.departure_airport = $1;`
	rows, err := db.Query(query, airport_code)

	if err != nil {
		return nil, err
	}

	var flights []*Flights
	for rows.Next() {
		var flight Flights
		err := rows.Scan(
			&flight.Flight_id,
			&flight.Scheduled_departure,
			&flight.Scheduled_arrival,
			&flight.Departure_airport,
			&flight.Arrival_airport,
			&flight.Aircraft_code,
		)

		if err != nil {
			return nil, err
		}
		flights = append(flights, &flight)
	}

	return flights, nil
}
func (c *Flights) GetArrivalFlights(airport_code string) ([]*Flights, error) {

	query := `SELECT f.flight_id, f.scheduled_departure, f.scheduled_arrival,ad.airport_name, 
f.arrival_airport, f.aircraft_code FROM flights f join airports_data ad on f.departure_airport = ad.airport_code
where  f.arrival_airport = $1;`
	rows, err := db.Query(query, airport_code)

	if err != nil {
		return nil, err
	}

	var flights []*Flights
	for rows.Next() {
		var flight Flights
		err := rows.Scan(
			&flight.Flight_id,
			&flight.Scheduled_departure,
			&flight.Scheduled_arrival,
			&flight.Departure_airport,
			&flight.Arrival_airport,
			&flight.Aircraft_code,
		)

		if err != nil {
			return nil, err
		}
		flights = append(flights, &flight)
	}

	return flights, nil
}

func (c *Aircrafts) GetAllAircrafts() ([]*Aircrafts, error) {

	query := `select aircraft_code,model,range from aircrafts_data`
	rows, err := db.Query(query)

	if err != nil {
		return nil, err
	}

	var airports []*Aircrafts
	for rows.Next() {
		var airport Aircrafts
		err := rows.Scan(
			&airport.Airport_code,
			&airport.Model,
			&airport.Range,
		)

		if err != nil {
			return nil, err
		}
		airports = append(airports, &airport)
	}

	return airports, nil
}
