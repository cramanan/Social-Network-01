"use client";

import React, { ReactNode, useEffect } from "react";
import { webSocketContext } from "./WebSocketContext";

export default function WebSocketProvider({
    children,
}: {
    children: ReactNode;
}) {
    const socket = new WebSocket(
        `${process.env.NEXT_PUBLIC_WEBSOCKET_URL}/api/socket`
    );

    const handleOpen = () => console.log("connected");

    useEffect(() => {
        socket.addEventListener("open", handleOpen);

        return () => socket.removeEventListener("open", handleOpen);
    }, []);

    return (
        <webSocketContext.Provider value={socket}>
            {children}
        </webSocketContext.Provider>
    );
}
