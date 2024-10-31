"use client";
import { useWebSocket } from "@/providers/WebSocketContext";
import { Chat } from "@/types/chat";
import React, { useState } from "react";

export default function Page({ params }: { params: { id: string } }) {
    const [chat, setChat] = useState<Chat>({
        recipientId: params.id,
        content: "Hello World",
    });
    const websocket = useWebSocket();

    if (!websocket) return <>No socket</>;

    return (
        <>
            <button onClick={() => websocket.ping()}>Ping</button>
            <form
                onSubmit={(e) => {
                    e.preventDefault();
                    websocket.sendChat(chat);
                }}
            >
                <div className="h-20">{chat.content}</div>
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
