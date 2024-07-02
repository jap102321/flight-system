import NextAuth from "next-auth";
import Credentials from "next-auth/providers/credentials"



const LogIn = async (formState : any) => {
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

      return data;
} 


export const { handlers, signIn, signOut, auth  } = NextAuth({
    providers: [
        Credentials({
            credentials:{
                email: {},
                password: {},
            },
            async authorize(credentials) {
                try{
                  let user = await LogIn(credentials);
                  return user;
                }catch(err){
                    console.error(err)
                }
            },
        })
    ]
})