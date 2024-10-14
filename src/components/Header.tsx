"use client";

import UserContext from "@/providers/UserContext";
import Link from "next/link";
import React, { useContext } from "react";
import { Search } from "@/components/icons/Search";
import { ProfileCircle } from "./icons/ProfileCircle";

export default function Header() {
    const user = useContext(UserContext);
    return (
        <>
            <header className="flex flex-col w-full h-12 bg-[#FFFFFF42] justify-center md:flex-row md:justify-between md:items-center" role="banner">
                <Link href="/" className="font-['Inria_Sans'] text-[32px] font-bold leading-[38.37px] text-white ml-5">
                    <h1>SocialNetwork</h1>
                </Link>

                <div className="hidden md:flex">
                    <div className="flex justify-center items-center w-12 h-8 border border-white rounded-l-3xl bg-white bg-opacity-40 border-r-0" aria-hidden="true">
                        <Search />
                    </div>
                    <label htmlFor="search-input" className="sr-only">Search</label>
                    <input id="search-input" type="search" className="w-80 h-8 border rounded-r-3xl border-l-0  border-white bg-white bg-opacity-40 focus:outline-none" placeholder="Search" aria-label="Search" />
                </div>

                <div className="items-center md:flex md:relative">
                    <div className="flex relative z-10 -m-10 w-11 h-11 bg-white border rounded-full">
                        <ProfileCircle />
                    </div>
                    <div className="w-36 relative justify-end bg-white rounded-3xl bg-opacity-40 mr-5">
                        <div className="flex flex-col items-center pl-5">
                            <div className="text-xs font-bold">{JSON.stringify(user)}</div>
                            <div className="text-xs">@nickname</div>
                        </div>
                    </div>
                </div>
            </header>
        </>
    );
}


