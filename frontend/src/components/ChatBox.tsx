"use client";

import React, { useEffect, useState } from "react";
import { BackIcon } from "./icons/BackIcon";
import { SendIcon } from "./icons/SendIcon";
import { EmoteIcon } from "./icons/EmoteIcon";
import { User } from "@/types/user";
import { ClientChat, ServerChat, SocketMessage } from "@/types/chat";
import { useWebSocket } from "@/hooks/useWebSocket";
import Image from "next/image";

interface ChatBoxProps {
    onClose?: () => void;
    recipient: User;
}

const ChatBox = ({ onClose, recipient }: ChatBoxProps) => {
    // Incoming messages state array
    const [messages, setMessages] = useState<ServerChat[]>([]);

    // outcoming message state object
    const [chat, setChat] = useState<ClientChat>({
        recipientId: recipient.id,
        content: "",
    });

    // fetch latest messages
    useEffect(() => {
        const fetchMessages = async () => {
            const response = await fetch(`/api/user/${recipient.id}/chats`);
            const data = await response.json();
            setMessages(data);
        };

        fetchMessages();
    }, [recipient.id]);

    // retrieve WebSocket from Context
    const websocket = useWebSocket();

    // Add event listener on mount
    useEffect(() => {
        const addMessage = (msg: MessageEvent) => {
            const message = JSON.parse(msg.data) as SocketMessage<ServerChat>;
            if (message.type !== "message") return;

            console.log(message);

            setMessages((prev) => [...prev, message.data]);
        };

        websocket.socket.addEventListener("message", addMessage);

        // Remove event listenet on unmount with a closure function
        return () =>
            websocket.socket.removeEventListener("message", addMessage);
    }, [websocket]);

    // if the socket is somehow null
    if (!websocket) return <>No socket</>;

    return (
        <>
            <div
                id="chatBox"
                className="flex flex-col w-full h-full relative xl:w-[343px] xl:rounded-[25px] xl:h-[642px] xl:bg-[#fbfbfb]"
            >
                <div className="flex flex-row w-full min-h-14 items-center justify-between border-b border-black px-3 xl:rounded-tl-[25px] xl:rounded-t-[25px] xl:bg-[#445ab3]/20 xl:w-[343px]">
                    {/* {isMobile ? (
                        <Link href="/chats" onClick={onClose}>
                            <BackIcon />
                        </Link>
                    ) : ( */}
                    <button onClick={onClose}>
                        <BackIcon />
                    </button>
                    {/* )} */}

                    <span>{recipient.nickname}</span>

                    <Image
                        src={recipient.image}
                        alt=""
                        width={40}
                        height={40}
                        className="w-9 h-9 border border-black rounded-full"
                    />
                </div>

                <ul className="flex flex-col flex-grow px-3 py-2 overflow-scroll no-scrollbar">
                    {messages.map((msg, index) => {
                        const isRecipient = msg.recipientId === recipient.id;
                        const timestamp = new Date(msg.timestamp);

                        return (
                            <li
                                key={index}
                                className={`flex flex-col ${
                                    isRecipient
                                        ? " self-end items-end"
                                        : " self-start"
                                }`}
                            >
                                <p
                                    className={`p-3 rounded-2xl w-fit ${
                                        isRecipient
                                            ? "bg-[#b88ee5] text-black"
                                            : "bg-[#4174e2] text-white"
                                    }`}
                                >
                                    {msg.content}
                                </p>
                                <div>
                                    {timestamp.toLocaleDateString()},
                                    {timestamp.toLocaleTimeString()}
                                </div>
                            </li>
                        );
                    })}
                </ul>

                <form
                    id="chatMessageForm"
                    onSubmit={(e) => {
                        e.preventDefault();
                        websocket.sendChat(chat);
                        setChat({ ...chat, content: "" });
                    }}
                    className="h-[50px] flex flex-row items-center m-5 bg-[#445ab3]/20 rounded-[25px] p-2 gap-2"
                >
                    <EmoteIcon />
                    <input
                        id="chatMessage"
                        type="text"
                        placeholder="Enter your message"
                        onChange={(e) =>
                            setChat({ ...chat, content: e.target.value })
                        }
                        value={chat.content}
                        className="bg-white/0 w-full placeholder:text-black"
                    ></input>
                    <button type="submit">
                        <SendIcon />
                    </button>
                </form>
            </div>
        </>
    );
};

export default ChatBox;
