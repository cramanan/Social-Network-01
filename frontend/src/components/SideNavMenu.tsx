"use client";

import React, { useState } from "react";
import { RequestIcon } from "./icons/RequestIcon";
import { HomeIcon } from "./icons/HomeIcon";
import { GroupsIcon } from "./icons/GroupsIcon";
import { NotificationsIcon } from "./icons/NotificationsIcon";
import { SettingIcon } from "./icons/SettingIcon";
import { ExitIcon } from "./icons/ExitIcon";
import { OpenSideMenuIcon } from "./icons/OpenSideMenuIcon";
import { CloseSideMenuIcon } from "./icons/CloseSideMenuIcon";
import { BookmarkMenuIcon } from "./icons/BookmarkMenuIcon";
import FriendInviteList from "./FriendInviteList";

const SideNavMenu = () => {
    const [isOpen, setIsOpen] = useState(false);
    const [isFriendInvListOpen, setFriendInvListOpen] = useState(false);

    const toggleSideNav = () => {
        setIsOpen(!isOpen);
        document
            .getElementById("sideNav")
            ?.classList.toggle("-translate-x-[182px]");
        document
            .getElementById("backIcon")
            ?.classList.toggle("translate-x-[182px]");
        if (!isOpen) {
            setFriendInvListOpen(false);
        }
    };

    const handleFriendInviteIcon = () => {
        setFriendInvListOpen(!isFriendInvListOpen);
        if (isOpen) {
            toggleSideNav();
        }
    };

    const menuItems = [
        { label: "Home", icon: <HomeIcon />, href: "/" },
        {
            label: "Request",
            icon: <RequestIcon />,
            onClick: handleFriendInviteIcon,
        },
        { label: "Groups", icon: <GroupsIcon />, href: "/group" },
        { label: "Notifications", icon: <NotificationsIcon /> },
        { label: "Bookmarks", icon: <BookmarkMenuIcon /> },
        { label: "Setting", icon: <SettingIcon /> },
        { label: "Exit", icon: <ExitIcon /> },
    ];

    return (
        <>
            <nav
                id="sideNav"
                className="w-[267px] h-[667px] relative bg-white/25 rounded-r-[25px] px-5 py-7 -translate-x-[182px] duration-300 ease-in-out select-none"
                aria-label="Side navigation"
            >
                {/* aria-expanded={isOpen} */}
                <ul className="h-full flex flex-col justify-between">
                    <li className={`flex flex-rowitems-center`}>
                        <button
                            id="backIcon"
                            className="translate-x-[182px] duration-300 ease-in-out"
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
                            <span className="text-white text-2xl font-semibold font-['Inter']">
                                {item.label}
                            </span>
                            {item.href ? (
                                <a aria-label={item.label} href={item.href}>
                                    {item.icon}
                                </a>
                            ) : (
                                <button
                                    aria-label={item.label}
                                    onClick={item.onClick}
                                >
                                    {item.icon}
                                </button>
                            )}
                        </li>
                    ))}
                </ul>
            </nav>

            <div
                id="friend-inv-list"
                className={`transition-all duration-300 ease-in-out
                    ${
                        isFriendInvListOpen
                            ? "opacity-100 -translate-x-44 pointer-events-auto"
                            : "opacity-0 -translate-x-40 pointer-events-none"
                    }`}
                aria-label="Friend invite list"
            >
                <FriendInviteList />
            </div>
        </>
    );
};

export default SideNavMenu;
