"use client";

import { useAuth } from "@/hooks/useAuth";
import useQueryParams from "@/hooks/useQueryParams";
import { ServerChat } from "@/types/chat";
import formatDate from "@/utils/formatDate";
import { useParams } from "next/navigation";
import React, { ChangeEvent, useEffect, useState } from "react";

export default function Page() {
    const { user, loading } = useAuth();
    const { id } = useParams();
    const [messages, setMessages] = useState<ServerChat[]>([]);
    const { limit, offset } = useQueryParams();
    const [content, setContent] = useState("");
    const [socket, setSocket] = useState<WebSocket | null>(null);

    const onMessage = (msg: MessageEvent) => {
        const message = JSON.parse(msg.data) as ServerChat;
        setMessages([...messages, message]);
    };

    const onSubmit = async (e: ChangeEvent<HTMLFormElement>) => {
        e.preventDefault();
        if (!user) return;
        setContent("");
        if (socket && socket.readyState === WebSocket.OPEN) {
            socket.send(JSON.stringify({ content }));
            setMessages([
                ...messages,
                {
                    senderid: user.id,
                    content,
                    timestamp: formatDate(new Date().toString()),
                    recipientId: "",
                },
            ]);
        }
    };

    useEffect(() => {
        const fetchMessages = async () => {
            try {
                const response = await fetch(
                    `/api/group/${id}/chats?limit=${limit}&offset=${offset}`
                );
                const data = await response.json();
                setMessages(data);
            } catch (error) {
                console.error(error);
            }
        };

        fetchMessages();

        const newSocket = new WebSocket(
            `ws://${process.env.NEXT_PUBLIC_API_URL}/api/group/${id}/chatroom`
        );

        newSocket.addEventListener("message", onMessage);

        setSocket(newSocket);

        return () => {
            newSocket.close();
        };
    }, []);

    if (loading) return <>Loading</>;

    return (
        <div>
            <ul>
                {messages.map(({ senderid, content, timestamp }, idx) => (
                    <li key={idx}>
                        <h2>{senderid}</h2>
                        <div>{content}</div>
                        <div>{timestamp}</div>
                    </li>
                ))}
            </ul>
            <form onSubmit={onSubmit}>
                <input
                    type="text"
                    value={content}
                    onChange={(e) => setContent(e.target.value)}
                />
                <button type="submit">Send</button>
            </form>
        </div>
    );
}
