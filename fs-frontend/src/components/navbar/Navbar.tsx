"use client"
import Image from 'next/image'
import React from 'react'
import styles from "./navbar.module.css"
import Link from 'next/link'
import { Nunito } from 'next/font/google'
import { useStore } from '@/utils/store'

const nunito = Nunito({style:["normal"],weight:["700"],subsets:["latin"]})


const Navbar = () => {
  const isLogged = useStore(state => state.isLoggedIn)
  const isAdmin = useStore(state => state.isAdmin)
  return (
    <nav className={styles.navbar}>
        <Image src={"/logo-plane.png"} width={50} height={50} alt='logo'/>
        {!isLogged ? (<ul className={styles.list}>
            <li><button className={`${styles.button}`}><Link href="/login" style={nunito.style}>Log In</Link></button></li>
            <li ><button className={`${styles.button} ${styles.signup}`} style={nunito.style}><Link href="/register">Sign Up</Link></button></li>
        </ul>) : <p>Welcome! see your <span style={{color:"blue"}}> <Link href="/">reservations</Link></span></p>}
          {
            isAdmin && <button><Link href={"/admin"}>Admin Panel</Link></button>
          }
    </nav>
  )
}

export default Navbar