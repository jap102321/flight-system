import React from "react";
import styles from "./Footer.module.css"
import LinkedInIcon from '@mui/icons-material/LinkedIn';
import TravelExploreIcon from '@mui/icons-material/TravelExplore';
import { GitHub } from "@mui/icons-material";
import Link from "next/link";

const Footer = () => {
  return <footer className={styles.footer}>
    <div className={styles.container}>
      <div className={styles.socials}>
        <Link href={"https://www.linkedin.com/in/jpabgonzalez"}><LinkedInIcon /></Link>
        <Link href={"https://github.com/jap102321"}><GitHub/></Link>
        <Link href={"/"}><TravelExploreIcon/></Link>
      </div>
       Flight System App
       <div className={styles.mockInfo}>
          <ul>
            <li>Terms & Conditions.</li>
            <li>Against discrimination.</li>
            <li>Contact Us.</li>
          </ul>
       </div>
    </div>
  </footer>;
};

export default Footer;
