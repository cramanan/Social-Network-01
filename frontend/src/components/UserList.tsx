"use client";

import { OnlineUser } from "@/types/user";
import React, { useCallback, useEffect, useState } from "react";
import Users from "./Users";
import { useWebSocket } from "@/hooks/useWebSocket";
import { SocketMessage } from "@/types/chat";

export default function UserList() {
    const [users, setUsers] = useState<OnlineUser[]>([]);
    const { socket } = useWebSocket();

    useEffect(() => {
        const fetchUsers = async () => {
            const response = await fetch("/api/profile/following");
            const data: OnlineUser[] = await response.json();

            setUsers(data);
        };

        fetchUsers();
    }, []);

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
        socket.addEventListener("message", callback);

        // Remove when the component is unmounted
        return () => socket.removeEventListener("message", callback);
    }, [socket, callback]);

    return (
        <>
            <div
                id="userList"
                className="flex flex-col w-full h-[calc(100vh-130px)] xl:w-72 xl:h-fit xl:bg-white/40 xl:rounded-3xl xl:py-3"
            >
                <h2 className="text-4xl text-white text-center py-5 xl:sr-only">
                    Follow List
                </h2>

                <div className="flex flex-col items-center gap-3 mx-5 overflow-scroll no-scrollbar xl:max-h-[65vh]">
                    {users.length > 0 ? (
                        users.map((user, idx) => (
                            <Users
                                key={idx}
                                user={user}
                                showLastMessage={false}
                            />
                        ))
                    ) : (
                        <p className="text-center font-bold">
                            No follow(s) found.
                        </p>
                    )}
                </div>
            </div>
        </>
    );
}
