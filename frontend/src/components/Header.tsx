"use client";

import Link from "next/link";
import { ProfileCircle } from "./icons/ProfileCircle";
import SearchBar from "./SearchBar";
import MobileNav from "./MobileNav";
import { useAuth } from "@/hooks/useAuth";

export default function Header() {
    const { user } = useAuth();
    return (
        <>
            <header
                className="flex flex-col w-full h-16 bg-[#FFFFFF42] justify-center md:flex-row md:justify-between md:items-center"
                role="banner"
            >
                <div
                    id="header-container"
                    className="flex flex-row w-full items-center justify-between px-5 xl:flex-row xl:px-10"
                >
                    <Link
                        href="/"
                        className="font-['Inria_Sans'] text-[32px] font-bold leading-[38.37px] text-white"
                    >
                        <h1>SocialNetwork</h1>
                    </Link>

                    <div className="hidden xl:flex">
                        <SearchBar id="search-bar-header" />
                    </div>

                    <a href="/profile" className="hidden items-center xl:flex xl:relative">
                        <div className="flex items-center justify-center relative z-10 -mr-9 w-11 h-11 bg-white border rounded-full">
                            <ProfileCircle />
                        </div>
                        <div className=" min-w-36 relative justify-end bg-white rounded-3xl bg-opacity-40">
                            <div className="flex flex-col items-center pl-5">
                                <div className="text-xs font-bold">
                                    {user?.firstName} {user?.lastName}
                                </div>
                                <div className="text-xs">@{user?.nickname}</div>
                            </div>
                        </div>
                    </a>

                    <MobileNav />
                </div>
            </header>
        </>
    );
}
