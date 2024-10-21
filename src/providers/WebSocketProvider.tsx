"use client";

import React, { ReactNode } from "react";
import { webSocketContext } from "./WebSocketContext";

export default function WebSocketProvider({
    children,
}: {
    children: ReactNode;
}) {
    const socket = new WebSocket(
        `${process.env.NEXT_PUBLIC_WEBSOCKET_URL}/api/socket`
    );

    return (
        <webSocketContext.Provider value={socket}>
            {children}
        </webSocketContext.Provider>
    );
}
