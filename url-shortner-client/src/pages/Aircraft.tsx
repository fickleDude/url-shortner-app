import axios from "axios";
import React, { Component, useEffect, useState } from "react";
import Flight from "../components/Flight";
import { redirect, useParams } from "react-router";
import { useNavigate } from "react-router-dom";
import Aircraft from "../components/Aircraft";

const api = axios.create({
  baseURL: "http://localhost:8080/aircrafts",
});

const heading = "AIRCRAFTS";

interface Aircrafts {
  airport_code: string;
  model: string;
  range: number;
}

const Aircrafts = () => {
  const [aircrafts, setAircrafts] = useState([]);

  useEffect(() => {
    api.get("").then((res) => {
      console.log(res.data);
      setAircrafts(Array.from(res.data.Aircrafts));
    });
  });
  return (
    <>
      <div className="col-md-4 order-md-2 mb-4">
        <h4 className="d-flex justify-content-between align-items-center mb-3">
          <span className="text-muted">{heading}</span>
        </h4>
        {aircrafts.length == 0 && <p>no items</p>}
        <ul className="list-group mb-3">
          {aircrafts.map((item: Aircrafts, index) => (
            <Aircraft
              model={item.model}
              airport_code={item.airport_code}
              range={item.range}
            ></Aircraft>
          ))}
        </ul>
      </div>
    </>
  );
};

export default Aircrafts;
