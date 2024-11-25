import { createBrowserRouter } from "react-router-dom";
import App from "../App";
import DepartureFlightsPage from "../pages/DepartureFlights";
import ArriveFlightsPage from "../pages/ArrivalFlights";
import HomePage from "../pages/Home";

export const router = createBrowserRouter(
[    {
        path: "/",
        element: <App/>,
        children: [
            {path: "airports", element: <HomePage></HomePage>},
            {path: "departure/:airport_code", element: <DepartureFlightsPage></DepartureFlightsPage>},
            {path: "arrive/:airport_code", element: <ArriveFlightsPage></ArriveFlightsPage>}
        ]
    }]
)