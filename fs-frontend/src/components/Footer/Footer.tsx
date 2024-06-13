import React from "react";
import styles from "./Footer.module.css"
import LinkedInIcon from '@mui/icons-material/LinkedIn';

const Footer = () => {
  return <footer className={styles.footer}>
    <div className={styles.container}>
      <div className={styles.socials}>
        <LinkedInIcon />
      </div>
       Flight System App
    </div>
  </footer>;
};

export default Footer;
