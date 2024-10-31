"use client";
import { useWebSocket } from "@/providers/WebSocketContext";
import { Chat } from "@/types/chat";
import React, { useEffect, useState } from "react";

export default function Page({ params }: { params: { id: string } }) {
    const [messages, setMessages] = useState<Chat[]>([]);
    const [chat, setChat] = useState<Chat>({
        recipientId: params.id,
        content: "",
    });
    const websocket = useWebSocket();

    useEffect(() => {
        const fetchMessages = async () => {
            const response = await fetch(`/api/user/${params.id}/chats`);
            const data = await response.json();
            setMessages(data);
        };

        fetchMessages();
    }, []);

    useEffect(() => {
        if (!websocket) return;

        const addMessage = (msg: MessageEvent) => {
            setMessages((prev) => [...prev, JSON.parse(msg.data)]);
        };

        websocket.socket.addEventListener("message", addMessage);

        return () =>
            websocket.socket.removeEventListener("message", addMessage);
    }, []);

    if (!websocket) return <>No socket</>;

    return (
        <>
            <h1 className="font-bold text-center">{params.id}</h1>
            <ul>
                {messages.map((msg, idx) => (
                    <li
                        key={idx}
                        className={`${
                            msg.recipientId === params.id
                                ? "bg-blue-200"
                                : "bg-red-200"
                        }`}
                    >
                        {msg.content}, {msg.recipientId}
                    </li>
                ))}
            </ul>
            <form
                onSubmit={(e) => {
                    e.preventDefault();
                    websocket.sendChat(chat);
                    setChat({ ...chat, content: "" });
                }}
            >
                <textarea
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
