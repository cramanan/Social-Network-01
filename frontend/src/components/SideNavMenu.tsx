"use client";

import React, { useState } from "react";
import { HomeIcon } from "./icons/HomeIcon";
import { GroupsIcon } from "./icons/GroupsIcon";
import { NotificationsIcon } from "./icons/NotificationsIcon";
import { SettingIcon } from "./icons/SettingIcon";
import { ExitIcon } from "./icons/ExitIcon";
import { OpenSideMenuIcon } from "./icons/OpenSideMenuIcon";
import { CloseSideMenuIcon } from "./icons/CloseSideMenuIcon";

const SideNavMenu = () => {
    const [isOpen, setIsOpen] = useState(false);

    const toggleSideNav = () => setIsOpen(!isOpen);

    const menuItems = [
        { label: "Home", icon: <HomeIcon />, href: "/" },
        { label: "Groups", icon: <GroupsIcon />, href: "/group" },
        { label: "Inbox", icon: <NotificationsIcon />, href: "/inbox" },
        { label: "Setting", icon: <SettingIcon />, href: "/profile/settings" },
        { label: "Exit", icon: <ExitIcon /> },
    ];

    return (
        <>
            <nav
                id="sideNav"
                className={`w-[267px] h-[467px] relative bg-white/25 rounded-r-[25px] px-5 py-7 ${!isOpen && "-translate-x-[182px]"
                    } duration-300 ease-in-out select-none`}
                aria-label="Side navigation"
            >
                {" "}
                {/* aria-expanded={isOpen} */}
                <ul className="h-full flex flex-col justify-between">
                    <li className={`flex flex-rowitems-center`}>
                        <button
                            id="backIcon"
                            className={`${!isOpen && "translate-x-[182px]"
                                } duration-300 ease-in-out`}
                            aria-label={isOpen ? "Close menu" : "Open menu"}
                            onClick={toggleSideNav}
                        >
                            {isOpen ? (
                                <CloseSideMenuIcon />
                            ) : (
                                <OpenSideMenuIcon />
                            )}
                        </button>
                    </li>

                    {menuItems.map((item, index) => (
                        <li
                            key={index}
                            className="flex flex-row justify-between items-center"
                        >
                            <div className="flex justify-between w-full text-white text-2xl font-semibold font-['Inter']">
                                {item.label}
                                <a aria-label={item.label} href={item.href}>
                                    {item.icon}
                                </a>
                            </div>
                        </li>
                    ))}
                </ul>
            </nav>
        </>
    );
};

export default SideNavMenu;
