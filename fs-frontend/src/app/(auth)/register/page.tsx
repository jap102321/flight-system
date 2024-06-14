"use client"
import RegisterForm from "@/components/AuthForms/RegisterForm/RegisterForm";
import React, { useEffect, useRef, useState } from "react";
import styles from "./register.module.css"
import { Nunito } from "next/font/google";
import Link from "next/link";


const nunito = Nunito({style:["normal"],subsets:["latin"]})


const SignUp = () => {
  const [formState, setFormState] = useState();
  const formRef = useRef<HTMLFormElement>(null)

  const sendData = async () => {
    const res = await fetch('http://localhost:8080/user/register', {
      method: "POST",
      headers:{
        "Content-Type" : "application/json"
      },
      body: JSON.stringify(formState)
    })
    if (!res.ok){
      throw new Error("Could not sign up")
    }
    
    console.log(res.json())
    return res.json()
  }

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
    <h1 style={{fontSize:"25px"}}>Register</h1>
    <form className={styles.form} ref={formRef} onSubmit={(evt)=> {
      getInputData(evt)
      sendData()
      }}>
        <RegisterForm/>
        <button type="submit" className={`${styles.submitButton} ${nunito.className}`}>Register</button>
    </form>
    <p>Already have an <Link href="/login" style={{color:"blue"}}>account?</Link> </p>
  </div>;
};

export default SignUp;
