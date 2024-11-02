"use client";

import { useWebSocket } from "@/providers/WebSocketContext";
import { ClientChat, ServerChat, SocketMessage } from "@/types/chat";
import { User } from "@/types/user";
import React, { useEffect, useState } from "react";

export default function ChatRoom({ recipient }: { recipient: User }) {
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
        if (!websocket) return;

        const addMessage = (msg: MessageEvent) => {
            const message = JSON.parse(msg.data) as SocketMessage<ServerChat>;
            if (message.type !== "message") return;

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
            <ul className="w-4/5 m-auto flex flex-col px-3 overflow-auto">
                {messages.map((msg, idx) => {
                    const isRecipient = msg.recipientId === recipient.id;
                    const timestamp = new Date(msg.timestamp);

                    return (
                        <li
                            key={idx}
                            className={`flex flex-col w-fit ${
                                isRecipient
                                    ? " self-end items-end"
                                    : " self-start"
                            }`}
                        >
                            <p
                                className={`p-3 rounded-2xl ${
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
                className="w-fit m-auto"
                onSubmit={(e) => {
                    e.preventDefault();
                    websocket.sendChat(chat);
                    setChat({ ...chat, content: "" });
                }}
            >
                <textarea
                    className="resize-none"
                    onChange={(e) =>
                        setChat({ ...chat, content: e.target.value })
                    }
                    value={chat.content}
                />
                <button type="submit">Send</button>
            </form>
        </>
    );
}
