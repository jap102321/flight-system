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


const FlightSearchForm = () => {
  const [formResult, setFormResult] = useState({})

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

  useEffect(() => {
    console.log(formResult)
  }, [formResult]);

  return(
   <div className={styles.container}>
    <div className={styles.oneway}>
      <input type="checkbox" name="isOneWay" checked={isOneWay} onChange={()=> setOneWay(prevVal => !prevVal)}/>
      <label htmlFor="isOneWay">One way flight?</label>
    </div>
    <form ref={formRef} onSubmit={(evt)=> handleSubmitData(evt)} className={styles.searchBarForm}>
        <FlightSearchBar label="From" inputName="origin"/>
        <FlightSearchBar label="To" inputName="destiny"/>
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
