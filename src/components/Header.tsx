"use client";

import UserContext from "@/providers/UserContext";
import Link from "next/link";
import React ,{ useContext } from "react";

export default function Header() {
    const user = useContext(UserContext);
    return (
        <>
            <header className="flex w-full h-12 justify-between items-center px-16 bg-[#FFFFFF42]">
                <Link href="/" className="font-['Inria_Sans'] text-[32px] font-bold leading-[38.37px] text-white">SocialNetwork</Link>
                <div className="flex">
                  <div className="border rounded-l-3xl bg-white bg-opacity-40 pl-20"></div>
                <input className="w-96 h-7 border rounded-r-3xl bg-white bg-opacity-40 focus:outline-none" placeholder="Search"></input>  
                </div>
                <div className="flex items-center">
                        <image className="relative z-10 -m-10 w-11 h-11 bg-white border rounded-full" />
                    <div className="flex w-36 h-7 items-center justify-end bg-white rounded-3xl pr-7 bg-opacity-40">{JSON.stringify(user)}</div>
                </div>
            </header>
        </>
    );
}


