'use client'

import { OnlineUser } from "@/types/user";
import React, { useEffect, useState } from 'react'

export const MemberGroupList = () => {
    const [users, setUsers] = useState<OnlineUser[]>([]);

    useEffect(() => {
        const fetchUsers = async () => {
            const response = await fetch("/api/follow-list"); //TODO: change fetch route
            const data: OnlineUser[] = await response.json();

            setUsers(data);
        };
        fetchUsers();
    }, []);

    return (
        <>
            <div>MemberGroupList</div>
        </>
    )
}
