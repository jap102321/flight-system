import React from "react";
import styles from "./DestinyCards.module.css"
import DestinyCard from "./DestinyCard/DestinyCard";

const DestinyCards = () => { 
  return <div className={styles.cardContainer} >
    <DestinyCard destinyName="Singapore" imgSrc="https://images.pexels.com/photos/1842332/pexels-photo-1842332.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"/>
    <DestinyCard destinyName="Singapore" imgSrc="https://images.pexels.com/photos/1842332/pexels-photo-1842332.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"/>
    <DestinyCard destinyName="Singapore" imgSrc="https://images.pexels.com/photos/1842332/pexels-photo-1842332.jpeg?auto=compress&cs=tinysrgb&w=1260&h=750&dpr=1"/>
  </div>;
};

export default DestinyCards;
