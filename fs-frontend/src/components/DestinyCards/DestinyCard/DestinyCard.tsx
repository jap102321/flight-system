import Image from "next/image";
import React from "react";
import styles from "./DestinyCard.module.css"

interface DestinyCardProps {
    destinyName : string;
    imgSrc: string;

}

const DestinyCard = ({destinyName, imgSrc} : DestinyCardProps) => {
  return (
    <div className={styles.card}>
        <h3>{destinyName}</h3>
        <Image src={`${imgSrc}`} height={220} width={300} alt={`destiny-${destinyName}`}/>
        <div className={styles.flightInfo}>
            <p className={styles.airline}>Airline</p>
            <p className={styles.date}>Date</p>
            <p className={styles.price}>Price</p>
        </div>
    </div>
  );
};

export default DestinyCard;
