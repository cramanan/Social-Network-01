"use client";
import { useWebSocket } from "@/providers/WebSocketContext";
import { Chat, ServerChat } from "@/types/chat";
import Link from "next/link";
import React, { useEffect, useState } from "react";

export default function Page({ params }: { params: { id: string } }) {
    const [messages, setMessages] = useState<ServerChat[]>([]);
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
    }, [websocket, params.id]);

    useEffect(() => {
        if (!websocket) return;

        const addMessage = (msg: MessageEvent) => {
            setMessages((prev) => [...prev, JSON.parse(msg.data)]);
        };

        websocket.socket.addEventListener("message", addMessage);

        return () =>
            websocket.socket.removeEventListener("message", addMessage);
    }, [websocket]);

    if (!websocket) return <>No socket</>;

    return (
        <>
            <h1 className="flex justify-between font-bold p-2">
                <Link href="/chats" className="">
                    &lt;-
                </Link>
                <span>{params.id}</span>
                <span />
            </h1>
            <ul className="w-4/5 m-auto flex flex-col px-3 overflow-auto">
                {messages.map((msg, idx) => (
                    <>
                        <li
                            key={idx}
                            className={`flex flex-col w-fit p-2 rounded-3xl ${
                                msg.recipientId === params.id
                                    ? "bg-[#b88ee5] text-black self-end items-end"
                                    : "bg-[#4174e2] text-white self-start"
                            }`}
                        >
                            <div>{msg.recipientId}</div>
                            <p>{msg.content}</p>
                        </li>
                        <div
                            className={`flex flex-col w-fit ${
                                msg.recipientId === params.id
                                    ? "self-end items-end"
                                    : "self-start"
                            }`}
                        >
                            {new Date(msg.timestamp).toLocaleTimeString()}
                        </div>
                    </>
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
