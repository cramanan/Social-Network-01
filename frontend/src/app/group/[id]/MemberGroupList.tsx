"use client";

import { OnlineUser } from "@/types/user";
import React, { useEffect, useState } from "react";

export const MemberGroupList = ({ groupId }: { groupId: string }) => {
    const [users, setUsers] = useState<OnlineUser[]>([]);

    useEffect(() => {
        const fetchUsers = async () => {
            const response = await fetch(`/api/groups/${groupId}/members`);
            const data: OnlineUser[] = await response.json();

            setUsers(data);
        };
        fetchUsers();
    }, [groupId]);

    return (
        <>
            <ul>
                {users.map((user, idx) => (
                    <li key={idx}>{user.nickname}</li>
                ))}
            </ul>
        </>
    );
};
