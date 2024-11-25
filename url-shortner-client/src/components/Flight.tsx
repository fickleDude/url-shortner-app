interface Props{
    time: string,
    city: string,
    aircraft_code: string
}
const Flight = function(props:Props){
  return <>
        <th scope="row">{props.time}</th>
      <td>{props.city}</td>
      <td>{props.aircraft_code}</td>
  </>


};

export default Flight;