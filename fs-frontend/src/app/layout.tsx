import type { Metadata } from "next";
import { Inter, Nunito } from "next/font/google";
import "./globals.css";
import Footer from "@/components/Footer/Footer";


const nunito = Nunito({style:["normal"],subsets:["latin"]})

export const metadata: Metadata = {
  title: "Flight system",
  description: "Flight system demo app created for portfolio.",
};

export default function RootLayout({
  children,
}: Readonly<{
  children: React.ReactNode;
}>) {
  
  return (
    <html lang="en">
      <body className={nunito.className}>
        <div className="container">
          {children}
          <Footer/>
        </div>
      </body>
    </html>
  );
}
