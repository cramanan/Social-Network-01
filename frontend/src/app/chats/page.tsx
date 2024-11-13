"use client";

import ChatList from "@/components/ChatList";
import useQueryParams from "@/hooks/useQueryParams";
import { useWebSocket } from "@/hooks/useWebSocket";
import HomeProfileLayout from "@/layouts/HomeProfileLayout";
import { SocketMessage } from "@/types/chat";
import { OnlineUser } from "@/types/user";
import React, { useCallback, useEffect, useState } from "react";

export default function Page() {
    // Online users
    const [users, setUsers] = useState<OnlineUser[]>([]);
    const { limit, offset } = useQueryParams();

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
            <HomeProfileLayout>
                <ChatList />
            </HomeProfileLayout>
        </>
    );
}
