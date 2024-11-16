"use client";

import { useWebSocket } from "@/hooks/useWebSocket";
import { Children } from "@/utils/types";
import React, { useEffect } from "react";

export default function NotificationLayout({ children }: Children) {
    const websocket = useWebSocket();

    const handleMessages = (msg: MessageEvent) => {
        console.log(msg.data);
    };

    useEffect(() => {
        if (!websocket || websocket.socket.readyState !== WebSocket.OPEN)
            return;

        websocket.socket.addEventListener("message", handleMessages);
        return () =>
            websocket.socket.removeEventListener("message", handleMessages);
    }, [websocket, handleMessages]);

    return <div>{children}</div>;
}
