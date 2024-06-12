import React from "react";
import styles from "./fSearch.module.css"

interface SearchBarProps{
  label : string;
  inputName: string;
}

const FlightSearchBar = ({label, inputName} : SearchBarProps) => {
  return(
    <div className={styles.input}>
      <label htmlFor={inputName}>{label}</label>
      <input name={inputName} type="text" autoComplete={"off"}/>
    </div>
  );
};

export default FlightSearchBar;
