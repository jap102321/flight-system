"use client"
import React, { useEffect, useRef, useState } from "react";
import FlightSearchBar from "../FlightSearchBar/FlightSearchBar";
import styles from "./fSearchForm.module.css"
import Datepicker from "../FlightSearchBar/Datepicker/Datepicker";
import { Nunito } from "next/font/google";

interface handleSubmitType {
  origin:  string;
  destiny: string;
  departureDate: any;
  returnDate?: any;
}


const nunito = Nunito({style:["normal"],weight:["700"],subsets:["latin"]})

const PlacesArray = [{
  key:"Medellín",
  value: "medellin"
},
{
  key:"Pereira",
  value: "pereira"
},
{
  key:"Barranquilla",
  value: "barranquilla"
},
{
  key:"Bogotá",
  value: "bogota"
}
]

const FlightSearchForm = () => {
  const [formResult, setFormResult] = useState({})
  const [selectedCity, setSelectedCity] = useState({
    origin: "",
    destiny:"",
  });

  const filterOriginOptions = PlacesArray;
  const filterDestinyOptions = PlacesArray.filter(
    ({ value }) => value !== selectedCity.origin
  );

  const formRef = useRef<HTMLFormElement>(null)
  const [isOneWay, setOneWay] = useState<boolean>(false)

  const handleSubmitData = (event: React.FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (formRef.current) {
      const formData = new FormData(formRef.current);
      const newFormResult : any = {}
      formData.forEach((value, key) => {
        newFormResult[key] = value
      });
      setFormResult(newFormResult)
    } else {
      console.log('Form reference is null');
    }
  }


  return(
   <div className={styles.container}>
    <div className={styles.oneway}>
      <input type="checkbox" name="isOneWay" checked={isOneWay} onChange={()=> setOneWay(prevVal => !prevVal)}/>
      <label htmlFor="isOneWay">One way flight?</label>
    </div>
    <form ref={formRef} onSubmit={(evt)=> handleSubmitData(evt)} className={styles.searchBarForm}>
        <FlightSearchBar label="From" inputName="origin" options={filterOriginOptions} setSelectedCity={setSelectedCity}/>
        <FlightSearchBar label="To" inputName="destiny" options={filterDestinyOptions} setSelectedCity={setSelectedCity}/>
        <Datepicker label="Departure Date"/>
        {
          !isOneWay && <Datepicker label="Return Date"/>
        }
      <button className={`${styles.submitButton} ${nunito.className}`} type="submit">Search</button>
    </form>
  </div>
  );
};

export default FlightSearchForm;
