'use client'
import React from "react";
import type { DatePickerProps } from 'antd';
import { DatePicker } from 'antd';
import styles from "./Datepicker.module.css"
import dayjs from "dayjs";

interface DatepickerProps{
    label : string;
  }


const Datepicker = ({label}: DatepickerProps) => {
    const now = dayjs()

  return (
    <div className={styles.datePickerContainer}>
        <label htmlFor={label}>{label}</label>
        <DatePicker name={label} style={{height:'28px'}} minDate={now}/>
    </div>
  )
};

export default Datepicker;
