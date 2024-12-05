"use client";

import useQueryParams from "@/hooks/useQueryParams";
import { useWebSocket } from "@/hooks/useWebSocket";
import { SocketMessage } from "@/types/chat";
import { OnlineUser } from "@/types/user";
import React, { useCallback, useEffect, useState } from "react";
import Users from "./Users";
import ChatBox from "./ChatBox";

const ChatList = () => {
    // Online users
    const [users, setUsers] = useState<OnlineUser[]>([]);
    const { limit, offset } = useQueryParams();
    const [selectedUser, setSelectedUser] = useState<OnlineUser | null>(null);
    const [ShowUserList, setShowUserList] = useState(true);

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
        websocket.socket.addEventListener("message", callback);

        // Remove when the component is unmounted
        return () => websocket.socket.removeEventListener("message", callback);
    }, [websocket, callback]);

    const handleUserSelect = (user: OnlineUser) => {
        setSelectedUser(user);
        setShowUserList(!ShowUserList);
    };

    const handleCloseChatBox = () => {
        setSelectedUser(null);
        setShowUserList(!ShowUserList);
    };

    return (
        <>
            {ShowUserList && (
                <div
                    className="relative flex flex-col w-full h-[calc(100vh-185px)] xl:w-72 xl:h-fit xl:rounded-3xl xl:py-3 xl:bg-white/40"
                    aria-label="Chat list"
                    role="region"
                >
                    <h2 className="text-4xl text-white text-center py-5 xl:sr-only">
                        Chat List
                    </h2>

                    <div className="flex flex-col gap-3 mx-5 overflow-scroll no-scrollbar xl:max-h-[65vh]">
                        {users.length > 0 ?
                            users.map((user, index) => (
                                <Users
                                    key={index}
                                    user={user}
                                    onUserSelect={handleUserSelect}
                                    showLastMessage={true}
                                />
                            ))
                            :
                            <p className="text-center font-bold">
                                No Conversation(s) found.
                            </p>
                        }
                    </div>
                </div>
            )}

            {selectedUser && (
                <div className="w-full h-full xl:w-fit xl:h-fit xl:absolute xl:right-0">
                    <ChatBox
                        recipient={selectedUser}
                        onClose={handleCloseChatBox}
                    />
                </div>
            )}
        </>
    );
};

export default ChatList;
