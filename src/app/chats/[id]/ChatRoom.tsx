"use client";

import { useWebSocket } from "@/providers/WebSocketContext";
import { ClientChat, ServerChat, SocketMessage } from "@/types/chat";
import { User } from "@/types/user";
import React, { useEffect, useState } from "react";

export default function ChatRoom({ recipient }: { recipient: User }) {
    const [messages, setMessages] = useState<ServerChat[]>([]);
    const [chat, setChat] = useState<ClientChat>({
        recipientId: recipient.id,
        content: "",
    });

    const websocket = useWebSocket();

    useEffect(() => {
        const fetchMessages = async () => {
            const response = await fetch(`/api/user/${recipient.id}/chats`);
            const data = await response.json();
            setMessages(data);
        };

        fetchMessages();
    }, [recipient.id]);

    useEffect(() => {
        if (!websocket) return;

        const addMessage = (msg: MessageEvent) => {
            const message = JSON.parse(msg.data) as SocketMessage<ServerChat>;
            if (message.type != "message") return;

            setMessages((prev) => [...prev, message.data]);
        };

        websocket.socket.addEventListener("message", addMessage);

        return () =>
            websocket.socket.removeEventListener("message", addMessage);
    }, [websocket]);

    if (!websocket) return <>No socket</>;

    return (
        <>
            <ul className="w-4/5 m-auto flex flex-col px-3 overflow-auto">
                {messages.map((msg, idx) => (
                    <li
                        key={idx}
                        className={`flex flex-col w-fit ${
                            msg.recipientId === recipient.id
                                ? " self-end items-end"
                                : " self-start"
                        }`}
                    >
                        <p
                            className={`p-3 rounded-2xl ${
                                msg.recipientId === recipient.id
                                    ? "bg-[#b88ee5] text-black"
                                    : "bg-[#4174e2] text-white"
                            }`}
                        >
                            {msg.content}
                        </p>
                        <div>
                            {new Date(msg.timestamp).toLocaleDateString()},
                            {new Date(msg.timestamp).toLocaleTimeString()}
                        </div>
                    </li>
                ))}
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
