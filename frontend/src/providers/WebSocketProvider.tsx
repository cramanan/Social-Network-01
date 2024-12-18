"use client";

import React, { PropsWithChildren } from "react";
import { webSocketContext } from "./WebSocketContext";
import { ClientChat } from "@/types/chat";

export default function WebSocketProvider({ children }: PropsWithChildren) {
    const socket = new WebSocket(
        `ws://${process.env.NEXT_PUBLIC_API_URL}/api/socket`
    );

    const sendChat = (chat: ClientChat) => {
        if (socket.readyState !== WebSocket.OPEN) return;

        socket.send(JSON.stringify({ type: "message", data: chat }));
    };

    return (
        <webSocketContext.Provider value={{ socket, sendChat }}>
            {children}
        </webSocketContext.Provider>
    );
}
