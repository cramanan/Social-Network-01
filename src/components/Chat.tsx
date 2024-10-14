'use client'

import React, { useState } from "react";
import { ChatBubbles } from "./icons/ChatBubbles";
import { Icon } from "./icons/Icon";
import { UserListIcon } from "./icons/UserListIcon";
import UserList from "./UserList";

const Chat = () => {
    // const [isChatBox, setIsChatBoxOpen] = useState("false")
    // const [isMessageListOpen, setIsMessageListOpen] = useState("false")
    const [isUserListOpen, setIsUserListOpen] = useState(false)

    const handleUserListIconClick = () => {
        setIsUserListOpen(!isUserListOpen)
    };
    const handleChatIconClick = () => {

    };
    return (
        <>
            <nav id="chat-nav" className="w-72 h-12 flex bg-white bg-opacity-40 m-3 border border-neutral-400 rounded-b-3xl px-7" aria-label="Chat Navigation">
                <ul className="w-full flex flex-row justify-between items-center">
                    <li>
                        <button onClick={handleChatIconClick} aria-label="Toggle chat box">
                            <span className="sr-only">Chat</span>
                            <ChatBubbles />
                        </button>
                    </li>

                    <li>
                        <button aria-label="Toggle Message list">
                            <span className="sr-only">Message list</span>
                            <Icon />
                        </button>
                    </li>

                    <li>
                        <button onClick={handleUserListIconClick} aria-label="Toggle User list" aria-expanded={isUserListOpen}>
                            <span className="sr-only">User list</span>
                            <UserListIcon />
                        </button>
                    </li>
                </ul>
            </nav>

            <div id="user-list" className={`
                    absolute right-3 transition-all duration-300 ease-in-out
                    ${isUserListOpen
                    ? 'opacity-100 translate-y-0 pointer-events-auto'
                    : 'opacity-0 translate-y-5 pointer-events-none'}
                `}>
                <UserList />
            </div>
        </>
    );
}

export default Chat