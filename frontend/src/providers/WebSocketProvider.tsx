"use client";

import React from "react";
import { webSocketContext } from "./WebSocketContext";
import { ClientChat } from "@/types/chat";
import { Children } from "@/utils/types";

export default function WebSocketProvider({ children }: Children) {
    const socket = new WebSocket(
        `${process.env.NEXT_PUBLIC_API_URL}/api/socket`
    );

    const sendChat = (chat: ClientChat) => {
        if (!socket || socket.readyState !== WebSocket.OPEN) return;

        socket.send(JSON.stringify({ type: "message", data: chat }));
    };

    return (
        <webSocketContext.Provider value={{ socket, sendChat }}>
            {children}
        </webSocketContext.Provider>
    );
}
