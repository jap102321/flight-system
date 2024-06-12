import { create } from "zustand";


interface PassengerState {
    passengers: number
    setPassengerNumber : (passenger : number) => void
  }
  
  const usePassengerStore = create<PassengerState>()((set) => ({
    passengers: 0,
    setPassengerNumber: (passenger) => set(() => ({passengers: passenger}))
  }))