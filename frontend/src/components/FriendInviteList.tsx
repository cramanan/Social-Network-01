import React, { useEffect, useState } from "react";
import FriendInvite from "./FriendInvite";
import { User } from "@/types/user";

const FriendInviteList = () => {
    const [, setUsers] = useState<User[]>([]);

    useEffect(() => {
        const fetchUsers = async () => {
            const response = await fetch("/api/friend-requests");
            const data: User[] = await response.json();

            setUsers(data);
        };
        fetchUsers();
    }, []);

    return (
        <div
            id="friendInviteList"
            className="relative flex flex-col w-full h-full xl:w-fit xl:h-fit xl:bg-white/25 xl:rounded-[30px] xl:px-2 xl:py-5"
        >
            <h2 className="text-4xl text-white text-center py-5 xl:sr-only">
                Friend Request List
            </h2>

            <div className="flex flex-col h-fit items-center gap-3 mx-2 overflow-scroll no-scrollbar xl:max-h-[68vh] xl:gap-1">
                {users.length > 0 ? (
                    users.map((user, idx) =>
                        <FriendInvite key={idx} {...user} />
                    )
                ) : (
                    <p className="text-center font-bold">No invite(s) found.</p>
                )}

            </div>
        </div>
    );
};

export default FriendInviteList;
