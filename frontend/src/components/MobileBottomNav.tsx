"use client";

import { HomeIcon } from "./icons/HomeIcon";
import { FindUserIcon } from "./icons/FindUserIcon";
import { UserListIcon } from "./icons/UserListIcon";
import ChatIcon from "./icons/ChatIcon";
import ChatList from "./ChatList";
import { useState } from "react";
import UserList from "./UserList";

const MobileBottomNav = () => {
    const [showChatList, setShowChatList] = useState(false);
    const [showFollowList, setShowFollowList] = useState(false);

    const handleChatListClick = () => {
        setShowChatList(!showChatList);
        setShowFollowList(false);
    };
    const handleFollowListClick = () => {
        setShowFollowList(!showFollowList);
        setShowChatList(false);
    };

    return (
        <>
            <div className="relative w-full h-full z-60">
                <nav
                    className="relative flex flex-row w-full h-16 bg-[#FFFFFF42] border-t border-white justify-between items-center"
                    aria-label="mobile bottom navigation"
                >
                    <ul className="flex flex-row w-full justify-evenly">
                        <li>
                            <a href="/">
                                <span className="sr-only">Home</span>
                                <HomeIcon />
                            </a>
                        </li>

                        <li>
                            <button onClick={handleFollowListClick}>
                                <span className="sr-only">FollowList</span>
                                <UserListIcon />
                            </button>
                        </li>

                        <li>
                            <button>
                                <span className="sr-only">FindUser</span>
                                <FindUserIcon />
                            </button>
                        </li>

                        <li>
                            {/* <a href="/chats">
                                <span className="sr-only">Chat</span><ChatIcon />
                            </a> */}
                            <button onClick={handleChatListClick}>
                                <span className="sr-only">Chat</span>
                                <ChatIcon />
                            </button>
                        </li>
                    </ul>
                </nav>

                <div
                    className={`w-full h-[calc(100vh-64px)] fixed bottom-16 right-0 border-gradient-test duration-300 ease-in-out z-50 
                    ${showChatList ? "translate-x-0" : "translate-x-full"}
                    `}
                >
                    <ChatList />
                </div>

                <div
                    className={`w-full h-[calc(100vh-64px)] fixed bottom-16 right-0 border-gradient-test duration-300 ease-in-out z-50 
                    ${showFollowList ? "translate-x-0" : "translate-x-full"}
                    `}
                >
                    <UserList />
                </div>
            </div>
        </>
    );
};

export default MobileBottomNav;
