import DestinyCards from "@/components/DestinyCards/DestinyCardContainer";
import styles from "./page.module.css"
import Navbar from "@/components/Navbar/Navbar";
import SearchFlight from "@/components/FlightSearchSection/SearchFlight";


export default function Home() {
  return (
    <>
      <div className={styles.planeBg}>
              <Navbar/>
              <SearchFlight/>
      </div>
      <div className={styles.mainHero}>   
        <h1>Explore new cheap destinies.</h1>
        <h3>Get the best offers that the market can find</h3>
        <DestinyCards/>
      </div>
    </>
  );
}
