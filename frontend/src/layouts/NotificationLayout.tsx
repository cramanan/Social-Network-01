"use client";

import { useWebSocket } from "@/hooks/useWebSocket";
import { ServerChat, SocketMessage, SocketMessageType } from "@/types/chat";
import React, { PropsWithChildren, useEffect, useState } from "react";

type Toast = {
    type: SocketMessageType;
    message: string;
};

export default function NotificationLayout({ children }: PropsWithChildren) {
    const { socket } = useWebSocket();
    const [toasts, setToasts] = useState<Toast[]>([]);

    const handleMessages = (msg: MessageEvent) => {
        try {
            const message = JSON.parse(msg.data) as SocketMessage<ServerChat>;

            const toast: Toast = {
                type: message.type,
                message: "",
            };
            switch (message.type) {
                case "message":
                    toast.message = "New Message";
                    break;

                case "ping":
                    toast.message = "Ping";
                    break;

                case "friend-request":
                    toast.message = "New Friend";
                    break;

                default:
                    break;
            }
            setToasts((prevToasts) => [...prevToasts, toast]);
            setTimeout(
                () => setToasts((prev) => prev.filter((t) => t !== toast)),
                1000
            );
        } catch (error) {
            console.error(error);
        }
    };
    socket.addEventListener("open", () => {
        console.log("opened");
        socket.addEventListener("message", handleMessages);
    });

    return (
        <>
            <div className="fixed">
                {toasts.map((toast, idx) => (
                    <div key={idx} className="bg-white">
                        <h1>Notification</h1>
                        <p>{toast.message}</p>
                    </div>
                ))}
            </div>
            {children}
        </>
    );
}
