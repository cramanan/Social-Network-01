"use client";

import { useWebSocket } from "@/providers/WebSocketContext";
import React from "react";

export default function Page() {
    const socket = useWebSocket();

    if (!socket) return <>No Socket</>;

    return <>{socket.readyState}</>;
}
