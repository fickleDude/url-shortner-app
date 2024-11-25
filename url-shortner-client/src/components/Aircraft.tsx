interface Props{
    model: string,
    airport_code: string,
    range: number
}
const Aircraft = function(props:Props){
  return <>
            <li className="list-group-item d-flex justify-content-between lh-condensed">
              <div>
                <h6 className="my-0">{props.model}</h6>
                <small className="text-muted">range: {props.range}</small>
              </div>
              <span className="text-muted">{props.airport_code}</span>
            </li>
  </>


};

export default Aircraft;