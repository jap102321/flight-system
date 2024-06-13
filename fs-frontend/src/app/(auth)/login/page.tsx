"use client"
import React, { useRef, useState } from "react";
import styles from "./login.module.css"
import { Nunito } from "next/font/google";
import LoginForm from "@/components/AuthForms/LoginForm/LoginForm";
import Link from "next/link";


const nunito = Nunito({style:["normal"],subsets:["latin"]})


const LogIn = () => {
  const [formState, setFormState] = useState();
  const formRef = useRef<HTMLFormElement>(null)

  const getInputData = (evt : React.FormEvent<HTMLFormElement>) => {
    evt.preventDefault();
    
    if(formRef.current){
      const formData = new FormData(formRef.current)
      const newFormResult : any = {}
      formData.forEach((value, key) => {
        newFormResult[key] = value
      })
      setFormState(newFormResult);
    }else{
        console.error("Could not find the ref")
    }
  }
  
  return <div className={styles.container}>
    <h1 style={{fontSize:"25px"}}>Login</h1>
    <form className={styles.form} ref={formRef} onSubmit={()=> getInputData}>
        <LoginForm/>
        <button type="submit" className={`${styles.submitButton} ${nunito.className}`}>Log In</button>
    </form>
    <p>Don't have an <Link href="/register" style={{color:"blue"}}>account?</Link> </p>
  </div>;
};

export default LogIn;
