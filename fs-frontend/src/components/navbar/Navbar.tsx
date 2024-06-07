import Image from 'next/image'
import React from 'react'
import styles from "./navbar.module.css"

const Navbar = () => {
  return (
    <nav className={styles.navbar}>
        <Image src={"/logo-plane.png"} width={50} height={50} alt='logo'/>
        <ul className={styles.list}>
            <li><button className={`${styles.button}`}>Log In</button></li>
            <li ><button className={`${styles.button} ${styles.signup}`}>Register</button></li>
        </ul>
    </nav>
  )
}

export default Navbar