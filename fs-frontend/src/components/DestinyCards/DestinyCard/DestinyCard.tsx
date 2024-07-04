import Image from "next/image";
import React from "react";
import styles from "./DestinyCard.module.css"

interface DestinyCardProps {
    destinyName : string;
    imgSrc: string;
    airline : string;
    date: string;
    price : number
}

const DestinyCard = ({destinyName, imgSrc, airline, date, price} : DestinyCardProps) => {
  return (
    <div className={styles.card}>
        <h3>{destinyName}</h3>
        <Image src={`${imgSrc}`} height={220} width={300} alt={`destiny-${destinyName}`}/>
        <div className={styles.flightInfo}>
            <p className={styles.airline}>{airline}</p>
            <p className={styles.date}>{date}</p>
            <p className={styles.price}>{price ? price : "TBD"}</p>
        </div>
    </div>
  );
};

export default DestinyCard;
