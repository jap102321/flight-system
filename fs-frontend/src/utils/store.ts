import { create } from "zustand";


interface FlightAppGlobalStates {
    jwtToken : string
    setJwtToken : (token : string) => void
    isLoggedIn : boolean,
    setIsLoggedIn : (isLogged : boolean) => void;
    isAdmin : boolean, 
    setIsAdmin : (isAdmin : boolean) => void;
  }
  
  export const useStore = create<FlightAppGlobalStates>()((set) => ({
    jwtToken : "",
    setJwtToken : (token) => set(()=> ({jwtToken: token})),
    isLoggedIn : false,
    setIsLoggedIn : () => set(()=> ({isLoggedIn : true})),
    isAdmin : false,
    setIsAdmin : () => set(()=> ({isAdmin : true})),
  }))