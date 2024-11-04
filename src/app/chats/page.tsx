"use client";

import useQueryParams from "@/hooks/useQueryParams";
import { useWebSocket } from "@/providers/WebSocketContext";
import { SocketMessage } from "@/types/chat";
import { OnlineUser } from "@/types/user";
import Image from "next/image";
import Link from "next/link";
import React, { useCallback, useEffect, useState } from "react";

export default function Page() {
    // Online users
    const [users, setUsers] = useState<OnlineUser[]>([]);
    const { limit, offset, next, previous } = useQueryParams();

    useEffect(() => {
        // Fetch online users
        const fetchUsers = async () => {
            const response = await fetch(
                `/api/online?limit=${limit}&offset=${offset}`
            );
            if (!response.ok) return;
            const users = await response.json();
            setUsers(users);
        };

        fetchUsers();
    }, [limit, offset]);

    const websocket = useWebSocket();

    // Handle incoming connection and disconnection
    const handleConnection = (msg: MessageEvent) => {
        const message = JSON.parse(msg.data) as SocketMessage<OnlineUser>; // Parse the message to get type
        if (message.type !== "ping") return; // ignore if it is not of type "ping"

        // Check for each displayed user if the user is incoming/outcoming
        const update = users.map((user) => {
            if (user.id === message.data.id) user.online = message.data.online;
            return user;
        });

        setUsers(update); // Update displayed users
    };

    // Memoize the function
    const callback = useCallback(handleConnection, [users]);

    // Add event listener
    useEffect(() => {
        if (!websocket) return;

        websocket.socket.addEventListener("message", callback);

        // Remove when the component is unmounted
        return () => websocket.socket.removeEventListener("message", callback);
    }, [websocket, callback]);

    // If the websocket somehow couldn't load
    if (!websocket) return <>No socket</>;

    return (
        <>
            <div>
                {users.map((user, idx) => (
                    <Link
                        key={idx}
                        href={`/chats/${user.id}`}
                        className="flex items-center gap-2"
                    >
                        <div className="relative">
                            <span
                                className={`h-3 w-3 block absolute bottom-0 right-0 rounded-full bg-${
                                    user.online ? "green" : "red"
                                }-500 z-10`}
                            />
                            <Image
                                src={user.image}
                                width={40}
                                height={40}
                                alt=""
                            />
                            {user.online}
                        </div>
                        <span>{user.nickname}</span>
                    </Link>
                ))}
            </div>
            <button className="w-fit" onClick={next}>
                next
            </button>
            <button className="w-fit" onClick={previous}>
                previous
            </button>
        </>
    );
}
