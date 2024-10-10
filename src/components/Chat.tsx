'use client'

import React from "react";
import { ChatBubbles } from "./icons/ChatBubbles";
import { Icon } from "./icons/Icon";
import { Vector } from "./icons/Vector";

const Chat = () => {
    const handleVectorClick = () => {
        const userList = document.getElementById("userList")
        userList?.classList.toggle("hidden")
    };
    return (
        <>
            <div className="w-64 h-9 bg-white bg-opacity-40 m-3 border border-neutral-400 rounded-b-3xl flex justify-between items-center px-7">
                <ChatBubbles />
                <Icon />
                <button onClick={handleVectorClick}>
                    <Vector />
                </button>
            </div>
        </>
    );
}

export default Chat