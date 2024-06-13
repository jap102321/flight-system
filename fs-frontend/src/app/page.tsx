import DestinyCards from "@/components/DestinyCards/DestinyCardContainer";
import styles from "./page.module.css"


export default function Home() {
  return (
    <div className={styles.mainHero}>   
      <h1>Explore new cheap destinies.</h1>
      <h3>Get the best offers that the market can find</h3>
      <DestinyCards/>
    </div>
  );
}
