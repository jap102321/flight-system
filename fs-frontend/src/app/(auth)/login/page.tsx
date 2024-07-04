"use client"
import React, { useRef, useState } from "react";
import styles from "./login.module.css"
import { Nunito } from "next/font/google";
import LoginForm from "@/components/AuthForms/LoginForm/LoginForm";
import Link from "next/link";
import { useStore } from "@/utils/store";
import { redirect, useRouter } from "next/navigation";
import { signIn } from "next-auth/react";



const nunito = Nunito({style:["normal"],subsets:["latin"]})


const LogIn = () => {
  const [formState, setFormState] = useState();
  const formRef = useRef<HTMLFormElement>(null)
  const [status, setStatus] = useState({
    loading: false,
    message: '',
    isError: false,
  });
  const setJWTToken = useStore(state => state.setJwtToken)
  const setUserLogged = useStore(state => state.setIsLoggedIn)
  const router = useRouter()

  const sendData = async () => {
    try{
      setStatus(prevVal => ({...prevVal, loading : true}))
      const res = await fetch("http://localhost:8080/user/login", {
        method: "POST",
        headers: {
          "Content-Type" : "application/json",
        },
        body: JSON.stringify(formState),
      })
      if(!res.ok){
        throw new Error("Could not log in, check your credentials.")
      }

      const data = await res.json()

      setJWTToken(data?.token)
      setUserLogged(true)
      setStatus({ loading: false, message: 'Login successful', isError: false });
      router.push("/")
      return data;
    }catch(err){
      setStatus({ loading: false, message: String(err), isError: true });
    }finally{
      setStatus(prev => ({ ...prev, loading: false }));
    }
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
    <h1 style={{fontSize:"25px"}}>Login</h1>
    <form className={styles.form} ref={formRef} onSubmit={(evt)=> {
      getInputData(evt)
      sendData()
    }}
      >
        <LoginForm/>
        <button type="submit" className={`${styles.submitButton} ${nunito.className}`}>Log In</button>
    </form>
    {status.isError && <p>{status.message}</p>}
    <p>Don't have an <Link href="/register" style={{color:"blue"}}>account?</Link> </p>
  </div>;
};

export default LogIn;
