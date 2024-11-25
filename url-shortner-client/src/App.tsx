import { Outlet } from "react-router";
import "./App.css";

function App() {
  return (
    <>
      <div className="container pt-3">
        <div className="py-5 text-center">
          <h2>WELCOME</h2>
          <p className="lead">
            Have a nice flight!
          </p>
        </div>
        <Outlet></Outlet>
      </div>
    </>
  );
}

export default App;
