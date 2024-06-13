import React from "react";
import styles from "./RegisterForm.module.css"

const RegisterForm = () => {
  return (
    <>
        <div className={styles.input}>
            <label htmlFor="email">Email</label>
            <input name="email" autoComplete="off" type="text"/>
        </div>
        <div className={styles.input}>
            <label htmlFor="password">Password</label>
            <input name="password" type="password"/>
        </div>
    </>
  );
};

export default RegisterForm;
