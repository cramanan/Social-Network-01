"use client";

import { useWebSocket } from "@/hooks/useWebSocket";
import { SocketMessage, SocketMessageType } from "@/types/chat";
import React, { PropsWithChildren, useState } from "react";

type Toast = {
    type: SocketMessageType;
    message: string;
};

export default function NotificationLayout({ children }: PropsWithChildren) {
    const { socket } = useWebSocket();
    const [toasts, setToasts] = useState<Toast[]>([]);

    const handleMessages = (msg: MessageEvent) => {
        try {
            const message = JSON.parse(msg.data) as SocketMessage;

            const toast: Toast = {
                type: message.type,
                message: "",
            };

            switch (message.type) {
                case "event":
                case "group-request":
                case "group-invite":
                case "follow-request":
                    toast.message = message.data as string;
                    break;

                default:
                    return;
            }

            setToasts((prevToasts) => [...prevToasts, toast]);
            setTimeout(
                () => setToasts((prev) => prev.filter((t) => t !== toast)),
                3000
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
                {toasts.map(({}, idx) => {
                    const { message } = toasts[toasts.length - 1 - idx];
                    return (
                        <div key={idx} className="bg-white m-2 p-3 rounded-3xl">
                            <h1>Notification</h1>
                            <p>{message}</p>
                        </div>
                    );
                })}
            </div>
            {children}
        </>
    );
}
