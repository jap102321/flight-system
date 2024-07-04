"use client"
import React, { useEffect, useState } from "react";
import styles from "./DestinyCards.module.css"
import DestinyCard from "./DestinyCard/DestinyCard";

const DestinyCards = () => { 
  const [posts, setPosts] = useState([])

  useEffect(() => {
    const fetchPosts = async () => {
      try{
        const res = await fetch("http://localhost:8080/flights/all-flights")
        if(!res.ok){
          throw new Error("Error fetching flights.")
        }
        const data = await res.json()

        setPosts(data?.flights)
        return data;
      }catch(err){
        console.log("error!")
        throw new Error(String(err))
      }
    }

    fetchPosts()
    
  }, []);
  

  return <div className={styles.cardContainer} >
    {
      posts?.map(({_id, destiny, airline, date_of_departure, price})=> {

      return (
          <DestinyCard key={_id} price={price} airline={airline} date={date_of_departure} destinyName={destiny} imgSrc="https://images.pexels.com/photos/1842332/pexels-photo-1842332.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"/>
      )
      })
    }
  </div>;
};

export default DestinyCards;
