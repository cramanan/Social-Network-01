"use client";

import { useWebSocket } from "@/hooks/useWebSocket";
import { ServerChat } from "@/types/chat";
import React, { PropsWithChildren, useEffect, useState } from "react";

type Toast = {
    type: string;
    message: string;
};

export default function NotificationLayout({ children }: PropsWithChildren) {
    const websocket = useWebSocket();
    const [toast] = useState<Toast[]>([]);

    useEffect(() => {
        if (!websocket || websocket.socket.readyState !== WebSocket.OPEN)
            return;

        const handleMessages = (msg: MessageEvent<ServerChat>) => {
            console.log(msg.data);
        };

        websocket.socket.addEventListener("message", handleMessages);
        return () =>
            websocket.socket.removeEventListener("message", handleMessages);
    }, [websocket]);

    return (
        <>
            {toast.map((toast, idx) => (
                <span key={idx}>{JSON.stringify(toast)}</span>
            ))}
            {children}
        </>
    );
}
