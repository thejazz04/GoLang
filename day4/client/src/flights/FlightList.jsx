import { useState } from "react";
import PageHeader from "../header/PageHeader";

function FlightList() {
    const [flights,setFlights]=useState([
        {id:"1010",number:"AI 400",airline_name:"AIR INDIA",source:"BNLR",destination:"MYS",capacity:180,price:5000.00},
        {id:"1001",number:"AI 404",airline_name:"AIR INDIA",source:"MYS",destination:"BNLR",capacity:180,price:5000.00}
    ]);
    return (
        <>
            <PageHeader PageNumber={1}/>
            <h3>List of Flights</h3>
            <div className="container">
                <table className="table table-primary table-striped">
                    <thead className="table-dark">
                        <tr>
                            <th scope="col">Flight Number</th>
                            <th scope="col">Airline Name</th>
                            <th scope="col">Source</th>
                            <th scope="col">Destination</th>
                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                    {flights.map((flight)=>{
                        return (
                            <tr>
                            <th scope="row">{flight.number}</th>
                            <td>{flight.airline_name}</td>
                            <td>{flight.source}</td>
                            <td>{flight.destination}</td>
                            <td>
                                <a href="/flights/edit/1023459870" className="btn btn-warning">Edit Price</a>
                                <button className="btn btn-danger">Delete</button>
                            </td>
                        </tr>
                        );
                    })
                    }
                        
                    </tbody>
                </table>
            </div>
        </>
    );
}

export default FlightList;
