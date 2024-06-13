"use client"
import React from "react";
import styles from "./searchbar.module.css"
import FlightSearchForm from "../FlightSearchForm/FlightSearchForm";
import { usePathname } from "next/navigation";

const SearchFlight = () => {
  const path = usePathname()
  return (
     <>
      {path === "/" && (
        <div className={styles.searchFlight}>
            <h1 className={styles.title}>Discover Your Flight Experience</h1>
            <div className={styles.flightSearchContainer}>
              <FlightSearchForm/>
            </div>
        </div>
      )}
     </>
    )
};

export default SearchFlight;

