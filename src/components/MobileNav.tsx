'use client'

import React, { useState } from 'react'
import SearchBar from "./SearchBar";
import { HomeIcon } from "./icons/HomeIcon";

const MobileNav = () => {
    const [isOpen, setIsOpen] = useState(false)


    const handleClick = () => {
        setIsOpen(!isOpen);
    };
    const menuLine = `h-2 w-10 my-1 rounded-full bg-white ease transform duration-300`;
    return (
        <>
            <button
                className="flex flex-col h-12 w-12 z-50 xl:hidden" onClick={handleClick}
            >
                <div className={`${menuLine} ${isOpen ? "rotate-45 translate-y-4" : ""}`} />
                <div className={`${menuLine} ${isOpen ? "opacity-0" : ""}`} />
                <div className={`${menuLine} ${isOpen ? "-rotate-45 -translate-y-4" : ""}`} />
            </button>

            <nav className={`fixed top-0 right-0 w-full h-full border-gradient-test duration-300 ease-in-out z-40
            ${isOpen ? "opacity-100 -translate-x-0 pointer-events-auto" : "translate-x-1/2 opacity-0 pointer-events-none"}
            `}
            >
                <ul className="flex flex-col relative w-full h-full mt-14">
                    <li className="w-full flex justify-center border-b border-black p-2"><SearchBar id={"mobile-search-nav"} /></li>
                    <li className="border-b border-black p-2">
                        <a href="" className="w-full flex flex-row items-center justify-center gap-5" >
                            <HomeIcon />
                            <span className="font-bold text-white font-['Inter'] text-2xl">Home</span>
                        </a>
                    </li>
                </ul>
            </nav>
        </>
    )
}

export default MobileNav