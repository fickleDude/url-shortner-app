package controllers

import (
	"encoding/json"
	"net/http"

	helpers "example.com/url-shortner/helper"
	services "example.com/url-shortner/service"
	"github.com/go-chi/chi/v5"
)

var Airports services.Airports

var Flights services.Flights

var Aircrafts services.Aircrafts

func GetAllAirports(w http.ResponseWriter, r *http.Request) {
	all, err := Airports.GetAllAirports()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"Airports": all})
}

func GetDepartureFlights(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "airport_code")
	all, err := Flights.GetDepartureFlights(code)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"Flights": all})
}

func GetArrivalFlights(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "airport_code")
	all, err := Flights.GetArrivalFlights(code)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"Flights": all})
}

func GetAllAircrafts(w http.ResponseWriter, r *http.Request) {
	all, err := Aircrafts.GetAllAircrafts()
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"Aircrafts": all})
}

// GET//Airportss/Airports/{id}
func GetAirportByAirportCode(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "airport_code")
	Airports, err := Airports.GetAirportsByAirportCode(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, Airports)
}

// POST/Airportss/Airports
func CreateAirports(w http.ResponseWriter, r *http.Request) {
	var AirportsData services.Airports
	err := json.NewDecoder(r.Body).Decode(&AirportsData)

	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}

	AirportsCreated, err := Airports.CreateAirport(AirportsData)
	// CHECK
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
		return
	}
	helpers.WriteJSON(w, http.StatusOK, AirportsCreated)
}

// PUT/Airportss/Airports/{code}
func UpdateAirports(w http.ResponseWriter, r *http.Request) {
	var AirportsData services.Airports
	code := chi.URLParam(r, "airport_code")
	err := json.NewDecoder(r.Body).Decode(&AirportsData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	airportUpdated, err := Airports.UpdateAirports(code, AirportsData)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, airportUpdated)
}

// DELETE/Airportss/Airports/{id}
func DeleteAirports(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "airport_code")
	err := Airports.DeleteAirports(id)
	if err != nil {
		helpers.MessageLogs.ErrorLog.Println(err)
	}
	helpers.WriteJSON(w, http.StatusOK, helpers.Envelope{"message": "successfull deletion"})
}
