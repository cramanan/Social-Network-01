'use client'

import React, { useState } from 'react'
import { BackSideBarRight } from './icons/BackSideBarRight';
import { BackSideBarLeft } from './icons/BackSideBarLeft';
import { RequestIcon } from './icons/RequestIcon';
import { HomeIcon } from './icons/HomeIcon';
import { GroupsIcon } from './icons/GroupsIcon';
import { NotificationsIcon } from './icons/NotificationsIcon';
import { BookmarkIcon } from './icons/BookmarkIcon';
import { SettingIcon } from './icons/SettingIcon';
import { ExitIcon } from './icons/ExitIcon';


const SideNavMenu = () => {
    const [isOpen, setIsOpen] = useState(false)

    const toggleSideNav = () => {
        setIsOpen(!isOpen)
        document.getElementById("sideNav")?.classList.toggle("-translate-x-44")
        document.getElementById("backIcon")?.classList.toggle("translate-x-44")
        if (!isOpen) {
            document.getElementById("friendInviteList")?.classList.add("hidden")
        }
    };

    const handleFriendInviteIcon = () => {
        const friendInviteList = document.getElementById("friendInviteList")
        friendInviteList?.classList.toggle("hidden")
        friendInviteList?.classList.toggle("flex")
        if (isOpen) {
            toggleSideNav();
        }
    }

    const menuItems = [
        { label: "Request", icon: <RequestIcon />, onClick: handleFriendInviteIcon },
        { label: "Home", icon: <HomeIcon />, href: "/" },
        { label: "Groups", icon: <GroupsIcon /> },
        { label: "Notifications", icon: <NotificationsIcon /> },
        { label: "Bookmarks", icon: <BookmarkIcon /> },
        { label: "Setting", icon: <SettingIcon /> },
        { label: "Exit", icon: <ExitIcon /> },
    ];

    return (
        <>
            <nav id='sideNav' className='w-[267px] h-[667px] relative bg-white/25 rounded-r-[25px] px-5 py-7 -translate-x-44 duration-300 ease-in-out select-none' aria-label="Side navigation" aria-expanded={isOpen}>
                <ul className="h-full flex flex-col justify-between">

                    <li>
                        <button id='backIcon' className='w-[51px] translate-x-44 duration-300 ease-in-out' aria-label={isOpen ? "Close menu" : "Open menu"} onClick={toggleSideNav}>{isOpen ? <BackSideBarLeft /> : <BackSideBarRight />}</button>
                    </li>

                    {menuItems.map((item, index) => (
                        <li key={index} className="flex flex-row justify-between items-center">
                            <span className="text-white text-2xl font-semibold font-['Inter']">{item.label}</span>
                            {item.href ? (
                                <a aria-label={item.label} href={item.href}>{item.icon}</a>
                            ) : (
                                <button aria-label={item.label} onClick={item.onClick}>{item.icon}</button>
                            )}
                        </li>
                    ))}
                </ul>
            </nav >
        </>
    )
}

export default SideNavMenu