"use client"
import React, { useEffect, useState } from "react";
import styles from "./fSearch.module.css"

interface SearchBarProps{
  label : string;
  inputName: string;
  options : any[];
  setSelectedCity:  React.Dispatch<React.SetStateAction<string>>;
}



const FlightSearchBar = ({label, inputName, options, setSelectedCity} : SearchBarProps) => {

  return(
    <div className={styles.input}>
      <label htmlFor={inputName}>{label}</label>
      <select onChange={(evt) => setSelectedCity(evt.target.value)}>
      <option selected disabled>Select {inputName}</option>
          {
            options.map((option)=> {
              return (
                <option key={option.value} value={option.value}>{option.key}</option>
              )
            })
          }
      </select>
    </div>
  );
};

export default FlightSearchBar;
