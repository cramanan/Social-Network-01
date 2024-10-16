"use client";

import Link from "next/link";
import React from "react";
import { ProfileCircle } from "./icons/ProfileCircle";
import SearchBar from "./SearchBar";
import { useAuth } from "@/providers/AuthContext";

export default function Header() {
    const { user } = useAuth();
    return (
        <>
            <header
                className="flex items-center h-14 bg-[#FFFFFF42]"
                role="banner"
            >
                <div
                    id="header-container"
                    className="flex flex-col w-full justify-center md:flex-row md:justify-between md:items-center px-10"
                >
                    <Link
                        href="/"
                        className="font-['Inria_Sans'] text-[32px] font-bold leading-[38.37px] text-white"
                    >
                        <h1>SocialNetwork</h1>
                    </Link>

                    <SearchBar id="search-bar-header" />

                    <div className="items-center md:flex md:relative">
                        <div className="flex relative z-10 -m-10 w-11 h-11 bg-white border rounded-full">
                            <ProfileCircle />
                        </div>
                        <div className="w-36 relative justify-end bg-white rounded-3xl bg-opacity-40">
                            <div className="flex flex-col items-center pl-5">
                                <div className="text-xs font-bold">
                                    {JSON.stringify(user)}
                                </div>
                                <div className="text-xs">@nickname</div>
                            </div>
                        </div>
                    </div>
                </div>
            </header>
        </>
    );
}
