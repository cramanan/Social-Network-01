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
    const handleShowClick = () => {
        setIsOpen(!isOpen)
        document.getElementById("sideNav")?.classList.toggle("-translate-x-44")
        const backIcon = document.getElementById("backIcon")
        backIcon?.classList.toggle("ml-44")
    };
    return (
        <div id='sideNav' className='w-[267px] h-[667px] relative bg-white/25 rounded-r-[25px] border-r-2 border-y-2 border-[#bfbfbf] flex flex-col justify-between px-5 py-7 -translate-x-44 duration-300 ease-in-out'>
            <div id='backIcon' className='ml-44' onClick={handleShowClick}>{isOpen ? <BackSideBarLeft /> : <BackSideBarRight />}</div>
            <div className='flex flex-row justify-between items-center'>
                <span className="text-white text-2xl font-semibold font-['Inter']">Request</span>
                <RequestIcon />
            </div>
            <div className='flex flex-row justify-between items-center'>
                <span className="text-white text-2xl font-semibold font-['Inter']">Home</span>
                <HomeIcon />
            </div>
            <div className='flex flex-row justify-between items-center'>
                <span className="text-white text-2xl font-semibold font-['Inter']">Groups</span>
                <GroupsIcon />
            </div>
            <div className='flex flex-row justify-between items-center'>
                <span className="text-white text-2xl font-semibold font-['Inter']">Notifications</span>
                <NotificationsIcon />
            </div>
            <div className='flex flex-row justify-between items-center'>
                <span className="text-white text-2xl font-semibold font-['Inter']">Bookmarks</span>
                <BookmarkIcon />
            </div>
            <div className='flex flex-row justify-between items-center'>
                <span className="text-white text-2xl font-semibold font-['Inter']">Setting</span>
                <SettingIcon />
            </div>
            <div className='flex flex-row justify-between items-center'>
                <span className="text-white text-2xl font-semibold font-['Inter']">Exit</span>
                <ExitIcon />
            </div>
        </div>
    )
}

export default SideNavMenu