import { Outlet } from "react-router";
import Aircraft from "./Aircraft";
import Airports from "./Airports";

const HomePage = () => {
  return (
    <>
        <div className="row">
        <Airports></Airports>
        <Aircraft></Aircraft>
        </div>
    </>
  );
};

export default HomePage;
