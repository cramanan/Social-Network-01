"use client";

import React, { useState } from "react";
import { UserListIcon } from "./icons/UserListIcon";
import UserList from "./UserList";
import FindUser from "./FindUser";
import ChatList from "./ChatList";
import ChatIcon from "./icons/ChatIcon";
import { FindUserIcon } from "./icons/FindUserIcon";

const Chat = () => {
    const [windows, setWindows] = useState([false, false, false]);
    const handleClick = (i: number) => () => {
        const next = [false, false, false];
        if (!windows[i]) next[i] = true;
        setWindows(next);
    };

    const navItems = [
        {
            icon: <ChatIcon />,
            label: "Chat list",
        },
        {
            icon: <FindUserIcon />,
            label: "FindUser list",
        },
        {
            icon: <UserListIcon />,
            label: "User list",
        },
    ];

    const navBody = [ChatList, FindUser, UserList];

    return (
        <div>
            <nav
                id="chat-nav"
                className="w-72 h-12 flex bg-white bg-opacity-40 mr-3 rounded-b-3xl px-7"
                aria-label="Chat Navigation"
            >
                <ul className="w-full flex flex-row justify-between items-center">
                    {navItems.map((items, index) => (
                        <li key={index}>
                            <button
                                aria-label={`Toggle ${items.label}`}
                                aria-expanded={windows[index]}
                                onClick={handleClick(index)}
                            >
                                <span className="sr-only">{items.label}</span>
                                {items.icon}
                            </button>
                        </li>
                    ))}
                </ul>
            </nav>
            <ul className="pt-2">
                {navBody.map((Component, idx) => (
                    <li
                        key={idx}
                        className={`absolute w-72 transition-all duration-300 ease-in-out ${
                            windows[idx]
                                ? "opacity-100 translate-y-0 pointer-events-auto"
                                : "opacity-0 translate-y-5 pointer-events-none"
                        }`}
                    >
                        <Component />
                    </li>
                ))}
            </ul>
        </div>
    );
};

export default Chat;
