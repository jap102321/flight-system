import Image from 'next/image'
import React from 'react'
import styles from "./navbar.module.css"
import Link from 'next/link'
import { Nunito } from 'next/font/google'

const nunito = Nunito({style:["normal"],weight:["700"],subsets:["latin"]})


const Navbar = () => {
  return (
    <nav className={styles.navbar}>
        <Image src={"/logo-plane.png"} width={50} height={50} alt='logo'/>
        <ul className={styles.list}>
            <li><button className={`${styles.button}`}><Link href="/login" style={nunito.style}>Log In</Link></button></li>
            <li ><button className={`${styles.button} ${styles.signup}`} style={nunito.style}><Link href="/register">Sign Up</Link></button></li>
        </ul>
    </nav>
  )
}

export default Navbar