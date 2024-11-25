import { redirect } from "react-router";

interface Props{
        airport_code: string,
        airport_name: string,
        city:         string 
}
const Airport = function(props:Props){
    return <div className="card">
    <div className="card-body">
      <h5 className="card-title">{props.airport_name}</h5>
      <h6 className="card-subtitle mb-2 text-muted">{props.airport_code}</h6>
      <p className="card-text">{props.city}</p>
      <a href={`departure/${props.airport_code}`} className="card-link" onClick={()=>redirect(`departure/${props.airport_code}`)}>DEPARTURES</a>
      <a href={`arrive/${props.airport_code}`} className="card-link" onClick={()=>redirect(`arrive/${props.airport_code}`)}>ARRIVALS</a>
    </div>
  </div>;
};

export default Airport;