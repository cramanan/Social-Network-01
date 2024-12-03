import { OnlineUser } from "@/types/user";
import React, { useEffect, useState } from 'react'

export const FollowersList = ({ groupId }: { groupId: string }) => {
    const [users, setUsers] = useState<OnlineUser[]>([]);

    useEffect(() => {
        const fetchUsers = async () => {
            const response = await fetch("/api/profile/following");
            const data: OnlineUser[] = await response.json();

            setUsers(data);
        };
        fetchUsers();
    }, []);

    const handleInvitation = (userId: string) => {
        fetch(`/api/groups/${groupId}/invite`, {
            method: "POST",
            headers: {
                "Content-Type": "application/json"
            },
            body: JSON.stringify({ groupId: groupId, userId: userId })
        });
    }

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
                            <>
                                {/* <Users
                                    key={idx}
                                    user={user}
                                    showLastMessage={false}
                                /> */}
                                <li key={idx} className="flex flex-row relative w-full justify-between z-30">
                                    {user.nickname}
                                    <input type="button" value="Send Invite" onClick={() => handleInvitation(user.id)} className="cursor-pointer" />
                                </li>
                            </>
                        ))
                    ) : (
                        <p className="text-center font-bold">
                            No follow(s) found.
                        </p>
                    )}
                </div>
            </div>
        </>
    )
}
