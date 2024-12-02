"use client";

import { OnlineUser } from "@/types/user";
import React, { useEffect, useState } from "react";
import Users from "./Users";

export default function UserList() {
    const [users, setUsers] = useState<OnlineUser[]>([]);

    useEffect(() => {
        const fetchUsers = async () => {
            const response = await fetch("/api/follow-list");
            const data: OnlineUser[] = await response.json();

            setUsers(data);
        };
        fetchUsers();
    }, []);

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
                            No follower(s) found.
                        </p>
                    )}
                </div>
            </div>
        </>
    );
}
