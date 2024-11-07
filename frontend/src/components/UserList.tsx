"use client";

import React, { useEffect, useState } from "react";
import Users from "./Users";
import useQueryParams from "@/hooks/useQueryParams";
import { OnlineUser } from "@/types/user";


export default function UserList() {
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

    return (
        <>
            <div
                id="userList"
                className="flex flex-col w-full h-[calc(100vh-130px)] xl:w-60 xl:h-fit xl:bg-white/40 xl:rounded-3xl xl:py-3"
            >
                <h2 className="text-4xl text-white text-center py-5 xl:sr-only">Friend List</h2>

                <div className="flex flex-col items-center gap-3 mx-5 overflow-scroll no-scrollbar xl:max-h-[65vh]">
                    {users.map((user, index) => (
                        <Users key={index} user={user} />
                    ))}
                </div>
                <div className="w-full h-10 flex flex-row justify-center gap-10 mt-2">
                    <button className="w-fit" onClick={next}>
                        next
                    </button>
                    <button className="w-fit" onClick={previous}>
                        previous
                    </button>
                </div>
            </div>
        </>
    );
}
