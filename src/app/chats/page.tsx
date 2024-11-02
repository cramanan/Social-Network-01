"use client";

import { useWebSocket } from "@/providers/WebSocketContext";
import { ServerChat, SocketMessage } from "@/types/chat";
import { User } from "@/types/user";
import Image from "next/image";
import Link from "next/link";
import React, { useEffect, useState } from "react";

type OnlineUser = User & { online: boolean };

export default function Page() {
    // Online users
    const [users, setUsers] = useState<OnlineUser[]>([]);

    const websocket = useWebSocket();

    useEffect(() => {
        // Fetch online users
        const fetchUsers = async () => {
            const response = await fetch("/api/online");
            if (!response.ok) return;
            const users = await response.json();
            setUsers(users);
        };

        fetchUsers();
    }, []);

    // Handle incoming connection and disconnection
    const handleConnection = (msg: MessageEvent) => {
        const message = JSON.parse(msg.data) as SocketMessage<OnlineUser>; // Parse the message to get type
        if (message.type !== "ping") return;

        const update = users.map((user) => {
            if (user.id === message.data.id) user.online = message.data.online;
            return user;
        });

        console.log(update);

        setUsers(update);
    };

    useEffect(() => {
        if (!websocket) return;

        websocket.socket.addEventListener("message", handleConnection);

        return () =>
            websocket.socket.removeEventListener("message", handleConnection);
    }, [users]);

    if (!websocket) return <>No socket</>;

    return (
        <div>
            {users.map((user, idx) => (
                <Link
                    key={idx}
                    href={`/chats/${user.id}`}
                    className="flex items-center gap-2"
                >
                    <div className="relative">
                        <Image src={user.image} width={40} height={40} alt="" />
                        <span
                            className={`h-3 w-3 block absolute bottom-0 right-0 rounded-full bg-${
                                user.online ? "green" : "red"
                            }-500`}
                        />
                        {user.online}
                    </div>
                    <span>{user.id}</span>
                </Link>
            ))}
        </div>
    );
}
