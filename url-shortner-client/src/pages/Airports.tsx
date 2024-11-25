import axios from "axios";
import React, { Component, useEffect, useState } from "react";
import Airport from "../components/Airport";
import Aircraft from "./Aircraft";

const api = axios.create({ baseURL: "http://localhost:8080/airports" });

const heading = "AIRPORTS";

type Airport = {
  airport_code: string;
  airport_name: string;
  city: string;
};

class Airports extends Component {
  state = { airports: [], selectedIndex: -1 };
  constructor(props: any) {
    super(props);
  }
  componentDidMount() {
    api.get("").then((res) => {
      console.log(res.data);
      this.setState({ airports: Array.from(res.data.Airports) });
      console.log(this.state.airports);
    });
  }
  render(): React.ReactNode {
    return (
      <>
        <div className="col-md-8 order-md-1">
        <h4 className="d-flex justify-content-between align-items-center mb-3">
          <span className="text-muted">{heading}</span>
        </h4>
          {this.state.airports.length == 0 && <p>no items</p>}
          <ul className="list-group">
            {this.state.airports.map((item: Airport, index) => (
              <div className="container">
                <li
                  className={
                    this.state.selectedIndex == index
                      ? "list-group-item active"
                      : "list-group-item"
                  }
                  onClick={() => {
                    this.setState({ selectedIndex: index });
                  }}
                  key={index}
                >
                  <Airport
                    airport_code={item.airport_code}
                    airport_name={item.airport_name}
                    city={item.city}
                  ></Airport>
                </li>
              </div>
            ))}
          </ul>
        </div>
      </>
    );
  }
}

export default Airports;
