"use client"
import React, { useEffect, useState } from "react";
import styles from "./fSearch.module.css"

interface SearchBarProps{
  label : string;
  inputName: string;
  options : any[];
  setSelectedCity:  React.Dispatch<React.SetStateAction<{
    origin?: string;
    destiny?: string;
}>>;
}



const FlightSearchBar = ({label, inputName, options, setSelectedCity} : SearchBarProps) => {

  const handleChange = (evt: React.ChangeEvent<HTMLSelectElement>) => {
    const value = evt.target.value;
    setSelectedCity(prevVal => ({
      ...prevVal,
      [inputName]: value
    }));
  };

  return(
    <div className={styles.input}>
      <label htmlFor={inputName}>{label}</label>
      <select onChange={handleChange}>
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
