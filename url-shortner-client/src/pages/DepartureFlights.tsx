import axios from "axios";
import React, { Component, useEffect, useState } from "react";
import Flight from "../components/Flight";
import { redirect, useParams } from "react-router";
import { useNavigate } from "react-router-dom";

const api = axios.create({
  baseURL: "http://localhost:8080/departure",
});

const heading = "DEPARTURES";

interface Flights {
  flight_id: number;
  scheduled_departure: string;
  scheduled_arrival: string;
  departure_airport: string;
  arrival_airport: string;
  aircraft_code: string;
}
const DepartureFlightsPage = () => {
  const [flights, setFlights] = useState([]);
  const params = useParams();
  const history = useNavigate();
  useEffect(() => {
    console.log(params.airport_code);
    api.get(`/${params.airport_code}`).then((res) => {
      console.log(res.data);
      setFlights(Array.from(res.data.Flights));
    });
  });
  return (
    <>
      <h4 className="d-flex justify-content-between align-items-center mb-3">
        <span className="text-muted">{heading}</span>
      </h4>
      <table className="table table-bordered">
        <thead className="table-primary">
          <tr>
            <th scope="col-lg-2">TIME</th>
            <th scope="col-sm-8">TO</th>
            <th scope="col-sm-2">AIRCRAFT</th>
          </tr>
        </thead>
        <tbody>
          {flights.length == 0 && <p>no items</p>}
          {flights.map((item: Flights, index) => (
            <tr>
              <Flight
                key={index}
                time={item.scheduled_departure.substring(11, 16)}
                city={item.arrival_airport}
                aircraft_code={item.aircraft_code}
              ></Flight>
            </tr>
          ))}
        </tbody>
      </table>
      <a
        href="../airports"
        className="btn btn-link"
        onClick={() => history("/airports")}
      >
        BACK
      </a>
    </>
  );
};

export default DepartureFlightsPage;
