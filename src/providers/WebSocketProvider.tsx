"use client";

import React, { ReactNode } from "react";
import { webSocketContext } from "./WebSocketContext";
import { ClientChat } from "@/types/chat";

export default function WebSocketProvider({
    children,
}: {
    children: ReactNode;
}) {
    const socket = new WebSocket(
        `${process.env.NEXT_PUBLIC_WEBSOCKET_URL}/api/socket`
    );

    const sendChat = (chat: ClientChat) => {
        if (!socket) return;
        if (socket.readyState !== WebSocket.OPEN) return;

        socket.send(JSON.stringify({ type: "message", data: chat }));
    };

    return (
        <webSocketContext.Provider value={{ socket, sendChat }}>
            {children}
        </webSocketContext.Provider>
    );
}
